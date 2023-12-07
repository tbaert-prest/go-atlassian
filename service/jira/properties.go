package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

/*
IssuePropertyConnector represents issue properties, which provides for storing custom data against an issue.

Use it to get, set, and delete issue properties as well as obtain details of all properties on an issue.

Operations to bulk update and delete issue properties are also provided.
*/
type IssuePropertyConnector interface {

	/*
		Gets returns the URLs and keys of an issue's properties.
			- This operation can be accessed anonymously.

		Permissions required:
			- Browse projects project permission for the project containing the issue.
			- If issue-level security is configured, issue-level security permission to view the issue.

		Endpoint: GET /rest/api/{apiVersion}/issue/{issueIdOrKey}/properties

		You can refer to the documentation: [Get issue property keys]

		[Get issue property keys]: https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
	*/
	Gets(ctx context.Context, issueIdOrKey string) (*model.PropertyPageScheme, *model.ResponseScheme, error)

	/*
		Get returns the key and value of an issue's property.
			- This operation can be accessed anonymously.

		Permissions required:
			- Browse projects project permission for the project containing the issue.
			- If issue-level security is configured, issue-level security permission to view the issue.

		Endpoint: GET /rest/api/{apiVersion}/issue/{issueIdOrKey}/properties/{propertyKey}

		You can refer to the documentation: [Get issue property]

		[Get issue property]: https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
	*/
	Get(ctx context.Context, issueKey, propertyKey string) (*model.EntityPropertyScheme, *model.ResponseScheme, error)

	/*
		Set sets the value of an issue's property. Use this resource to store custom data against an issue.
			- The value of the request body must be a valid, non-empty JSON blob. The maximum length is 32768 characters.
			- This operation can be accessed anonymously.

		Permissions required:
			- Browse projects and Edit issues project permissions for the project containing the issue.
			- If issue-level security is configured, issue-level security permission to view the issue.

		Endpoint: PUT /rest/api/{apiVersion}/issue/{issueIdOrKey}/properties/{propertyKey}

		You can refer to the documentation: [Set issue property]

		[Set issue property]: https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
	*/
	Set(ctx context.Context, issueKey, propertyKey string, payload interface{}) (*model.ResponseScheme, error)

	/*
		Delete deletes an issue's property.
			- This operation can be accessed anonymously.

		Permissions required:
			- Browse projects and Edit issues project permissions for the project containing the issue.
			- If issue-level security is configured, issue-level security permission to view the issue.

		Endpoint: DELETE /rest/api/{apiVersion}/issue/{issueIdOrKey}/properties/{propertyKey}

		You can refer to the documentation: [Delete issue property]

		[Delete issue property]: https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
	*/
	Delete(ctx context.Context, issueKey, propertyKey string) (*model.ResponseScheme, error)
}
