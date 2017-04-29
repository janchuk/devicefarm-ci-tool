// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Access denied. Check your permissions.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeBatchWriteException for service response error code
	// "BatchWriteException".
	//
	// A BatchWrite exception has occurred.
	ErrCodeBatchWriteException = "BatchWriteException"

	// ErrCodeCannotListParentOfRootException for service response error code
	// "CannotListParentOfRootException".
	//
	// Cannot list the parents of a Directory root.
	ErrCodeCannotListParentOfRootException = "CannotListParentOfRootException"

	// ErrCodeDirectoryAlreadyExistsException for service response error code
	// "DirectoryAlreadyExistsException".
	//
	// Indicates that a Directory could not be created due to a naming conflict.
	// Choose a different name and try again.
	ErrCodeDirectoryAlreadyExistsException = "DirectoryAlreadyExistsException"

	// ErrCodeDirectoryDeletedException for service response error code
	// "DirectoryDeletedException".
	//
	// A directory that has been deleted has been attempted to be accessed. Note:
	// The requested resource will eventually cease to exist.
	ErrCodeDirectoryDeletedException = "DirectoryDeletedException"

	// ErrCodeDirectoryNotDisabledException for service response error code
	// "DirectoryNotDisabledException".
	//
	// An operation can only operate on a disabled directory.
	ErrCodeDirectoryNotDisabledException = "DirectoryNotDisabledException"

	// ErrCodeDirectoryNotEnabledException for service response error code
	// "DirectoryNotEnabledException".
	//
	// An operation can only operate on a directory that is not enabled.
	ErrCodeDirectoryNotEnabledException = "DirectoryNotEnabledException"

	// ErrCodeFacetAlreadyExistsException for service response error code
	// "FacetAlreadyExistsException".
	//
	// A facet with the same name already exists.
	ErrCodeFacetAlreadyExistsException = "FacetAlreadyExistsException"

	// ErrCodeFacetInUseException for service response error code
	// "FacetInUseException".
	//
	// Occurs when deleting a facet that contains an attribute which is a target
	// to an attribute reference in a different facet.
	ErrCodeFacetInUseException = "FacetInUseException"

	// ErrCodeFacetNotFoundException for service response error code
	// "FacetNotFoundException".
	//
	// The specified Facet could not be found.
	ErrCodeFacetNotFoundException = "FacetNotFoundException"

	// ErrCodeFacetValidationException for service response error code
	// "FacetValidationException".
	//
	// The Facet you provided was not well formed or could not be validated with
	// the schema.
	ErrCodeFacetValidationException = "FacetValidationException"

	// ErrCodeIndexedAttributeMissingException for service response error code
	// "IndexedAttributeMissingException".
	//
	// An object has been attempted to be attached to an object that does not have
	// the appropriate attribute value.
	ErrCodeIndexedAttributeMissingException = "IndexedAttributeMissingException"

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	//
	// Indicates a problem that must be resolved by Amazon Web Services. This might
	// be a transient error in which case you can retry your request until it succeeds.
	// Otherwise, go to the AWS Service Health Dashboard (http://status.aws.amazon.com/)
	// site to see if there are any operational issues with the service.
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidArnException for service response error code
	// "InvalidArnException".
	//
	// Indicates that the provided ARN value is not valid.
	ErrCodeInvalidArnException = "InvalidArnException"

	// ErrCodeInvalidAttachmentException for service response error code
	// "InvalidAttachmentException".
	//
	// Indicates that an attempt to attach an object with the same link name or
	// to apply a schema with same name has occurred. Rename the link or the schema
	// and then try again.
	ErrCodeInvalidAttachmentException = "InvalidAttachmentException"

	// ErrCodeInvalidFacetUpdateException for service response error code
	// "InvalidFacetUpdateException".
	//
	// An attempt to modify a Facet resulted in an invalid schema exception.
	ErrCodeInvalidFacetUpdateException = "InvalidFacetUpdateException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// Indicates that the NextToken value is not valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidRuleException for service response error code
	// "InvalidRuleException".
	//
	// Occurs when any of the rule parameter keys or values are invalid.
	ErrCodeInvalidRuleException = "InvalidRuleException"

	// ErrCodeInvalidSchemaDocException for service response error code
	// "InvalidSchemaDocException".
	//
	// Indicates that the provided SchemaDoc value is not valid.
	ErrCodeInvalidSchemaDocException = "InvalidSchemaDocException"

	// ErrCodeInvalidTaggingRequestException for service response error code
	// "InvalidTaggingRequestException".
	//
	// Can occur for multiple reasons such as when you tag a resource that doesn’t
	// exist or if you specify a higher number of tags for a resource than the allowed
	// limit. Allowed limit is 50 tags per resource.
	ErrCodeInvalidTaggingRequestException = "InvalidTaggingRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Indicates limits are exceeded. See Limits (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/limits.html)
	// for more information.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeLinkNameAlreadyInUseException for service response error code
	// "LinkNameAlreadyInUseException".
	//
	// Indicates that a link could not be created due to a naming conflict. Choose
	// a different name and then try again.
	ErrCodeLinkNameAlreadyInUseException = "LinkNameAlreadyInUseException"

	// ErrCodeNotIndexException for service response error code
	// "NotIndexException".
	//
	// Indicates the requested operation can only operate on index objects.
	ErrCodeNotIndexException = "NotIndexException"

	// ErrCodeNotNodeException for service response error code
	// "NotNodeException".
	//
	// Occurs when any invalid operations are performed on an object which is not
	// a node, such as calling ListObjectChildren for a leaf node object.
	ErrCodeNotNodeException = "NotNodeException"

	// ErrCodeNotPolicyException for service response error code
	// "NotPolicyException".
	//
	// Indicates the requested operation can only operate on policy objects.
	ErrCodeNotPolicyException = "NotPolicyException"

	// ErrCodeObjectAlreadyDetachedException for service response error code
	// "ObjectAlreadyDetachedException".
	//
	// Indicates the object is not attached to the index.
	ErrCodeObjectAlreadyDetachedException = "ObjectAlreadyDetachedException"

	// ErrCodeObjectNotDetachedException for service response error code
	// "ObjectNotDetachedException".
	//
	// Indicates the requested operation cannot be completed because the object
	// has not been detached from the tree.
	ErrCodeObjectNotDetachedException = "ObjectNotDetachedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeRetryableConflictException for service response error code
	// "RetryableConflictException".
	//
	// Occurs when a conflict with a previous successful write is detected. For
	// example, if a write operation occurs on an object and then an attempt is
	// made to read the object using “SERIALIZABLE” consistency, this exception
	// may result. This generally occurs when the previous write did not have time
	// to propagate to the host serving the current request. A retry (with appropriate
	// backoff logic) is the recommended response to this exception.
	ErrCodeRetryableConflictException = "RetryableConflictException"

	// ErrCodeSchemaAlreadyExistsException for service response error code
	// "SchemaAlreadyExistsException".
	//
	// Indicates that a schema could not be created due to a naming conflict. Please
	// select a different name and then try again.
	ErrCodeSchemaAlreadyExistsException = "SchemaAlreadyExistsException"

	// ErrCodeSchemaAlreadyPublishedException for service response error code
	// "SchemaAlreadyPublishedException".
	//
	// Indicates a schema is already published.
	ErrCodeSchemaAlreadyPublishedException = "SchemaAlreadyPublishedException"

	// ErrCodeStillContainsLinksException for service response error code
	// "StillContainsLinksException".
	//
	// The object could not be deleted because links still exist. Remove the links
	// and then try the operation again.
	ErrCodeStillContainsLinksException = "StillContainsLinksException"

	// ErrCodeUnsupportedIndexTypeException for service response error code
	// "UnsupportedIndexTypeException".
	//
	// Indicates the requested index type is not supported.
	ErrCodeUnsupportedIndexTypeException = "UnsupportedIndexTypeException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// Indicates your request is malformed in some manner. See the exception message.
	ErrCodeValidationException = "ValidationException"
)
