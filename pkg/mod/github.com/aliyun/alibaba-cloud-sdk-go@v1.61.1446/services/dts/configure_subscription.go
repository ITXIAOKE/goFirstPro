package dts

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ConfigureSubscription invokes the dts.ConfigureSubscription API synchronously
func (client *Client) ConfigureSubscription(request *ConfigureSubscriptionRequest) (response *ConfigureSubscriptionResponse, err error) {
	response = CreateConfigureSubscriptionResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureSubscriptionWithChan invokes the dts.ConfigureSubscription API asynchronously
func (client *Client) ConfigureSubscriptionWithChan(request *ConfigureSubscriptionRequest) (<-chan *ConfigureSubscriptionResponse, <-chan error) {
	responseChan := make(chan *ConfigureSubscriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureSubscription(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ConfigureSubscriptionWithCallback invokes the dts.ConfigureSubscription API asynchronously
func (client *Client) ConfigureSubscriptionWithCallback(request *ConfigureSubscriptionRequest, callback func(response *ConfigureSubscriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureSubscriptionResponse
		var err error
		defer close(result)
		response, err = client.ConfigureSubscription(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ConfigureSubscriptionRequest is the request struct for api ConfigureSubscription
type ConfigureSubscriptionRequest struct {
	*requests.RpcRequest
	SourceEndpointRegion            string           `position:"Query" name:"SourceEndpointRegion"`
	Checkpoint                      string           `position:"Query" name:"Checkpoint"`
	SubscriptionInstanceVSwitchId   string           `position:"Query" name:"SubscriptionInstanceVSwitchId"`
	DelayNotice                     requests.Boolean `position:"Query" name:"DelayNotice"`
	SourceEndpointInstanceID        string           `position:"Query" name:"SourceEndpointInstanceID"`
	SourceEndpointUserName          string           `position:"Query" name:"SourceEndpointUserName"`
	SourceEndpointOwnerID           string           `position:"Query" name:"SourceEndpointOwnerID"`
	DelayPhone                      string           `position:"Query" name:"DelayPhone"`
	SubscriptionDataTypeDML         requests.Boolean `position:"Query" name:"SubscriptionDataTypeDML"`
	SourceEndpointDatabaseName      string           `position:"Query" name:"SourceEndpointDatabaseName"`
	SourceEndpointIP                string           `position:"Query" name:"SourceEndpointIP"`
	ErrorPhone                      string           `position:"Query" name:"ErrorPhone"`
	Reserve                         string           `position:"Query" name:"Reserve"`
	DtsJobId                        string           `position:"Query" name:"DtsJobId"`
	DbList                          string           `position:"Query" name:"DbList"`
	SubscriptionInstanceNetworkType string           `position:"Query" name:"SubscriptionInstanceNetworkType"`
	SubscriptionDataTypeDDL         requests.Boolean `position:"Query" name:"SubscriptionDataTypeDDL"`
	SourceEndpointPassword          string           `position:"Query" name:"SourceEndpointPassword"`
	SourceEndpointPort              string           `position:"Query" name:"SourceEndpointPort"`
	SubscriptionInstanceVPCId       string           `position:"Query" name:"SubscriptionInstanceVPCId"`
	DelayRuleTime                   requests.Integer `position:"Query" name:"DelayRuleTime"`
	SourceEndpointInstanceType      string           `position:"Query" name:"SourceEndpointInstanceType"`
	DtsJobName                      string           `position:"Query" name:"DtsJobName"`
	SourceEndpointOracleSID         string           `position:"Query" name:"SourceEndpointOracleSID"`
	ErrorNotice                     requests.Boolean `position:"Query" name:"ErrorNotice"`
	SourceEndpointRole              string           `position:"Query" name:"SourceEndpointRole"`
	DtsInstanceId                   string           `position:"Query" name:"DtsInstanceId"`
	SourceEndpointEngineName        string           `position:"Query" name:"SourceEndpointEngineName"`
}

// ConfigureSubscriptionResponse is the response struct for api ConfigureSubscription
type ConfigureSubscriptionResponse struct {
	*responses.BaseResponse
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	DtsJobId       string `json:"DtsJobId" xml:"DtsJobId"`
	Success        string `json:"Success" xml:"Success"`
	DtsInstanceId  string `json:"DtsInstanceId" xml:"DtsInstanceId"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateConfigureSubscriptionRequest creates a request to invoke ConfigureSubscription API
func CreateConfigureSubscriptionRequest() (request *ConfigureSubscriptionRequest) {
	request = &ConfigureSubscriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ConfigureSubscription", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigureSubscriptionResponse creates a response to parse from ConfigureSubscription response
func CreateConfigureSubscriptionResponse() (response *ConfigureSubscriptionResponse) {
	response = &ConfigureSubscriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
