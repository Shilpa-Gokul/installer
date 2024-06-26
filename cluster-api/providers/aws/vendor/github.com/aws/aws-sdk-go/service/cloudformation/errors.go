// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

const (

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExistsException".
	//
	// The resource with the name requested already exists.
	ErrCodeAlreadyExistsException = "AlreadyExistsException"

	// ErrCodeCFNRegistryException for service response error code
	// "CFNRegistryException".
	//
	// An error occurred during a CloudFormation registry operation.
	ErrCodeCFNRegistryException = "CFNRegistryException"

	// ErrCodeChangeSetNotFoundException for service response error code
	// "ChangeSetNotFound".
	//
	// The specified change set name or ID doesn't exit. To view valid change sets
	// for a stack, use the ListChangeSets operation.
	ErrCodeChangeSetNotFoundException = "ChangeSetNotFound"

	// ErrCodeConcurrentResourcesLimitExceededException for service response error code
	// "ConcurrentResourcesLimitExceeded".
	//
	// No more than 5 generated templates can be in an InProgress or Pending status
	// at one time. This error is also returned if a generated template that is
	// in an InProgress or Pending status is attempted to be updated or deleted.
	ErrCodeConcurrentResourcesLimitExceededException = "ConcurrentResourcesLimitExceeded"

	// ErrCodeCreatedButModifiedException for service response error code
	// "CreatedButModifiedException".
	//
	// The specified resource exists, but has been changed.
	ErrCodeCreatedButModifiedException = "CreatedButModifiedException"

	// ErrCodeGeneratedTemplateNotFoundException for service response error code
	// "GeneratedTemplateNotFound".
	//
	// The generated template was not found.
	ErrCodeGeneratedTemplateNotFoundException = "GeneratedTemplateNotFound"

	// ErrCodeInsufficientCapabilitiesException for service response error code
	// "InsufficientCapabilitiesException".
	//
	// The template contains resources with capabilities that weren't specified
	// in the Capabilities parameter.
	ErrCodeInsufficientCapabilitiesException = "InsufficientCapabilitiesException"

	// ErrCodeInvalidChangeSetStatusException for service response error code
	// "InvalidChangeSetStatus".
	//
	// The specified change set can't be used to update the stack. For example,
	// the change set status might be CREATE_IN_PROGRESS, or the stack status might
	// be UPDATE_IN_PROGRESS.
	ErrCodeInvalidChangeSetStatusException = "InvalidChangeSetStatus"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// The specified operation isn't valid.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeInvalidStateTransitionException for service response error code
	// "InvalidStateTransition".
	//
	// Error reserved for use by the CloudFormation CLI (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	// CloudFormation doesn't return this error to users.
	ErrCodeInvalidStateTransitionException = "InvalidStateTransition"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The quota for the resource has already been reached.
	//
	// For information about resource and stack limitations, see CloudFormation
	// quotas (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cloudformation-limits.html)
	// in the CloudFormation User Guide.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNameAlreadyExistsException for service response error code
	// "NameAlreadyExistsException".
	//
	// The specified name is already in use.
	ErrCodeNameAlreadyExistsException = "NameAlreadyExistsException"

	// ErrCodeOperationIdAlreadyExistsException for service response error code
	// "OperationIdAlreadyExistsException".
	//
	// The specified operation ID already exists.
	ErrCodeOperationIdAlreadyExistsException = "OperationIdAlreadyExistsException"

	// ErrCodeOperationInProgressException for service response error code
	// "OperationInProgressException".
	//
	// Another operation is currently in progress for this stack set. Only one operation
	// can be performed for a stack set at a given time.
	ErrCodeOperationInProgressException = "OperationInProgressException"

	// ErrCodeOperationNotFoundException for service response error code
	// "OperationNotFoundException".
	//
	// The specified ID refers to an operation that doesn't exist.
	ErrCodeOperationNotFoundException = "OperationNotFoundException"

	// ErrCodeOperationStatusCheckFailedException for service response error code
	// "ConditionalCheckFailed".
	//
	// Error reserved for use by the CloudFormation CLI (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).
	// CloudFormation doesn't return this error to users.
	ErrCodeOperationStatusCheckFailedException = "ConditionalCheckFailed"

	// ErrCodeResourceScanInProgressException for service response error code
	// "ResourceScanInProgress".
	//
	// A resource scan is currently in progress. Only one can be run at a time for
	// an account in a Region.
	ErrCodeResourceScanInProgressException = "ResourceScanInProgress"

	// ErrCodeResourceScanLimitExceededException for service response error code
	// "ResourceScanLimitExceeded".
	//
	// The limit on resource scans has been exceeded. Reasons include:
	//
	//    * Exceeded the daily quota for resource scans.
	//
	//    * A resource scan recently failed. You must wait 10 minutes before starting
	//    a new resource scan.
	//
	//    * The last resource scan failed after exceeding 100,000 resources. When
	//    this happens, you must wait 24 hours before starting a new resource scan.
	ErrCodeResourceScanLimitExceededException = "ResourceScanLimitExceeded"

	// ErrCodeResourceScanNotFoundException for service response error code
	// "ResourceScanNotFound".
	//
	// The resource scan was not found.
	ErrCodeResourceScanNotFoundException = "ResourceScanNotFound"

	// ErrCodeStackInstanceNotFoundException for service response error code
	// "StackInstanceNotFoundException".
	//
	// The specified stack instance doesn't exist.
	ErrCodeStackInstanceNotFoundException = "StackInstanceNotFoundException"

	// ErrCodeStackNotFoundException for service response error code
	// "StackNotFoundException".
	//
	// The specified stack ARN doesn't exist or stack doesn't exist corresponding
	// to the ARN in input.
	ErrCodeStackNotFoundException = "StackNotFoundException"

	// ErrCodeStackSetNotEmptyException for service response error code
	// "StackSetNotEmptyException".
	//
	// You can't yet delete this stack set, because it still contains one or more
	// stack instances. Delete all stack instances from the stack set before deleting
	// the stack set.
	ErrCodeStackSetNotEmptyException = "StackSetNotEmptyException"

	// ErrCodeStackSetNotFoundException for service response error code
	// "StackSetNotFoundException".
	//
	// The specified stack set doesn't exist.
	ErrCodeStackSetNotFoundException = "StackSetNotFoundException"

	// ErrCodeStaleRequestException for service response error code
	// "StaleRequestException".
	//
	// Another operation has been performed on this stack set since the specified
	// operation was performed.
	ErrCodeStaleRequestException = "StaleRequestException"

	// ErrCodeTokenAlreadyExistsException for service response error code
	// "TokenAlreadyExistsException".
	//
	// A client request token already exists.
	ErrCodeTokenAlreadyExistsException = "TokenAlreadyExistsException"

	// ErrCodeTypeConfigurationNotFoundException for service response error code
	// "TypeConfigurationNotFoundException".
	//
	// The specified extension configuration can't be found.
	ErrCodeTypeConfigurationNotFoundException = "TypeConfigurationNotFoundException"

	// ErrCodeTypeNotFoundException for service response error code
	// "TypeNotFoundException".
	//
	// The specified extension doesn't exist in the CloudFormation registry.
	ErrCodeTypeNotFoundException = "TypeNotFoundException"
)
