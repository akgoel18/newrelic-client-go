// Code generated by tutone: DO NOT EDIT
package workflows

import (
	"encoding/json"
	"fmt"

	"github.com/newrelic/newrelic-client-go/v2/pkg/ai"
	"github.com/newrelic/newrelic-client-go/v2/pkg/nrtime"
)

// AiWorkflowsCreateErrorType - Type of create error
type AiWorkflowsCreateErrorType string

var AiWorkflowsCreateErrorTypeTypes = struct {
	// We couldn't find a channel with the given id
	CHANNEL_NOT_FOUND AiWorkflowsCreateErrorType
	// A workflow with this name already exists
	DUPLICATE AiWorkflowsCreateErrorType
	// One or more of the parameters you provided are incorrect
	INVALID_PARAMETER AiWorkflowsCreateErrorType
	// Reached the maximum number of workflows per account
	LIMIT_REACHED AiWorkflowsCreateErrorType
	// This account is missing the required entitlement(s) to perform this action
	MISSING_ENTITLEMENT AiWorkflowsCreateErrorType
	// This account in not authorized to perform this action
	UNAUTHORIZED_ACCOUNT AiWorkflowsCreateErrorType
	// The given channel id represents an unsupported channel type
	UNSUPPORTED_CHANNEL_TYPE AiWorkflowsCreateErrorType
	// The parameter provided does not have a valid form
	VALIDATION_ERROR AiWorkflowsCreateErrorType
}{
	// We couldn't find a channel with the given id
	CHANNEL_NOT_FOUND: "CHANNEL_NOT_FOUND",
	// A workflow with this name already exists
	DUPLICATE: "DUPLICATE",
	// One or more of the parameters you provided are incorrect
	INVALID_PARAMETER: "INVALID_PARAMETER",
	// Reached the maximum number of workflows per account
	LIMIT_REACHED: "LIMIT_REACHED",
	// This account is missing the required entitlement(s) to perform this action
	MISSING_ENTITLEMENT: "MISSING_ENTITLEMENT",
	// This account in not authorized to perform this action
	UNAUTHORIZED_ACCOUNT: "UNAUTHORIZED_ACCOUNT",
	// The given channel id represents an unsupported channel type
	UNSUPPORTED_CHANNEL_TYPE: "UNSUPPORTED_CHANNEL_TYPE",
	// The parameter provided does not have a valid form
	VALIDATION_ERROR: "VALIDATION_ERROR",
}

// AiWorkflowsDeleteErrorType - Type of delete error
type AiWorkflowsDeleteErrorType string

var AiWorkflowsDeleteErrorTypeTypes = struct {
	// One or more of the parameters you provided are incorrect
	INVALID_PARAMETER AiWorkflowsDeleteErrorType
	// This account in not authorized to perform this action
	UNAUTHORIZED_ACCOUNT AiWorkflowsDeleteErrorType
	// The parameter provided does not have a valid form
	VALIDATION_ERROR AiWorkflowsDeleteErrorType
}{
	// One or more of the parameters you provided are incorrect
	INVALID_PARAMETER: "INVALID_PARAMETER",
	// This account in not authorized to perform this action
	UNAUTHORIZED_ACCOUNT: "UNAUTHORIZED_ACCOUNT",
	// The parameter provided does not have a valid form
	VALIDATION_ERROR: "VALIDATION_ERROR",
}

// AiWorkflowsDestinationType - Type of Destination Configuration
type AiWorkflowsDestinationType string

var AiWorkflowsDestinationTypeTypes = struct {
	// Email Destination Configuration type
	EMAIL AiWorkflowsDestinationType
	// Event Bridge Destination Configuration type
	EVENT_BRIDGE AiWorkflowsDestinationType
	// Jira Destination Configuration type
	JIRA AiWorkflowsDestinationType
	// New Relic Mobile Push Destination Configuration type
	MOBILE_PUSH AiWorkflowsDestinationType
	// Pager Duty Destination Configuration type
	PAGERDUTY AiWorkflowsDestinationType
	// Pager Duty with account integration Destination Configuration type
	PAGERDUTY_ACCOUNT_INTEGRATION AiWorkflowsDestinationType
	// Pager Duty with service integration Destination Configuration type
	PAGERDUTY_SERVICE_INTEGRATION AiWorkflowsDestinationType
	// Service Now Destination Configuration type
	SERVICE_NOW AiWorkflowsDestinationType
	// Slack Destination Configuration type
	SLACK AiWorkflowsDestinationType
	// Slack legacy Destination Configuration type
	SLACK_LEGACY AiWorkflowsDestinationType
	// Webhook Destination Configuration type
	WEBHOOK AiWorkflowsDestinationType
}{
	// Email Destination Configuration type
	EMAIL: "EMAIL",
	// Event Bridge Destination Configuration type
	EVENT_BRIDGE: "EVENT_BRIDGE",
	// Jira Destination Configuration type
	JIRA: "JIRA",
	// New Relic Mobile Push Destination Configuration type
	MOBILE_PUSH: "MOBILE_PUSH",
	// Pager Duty Destination Configuration type
	PAGERDUTY: "PAGERDUTY",
	// Pager Duty with account integration Destination Configuration type
	PAGERDUTY_ACCOUNT_INTEGRATION: "PAGERDUTY_ACCOUNT_INTEGRATION",
	// Pager Duty with service integration Destination Configuration type
	PAGERDUTY_SERVICE_INTEGRATION: "PAGERDUTY_SERVICE_INTEGRATION",
	// Service Now Destination Configuration type
	SERVICE_NOW: "SERVICE_NOW",
	// Slack Destination Configuration type
	SLACK: "SLACK",
	// Slack legacy Destination Configuration type
	SLACK_LEGACY: "SLACK_LEGACY",
	// Webhook Destination Configuration type
	WEBHOOK: "WEBHOOK",
}

// AiWorkflowsEnrichmentType - Type of Enrichment
type AiWorkflowsEnrichmentType string

var AiWorkflowsEnrichmentTypeTypes = struct {
	// NRQL Enrichment type
	NRQL AiWorkflowsEnrichmentType
}{
	// NRQL Enrichment type
	NRQL: "NRQL",
}

// AiWorkflowsFilterType - Type of Filter
type AiWorkflowsFilterType string

var AiWorkflowsFilterTypeTypes = struct {
	// Standard Filter type
	FILTER AiWorkflowsFilterType
	// View Filter type
	VIEW AiWorkflowsFilterType
}{
	// Standard Filter type
	FILTER: "FILTER",
	// View Filter type
	VIEW: "VIEW",
}

// AiWorkflowsMutingRulesHandling - The wanted behavior for muted issues in the workflow
type AiWorkflowsMutingRulesHandling string

var AiWorkflowsMutingRulesHandlingTypes = struct {
	// Notify only about partially muted and unmuted issues
	DONT_NOTIFY_FULLY_MUTED_ISSUES AiWorkflowsMutingRulesHandling
	// Notify only about unmuted issues
	DONT_NOTIFY_FULLY_OR_PARTIALLY_MUTED_ISSUES AiWorkflowsMutingRulesHandling
	// Notify about all issues
	NOTIFY_ALL_ISSUES AiWorkflowsMutingRulesHandling
}{
	// Notify only about partially muted and unmuted issues
	DONT_NOTIFY_FULLY_MUTED_ISSUES: "DONT_NOTIFY_FULLY_MUTED_ISSUES",
	// Notify only about unmuted issues
	DONT_NOTIFY_FULLY_OR_PARTIALLY_MUTED_ISSUES: "DONT_NOTIFY_FULLY_OR_PARTIALLY_MUTED_ISSUES",
	// Notify about all issues
	NOTIFY_ALL_ISSUES: "NOTIFY_ALL_ISSUES",
}

// AiWorkflowsNotificationTrigger - Notification Triggers for the Destination Configuration
type AiWorkflowsNotificationTrigger string

var AiWorkflowsNotificationTriggerTypes = struct {
	// Send a notification when the issue is acknowledged
	ACKNOWLEDGED AiWorkflowsNotificationTrigger
	// Send a notification when the issue is activated
	ACTIVATED AiWorkflowsNotificationTrigger
	// Send a notification when the issue is closed
	CLOSED AiWorkflowsNotificationTrigger
	// Sends notification when the issue has other updates
	OTHER_UPDATES AiWorkflowsNotificationTrigger
	// Send a notification when the issue's priority has changed
	PRIORITY_CHANGED AiWorkflowsNotificationTrigger
}{
	// Send a notification when the issue is acknowledged
	ACKNOWLEDGED: "ACKNOWLEDGED",
	// Send a notification when the issue is activated
	ACTIVATED: "ACTIVATED",
	// Send a notification when the issue is closed
	CLOSED: "CLOSED",
	// Sends notification when the issue has other updates
	OTHER_UPDATES: "OTHER_UPDATES",
	// Send a notification when the issue's priority has changed
	PRIORITY_CHANGED: "PRIORITY_CHANGED",
}

// AiWorkflowsOperator - Type of Filter
type AiWorkflowsOperator string

var AiWorkflowsOperatorTypes = struct {
	// String or list attribute contains this value
	CONTAINS AiWorkflowsOperator
	// String or list attribute does not contain this value
	DOES_NOT_CONTAIN AiWorkflowsOperator
	// String or Numeric attribute does not equal this value
	DOES_NOT_EQUAL AiWorkflowsOperator
	// Element in list attribute does not exactly match this value
	DOES_NOT_EXACTLY_MATCH AiWorkflowsOperator
	// String attribute ends with this value
	ENDS_WITH AiWorkflowsOperator
	// String or Numeric attribute equals this value
	EQUAL AiWorkflowsOperator
	// Element in list attribute exactly matches this value
	EXACTLY_MATCHES AiWorkflowsOperator
	// Numeric attribute is greater or equal to this value
	GREATER_OR_EQUAL AiWorkflowsOperator
	// Numeric attribute is greater than this value
	GREATER_THAN AiWorkflowsOperator
	// Boolean attribute equals value
	IS AiWorkflowsOperator
	// Boolean attribute does not equal value
	IS_NOT AiWorkflowsOperator
	// Numeric attribute is less or equal to this value
	LESS_OR_EQUAL AiWorkflowsOperator
	// Numeric attribute is less than this value
	LESS_THAN AiWorkflowsOperator
	// String attribute starts with this value
	STARTS_WITH AiWorkflowsOperator
}{
	// String or list attribute contains this value
	CONTAINS: "CONTAINS",
	// String or list attribute does not contain this value
	DOES_NOT_CONTAIN: "DOES_NOT_CONTAIN",
	// String or Numeric attribute does not equal this value
	DOES_NOT_EQUAL: "DOES_NOT_EQUAL",
	// Element in list attribute does not exactly match this value
	DOES_NOT_EXACTLY_MATCH: "DOES_NOT_EXACTLY_MATCH",
	// String attribute ends with this value
	ENDS_WITH: "ENDS_WITH",
	// String or Numeric attribute equals this value
	EQUAL: "EQUAL",
	// Element in list attribute exactly matches this value
	EXACTLY_MATCHES: "EXACTLY_MATCHES",
	// Numeric attribute is greater or equal to this value
	GREATER_OR_EQUAL: "GREATER_OR_EQUAL",
	// Numeric attribute is greater than this value
	GREATER_THAN: "GREATER_THAN",
	// Boolean attribute equals value
	IS: "IS",
	// Boolean attribute does not equal value
	IS_NOT: "IS_NOT",
	// Numeric attribute is less or equal to this value
	LESS_OR_EQUAL: "LESS_OR_EQUAL",
	// Numeric attribute is less than this value
	LESS_THAN: "LESS_THAN",
	// String attribute starts with this value
	STARTS_WITH: "STARTS_WITH",
}

// AiWorkflowsTestErrorType - Type of test error
type AiWorkflowsTestErrorType string

var AiWorkflowsTestErrorTypeTypes = struct {
	// We couldn't find a channel with the given id
	CHANNEL_NOT_FOUND AiWorkflowsTestErrorType
	// Failed running test workflow
	FAILED_RUNNING_TEST AiWorkflowsTestErrorType
	// This account is missing the required entitlement(s) to perform this action
	MISSING_ENTITLEMENT AiWorkflowsTestErrorType
	// This account is not allowed to preform this action
	UNAUTHORIZED_ACCOUNT AiWorkflowsTestErrorType
	// The given channel id represents an unsupported channel type
	UNSUPPORTED_CHANNEL_TYPE AiWorkflowsTestErrorType
	// The parameter provided does not have a valid form
	VALIDATION_ERROR AiWorkflowsTestErrorType
	// Failed to send a notification to the channel
	WARNING_FAILED_SENDING_NOTIFICATION AiWorkflowsTestErrorType
	// There are no issues that match this filter
	WARNING_NO_FILTERED_ISSUE_FOUND AiWorkflowsTestErrorType
	// There are no issues that match these dynamic variables
	WARNING_NO_MATCHING_DYNAMIC_VARIABLES_FOUND AiWorkflowsTestErrorType
}{
	// We couldn't find a channel with the given id
	CHANNEL_NOT_FOUND: "CHANNEL_NOT_FOUND",
	// Failed running test workflow
	FAILED_RUNNING_TEST: "FAILED_RUNNING_TEST",
	// This account is missing the required entitlement(s) to perform this action
	MISSING_ENTITLEMENT: "MISSING_ENTITLEMENT",
	// This account is not allowed to preform this action
	UNAUTHORIZED_ACCOUNT: "UNAUTHORIZED_ACCOUNT",
	// The given channel id represents an unsupported channel type
	UNSUPPORTED_CHANNEL_TYPE: "UNSUPPORTED_CHANNEL_TYPE",
	// The parameter provided does not have a valid form
	VALIDATION_ERROR: "VALIDATION_ERROR",
	// Failed to send a notification to the channel
	WARNING_FAILED_SENDING_NOTIFICATION: "WARNING_FAILED_SENDING_NOTIFICATION",
	// There are no issues that match this filter
	WARNING_NO_FILTERED_ISSUE_FOUND: "WARNING_NO_FILTERED_ISSUE_FOUND",
	// There are no issues that match these dynamic variables
	WARNING_NO_MATCHING_DYNAMIC_VARIABLES_FOUND: "WARNING_NO_MATCHING_DYNAMIC_VARIABLES_FOUND",
}

// AiWorkflowsUpdateErrorType - Type of update error
type AiWorkflowsUpdateErrorType string

var AiWorkflowsUpdateErrorTypeTypes = struct {
	// We couldn't find a channel with the given id
	CHANNEL_NOT_FOUND AiWorkflowsUpdateErrorType
	// A workflow with this name already exists
	DUPLICATE AiWorkflowsUpdateErrorType
	// One or more of the parameters you provided are incorrect
	INVALID_PARAMETER AiWorkflowsUpdateErrorType
	// This account is missing the required entitlement(s) to perform this action
	MISSING_ENTITLEMENT AiWorkflowsUpdateErrorType
	// This account in not authorized to perform this action
	UNAUTHORIZED_ACCOUNT AiWorkflowsUpdateErrorType
	// The given channel id represents an unsupported channel type
	UNSUPPORTED_CHANNEL_TYPE AiWorkflowsUpdateErrorType
	// The parameter provided does not have a valid form
	VALIDATION_ERROR AiWorkflowsUpdateErrorType
}{
	// We couldn't find a channel with the given id
	CHANNEL_NOT_FOUND: "CHANNEL_NOT_FOUND",
	// A workflow with this name already exists
	DUPLICATE: "DUPLICATE",
	// One or more of the parameters you provided are incorrect
	INVALID_PARAMETER: "INVALID_PARAMETER",
	// This account is missing the required entitlement(s) to perform this action
	MISSING_ENTITLEMENT: "MISSING_ENTITLEMENT",
	// This account in not authorized to perform this action
	UNAUTHORIZED_ACCOUNT: "UNAUTHORIZED_ACCOUNT",
	// The given channel id represents an unsupported channel type
	UNSUPPORTED_CHANNEL_TYPE: "UNSUPPORTED_CHANNEL_TYPE",
	// The parameter provided does not have a valid form
	VALIDATION_ERROR: "VALIDATION_ERROR",
}

// Account - The `Account` object provides general data about the account, as well as
// being the entry point into more detailed data about a single account.
//
// Account configuration data is queried through this object, as well as
// telemetry data that is specific to a single account.
type Account struct {
	// This field provides access to AiWorkflows data.
	AiWorkflows AiWorkflowsAccountStitchedFields `json:"aiWorkflows,omitempty"`
	//
	ID int `json:"id,omitempty"`
	//
	LicenseKey string `json:"licenseKey,omitempty"`
	//
	Name string `json:"name,omitempty"`
}

// Actor - The `Actor` object contains fields that are scoped to the API user's access level.
type Actor struct {
	// The `account` field is the entry point into data that is scoped to a single account.
	Account Account `json:"account,omitempty"`
}

// AiWorkflowsAccountStitchedFields -
type AiWorkflowsAccountStitchedFields struct {
	// Returns a list of workflows with pagination cursor according to account id and filters
	Workflows AiWorkflowsWorkflows `json:"workflows,omitempty"`
}

// AiWorkflowsCreateResponseError - Create error description
type AiWorkflowsCreateResponseError struct {
	// The error description
	Description string `json:"description"`
	// The error type
	Type AiWorkflowsCreateErrorType `json:"type"`
}

func (x *AiWorkflowsCreateResponseError) ImplementsAiWorkflowsResponseError() {}

// AiWorkflowsCreateWorkflowInput - Workflow input object
type AiWorkflowsCreateWorkflowInput struct {
	// destinationConfigurations
	DestinationConfigurations []AiWorkflowsDestinationConfigurationInput `json:"destinationConfigurations,omitempty"`
	// destinationsEnabled
	DestinationsEnabled bool `json:"destinationsEnabled,omitempty"`
	// enrichments
	Enrichments *AiWorkflowsEnrichmentsInput `json:"enrichments,omitempty"`
	// enrichmentsEnabled
	EnrichmentsEnabled bool `json:"enrichmentsEnabled,omitempty"`
	// issuesFilter
	IssuesFilter AiWorkflowsFilterInput `json:"issuesFilter,omitempty"`
	// mutingRulesHandling
	MutingRulesHandling AiWorkflowsMutingRulesHandling `json:"mutingRulesHandling"`
	// name
	Name string `json:"name"`
	// workflowEnabled
	WorkflowEnabled bool `json:"workflowEnabled,omitempty"`
}

// AiWorkflowsCreateWorkflowResponse - Create workflow mutation response including errors
type AiWorkflowsCreateWorkflowResponse struct {
	// A list of errors that occurred while performing the create workflow action
	Errors []AiWorkflowsCreateResponseError `json:"errors"`
	// Successfully created workflow
	Workflow AiWorkflowsWorkflow `json:"workflow,omitempty"`
}

// AiWorkflowsDeleteResponseError - Delete error description
type AiWorkflowsDeleteResponseError struct {
	// The error description
	Description string `json:"description"`
	// The error type
	Type AiWorkflowsDeleteErrorType `json:"type"`
}

func (x *AiWorkflowsDeleteResponseError) ImplementsAiWorkflowsResponseError() {}

// AiWorkflowsDeleteWorkflowResponse - Delete workflow mutation response including errors
type AiWorkflowsDeleteWorkflowResponse struct {
	// A list of errors that occurred while performing the delete workflow action
	Errors []AiWorkflowsDeleteResponseError `json:"errors"`
	// Id of the successfully deleted workflow
	ID string `json:"id,omitempty"`
}

// AiWorkflowsDestinationConfiguration - Destination Configuration Object
type AiWorkflowsDestinationConfiguration struct {
	// Channel Id of the Destination Configuration
	ChannelId string `json:"channelId"`
	// Name of the Destination Configuration
	Name string `json:"name"`
	// Notification triggers of the Destination Configuration
	NotificationTriggers []AiWorkflowsNotificationTrigger `json:"notificationTriggers"`
	// Type of the Destination Configuration
	Type AiWorkflowsDestinationType `json:"type"`
}

// AiWorkflowsDestinationConfigurationInput - Destination Configuration input object
type AiWorkflowsDestinationConfigurationInput struct {
	// channelId
	ChannelId string `json:"channelId"`
	// notificationTriggers
	NotificationTriggers []AiWorkflowsNotificationTrigger `json:"notificationTriggers"`
}

// AiWorkflowsEnrichment - Makes it possible to augment the notification with additional data from the New Relic platform
type AiWorkflowsEnrichment struct {
	// Account Id of the Enrichment
	AccountID int `json:"accountId"`
	// List of configurations for the enrichment
	Configurations []ai.AiWorkflowsConfiguration `json:"configurations"`
	// The time the Enrichment was created
	CreatedAt nrtime.DateTime `json:"createdAt"`
	// Enrichment Id
	ID string `json:"id"`
	// Name of the Enrichment
	Name string `json:"name"`
	// Type of the Enrichment
	Type AiWorkflowsEnrichmentType `json:"type"`
	// The time the Enrichment was last updated
	UpdatedAt nrtime.DateTime `json:"updatedAt"`
}

// AiWorkflowsEnrichmentsInput - Enrichment input object
type AiWorkflowsEnrichmentsInput struct {
	// nrql
	NRQL []AiWorkflowsNRQLEnrichmentInput `json:"nrql,omitempty"`
}

// AiWorkflowsFilter - Filter Object
type AiWorkflowsFilter struct {
	// Account Id of this Filter
	AccountID int `json:"accountId"`
	// Filter Id
	ID string `json:"id"`
	// Name of the Filter
	Name string `json:"name"`
	// Expressions that determine which issues will be handled
	Predicates []AiWorkflowsPredicate `json:"predicates"`
	// The type of the Filter
	Type AiWorkflowsFilterType `json:"type"`
}

// AiWorkflowsFilterInput - Filter input object
type AiWorkflowsFilterInput struct {
	// name
	Name string `json:"name,omitempty"`
	// predicates
	Predicates []AiWorkflowsPredicateInput `json:"predicates,omitempty"`
	// type
	Type AiWorkflowsFilterType `json:"type"`
}

// AiWorkflowsFilters - Filter on the workflow objects
type AiWorkflowsFilters struct {
	// channelId
	ChannelId string `json:"channelId,omitempty"`
	// destinationType
	DestinationType AiWorkflowsDestinationType `json:"destinationType,omitempty"`
	// enrichmentId
	EnrichmentId string `json:"enrichmentId,omitempty"`
	// filterId
	FilterId string `json:"filterId,omitempty"`
	// id
	ID string `json:"id,omitempty"`
	// name
	Name string `json:"name,omitempty"`
	// nameLike
	NameLike string `json:"nameLike,omitempty"`
	// workflowEnabled
	WorkflowEnabled bool `json:"workflowEnabled,omitempty"`
}

// AiWorkflowsNRQLConfigurationInput - NRQL type configuration input object
type AiWorkflowsNRQLConfigurationInput struct {
	// query
	Query string `json:"query"`
}

// AiWorkflowsNRQLEnrichmentInput - NRQL type enrichment input object
type AiWorkflowsNRQLEnrichmentInput struct {
	// configuration
	Configuration []AiWorkflowsNRQLConfigurationInput `json:"configuration,omitempty"`
	// name
	Name string `json:"name"`
}

// AiWorkflowsNRQLUpdateEnrichmentInput - NRQL type update enrichment input object
type AiWorkflowsNRQLUpdateEnrichmentInput struct {
	// configuration
	Configuration []AiWorkflowsNRQLConfigurationInput `json:"configuration,omitempty"`
	// id
	ID string `json:"id,omitempty"`
	// name
	Name string `json:"name"`
}

// AiWorkflowsPredicate - Predicate Object
type AiWorkflowsPredicate struct {
	// Field name in the issue event
	Attribute string `json:"attribute"`
	// Type of operator used to match the values
	Operator AiWorkflowsOperator `json:"operator"`
	// Values to compare
	Values []string `json:"values"`
}

// AiWorkflowsPredicateInput - PredicateInput input object
type AiWorkflowsPredicateInput struct {
	// attribute
	Attribute string `json:"attribute"`
	// operator
	Operator AiWorkflowsOperator `json:"operator"`
	// values
	Values []string `json:"values"`
}

// AiWorkflowsResponseError - Error description
type AiWorkflowsResponseError struct {
	// The error description
	Description string `json:"description"`
}

func (x *AiWorkflowsResponseError) ImplementsAiWorkflowsResponseError() {}

// AiWorkflowsTestResponseError - Test error description
type AiWorkflowsTestResponseError struct {
	// The error description
	Description string `json:"description"`
	// The error type
	Type AiWorkflowsTestErrorType `json:"type"`
}

func (x *AiWorkflowsTestResponseError) ImplementsAiWorkflowsResponseError() {}

// AiWorkflowsUpdateEnrichmentsInput - Update Enrichment input object
type AiWorkflowsUpdateEnrichmentsInput struct {
	// nrql
	NRQL []AiWorkflowsNRQLUpdateEnrichmentInput `json:"nrql"`
}

// AiWorkflowsUpdateResponseError - Update error description
type AiWorkflowsUpdateResponseError struct {
	// The error description
	Description string `json:"description"`
	// The error type
	Type AiWorkflowsUpdateErrorType `json:"type"`
}

func (x *AiWorkflowsUpdateResponseError) ImplementsAiWorkflowsResponseError() {}

// AiWorkflowsUpdateWorkflowInput - Update Workflow input object
type AiWorkflowsUpdateWorkflowInput struct {
	// destinationConfigurations
	DestinationConfigurations *[]AiWorkflowsDestinationConfigurationInput `json:"destinationConfigurations,omitempty"`
	// destinationsEnabled
	DestinationsEnabled *bool `json:"destinationsEnabled,omitempty"`
	// enrichments
	Enrichments *AiWorkflowsUpdateEnrichmentsInput `json:"enrichments,omitempty"`
	// enrichmentsEnabled
	EnrichmentsEnabled *bool `json:"enrichmentsEnabled,omitempty"`
	// id
	ID string `json:"id"`
	// issuesFilter
	IssuesFilter *AiWorkflowsUpdatedFilterInput `json:"issuesFilter,omitempty"`
	// mutingRulesHandling
	MutingRulesHandling AiWorkflowsMutingRulesHandling `json:"mutingRulesHandling,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// workflowEnabled
	WorkflowEnabled *bool `json:"workflowEnabled,omitempty"`
}

// AiWorkflowsUpdateWorkflowResponse - Update workflow mutation response including errors
type AiWorkflowsUpdateWorkflowResponse struct {
	// A list of errors that occurred while performing the update workflow action
	Errors []AiWorkflowsUpdateResponseError `json:"errors"`
	// Successfully updated workflow
	Workflow AiWorkflowsWorkflow `json:"workflow,omitempty"`
}

// AiWorkflowsUpdatedFilterInput - Update Filter input object
type AiWorkflowsUpdatedFilterInput struct {
	// filterInput
	FilterInput AiWorkflowsFilterInput `json:"filterInput,omitempty"`
	// id
	ID string `json:"id,omitempty"`
}

// AiWorkflowsWorkflow - Workflow object
type AiWorkflowsWorkflow struct {
	// Account Id of this Workflow
	AccountID int `json:"accountId"`
	// The time this workflow was created
	CreatedAt nrtime.DateTime `json:"createdAt"`
	// Specifies where to send the notifications
	DestinationConfigurations []AiWorkflowsDestinationConfiguration `json:"destinationConfigurations"`
	// Are Destinations enabled
	DestinationsEnabled bool `json:"destinationsEnabled"`
	// List of enrichments that are attached to the notifications
	Enrichments []AiWorkflowsEnrichment `json:"enrichments"`
	// Are Enrichments enabled
	EnrichmentsEnabled bool `json:"enrichmentsEnabled"`
	// Workflow Id
	ID string `json:"id"`
	// Specifies which issues the workflow will handle
	IssuesFilter AiWorkflowsFilter `json:"issuesFilter"`
	// Last time a notification was sent regarding this workflow
	LastRun nrtime.DateTime `json:"lastRun,omitempty"`
	// Describes how to handle muted issues
	MutingRulesHandling AiWorkflowsMutingRulesHandling `json:"mutingRulesHandling"`
	// Name of the Workflow
	Name string `json:"name"`
	// The time this workflow was updated
	UpdatedAt nrtime.DateTime `json:"updatedAt"`
	// Is Workflow enabled
	WorkflowEnabled bool `json:"workflowEnabled"`
}

// AiWorkflowsWorkflows - Workflows query response
type AiWorkflowsWorkflows struct {
	// List of all workflows
	Entities []AiWorkflowsWorkflow `json:"entities"`
	// Cursor to get the next batch of results
	NextCursor string `json:"nextCursor,omitempty"`
	// Total count of all workflows
	TotalCount int `json:"totalCount"`
}

type workflowsResponse struct {
	Actor Actor `json:"actor"`
}

// AiWorkflowsResponseError - Error description
type AiWorkflowsResponseErrorInterface interface {
	ImplementsAiWorkflowsResponseError()
}

// UnmarshalAiWorkflowsResponseErrorInterface unmarshals the interface into the correct type
// based on __typename provided by GraphQL
func UnmarshalAiWorkflowsResponseErrorInterface(b []byte) (*AiWorkflowsResponseErrorInterface, error) {
	var err error

	var rawMessageAiWorkflowsResponseError map[string]*json.RawMessage
	err = json.Unmarshal(b, &rawMessageAiWorkflowsResponseError)
	if err != nil {
		return nil, err
	}

	// Nothing to unmarshal
	if len(rawMessageAiWorkflowsResponseError) < 1 {
		return nil, nil
	}

	var typeName string

	if rawTypeName, ok := rawMessageAiWorkflowsResponseError["__typename"]; ok {
		err = json.Unmarshal(*rawTypeName, &typeName)
		if err != nil {
			return nil, err
		}

		switch typeName {
		case "AiWorkflowsCreateResponseError":
			var interfaceType AiWorkflowsCreateResponseError
			err = json.Unmarshal(b, &interfaceType)
			if err != nil {
				return nil, err
			}

			var xxx AiWorkflowsResponseErrorInterface = &interfaceType

			return &xxx, nil
		case "AiWorkflowsDeleteResponseError":
			var interfaceType AiWorkflowsDeleteResponseError
			err = json.Unmarshal(b, &interfaceType)
			if err != nil {
				return nil, err
			}

			var xxx AiWorkflowsResponseErrorInterface = &interfaceType

			return &xxx, nil
		case "AiWorkflowsTestResponseError":
			var interfaceType AiWorkflowsTestResponseError
			err = json.Unmarshal(b, &interfaceType)
			if err != nil {
				return nil, err
			}

			var xxx AiWorkflowsResponseErrorInterface = &interfaceType

			return &xxx, nil
		case "AiWorkflowsUpdateResponseError":
			var interfaceType AiWorkflowsUpdateResponseError
			err = json.Unmarshal(b, &interfaceType)
			if err != nil {
				return nil, err
			}

			var xxx AiWorkflowsResponseErrorInterface = &interfaceType

			return &xxx, nil
		}
	} else {
		keys := []string{}
		for k := range rawMessageAiWorkflowsResponseError {
			keys = append(keys, k)
		}
		return nil, fmt.Errorf("interface AiWorkflowsResponseError did not include a __typename field for inspection: %s", keys)
	}

	return nil, fmt.Errorf("interface AiWorkflowsResponseError was not matched against all PossibleTypes: %s", typeName)
}
