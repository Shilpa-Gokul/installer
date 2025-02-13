// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkAddressGroupAddMember network address group add member
//
// swagger:model NetworkAddressGroupAddMember
type NetworkAddressGroupAddMember struct {

	// The member to add in CIDR format
	// Required: true
	Cidr *string `json:"cidr"`
}

// Validate validates this network address group add member
func (m *NetworkAddressGroupAddMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkAddressGroupAddMember) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network address group add member based on context it is used
func (m *NetworkAddressGroupAddMember) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkAddressGroupAddMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkAddressGroupAddMember) UnmarshalBinary(b []byte) error {
	var res NetworkAddressGroupAddMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
