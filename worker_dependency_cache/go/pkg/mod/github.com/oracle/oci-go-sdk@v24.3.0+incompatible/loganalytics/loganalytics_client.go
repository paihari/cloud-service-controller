// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

//LogAnalyticsClient a client for LogAnalytics
type LogAnalyticsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewLogAnalyticsClientWithConfigurationProvider Creates a new default LogAnalytics client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewLogAnalyticsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client LogAnalyticsClient, err error) {
	baseClient, err := common.NewClientWithConfig(configProvider)
	if err != nil {
		return
	}

	return newLogAnalyticsClientFromBaseClient(baseClient, configProvider)
}

// NewLogAnalyticsClientWithOboToken Creates a new default LogAnalytics client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewLogAnalyticsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client LogAnalyticsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return
	}

	return newLogAnalyticsClientFromBaseClient(baseClient, configProvider)
}

func newLogAnalyticsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client LogAnalyticsClient, err error) {
	client = LogAnalyticsClient{BaseClient: baseClient}
	client.BasePath = "20200601"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *LogAnalyticsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("loganalytics", "https://loganalytics.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *LogAnalyticsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *LogAnalyticsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddEntityAssociation Adds association between input source log analytics entity and destination entities.
func (client LogAnalyticsClient) AddEntityAssociation(ctx context.Context, request AddEntityAssociationRequest) (response AddEntityAssociationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.addEntityAssociation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddEntityAssociationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddEntityAssociationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddEntityAssociationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddEntityAssociationResponse")
	}
	return
}

// addEntityAssociation implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) addEntityAssociation(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/actions/addEntityAssociations")
	if err != nil {
		return nil, err
	}

	var response AddEntityAssociationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// BatchGetBasicInfo get basic information about a specified set of labels
func (client LogAnalyticsClient) BatchGetBasicInfo(ctx context.Context, request BatchGetBasicInfoRequest) (response BatchGetBasicInfoResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.batchGetBasicInfo, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = BatchGetBasicInfoResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = BatchGetBasicInfoResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(BatchGetBasicInfoResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into BatchGetBasicInfoResponse")
	}
	return
}

// batchGetBasicInfo implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) batchGetBasicInfo(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/labels/actions/basicInfo")
	if err != nil {
		return nil, err
	}

	var response BatchGetBasicInfoResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CancelQueryWorkRequest Cancel/Remove query job work request.
func (client LogAnalyticsClient) CancelQueryWorkRequest(ctx context.Context, request CancelQueryWorkRequestRequest) (response CancelQueryWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.cancelQueryWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelQueryWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelQueryWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelQueryWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelQueryWorkRequestResponse")
	}
	return
}

// cancelQueryWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) cancelQueryWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/queryWorkRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response CancelQueryWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeLogAnalyticsEntityCompartment Update the compartment of the log analytics entity with the given id.
func (client LogAnalyticsClient) ChangeLogAnalyticsEntityCompartment(ctx context.Context, request ChangeLogAnalyticsEntityCompartmentRequest) (response ChangeLogAnalyticsEntityCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeLogAnalyticsEntityCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogAnalyticsEntityCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogAnalyticsEntityCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogAnalyticsEntityCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogAnalyticsEntityCompartmentResponse")
	}
	return
}

// changeLogAnalyticsEntityCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeLogAnalyticsEntityCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeLogAnalyticsEntityCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeLogAnalyticsLogGroupCompartment Updates the compartment of the Log-Analytics group with the given id.
func (client LogAnalyticsClient) ChangeLogAnalyticsLogGroupCompartment(ctx context.Context, request ChangeLogAnalyticsLogGroupCompartmentRequest) (response ChangeLogAnalyticsLogGroupCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeLogAnalyticsLogGroupCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogAnalyticsLogGroupCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogAnalyticsLogGroupCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogAnalyticsLogGroupCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogAnalyticsLogGroupCompartmentResponse")
	}
	return
}

// changeLogAnalyticsLogGroupCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeLogAnalyticsLogGroupCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeLogAnalyticsLogGroupCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeLogAnalyticsObjectCollectionRuleCompartment Move the rule from it's current compartment to given compartment.
func (client LogAnalyticsClient) ChangeLogAnalyticsObjectCollectionRuleCompartment(ctx context.Context, request ChangeLogAnalyticsObjectCollectionRuleCompartmentRequest) (response ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.changeLogAnalyticsObjectCollectionRuleCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse")
	}
	return
}

// changeLogAnalyticsObjectCollectionRuleCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeLogAnalyticsObjectCollectionRuleCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeLogAnalyticsObjectCollectionRuleCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeScheduledTaskCompartment Move the scheduled task into a different compartment within the same tenancy.
func (client LogAnalyticsClient) ChangeScheduledTaskCompartment(ctx context.Context, request ChangeScheduledTaskCompartmentRequest) (response ChangeScheduledTaskCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeScheduledTaskCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeScheduledTaskCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeScheduledTaskCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeScheduledTaskCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeScheduledTaskCompartmentResponse")
	}
	return
}

// changeScheduledTaskCompartment implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) changeScheduledTaskCompartment(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/changeCompartment")
	if err != nil {
		return nil, err
	}

	var response ChangeScheduledTaskCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Clean Clean accumulated acceleration data stored for the accelerated saved search.
// The ScheduledTask taskType must be ACCELERATION.
func (client LogAnalyticsClient) Clean(ctx context.Context, request CleanRequest) (response CleanResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.clean, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CleanResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CleanResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CleanResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CleanResponse")
	}
	return
}

// clean implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) clean(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/clean")
	if err != nil {
		return nil, err
	}

	var response CleanResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLogAnalyticsEntity Create a new log analytics entity.
func (client LogAnalyticsClient) CreateLogAnalyticsEntity(ctx context.Context, request CreateLogAnalyticsEntityRequest) (response CreateLogAnalyticsEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsEntityResponse")
	}
	return
}

// createLogAnalyticsEntity implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntities")
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLogAnalyticsEntityType Add custom log analytics entity type.
func (client LogAnalyticsClient) CreateLogAnalyticsEntityType(ctx context.Context, request CreateLogAnalyticsEntityTypeRequest) (response CreateLogAnalyticsEntityTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsEntityType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsEntityTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsEntityTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsEntityTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsEntityTypeResponse")
	}
	return
}

// createLogAnalyticsEntityType implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsEntityType(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntityTypes")
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsEntityTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLogAnalyticsLogGroup Creates a new Log-Analytics group.
func (client LogAnalyticsClient) CreateLogAnalyticsLogGroup(ctx context.Context, request CreateLogAnalyticsLogGroupRequest) (response CreateLogAnalyticsLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsLogGroupResponse")
	}
	return
}

// createLogAnalyticsLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsLogGroups")
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateLogAnalyticsObjectCollectionRule Create a configuration to collect logs from object storage bucket.
func (client LogAnalyticsClient) CreateLogAnalyticsObjectCollectionRule(ctx context.Context, request CreateLogAnalyticsObjectCollectionRuleRequest) (response CreateLogAnalyticsObjectCollectionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.createLogAnalyticsObjectCollectionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateLogAnalyticsObjectCollectionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateLogAnalyticsObjectCollectionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateLogAnalyticsObjectCollectionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateLogAnalyticsObjectCollectionRuleResponse")
	}
	return
}

// createLogAnalyticsObjectCollectionRule implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createLogAnalyticsObjectCollectionRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules")
	if err != nil {
		return nil, err
	}

	var response CreateLogAnalyticsObjectCollectionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateScheduledTask Schedule a task as specified and return task info.
func (client LogAnalyticsClient) CreateScheduledTask(ctx context.Context, request CreateScheduledTaskRequest) (response CreateScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateScheduledTaskResponse")
	}
	return
}

// createScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) createScheduledTask(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks")
	if err != nil {
		return nil, err
	}

	var response CreateScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAssociations delete associations
func (client LogAnalyticsClient) DeleteAssociations(ctx context.Context, request DeleteAssociationsRequest) (response DeleteAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAssociationsResponse")
	}
	return
}

// deleteAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/associations/actions/delete")
	if err != nil {
		return nil, err
	}

	var response DeleteAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteField delete field with specified name
func (client LogAnalyticsClient) DeleteField(ctx context.Context, request DeleteFieldRequest) (response DeleteFieldResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteField, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteFieldResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteFieldResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteFieldResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteFieldResponse")
	}
	return
}

// deleteField implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteField(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/fields/{fieldName}")
	if err != nil {
		return nil, err
	}

	var response DeleteFieldResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLabel delete a label
func (client LogAnalyticsClient) DeleteLabel(ctx context.Context, request DeleteLabelRequest) (response DeleteLabelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteLabel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLabelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLabelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLabelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLabelResponse")
	}
	return
}

// deleteLabel implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLabel(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/labels/{labelName}")
	if err != nil {
		return nil, err
	}

	var response DeleteLabelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLogAnalyticsEntity Delete log analytics entity with the given id.
func (client LogAnalyticsClient) DeleteLogAnalyticsEntity(ctx context.Context, request DeleteLogAnalyticsEntityRequest) (response DeleteLogAnalyticsEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsEntityResponse")
	}
	return
}

// deleteLogAnalyticsEntity implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLogAnalyticsEntityType Delete the log analytics entity type with the given name.
func (client LogAnalyticsClient) DeleteLogAnalyticsEntityType(ctx context.Context, request DeleteLogAnalyticsEntityTypeRequest) (response DeleteLogAnalyticsEntityTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsEntityType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsEntityTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsEntityTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsEntityTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsEntityTypeResponse")
	}
	return
}

// deleteLogAnalyticsEntityType implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsEntityType(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsEntityTypes/{entityTypeName}")
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsEntityTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLogAnalyticsLogGroup Deletes the Log-Analytics group with the given id.
func (client LogAnalyticsClient) DeleteLogAnalyticsLogGroup(ctx context.Context, request DeleteLogAnalyticsLogGroupRequest) (response DeleteLogAnalyticsLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsLogGroupResponse")
	}
	return
}

// deleteLogAnalyticsLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteLogAnalyticsObjectCollectionRule Deletes a configured object storage bucket based collection rule to stop the log collection of the configured bucket .
// It will not delete the already collected log data from the configured bucket.
func (client LogAnalyticsClient) DeleteLogAnalyticsObjectCollectionRule(ctx context.Context, request DeleteLogAnalyticsObjectCollectionRuleRequest) (response DeleteLogAnalyticsObjectCollectionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteLogAnalyticsObjectCollectionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteLogAnalyticsObjectCollectionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteLogAnalyticsObjectCollectionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteLogAnalyticsObjectCollectionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteLogAnalyticsObjectCollectionRuleResponse")
	}
	return
}

// deleteLogAnalyticsObjectCollectionRule implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteLogAnalyticsObjectCollectionRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}")
	if err != nil {
		return nil, err
	}

	var response DeleteLogAnalyticsObjectCollectionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteParser delete parser with specified name
func (client LogAnalyticsClient) DeleteParser(ctx context.Context, request DeleteParserRequest) (response DeleteParserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteParser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteParserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteParserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteParserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteParserResponse")
	}
	return
}

// deleteParser implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteParser(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/parsers/{parserName}")
	if err != nil {
		return nil, err
	}

	var response DeleteParserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteScheduledTask Delete the scheduled task.
func (client LogAnalyticsClient) DeleteScheduledTask(ctx context.Context, request DeleteScheduledTaskRequest) (response DeleteScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteScheduledTaskResponse")
	}
	return
}

// deleteScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteScheduledTask(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}")
	if err != nil {
		return nil, err
	}

	var response DeleteScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSource delete source with specified ID
func (client LogAnalyticsClient) DeleteSource(ctx context.Context, request DeleteSourceRequest) (response DeleteSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.deleteSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSourceResponse")
	}
	return
}

// deleteSource implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/sources/{sourceName}")
	if err != nil {
		return nil, err
	}

	var response DeleteSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUpload Deletes an Upload by its reference.
// It deletes all the logs in storage asscoiated with the upload and the corresponding upload metadata.
func (client LogAnalyticsClient) DeleteUpload(ctx context.Context, request DeleteUploadRequest) (response DeleteUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUploadResponse")
	}
	return
}

// deleteUpload implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteUpload(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/uploads/{uploadReference}")
	if err != nil {
		return nil, err
	}

	var response DeleteUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUploadFile Deletes a specific log file inside an upload by providing upload file reference.
// It deletes all the logs in storage asscoiated with the upload file and the corresponding upload metadata.
func (client LogAnalyticsClient) DeleteUploadFile(ctx context.Context, request DeleteUploadFileRequest) (response DeleteUploadFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUploadFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUploadFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUploadFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUploadFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUploadFileResponse")
	}
	return
}

// deleteUploadFile implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteUploadFile(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/uploads/{uploadReference}/files/{fileReference}")
	if err != nil {
		return nil, err
	}

	var response DeleteUploadFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteUploadWarning Suppresses a specific warning inside an upload.
func (client LogAnalyticsClient) DeleteUploadWarning(ctx context.Context, request DeleteUploadWarningRequest) (response DeleteUploadWarningResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteUploadWarning, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteUploadWarningResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteUploadWarningResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteUploadWarningResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteUploadWarningResponse")
	}
	return
}

// deleteUploadWarning implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) deleteUploadWarning(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/namespaces/{namespaceName}/uploads/{uploadReference}/warnings/{warningReference}")
	if err != nil {
		return nil, err
	}

	var response DeleteUploadWarningResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DisableArchiving disable archiving
func (client LogAnalyticsClient) DisableArchiving(ctx context.Context, request DisableArchivingRequest) (response DisableArchivingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.disableArchiving, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DisableArchivingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DisableArchivingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DisableArchivingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DisableArchivingResponse")
	}
	return
}

// disableArchiving implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) disableArchiving(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/disableArchiving")
	if err != nil {
		return nil, err
	}

	var response DisableArchivingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EnableArchiving enable archiving.
func (client LogAnalyticsClient) EnableArchiving(ctx context.Context, request EnableArchivingRequest) (response EnableArchivingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.enableArchiving, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EnableArchivingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EnableArchivingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EnableArchivingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EnableArchivingResponse")
	}
	return
}

// enableArchiving implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) enableArchiving(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/enableArchiving")
	if err != nil {
		return nil, err
	}

	var response EnableArchivingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// EstimatePurgeDataSize estimate the size of data to be purged based on query parameters.
func (client LogAnalyticsClient) EstimatePurgeDataSize(ctx context.Context, request EstimatePurgeDataSizeRequest) (response EstimatePurgeDataSizeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.estimatePurgeDataSize, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = EstimatePurgeDataSizeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = EstimatePurgeDataSizeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(EstimatePurgeDataSizeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into EstimatePurgeDataSizeResponse")
	}
	return
}

// estimatePurgeDataSize implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) estimatePurgeDataSize(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/estimatePurgeDataSize")
	if err != nil {
		return nil, err
	}

	var response EstimatePurgeDataSizeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportCustomContent export
func (client LogAnalyticsClient) ExportCustomContent(ctx context.Context, request ExportCustomContentRequest) (response ExportCustomContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.exportCustomContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportCustomContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportCustomContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportCustomContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportCustomContentResponse")
	}
	return
}

// exportCustomContent implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) exportCustomContent(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/contents/actions/exportCustomContent")
	if err != nil {
		return nil, err
	}

	var response ExportCustomContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportQueryResult Export data based on query. Endpoint returns a stream of data. Endpoint is synchronous. Queries must deliver first result within 60 seconds or calls are subject to timeout.
func (client LogAnalyticsClient) ExportQueryResult(ctx context.Context, request ExportQueryResultRequest) (response ExportQueryResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.exportQueryResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportQueryResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportQueryResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportQueryResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportQueryResultResponse")
	}
	return
}

// exportQueryResult implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) exportQueryResult(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/export")
	if err != nil {
		return nil, err
	}

	var response ExportQueryResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExtractStructuredLogFieldPaths structured log fieldpaths
func (client LogAnalyticsClient) ExtractStructuredLogFieldPaths(ctx context.Context, request ExtractStructuredLogFieldPathsRequest) (response ExtractStructuredLogFieldPathsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.extractStructuredLogFieldPaths, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExtractStructuredLogFieldPathsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExtractStructuredLogFieldPathsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExtractStructuredLogFieldPathsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExtractStructuredLogFieldPathsResponse")
	}
	return
}

// extractStructuredLogFieldPaths implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) extractStructuredLogFieldPaths(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/parsers/actions/extractLogFieldPaths")
	if err != nil {
		return nil, err
	}

	var response ExtractStructuredLogFieldPathsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExtractStructuredLogHeaderPaths structured log header paths
func (client LogAnalyticsClient) ExtractStructuredLogHeaderPaths(ctx context.Context, request ExtractStructuredLogHeaderPathsRequest) (response ExtractStructuredLogHeaderPathsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.extractStructuredLogHeaderPaths, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExtractStructuredLogHeaderPathsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExtractStructuredLogHeaderPathsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExtractStructuredLogHeaderPathsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExtractStructuredLogHeaderPathsResponse")
	}
	return
}

// extractStructuredLogHeaderPaths implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) extractStructuredLogHeaderPaths(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/parsers/actions/extractLogHeaderPaths")
	if err != nil {
		return nil, err
	}

	var response ExtractStructuredLogHeaderPathsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Filter Each filter specifies an operator, a field and one or more values.
func (client LogAnalyticsClient) Filter(ctx context.Context, request FilterRequest) (response FilterResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.filter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = FilterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = FilterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(FilterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into FilterResponse")
	}
	return
}

// filter implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) filter(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/filter")
	if err != nil {
		return nil, err
	}

	var response FilterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAssociationSummary association summary
func (client LogAnalyticsClient) GetAssociationSummary(ctx context.Context, request GetAssociationSummaryRequest) (response GetAssociationSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAssociationSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAssociationSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAssociationSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAssociationSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAssociationSummaryResponse")
	}
	return
}

// getAssociationSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getAssociationSummary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/associationSummary")
	if err != nil {
		return nil, err
	}

	var response GetAssociationSummaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetColumnNames extract column names from SQL query
func (client LogAnalyticsClient) GetColumnNames(ctx context.Context, request GetColumnNamesRequest) (response GetColumnNamesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.getColumnNames, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetColumnNamesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetColumnNamesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetColumnNamesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetColumnNamesResponse")
	}
	return
}

// getColumnNames implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getColumnNames(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/sqlColumnNames")
	if err != nil {
		return nil, err
	}

	var response GetColumnNamesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetConfigWorkRequest association summary by source
func (client LogAnalyticsClient) GetConfigWorkRequest(ctx context.Context, request GetConfigWorkRequestRequest) (response GetConfigWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConfigWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConfigWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConfigWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConfigWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConfigWorkRequestResponse")
	}
	return
}

// getConfigWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getConfigWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/configWorkRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetConfigWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetField get common field with specified name
func (client LogAnalyticsClient) GetField(ctx context.Context, request GetFieldRequest) (response GetFieldResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getField, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFieldResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFieldResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFieldResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFieldResponse")
	}
	return
}

// getField implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getField(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/fields/{fieldName}")
	if err != nil {
		return nil, err
	}

	var response GetFieldResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetFieldsSummary get field summary
func (client LogAnalyticsClient) GetFieldsSummary(ctx context.Context, request GetFieldsSummaryRequest) (response GetFieldsSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getFieldsSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetFieldsSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetFieldsSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetFieldsSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetFieldsSummaryResponse")
	}
	return
}

// getFieldsSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getFieldsSummary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/fieldSummary")
	if err != nil {
		return nil, err
	}

	var response GetFieldsSummaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLabel get label with specified name
func (client LogAnalyticsClient) GetLabel(ctx context.Context, request GetLabelRequest) (response GetLabelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLabel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLabelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLabelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLabelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLabelResponse")
	}
	return
}

// getLabel implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLabel(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labels/{labelName}")
	if err != nil {
		return nil, err
	}

	var response GetLabelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLabelSummary get total count
func (client LogAnalyticsClient) GetLabelSummary(ctx context.Context, request GetLabelSummaryRequest) (response GetLabelSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLabelSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLabelSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLabelSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLabelSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLabelSummaryResponse")
	}
	return
}

// getLabelSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLabelSummary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labelSummary")
	if err != nil {
		return nil, err
	}

	var response GetLabelSummaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogAnalyticsEntitiesSummary Returns log analytics entities count summary.
func (client LogAnalyticsClient) GetLogAnalyticsEntitiesSummary(ctx context.Context, request GetLogAnalyticsEntitiesSummaryRequest) (response GetLogAnalyticsEntitiesSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsEntitiesSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsEntitiesSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsEntitiesSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsEntitiesSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsEntitiesSummaryResponse")
	}
	return
}

// getLogAnalyticsEntitiesSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsEntitiesSummary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities/entitySummary")
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsEntitiesSummaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogAnalyticsEntity Retrieve the log analytics entity with the given id.
func (client LogAnalyticsClient) GetLogAnalyticsEntity(ctx context.Context, request GetLogAnalyticsEntityRequest) (response GetLogAnalyticsEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsEntityResponse")
	}
	return
}

// getLogAnalyticsEntity implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}")
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogAnalyticsEntityType Retrieve the log analytics entity type with the given name.
func (client LogAnalyticsClient) GetLogAnalyticsEntityType(ctx context.Context, request GetLogAnalyticsEntityTypeRequest) (response GetLogAnalyticsEntityTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsEntityType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsEntityTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsEntityTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsEntityTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsEntityTypeResponse")
	}
	return
}

// getLogAnalyticsEntityType implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsEntityType(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntityTypes/{entityTypeName}")
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsEntityTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogAnalyticsLogGroup Retrieves the Log-Analytics group with the given id.
func (client LogAnalyticsClient) GetLogAnalyticsLogGroup(ctx context.Context, request GetLogAnalyticsLogGroupRequest) (response GetLogAnalyticsLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsLogGroupResponse")
	}
	return
}

// getLogAnalyticsLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}")
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogAnalyticsLogGroupsSummary Returns a count of Log-Analytics groups.
func (client LogAnalyticsClient) GetLogAnalyticsLogGroupsSummary(ctx context.Context, request GetLogAnalyticsLogGroupsSummaryRequest) (response GetLogAnalyticsLogGroupsSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsLogGroupsSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsLogGroupsSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsLogGroupsSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsLogGroupsSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsLogGroupsSummaryResponse")
	}
	return
}

// getLogAnalyticsLogGroupsSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsLogGroupsSummary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsLogGroupsSummary")
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsLogGroupsSummaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetLogAnalyticsObjectCollectionRule Gets a configured object storage based collection rule by given id
func (client LogAnalyticsClient) GetLogAnalyticsObjectCollectionRule(ctx context.Context, request GetLogAnalyticsObjectCollectionRuleRequest) (response GetLogAnalyticsObjectCollectionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getLogAnalyticsObjectCollectionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetLogAnalyticsObjectCollectionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetLogAnalyticsObjectCollectionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetLogAnalyticsObjectCollectionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetLogAnalyticsObjectCollectionRuleResponse")
	}
	return
}

// getLogAnalyticsObjectCollectionRule implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getLogAnalyticsObjectCollectionRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}")
	if err != nil {
		return nil, err
	}

	var response GetLogAnalyticsObjectCollectionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetNamespace Get Namespace of a tenancy already onboarded in Log Analytics Application
func (client LogAnalyticsClient) GetNamespace(ctx context.Context, request GetNamespaceRequest) (response GetNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetNamespaceResponse")
	}
	return
}

// getNamespace implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getNamespace(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}")
	if err != nil {
		return nil, err
	}

	var response GetNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetParser get parser with fields by Name
func (client LogAnalyticsClient) GetParser(ctx context.Context, request GetParserRequest) (response GetParserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getParser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetParserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetParserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetParserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetParserResponse")
	}
	return
}

// getParser implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getParser(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parsers/{parserName}")
	if err != nil {
		return nil, err
	}

	var response GetParserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetParserSummary parser summary
func (client LogAnalyticsClient) GetParserSummary(ctx context.Context, request GetParserSummaryRequest) (response GetParserSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getParserSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetParserSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetParserSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetParserSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetParserSummaryResponse")
	}
	return
}

// getParserSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getParserSummary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parsersSummary")
	if err != nil {
		return nil, err
	}

	var response GetParserSummaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetQueryResult Returns the intermediate results for a query that was specified to run asynchronously if the query has not completed,
// otherwise the final query results identified by a queryWorkRequestId returned when submitting the query execute asynchronously.
func (client LogAnalyticsClient) GetQueryResult(ctx context.Context, request GetQueryResultRequest) (response GetQueryResultResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQueryResult, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetQueryResultResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetQueryResultResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQueryResultResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQueryResultResponse")
	}
	return
}

// getQueryResult implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getQueryResult(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/search/actions/query")
	if err != nil {
		return nil, err
	}

	var response GetQueryResultResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetQueryWorkRequest Retrieve work request details by workRequestId. This endpoint can be polled for status tracking of work request. Clients should poll using the interval returned in the retry-after header.
func (client LogAnalyticsClient) GetQueryWorkRequest(ctx context.Context, request GetQueryWorkRequestRequest) (response GetQueryWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getQueryWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetQueryWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetQueryWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetQueryWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetQueryWorkRequestResponse")
	}
	return
}

// getQueryWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getQueryWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/queryWorkRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetQueryWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetScheduledTask Get the scheduled task for the specified task identifier.
func (client LogAnalyticsClient) GetScheduledTask(ctx context.Context, request GetScheduledTaskRequest) (response GetScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetScheduledTaskResponse")
	}
	return
}

// getScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getScheduledTask(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}")
	if err != nil {
		return nil, err
	}

	var response GetScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSource get source with specified name
func (client LogAnalyticsClient) GetSource(ctx context.Context, request GetSourceRequest) (response GetSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSourceResponse")
	}
	return
}

// getSource implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}")
	if err != nil {
		return nil, err
	}

	var response GetSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSourceSummary source summary
func (client LogAnalyticsClient) GetSourceSummary(ctx context.Context, request GetSourceSummaryRequest) (response GetSourceSummaryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSourceSummary, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSourceSummaryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSourceSummaryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSourceSummaryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSourceSummaryResponse")
	}
	return
}

// getSourceSummary implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getSourceSummary(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceSummary")
	if err != nil {
		return nil, err
	}

	var response GetSourceSummaryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStorage Storage configuration and status.
func (client LogAnalyticsClient) GetStorage(ctx context.Context, request GetStorageRequest) (response GetStorageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStorage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStorageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStorageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStorageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStorageResponse")
	}
	return
}

// getStorage implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getStorage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storage")
	if err != nil {
		return nil, err
	}

	var response GetStorageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStorageUsage Storage usage info includes active, archived or recalled data.  The unit of return value is in bytes.
func (client LogAnalyticsClient) GetStorageUsage(ctx context.Context, request GetStorageUsageRequest) (response GetStorageUsageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStorageUsage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStorageUsageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStorageUsageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStorageUsageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStorageUsageResponse")
	}
	return
}

// getStorageUsage implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getStorageUsage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storage/usage")
	if err != nil {
		return nil, err
	}

	var response GetStorageUsageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetStorageWorkRequest Retrieve work request details by key. This endpoint can be polled for status tracking of work request.
// Clients should poll using the interval returned in retry-after header.
func (client LogAnalyticsClient) GetStorageWorkRequest(ctx context.Context, request GetStorageWorkRequestRequest) (response GetStorageWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getStorageWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetStorageWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetStorageWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetStorageWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetStorageWorkRequestResponse")
	}
	return
}

// getStorageWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getStorageWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storageWorkRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetStorageWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetUpload Gets an On-Demand Upload info by reference
func (client LogAnalyticsClient) GetUpload(ctx context.Context, request GetUploadRequest) (response GetUploadResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getUpload, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetUploadResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetUploadResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetUploadResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetUploadResponse")
	}
	return
}

// getUpload implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getUpload(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/uploads/{uploadReference}")
	if err != nil {
		return nil, err
	}

	var response GetUploadResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetWorkRequest Gets the status of the work request with the given ID.
func (client LogAnalyticsClient) GetWorkRequest(ctx context.Context, request GetWorkRequestRequest) (response GetWorkRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getWorkRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetWorkRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetWorkRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetWorkRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetWorkRequestResponse")
	}
	return
}

// getWorkRequest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) getWorkRequest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/workRequests/{workRequestId}")
	if err != nil {
		return nil, err
	}

	var response GetWorkRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportCustomContent register custom content
func (client LogAnalyticsClient) ImportCustomContent(ctx context.Context, request ImportCustomContentRequest) (response ImportCustomContentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importCustomContent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportCustomContentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportCustomContentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportCustomContentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportCustomContentResponse")
	}
	return
}

// importCustomContent implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) importCustomContent(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/contents/actions/importCustomContent")
	if err != nil {
		return nil, err
	}

	var response ImportCustomContentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAssociatedEntities list of entities that have been associated to at least one source
func (client LogAnalyticsClient) ListAssociatedEntities(ctx context.Context, request ListAssociatedEntitiesRequest) (response ListAssociatedEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAssociatedEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAssociatedEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAssociatedEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAssociatedEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAssociatedEntitiesResponse")
	}
	return
}

// listAssociatedEntities implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listAssociatedEntities(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/associatedEntities")
	if err != nil {
		return nil, err
	}

	var response ListAssociatedEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListConfigWorkRequests association summary by source
func (client LogAnalyticsClient) ListConfigWorkRequests(ctx context.Context, request ListConfigWorkRequestsRequest) (response ListConfigWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConfigWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConfigWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConfigWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConfigWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConfigWorkRequestsResponse")
	}
	return
}

// listConfigWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listConfigWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/configWorkRequests")
	if err != nil {
		return nil, err
	}

	var response ListConfigWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEntityAssociations Return a list of log analytics entities associated with input source log analytics entity.
func (client LogAnalyticsClient) ListEntityAssociations(ctx context.Context, request ListEntityAssociationsRequest) (response ListEntityAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntityAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntityAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntityAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntityAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntityAssociationsResponse")
	}
	return
}

// listEntityAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listEntityAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/entityAssociations")
	if err != nil {
		return nil, err
	}

	var response ListEntityAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEntitySourceAssociations entity associations summary
func (client LogAnalyticsClient) ListEntitySourceAssociations(ctx context.Context, request ListEntitySourceAssociationsRequest) (response ListEntitySourceAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEntitySourceAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEntitySourceAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEntitySourceAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEntitySourceAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEntitySourceAssociationsResponse")
	}
	return
}

// listEntitySourceAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listEntitySourceAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/entityAssociations")
	if err != nil {
		return nil, err
	}

	var response ListEntitySourceAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListFields get all common field with specified display name and description
func (client LogAnalyticsClient) ListFields(ctx context.Context, request ListFieldsRequest) (response ListFieldsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listFields, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListFieldsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListFieldsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListFieldsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListFieldsResponse")
	}
	return
}

// listFields implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listFields(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/fields")
	if err != nil {
		return nil, err
	}

	var response ListFieldsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLabelPriorities get list of priorities
func (client LogAnalyticsClient) ListLabelPriorities(ctx context.Context, request ListLabelPrioritiesRequest) (response ListLabelPrioritiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLabelPriorities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLabelPrioritiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLabelPrioritiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLabelPrioritiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLabelPrioritiesResponse")
	}
	return
}

// listLabelPriorities implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLabelPriorities(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labelPriorities")
	if err != nil {
		return nil, err
	}

	var response ListLabelPrioritiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLabelSourceDetails get details of sources using the label
func (client LogAnalyticsClient) ListLabelSourceDetails(ctx context.Context, request ListLabelSourceDetailsRequest) (response ListLabelSourceDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLabelSourceDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLabelSourceDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLabelSourceDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLabelSourceDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLabelSourceDetailsResponse")
	}
	return
}

// listLabelSourceDetails implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLabelSourceDetails(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labelSourceDetails")
	if err != nil {
		return nil, err
	}

	var response ListLabelSourceDetailsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLabels get labels passing specified filter
func (client LogAnalyticsClient) ListLabels(ctx context.Context, request ListLabelsRequest) (response ListLabelsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLabels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLabelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLabelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLabelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLabelsResponse")
	}
	return
}

// listLabels implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLabels(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/labels")
	if err != nil {
		return nil, err
	}

	var response ListLabelsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLogAnalyticsEntities Return a list of log analytics entities.
func (client LogAnalyticsClient) ListLogAnalyticsEntities(ctx context.Context, request ListLogAnalyticsEntitiesRequest) (response ListLogAnalyticsEntitiesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsEntities, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsEntitiesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsEntitiesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsEntitiesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsEntitiesResponse")
	}
	return
}

// listLogAnalyticsEntities implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsEntities(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntities")
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsEntitiesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLogAnalyticsEntityTypes Return a list of log analytics entity types.
func (client LogAnalyticsClient) ListLogAnalyticsEntityTypes(ctx context.Context, request ListLogAnalyticsEntityTypesRequest) (response ListLogAnalyticsEntityTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsEntityTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsEntityTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsEntityTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsEntityTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsEntityTypesResponse")
	}
	return
}

// listLogAnalyticsEntityTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsEntityTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsEntityTypes")
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsEntityTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLogAnalyticsLogGroups Returns a list of Log-Analytics groups.
func (client LogAnalyticsClient) ListLogAnalyticsLogGroups(ctx context.Context, request ListLogAnalyticsLogGroupsRequest) (response ListLogAnalyticsLogGroupsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsLogGroups, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsLogGroupsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsLogGroupsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsLogGroupsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsLogGroupsResponse")
	}
	return
}

// listLogAnalyticsLogGroups implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsLogGroups(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsLogGroups")
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsLogGroupsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListLogAnalyticsObjectCollectionRules Gets list of configuration details of Object Storage based collection rules.
func (client LogAnalyticsClient) ListLogAnalyticsObjectCollectionRules(ctx context.Context, request ListLogAnalyticsObjectCollectionRulesRequest) (response ListLogAnalyticsObjectCollectionRulesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listLogAnalyticsObjectCollectionRules, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListLogAnalyticsObjectCollectionRulesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListLogAnalyticsObjectCollectionRulesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListLogAnalyticsObjectCollectionRulesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListLogAnalyticsObjectCollectionRulesResponse")
	}
	return
}

// listLogAnalyticsObjectCollectionRules implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listLogAnalyticsObjectCollectionRules(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules")
	if err != nil {
		return nil, err
	}

	var response ListLogAnalyticsObjectCollectionRulesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListMetaSourceTypes get all meta source types
func (client LogAnalyticsClient) ListMetaSourceTypes(ctx context.Context, request ListMetaSourceTypesRequest) (response ListMetaSourceTypesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listMetaSourceTypes, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListMetaSourceTypesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListMetaSourceTypesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListMetaSourceTypesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListMetaSourceTypesResponse")
	}
	return
}

// listMetaSourceTypes implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listMetaSourceTypes(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceMetaTypes")
	if err != nil {
		return nil, err
	}

	var response ListMetaSourceTypesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListNamespaces List Namespaces.
func (client LogAnalyticsClient) ListNamespaces(ctx context.Context, request ListNamespacesRequest) (response ListNamespacesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listNamespaces, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListNamespacesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListNamespacesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListNamespacesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListNamespacesResponse")
	}
	return
}

// listNamespaces implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listNamespaces(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces")
	if err != nil {
		return nil, err
	}

	var response ListNamespacesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListParserFunctions get pre-process plugin instance
func (client LogAnalyticsClient) ListParserFunctions(ctx context.Context, request ListParserFunctionsRequest) (response ListParserFunctionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listParserFunctions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListParserFunctionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListParserFunctionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListParserFunctionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListParserFunctionsResponse")
	}
	return
}

// listParserFunctions implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listParserFunctions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parserFunctions")
	if err != nil {
		return nil, err
	}

	var response ListParserFunctionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListParserMetaPlugins get pre-process Meta plugins
func (client LogAnalyticsClient) ListParserMetaPlugins(ctx context.Context, request ListParserMetaPluginsRequest) (response ListParserMetaPluginsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listParserMetaPlugins, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListParserMetaPluginsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListParserMetaPluginsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListParserMetaPluginsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListParserMetaPluginsResponse")
	}
	return
}

// listParserMetaPlugins implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listParserMetaPlugins(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parserMetaPlugins")
	if err != nil {
		return nil, err
	}

	var response ListParserMetaPluginsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListParsers List parsers passing specified filter
func (client LogAnalyticsClient) ListParsers(ctx context.Context, request ListParsersRequest) (response ListParsersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listParsers, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListParsersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListParsersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListParsersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListParsersResponse")
	}
	return
}

// listParsers implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listParsers(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/parsers")
	if err != nil {
		return nil, err
	}

	var response ListParsersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListQueryWorkRequests List active asynchronous queries.
func (client LogAnalyticsClient) ListQueryWorkRequests(ctx context.Context, request ListQueryWorkRequestsRequest) (response ListQueryWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listQueryWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListQueryWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListQueryWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListQueryWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListQueryWorkRequestsResponse")
	}
	return
}

// listQueryWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listQueryWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/queryWorkRequests")
	if err != nil {
		return nil, err
	}

	var response ListQueryWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListScheduledTasks Lists scheduled tasks.
func (client LogAnalyticsClient) ListScheduledTasks(ctx context.Context, request ListScheduledTasksRequest) (response ListScheduledTasksResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listScheduledTasks, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListScheduledTasksResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListScheduledTasksResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListScheduledTasksResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListScheduledTasksResponse")
	}
	return
}

// listScheduledTasks implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listScheduledTasks(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/scheduledTasks")
	if err != nil {
		return nil, err
	}

	var response ListScheduledTasksResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSourceAssociations association summary by source
func (client LogAnalyticsClient) ListSourceAssociations(ctx context.Context, request ListSourceAssociationsRequest) (response ListSourceAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceAssociationsResponse")
	}
	return
}

// listSourceAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceAssociations")
	if err != nil {
		return nil, err
	}

	var response ListSourceAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSourceExtendedFieldDefinitions get source extended fields for source with specified Id
func (client LogAnalyticsClient) ListSourceExtendedFieldDefinitions(ctx context.Context, request ListSourceExtendedFieldDefinitionsRequest) (response ListSourceExtendedFieldDefinitionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceExtendedFieldDefinitions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceExtendedFieldDefinitionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceExtendedFieldDefinitionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceExtendedFieldDefinitionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceExtendedFieldDefinitionsResponse")
	}
	return
}

// listSourceExtendedFieldDefinitions implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceExtendedFieldDefinitions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}/extendedFieldDefinitions")
	if err != nil {
		return nil, err
	}

	var response ListSourceExtendedFieldDefinitionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSourceLabelOperators list source label operators
func (client LogAnalyticsClient) ListSourceLabelOperators(ctx context.Context, request ListSourceLabelOperatorsRequest) (response ListSourceLabelOperatorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceLabelOperators, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceLabelOperatorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceLabelOperatorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceLabelOperatorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceLabelOperatorsResponse")
	}
	return
}

// listSourceLabelOperators implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceLabelOperators(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceLabelOperators")
	if err != nil {
		return nil, err
	}

	var response ListSourceLabelOperatorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSourceMetaFunctions get source meta functions
func (client LogAnalyticsClient) ListSourceMetaFunctions(ctx context.Context, request ListSourceMetaFunctionsRequest) (response ListSourceMetaFunctionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourceMetaFunctions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourceMetaFunctionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourceMetaFunctionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourceMetaFunctionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourceMetaFunctionsResponse")
	}
	return
}

// listSourceMetaFunctions implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourceMetaFunctions(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sourceMetaFunctions")
	if err != nil {
		return nil, err
	}

	var response ListSourceMetaFunctionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSourcePatterns get source patterns for source with specified Id
func (client LogAnalyticsClient) ListSourcePatterns(ctx context.Context, request ListSourcePatternsRequest) (response ListSourcePatternsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSourcePatterns, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourcePatternsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourcePatternsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourcePatternsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourcePatternsResponse")
	}
	return
}

// listSourcePatterns implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSourcePatterns(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources/{sourceName}/patterns")
	if err != nil {
		return nil, err
	}

	var response ListSourcePatternsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSources source list
func (client LogAnalyticsClient) ListSources(ctx context.Context, request ListSourcesRequest) (response ListSourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSourcesResponse")
	}
	return
}

// listSources implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSources(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/sources")
	if err != nil {
		return nil, err
	}

	var response ListSourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListStorageWorkRequestErrors Retrieve work request errors if any
func (client LogAnalyticsClient) ListStorageWorkRequestErrors(ctx context.Context, request ListStorageWorkRequestErrorsRequest) (response ListStorageWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStorageWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStorageWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStorageWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStorageWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStorageWorkRequestErrorsResponse")
	}
	return
}

// listStorageWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listStorageWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storageWorkRequests/{workRequestId}/errors")
	if err != nil {
		return nil, err
	}

	var response ListStorageWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListStorageWorkRequests List non-expired storage manager work requests.
func (client LogAnalyticsClient) ListStorageWorkRequests(ctx context.Context, request ListStorageWorkRequestsRequest) (response ListStorageWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listStorageWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListStorageWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListStorageWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListStorageWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListStorageWorkRequestsResponse")
	}
	return
}

// listStorageWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listStorageWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/storageWorkRequests")
	if err != nil {
		return nil, err
	}

	var response ListStorageWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSupportedCharEncodings Gets the list of character encodings supported for log files.
func (client LogAnalyticsClient) ListSupportedCharEncodings(ctx context.Context, request ListSupportedCharEncodingsRequest) (response ListSupportedCharEncodingsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSupportedCharEncodings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSupportedCharEncodingsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSupportedCharEncodingsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSupportedCharEncodingsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSupportedCharEncodingsResponse")
	}
	return
}

// listSupportedCharEncodings implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSupportedCharEncodings(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/supportedCharEncodings")
	if err != nil {
		return nil, err
	}

	var response ListSupportedCharEncodingsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSupportedTimezones Gets timezones that are supported when performing uploads.
func (client LogAnalyticsClient) ListSupportedTimezones(ctx context.Context, request ListSupportedTimezonesRequest) (response ListSupportedTimezonesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSupportedTimezones, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSupportedTimezonesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSupportedTimezonesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSupportedTimezonesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSupportedTimezonesResponse")
	}
	return
}

// listSupportedTimezones implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listSupportedTimezones(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/supportedTimezones")
	if err != nil {
		return nil, err
	}

	var response ListSupportedTimezonesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUploadFiles Gets list of files in an upload.
func (client LogAnalyticsClient) ListUploadFiles(ctx context.Context, request ListUploadFilesRequest) (response ListUploadFilesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUploadFiles, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUploadFilesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUploadFilesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUploadFilesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUploadFilesResponse")
	}
	return
}

// listUploadFiles implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listUploadFiles(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/uploads/{uploadReference}/files")
	if err != nil {
		return nil, err
	}

	var response ListUploadFilesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUploadWarnings Gets list of warnings in an upload explaining the failures due to incorrect configuration.
func (client LogAnalyticsClient) ListUploadWarnings(ctx context.Context, request ListUploadWarningsRequest) (response ListUploadWarningsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUploadWarnings, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUploadWarningsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUploadWarningsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUploadWarningsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUploadWarningsResponse")
	}
	return
}

// listUploadWarnings implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listUploadWarnings(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/uploads/{uploadReference}/warnings")
	if err != nil {
		return nil, err
	}

	var response ListUploadWarningsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListUploads Gets a list of all On-demand uploads.
// To use this and other API operations, you must be authorized in an IAM policy.
func (client LogAnalyticsClient) ListUploads(ctx context.Context, request ListUploadsRequest) (response ListUploadsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listUploads, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListUploadsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListUploadsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListUploadsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListUploadsResponse")
	}
	return
}

// listUploads implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listUploads(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/uploads")
	if err != nil {
		return nil, err
	}

	var response ListUploadsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestErrors Return a (paginated) list of errors for a given work request.
func (client LogAnalyticsClient) ListWorkRequestErrors(ctx context.Context, request ListWorkRequestErrorsRequest) (response ListWorkRequestErrorsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestErrors, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestErrorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestErrorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestErrorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestErrorsResponse")
	}
	return
}

// listWorkRequestErrors implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listWorkRequestErrors(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/workRequests/{workRequestId}/errors")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestErrorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequestLogs Return a (paginated) list of logs for a given work request.
func (client LogAnalyticsClient) ListWorkRequestLogs(ctx context.Context, request ListWorkRequestLogsRequest) (response ListWorkRequestLogsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequestLogs, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestLogsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestLogsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestLogsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestLogsResponse")
	}
	return
}

// listWorkRequestLogs implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listWorkRequestLogs(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/workRequests/{workRequestId}/logs")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestLogsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListWorkRequests Lists the work requests in a compartment.
func (client LogAnalyticsClient) ListWorkRequests(ctx context.Context, request ListWorkRequestsRequest) (response ListWorkRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listWorkRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListWorkRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListWorkRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListWorkRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListWorkRequestsResponse")
	}
	return
}

// listWorkRequests implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) listWorkRequests(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodGet, "/namespaces/{namespaceName}/workRequests")
	if err != nil {
		return nil, err
	}

	var response ListWorkRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// OffboardNamespace Off-boards a tenant from Logging Analytics
func (client LogAnalyticsClient) OffboardNamespace(ctx context.Context, request OffboardNamespaceRequest) (response OffboardNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.offboardNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = OffboardNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = OffboardNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(OffboardNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into OffboardNamespaceResponse")
	}
	return
}

// offboardNamespace implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) offboardNamespace(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/actions/offboard")
	if err != nil {
		return nil, err
	}

	var response OffboardNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// OnboardNamespace On-boards a tenant to Logging Analytics.
func (client LogAnalyticsClient) OnboardNamespace(ctx context.Context, request OnboardNamespaceRequest) (response OnboardNamespaceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.onboardNamespace, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = OnboardNamespaceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = OnboardNamespaceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(OnboardNamespaceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into OnboardNamespaceResponse")
	}
	return
}

// onboardNamespace implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) onboardNamespace(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/actions/onboard")
	if err != nil {
		return nil, err
	}

	var response OnboardNamespaceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ParseQuery Describe query
func (client LogAnalyticsClient) ParseQuery(ctx context.Context, request ParseQueryRequest) (response ParseQueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.parseQuery, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ParseQueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ParseQueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ParseQueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ParseQueryResponse")
	}
	return
}

// parseQuery implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) parseQuery(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/parse")
	if err != nil {
		return nil, err
	}

	var response ParseQueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PurgeStorageData submit work requests to purge old data based on the type.
func (client LogAnalyticsClient) PurgeStorageData(ctx context.Context, request PurgeStorageDataRequest) (response PurgeStorageDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.purgeStorageData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PurgeStorageDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PurgeStorageDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PurgeStorageDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PurgeStorageDataResponse")
	}
	return
}

// purgeStorageData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) purgeStorageData(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/purgeData")
	if err != nil {
		return nil, err
	}

	var response PurgeStorageDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PutQueryWorkRequestBackground Put the work request specified by {workRequestId} into the background.
func (client LogAnalyticsClient) PutQueryWorkRequestBackground(ctx context.Context, request PutQueryWorkRequestBackgroundRequest) (response PutQueryWorkRequestBackgroundResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.putQueryWorkRequestBackground, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PutQueryWorkRequestBackgroundResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PutQueryWorkRequestBackgroundResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PutQueryWorkRequestBackgroundResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PutQueryWorkRequestBackgroundResponse")
	}
	return
}

// putQueryWorkRequestBackground implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) putQueryWorkRequestBackground(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/queryWorkRequests/{workRequestId}/actions/background")
	if err != nil {
		return nil, err
	}

	var response PutQueryWorkRequestBackgroundResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Query Performs a log analytics search, if shouldRunAsync is false returns the query results once they become available subject to 60 second timeout. If a query is subject to exceed that time then it should be run asynchronously. Asynchronous query submissions return the queryWorkRequestId to use for execution tracking, query submission lifecycle actions and to poll for query results.
func (client LogAnalyticsClient) Query(ctx context.Context, request QueryRequest) (response QueryResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.query, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = QueryResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = QueryResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(QueryResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into QueryResponse")
	}
	return
}

// query implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) query(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/query")
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RecallArchivedData submit work requests to recall archived data.
func (client LogAnalyticsClient) RecallArchivedData(ctx context.Context, request RecallArchivedDataRequest) (response RecallArchivedDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.recallArchivedData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RecallArchivedDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RecallArchivedDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RecallArchivedDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RecallArchivedDataResponse")
	}
	return
}

// recallArchivedData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) recallArchivedData(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/recallArchivedData")
	if err != nil {
		return nil, err
	}

	var response RecallArchivedDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RegisterLookup register lookup
func (client LogAnalyticsClient) RegisterLookup(ctx context.Context, request RegisterLookupRequest) (response RegisterLookupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.registerLookup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RegisterLookupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RegisterLookupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RegisterLookupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RegisterLookupResponse")
	}
	return
}

// registerLookup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) registerLookup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/lookups/actions/register")
	if err != nil {
		return nil, err
	}

	var response RegisterLookupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ReleaseRecalledData submit work requests to release recalled data.
func (client LogAnalyticsClient) ReleaseRecalledData(ctx context.Context, request ReleaseRecalledDataRequest) (response ReleaseRecalledDataResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.releaseRecalledData, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ReleaseRecalledDataResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ReleaseRecalledDataResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ReleaseRecalledDataResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ReleaseRecalledDataResponse")
	}
	return
}

// releaseRecalledData implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) releaseRecalledData(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/storage/actions/releaseRecalledData")
	if err != nil {
		return nil, err
	}

	var response ReleaseRecalledDataResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveEntityAssociations Delete association between input source log analytics entity and destination entities.
func (client LogAnalyticsClient) RemoveEntityAssociations(ctx context.Context, request RemoveEntityAssociationsRequest) (response RemoveEntityAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.removeEntityAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveEntityAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveEntityAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveEntityAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveEntityAssociationsResponse")
	}
	return
}

// removeEntityAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) removeEntityAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}/actions/removeEntityAssociations")
	if err != nil {
		return nil, err
	}

	var response RemoveEntityAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Run Execute the saved search acceleration task in the foreground.
// The ScheduledTask taskType must be ACCELERATION.
// Optionally specify time range (timeStart and timeEnd). The default is all time.
func (client LogAnalyticsClient) Run(ctx context.Context, request RunRequest) (response RunResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.run, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RunResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RunResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RunResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RunResponse")
	}
	return
}

// run implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) run(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}/actions/run")
	if err != nil {
		return nil, err
	}

	var response RunResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// Suggest Returns a context specific list of either commands, fields, or values to add to the end of the query string.
func (client LogAnalyticsClient) Suggest(ctx context.Context, request SuggestRequest) (response SuggestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.suggest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = SuggestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = SuggestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(SuggestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into SuggestResponse")
	}
	return
}

// suggest implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) suggest(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/search/actions/suggest")
	if err != nil {
		return nil, err
	}

	var response SuggestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// TestParser test parser
func (client LogAnalyticsClient) TestParser(ctx context.Context, request TestParserRequest) (response TestParserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.testParser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = TestParserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = TestParserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(TestParserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into TestParserResponse")
	}
	return
}

// testParser implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) testParser(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/parsers/actions/test")
	if err != nil {
		return nil, err
	}

	var response TestParserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLogAnalyticsEntity Update the log analytics entity with the given id.
func (client LogAnalyticsClient) UpdateLogAnalyticsEntity(ctx context.Context, request UpdateLogAnalyticsEntityRequest) (response UpdateLogAnalyticsEntityResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsEntity, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsEntityResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsEntityResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsEntityResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsEntityResponse")
	}
	return
}

// updateLogAnalyticsEntity implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsEntity(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}")
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsEntityResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLogAnalyticsEntityType Update custom log analytics entity type. Out of box entity types cannot be udpated.
func (client LogAnalyticsClient) UpdateLogAnalyticsEntityType(ctx context.Context, request UpdateLogAnalyticsEntityTypeRequest) (response UpdateLogAnalyticsEntityTypeResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsEntityType, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsEntityTypeResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsEntityTypeResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsEntityTypeResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsEntityTypeResponse")
	}
	return
}

// updateLogAnalyticsEntityType implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsEntityType(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsEntityTypes/{entityTypeName}")
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsEntityTypeResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLogAnalyticsLogGroup Updates the Log-Analytics group with the given id.
func (client LogAnalyticsClient) UpdateLogAnalyticsLogGroup(ctx context.Context, request UpdateLogAnalyticsLogGroupRequest) (response UpdateLogAnalyticsLogGroupResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsLogGroup, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsLogGroupResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsLogGroupResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsLogGroupResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsLogGroupResponse")
	}
	return
}

// updateLogAnalyticsLogGroup implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsLogGroup(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsLogGroups/{logAnalyticsLogGroupId}")
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsLogGroupResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateLogAnalyticsObjectCollectionRule Update the rule with the given id.
func (client LogAnalyticsClient) UpdateLogAnalyticsObjectCollectionRule(ctx context.Context, request UpdateLogAnalyticsObjectCollectionRuleRequest) (response UpdateLogAnalyticsObjectCollectionRuleResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateLogAnalyticsObjectCollectionRule, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateLogAnalyticsObjectCollectionRuleResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateLogAnalyticsObjectCollectionRuleResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateLogAnalyticsObjectCollectionRuleResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateLogAnalyticsObjectCollectionRuleResponse")
	}
	return
}

// updateLogAnalyticsObjectCollectionRule implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateLogAnalyticsObjectCollectionRule(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/logAnalyticsObjectCollectionRules/{logAnalyticsObjectCollectionRuleId}")
	if err != nil {
		return nil, err
	}

	var response UpdateLogAnalyticsObjectCollectionRuleResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateScheduledTask Update the scheduled task. Schedules may be updated only for taskType SAVED_SEARCH and PURGE.
func (client LogAnalyticsClient) UpdateScheduledTask(ctx context.Context, request UpdateScheduledTaskRequest) (response UpdateScheduledTaskResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateScheduledTask, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateScheduledTaskResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateScheduledTaskResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateScheduledTaskResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateScheduledTaskResponse")
	}
	return
}

// updateScheduledTask implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateScheduledTask(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/scheduledTasks/{scheduledTaskId}")
	if err != nil {
		return nil, err
	}

	var response UpdateScheduledTaskResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateStorage update the archiving configuration
func (client LogAnalyticsClient) UpdateStorage(ctx context.Context, request UpdateStorageRequest) (response UpdateStorageResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateStorage, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateStorageResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateStorageResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateStorageResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateStorageResponse")
	}
	return
}

// updateStorage implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) updateStorage(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPut, "/namespaces/{namespaceName}/storage")
	if err != nil {
		return nil, err
	}

	var response UpdateStorageResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UploadLogFile Accepts log data for processing by Log Analytics.
func (client LogAnalyticsClient) UploadLogFile(ctx context.Context, request UploadLogFileRequest) (response UploadLogFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.uploadLogFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UploadLogFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UploadLogFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UploadLogFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UploadLogFileResponse")
	}
	return
}

// uploadLogFile implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) uploadLogFile(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/actions/uploadLogFile")
	if err != nil {
		return nil, err
	}

	var response UploadLogFileResponse
	var httpResponse *http.Response
	var customSigner common.HTTPRequestSigner
	excludeBodySigningPredicate := func(r *http.Request) bool { return false }
	customSigner, err = common.NewSignerFromOCIRequestSigner(client.Signer, excludeBodySigningPredicate)

	//if there was an error overriding the signer, then use the signer from the client itself
	if err != nil {
		customSigner = client.Signer
	}

	//Execute the request with a custom signer
	httpResponse, err = client.CallWithDetails(ctx, &httpRequest, common.ClientCallDetails{Signer: customSigner})
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpsertAssociations create or update associations for a source
func (client LogAnalyticsClient) UpsertAssociations(ctx context.Context, request UpsertAssociationsRequest) (response UpsertAssociationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertAssociations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertAssociationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertAssociationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertAssociationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertAssociationsResponse")
	}
	return
}

// upsertAssociations implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertAssociations(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/associations/actions/upsert")
	if err != nil {
		return nil, err
	}

	var response UpsertAssociationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpsertField Defines or update a field.
func (client LogAnalyticsClient) UpsertField(ctx context.Context, request UpsertFieldRequest) (response UpsertFieldResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertField, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertFieldResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertFieldResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertFieldResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertFieldResponse")
	}
	return
}

// upsertField implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertField(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/fields/actions/upsert")
	if err != nil {
		return nil, err
	}

	var response UpsertFieldResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpsertLabel Define or update a label.
func (client LogAnalyticsClient) UpsertLabel(ctx context.Context, request UpsertLabelRequest) (response UpsertLabelResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertLabel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertLabelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertLabelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertLabelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertLabelResponse")
	}
	return
}

// upsertLabel implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertLabel(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/labels/actions/upsert")
	if err != nil {
		return nil, err
	}

	var response UpsertLabelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpsertParser Define or update parser
func (client LogAnalyticsClient) UpsertParser(ctx context.Context, request UpsertParserRequest) (response UpsertParserResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertParser, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertParserResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertParserResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertParserResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertParserResponse")
	}
	return
}

// upsertParser implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertParser(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/parsers/actions/upsert")
	if err != nil {
		return nil, err
	}

	var response UpsertParserResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpsertSource Define or update a source
func (client LogAnalyticsClient) UpsertSource(ctx context.Context, request UpsertSourceRequest) (response UpsertSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.upsertSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpsertSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpsertSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpsertSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpsertSourceResponse")
	}
	return
}

// upsertSource implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) upsertSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/actions/upsert")
	if err != nil {
		return nil, err
	}

	var response UpsertSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateAssociationParameters association parameter validation
func (client LogAnalyticsClient) ValidateAssociationParameters(ctx context.Context, request ValidateAssociationParametersRequest) (response ValidateAssociationParametersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.validateAssociationParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateAssociationParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateAssociationParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateAssociationParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateAssociationParametersResponse")
	}
	return
}

// validateAssociationParameters implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateAssociationParameters(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/associations/actions/validateParameters")
	if err != nil {
		return nil, err
	}

	var response ValidateAssociationParametersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateFile Validates a log file to check whether it is eligible to upload or not.
func (client LogAnalyticsClient) ValidateFile(ctx context.Context, request ValidateFileRequest) (response ValidateFileResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateFile, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateFileResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateFileResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateFileResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateFileResponse")
	}
	return
}

// validateFile implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateFile(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/uploads/actions/validateFile")
	if err != nil {
		return nil, err
	}

	var response ValidateFileResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateSource Pre-define or update a source
func (client LogAnalyticsClient) ValidateSource(ctx context.Context, request ValidateSourceRequest) (response ValidateSourceResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.validateSource, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateSourceResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateSourceResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateSourceResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateSourceResponse")
	}
	return
}

// validateSource implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateSource(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/actions/validate")
	if err != nil {
		return nil, err
	}

	var response ValidateSourceResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateSourceExtendedFieldDetails test extended fields
func (client LogAnalyticsClient) ValidateSourceExtendedFieldDetails(ctx context.Context, request ValidateSourceExtendedFieldDetailsRequest) (response ValidateSourceExtendedFieldDetailsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.validateSourceExtendedFieldDetails, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateSourceExtendedFieldDetailsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateSourceExtendedFieldDetailsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateSourceExtendedFieldDetailsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateSourceExtendedFieldDetailsResponse")
	}
	return
}

// validateSourceExtendedFieldDetails implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateSourceExtendedFieldDetails(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/sources/actions/validateExtendedFields")
	if err != nil {
		return nil, err
	}

	var response ValidateSourceExtendedFieldDetailsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ValidateSourceMapping Validates the source mapping for given file and provides match status and parsed representation of log data.
func (client LogAnalyticsClient) ValidateSourceMapping(ctx context.Context, request ValidateSourceMappingRequest) (response ValidateSourceMappingResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.validateSourceMapping, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ValidateSourceMappingResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ValidateSourceMappingResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ValidateSourceMappingResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ValidateSourceMappingResponse")
	}
	return
}

// validateSourceMapping implements the OCIOperation interface (enables retrying operations)
func (client LogAnalyticsClient) validateSourceMapping(ctx context.Context, request common.OCIRequest) (common.OCIResponse, error) {
	httpRequest, err := request.HTTPRequest(http.MethodPost, "/namespaces/{namespaceName}/uploads/actions/validateSourceMapping")
	if err != nil {
		return nil, err
	}

	var response ValidateSourceMappingResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
