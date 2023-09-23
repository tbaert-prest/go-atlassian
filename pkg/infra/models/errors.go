package models

import (
	"errors"
)

var (
	ErrNoSCIMComplexValueError             = errors.New("admin: no scim complex value set")
	ErrNoSCIMValueError                    = errors.New("admin: no scim value set")
	ErrNoSCIMPathError                     = errors.New("admin: no scim path set")
	ErrNoSCIMOperationError                = errors.New("admin: no scim operation value set")
	ErrNoAdminOrganizationError            = errors.New("admin: no organization id set")
	ErrNoAdminDomainIDError                = errors.New("admin: no domain id set")
	ErrNoEventIDError                      = errors.New("admin: no event id set")
	ErrNoAdminPolicyError                  = errors.New("admin: no organization policy id set")
	ErrNoAdminDirectoryIDError             = errors.New("admin: no directory id set")
	ErrNoAdminGroupIDError                 = errors.New("admin: no group id set")
	ErrNoAdminGroupNameError               = errors.New("admin: no group name set")
	ErrNoAdminUserIDError                  = errors.New("admin: no user id set")
	ErrNoAdminAccountIDError               = errors.New("admin: no account id set")
	ErrNoAdminUserTokenError               = errors.New("admin: no user token id set")
	ErrNoCustomerMailError                 = errors.New("sm: no customer mail set")
	ErrNoCustomerDisplayNameError          = errors.New("sm: no customer display name set")
	ErrNoKBQueryError                      = errors.New("sm: no knowledge base query set")
	ErrNoOrganizationNameError             = errors.New("sm: no organization name set")
	ErrNoOrganizationIDError               = errors.New("sm: no organization id set")
	ErrNoCommentBodyError                  = errors.New("sm/jira: no comment body set")
	ErrNoServiceDeskIDError                = errors.New("sm: no service desk id set")
	ErrNoQueueIDError                      = errors.New("sm: no service desk queue id set")
	ErrNoRequestTypeIDError                = errors.New("sm: no request type id set")
	ErrNoFileNameError                     = errors.New("sm: no file name set")
	ErrNoFileReaderError                   = errors.New("sm: no io.Reader set")
	ErrNoCustomRequestFieldsError          = errors.New("sm: no customer request fields set")
	ErrNoSLAMetricIDError                  = errors.New("sm: no sla metric id set")
	ErrNoContentAttachmentIDError          = errors.New("confluence: no attachment id set")
	ErrNoContentAttachmentNameError        = errors.New("confluence: no attachment filename set")
	ErrNoContentReaderError                = errors.New("confluence: no reader set")
	ErrNoContentIDError                    = errors.New("confluence: no content id set")
	ErrNoPageIDError                       = errors.New("confluence: no page id set")
	ErrNoSpaceIDError                      = errors.New("confluence: no space id set")
	ErrNoTargetIDError                     = errors.New("confluence: no target id set")
	ErrNoPositionError                     = errors.New("confluence: no position set")
	ErrInvalidPositionError                = errors.New("confluence: invalid position: (before, after, append)")
	ValidPositions                         = map[string]bool{"before": true, "after": true, "append": true}
	ErrNoLabelIDError                      = errors.New("confluence: no label id set")
	ErrNoCQLError                          = errors.New("confluence: no CQL query set")
	ErrNoContentTypeError                  = errors.New("confluence: no content type set")
	ErrNoEntityIDError                     = errors.New("confluence: no entity id set")
	ValidEntityValues                      = []string{"blogposts", "custom-content", "labels", "pages"}
	ErrNoEntityValue                       = errors.New("confluence: no valid entity id set")
	ErrInvalidContentTypeError             = errors.New("confluence: invalid content type: (page, comment, attachment)")
	ValidContentTypes                      = []string{"page", "comment", "attachment"}
	ErrNoContentLabelError                 = errors.New("confluence: no content label set")
	ErrNoContentPropertyError              = errors.New("confluence: no content property set")
	ErrNoSpaceNameError                    = errors.New("confluence: no space name set")
	ErrNoSpaceKeyError                     = errors.New("confluence: no space key set")
	ErrNoContentRestrictionKeyError        = errors.New("confluence: no content restriction operation key set")
	ErrNoConfluenceGroupError              = errors.New("confluence: no group id or name set")
	ErrNoLabelNameError                    = errors.New("confluence: no label name set")
	ErrNoBoardIDError                      = errors.New("agile: no board id set")
	ErrNoFilterIDError                     = errors.New("agile: no filter id set")
	ErrNoNotificationSchemeIDError         = errors.New("jira: no notification scheme id set")
	ErrNoNotificationIDError               = errors.New("jira: no notification id set")
	ErrNoEpicIDError                       = errors.New("agile: no epic id set")
	ErrNoSprintIDError                     = errors.New("agile: no sprint id set")
	ErrNoApplicationRoleError              = errors.New("jira: no application role key set")
	ErrNoDashboardIDError                  = errors.New("jira: no dashboard id set")
	ErrNoGroupNameError                    = errors.New("jira: no group name set")
	ErrNoGroupIDError                      = errors.New("jira: no group name set")
	ErrNoGroupsNameError                   = errors.New("jira: no groups names set")
	ErrNoIssueKeyOrIDError                 = errors.New("jira: no issue key/id set")
	ErrNoRemoteLinkIDError                 = errors.New("jira: no remote link id set")
	ErrNoRemoteLinkGlobalIDError           = errors.New("jira: no global remote link id set")
	ErrNoTransitionIDError                 = errors.New("jira: no transition id set")
	ErrNoAttachmentIDError                 = errors.New("jira: no attachment id set")
	ErrNoAttachmentNameError               = errors.New("jira: no attachment filename set")
	ErrNoReaderError                       = errors.New("jira: no reader set")
	ErrNoCommentIDError                    = errors.New("jira: no comment id set")
	ErrNoProjectIDError                    = errors.New("jira: no project id set")
	ErrNoProjectIDOrKeyError               = errors.New("jira: no project id or key set")
	ErrNoProjectRoleIDError                = errors.New("jira: no project role id set")
	ErrNoProjectCategoryIDError            = errors.New("jira: no project category id set")
	ErrNoPropertyKeyError                  = errors.New("jira: no property key set")
	ErrNoProjectFeatureKeyError            = errors.New("jira: no project feature key set")
	ErrNoProjectFeatureStateError          = errors.New("jira: no project state key set")
	ErrNoFieldIDError                      = errors.New("jira: no field id set")
	ErrNoEditOperatorError                 = errors.New("jira: no update operation set")
	ErrNoOperatorError                     = errors.New("jira: no operation set")
	ErrNoEditValueError                    = errors.New("jira: no update operation value set")
	ErrNoCustomFieldError                  = errors.New("jira: no custom-fields set")
	ErrNoCustomFieldIDError                = errors.New("jira: no custom-field id set")
	ErrNoWorkflowStatusesError             = errors.New("jira: no workflow statuses set")
	ErrNoWorkflowScopeError                = errors.New("jira: no workflow scope set")
	ErrNoWorkflowStatusNameOrIdError       = errors.New("jira: no workflow status name or id set")
	ErrNoFieldContextIDError               = errors.New("jira: no field context id set")
	ErrNoIssueTypesError                   = errors.New("jira: no issue types id's set")
	ErrNoProjectsError                     = errors.New("jira: no projects set")
	ErrNoContextOptionIDError              = errors.New("jira: no field context option id set")
	ErrNoTypeIDError                       = errors.New("jira: no link id set")
	ErrNoLinkTypeIDError                   = errors.New("jira: no link type id set")
	ErrNoPriorityIDError                   = errors.New("jira: no priority id set")
	ErrNoResolutionIDError                 = errors.New("jira: no resolution id set")
	ErrNoJQLError                          = errors.New("jira: no sql set")
	ErrNoIssueTypeIDError                  = errors.New("jira: no issue type id set")
	ErrNoIssueTypeScreenSchemeIDError      = errors.New("jira: no issue type screen scheme id set")
	ErrNoScreenSchemeIDError               = errors.New("jira: no screen scheme id set")
	ErrNoAccountIDError                    = errors.New("jira: no account id set")
	ErrNoWorklogIDError                    = errors.New("jira: no worklog id set")
	ErrNpWorklogsError                     = errors.New("jira: no worklog's id set")
	ErrNoPermissionSchemeIDError           = errors.New("jira: no permission scheme id set")
	ErrNoPermissionGrantIDError            = errors.New("jira: no permission grant id set")
	ErrNoPermissionKeysError               = errors.New("jira: no permission keys set")
	ErrNoComponentIDError                  = errors.New("jira: no component id set")
	ErrProjectTypeKeyError                 = errors.New("jira: no project type key set")
	ErrNoProjectNameError                  = errors.New("jira: no project name set")
	ErrNoVersionIDError                    = errors.New("jira: no version id set")
	ErrNoScreenNameError                   = errors.New("jira: no screen name set")
	ErrNoScreenTabNameError                = errors.New("jira: no screen tab name set")
	ErrNoAccountSliceError                 = errors.New("jira: no account id's set")
	ErrNoProjectKeySliceError              = errors.New("jira: no project key's set")
	ErrNoProjectIDsError                   = errors.New("jira: no project id's set")
	ErrNoWorkflowIDError                   = errors.New("jira: no workflow id set")
	ErrNoWorkflowSchemeIDError             = errors.New("jira: no workflow scheme id set")
	ErrNoScreenIDError                     = errors.New("jira: no screen id set")
	ErrNoScreenTabIDError                  = errors.New("jira: no screen tab id set")
	ErrNoFieldConfigurationNameError       = errors.New("jira: no field configuration name set")
	ErrNoFieldConfigurationIDError         = errors.New("jira: no field configuration id set")
	ErrNoFieldConfigurationSchemeNameError = errors.New("jira: no field configuration scheme name set")
	ErrNoFieldConfigurationSchemeIDError   = errors.New("jira: no field configuration scheme id set")
	ErrNoVersionProvided                   = errors.New("client: no module version set")
	ErrNoIssueTypeSchemeIDError            = errors.New("jira: no issue type scheme id set")
	ErrNoTaskIDError                       = errors.New("atlassian: no task id set")
	ErrNoApprovalIDError                   = errors.New("jira: no approval id set")
	ErrInvalidStatusCodeError              = errors.New("client: invalid http response status, please refer the response.body for more details")
	ErrNotFound                            = errors.New("client: no atlassian resource found")
	ErrUnauthorized                        = errors.New("client: atlassian insufficient permissions")
	ErrInternalError                       = errors.New("client: atlassian internal error")
	ErrBadRequestError                     = errors.New("client: atlassian invalid payload")
	ErrNoSiteError                         = errors.New("client: no atlassian site set")
	ErrNilPayloadError                     = errors.New("client: please provide the necessary payload struct")
	ErrNonPayloadPointerError              = errors.New("client: please provide a valid payload struct pointer (&)")
	ErrNoFieldInformationError             = errors.New("custom-field: no json fields object found")
	ErrNoFloatTypeError                    = errors.New("custom-field: no float type set")
	ErrNoCustomFieldUnmarshalError         = errors.New("custom-field: no valid json provided")
	ErrNoMultiSelectTypeError              = errors.New("custom-field: no multiselect type found")
	ErrNoSprintTypeError                   = errors.New("custom-field: no sprint type found")
	ErrNoMultiVersionTypeError             = errors.New("custom-field: no multiversion type found")
	ErrNoUrlTypeError                      = errors.New("custom-field: no url type set")
	ErrNoTextTypeError                     = errors.New("custom-field: no text type set")
	ErrNoDateTimeTypeError                 = errors.New("custom-field: no date-time type set")
	ErrNoDateTypeError                     = errors.New("custom-field: no date type set")
	ErrNoSelectTypeError                   = errors.New("custom-field: no select type set")
	ErrNoButtonTypeError                   = errors.New("custom-field: no button type set")
	ErrNoUserTypeError                     = errors.New("custom-field: no user type set")
	ErrNoLabelsTypeError                   = errors.New("custom-field: no labels type set")
	ErrNoMultiUserTypeError                = errors.New("custom-field: no multi-user type set")
	ErrNoCheckBoxTypeError                 = errors.New("custom-field: no check-box type set")
	ErrNoCascadingParentError              = errors.New("custom-field: no cascading parent value set")
	ErrNoCascadingChildError               = errors.New("custom-field: no cascading child value set")
	ErrNoAttachmentIdsError                = errors.New("sm: no attachment id's set")
	ErrNoLabelsError                       = errors.New("sm: no label names set")
	ErrNoComponentsError                   = errors.New("sm: no components set")
	ErrNoCreateIssuesError                 = errors.New("jira: no issues payload set")
	ErrNoIssueSchemeError                  = errors.New("jira: no issue instance set")
	ErrNoIssuesSliceError                  = errors.New("jira: no issues object set")
	ErrNoMapValuesError                    = errors.New("jira: no map values set")
)
