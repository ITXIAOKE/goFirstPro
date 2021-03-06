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

// ConfigureSubscriptionInstance invokes the dts.ConfigureSubscriptionInstance API synchronously
func (client *Client) ConfigureSubscriptionInstance(request *ConfigureSubscriptionInstanceRequest) (response *ConfigureSubscriptionInstanceResponse, err error) {
	response = CreateConfigureSubscriptionInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureSubscriptionInstanceWithChan invokes the dts.ConfigureSubscriptionInstance API asynchronously
func (client *Client) ConfigureSubscriptionInstanceWithChan(request *ConfigureSubscriptionInstanceRequest) (<-chan *ConfigureSubscriptionInstanceResponse, <-chan error) {
	responseChan := make(chan *ConfigureSubscriptionInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureSubscriptionInstance(request)
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

// ConfigureSubscriptionInstanceWithCallback invokes the dts.ConfigureSubscriptionInstance API asynchronously
func (client *Client) ConfigureSubscriptionInstanceWithCallback(request *ConfigureSubscriptionInstanceRequest, callback func(response *ConfigureSubscriptionInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureSubscriptionInstanceResponse
		var err error
		defer close(result)
		response, err = client.ConfigureSubscriptionInstance(request)
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

// ConfigureSubscriptionInstanceRequest is the request struct for api ConfigureSubscriptionInstance
type ConfigureSubscriptionInstanceRequest struct {
	*requests.RpcRequest
	SourceEndpointInstanceID        string           `position:"Query" name:"SourceEndpoint.InstanceID"`
	SourceEndpointOracleSID         string           `position:"Query" name:"SourceEndpoint.OracleSID"`
	SourceEndpointIP                string           `position:"Query" name:"SourceEndpoint.IP"`
	SubscriptionDataTypeDML         requests.Boolean `position:"Query" name:"SubscriptionDataType.DML"`
	SourceEndpointInstanceType      string           `position:"Query" name:"SourceEndpoint.InstanceType"`
	AccountId                       string           `position:"Query" name:"AccountId"`
	SubscriptionObject              string           `position:"Body" name:"SubscriptionObject"`
	SubscriptionInstanceVSwitchId   string           `position:"Query" name:"SubscriptionInstance.VSwitchId"`
	SourceEndpointUserName          string           `position:"Query" name:"SourceEndpoint.UserName"`
	SourceEndpointDatabaseName      string           `position:"Query" name:"SourceEndpoint.DatabaseName"`
	SourceEndpointPort              string           `position:"Query" name:"SourceEndpoint.Port"`
	SourceEndpointOwnerID           string           `position:"Query" name:"SourceEndpoint.OwnerID"`
	SubscriptionInstanceVPCId       string           `position:"Query" name:"SubscriptionInstance.VPCId"`
	SubscriptionInstanceNetworkType string           `position:"Query" name:"SubscriptionInstanceNetworkType"`
	SubscriptionInstanceId          string           `position:"Query" name:"SubscriptionInstanceId"`
	SourceEndpointRole              string           `position:"Query" name:"SourceEndpoint.Role"`
	OwnerId                         string           `position:"Query" name:"OwnerId"`
	SubscriptionDataTypeDDL         requests.Boolean `position:"Query" name:"SubscriptionDataType.DDL"`
	SourceEndpointPassword          string           `position:"Query" name:"SourceEndpoint.Password"`
	SubscriptionInstanceName        string           `position:"Query" name:"SubscriptionInstanceName"`
}

// ConfigureSubscriptionInstanceResponse is the response struct for api ConfigureSubscriptionInstance
type ConfigureSubscriptionInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	Success    string `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateConfigureSubscriptionInstanceRequest creates a request to invoke ConfigureSubscriptionInstance API
func CreateConfigureSubscriptionInstanceRequest() (request *ConfigureSubscriptionInstanceRequest) {
	request = &ConfigureSubscriptionInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ConfigureSubscriptionInstance", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigureSubscriptionInstanceResponse creates a response to parse from ConfigureSubscriptionInstance response
func CreateConfigureSubscriptionInstanceResponse() (response *ConfigureSubscriptionInstanceResponse) {
	response = &ConfigureSubscriptionInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
