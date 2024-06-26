package gophercloud

import (
	"context"
	"net/url"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// NormalizePathURL is used to convert rawPath to a fqdn, using basePath as
// a reference in the filesystem, if necessary. basePath is assumed to contain
// either '.' when first used, or the file:// type fqdn of the parent resource.
// e.g. myFavScript.yaml => file://opt/lib/myFavScript.yaml
func NormalizePathURL(basePath, rawPath string) (string, error) {
	u, err := url.Parse(rawPath)
	if err != nil {
		return "", err
	}
	// if a scheme is defined, it must be a fqdn already
	if u.Scheme != "" {
		return u.String(), nil
	}
	// if basePath is a url, then child resources are assumed to be relative to it
	bu, err := url.Parse(basePath)
	if err != nil {
		return "", err
	}
	var basePathSys, absPathSys string
	if bu.Scheme != "" {
		basePathSys = filepath.FromSlash(bu.Path)
		absPathSys = filepath.Join(basePathSys, rawPath)
		bu.Path = filepath.ToSlash(absPathSys)
		return bu.String(), nil
	}

	absPathSys = filepath.Join(basePath, rawPath)
	u.Path = filepath.ToSlash(absPathSys)
	if err != nil {
		return "", err
	}
	u.Scheme = "file"
	return u.String(), nil
}

// NormalizeURL is an internal function to be used by provider clients.
//
// It ensures that each endpoint URL has a closing `/`, as expected by
// ServiceClient's methods.
func NormalizeURL(url string) string {
	if !strings.HasSuffix(url, "/") {
		return url + "/"
	}
	return url
}

// RemainingKeys will inspect a struct and compare it to a map. Any struct
// field that does not have a JSON tag that matches a key in the map or
// a matching lower-case field in the map will be returned as an extra.
//
// This is useful for determining the extra fields returned in response bodies
// for resources that can contain an arbitrary or dynamic number of fields.
func RemainingKeys(s any, m map[string]any) (extras map[string]any) {
	extras = make(map[string]any)
	for k, v := range m {
		extras[k] = v
	}

	valueOf := reflect.ValueOf(s)
	typeOf := reflect.TypeOf(s)
	for i := 0; i < valueOf.NumField(); i++ {
		field := typeOf.Field(i)

		lowerField := strings.ToLower(field.Name)
		delete(extras, lowerField)

		if tagValue := field.Tag.Get("json"); tagValue != "" && tagValue != "-" {
			delete(extras, tagValue)
		}
	}

	return
}

// WaitFor polls a predicate function, once per second, up to a context cancellation.
// This is useful to wait for a resource to transition to a certain state.
// Resource packages will wrap this in a more convenient function that's
// specific to a certain resource, but it can also be useful on its own.
func WaitFor(ctx context.Context, predicate func(context.Context) (bool, error)) error {
	if done, err := predicate(ctx); done || err != nil {
		return err
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if done, err := predicate(ctx); done || err != nil {
				return err
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
