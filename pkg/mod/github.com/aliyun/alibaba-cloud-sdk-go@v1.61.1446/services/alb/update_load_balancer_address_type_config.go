package alb

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

// UpdateLoadBalancerAddressTypeConfig invokes the alb.UpdateLoadBalancerAddressTypeConfig API synchronously
func (client *Client) UpdateLoadBalancerAddressTypeConfig(request *UpdateLoadBalancerAddressTypeConfigRequest) (response *UpdateLoadBalancerAddressTypeConfigResponse, err error) {
	response = CreateUpdateLoadBalancerAddressTypeConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLoadBalancerAddressTypeConfigWithChan invokes the alb.UpdateLoadBalancerAddressTypeConfig API asynchronously
func (client *Client) UpdateLoadBalancerAddressTypeConfigWithChan(request *UpdateLoadBalancerAddressTypeConfigRequest) (<-chan *UpdateLoadBalancerAddressTypeConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLoadBalancerAddressTypeConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLoadBalancerAddressTypeConfig(request)
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

// UpdateLoadBalancerAddressTypeConfigWithCallback invokes the alb.UpdateLoadBalancerAddressTypeConfig API asynchronously
func (client *Client) UpdateLoadBalancerAddressTypeConfigWithCallback(request *UpdateLoadBalancerAddressTypeConfigRequest, callback func(response *UpdateLoadBalancerAddressTypeConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLoadBalancerAddressTypeConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLoadBalancerAddressTypeConfig(request)
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

// UpdateLoadBalancerAddressTypeConfigRequest is the request struct for api UpdateLoadBalancerAddressTypeConfig
type UpdateLoadBalancerAddressTypeConfigRequest struct {
	*requests.RpcRequest
	ClientToken    string                                             `position:"Query" name:"ClientToken"`
	AddressType    string                                             `position:"Query" name:"AddressType"`
	DryRun         string                                             `position:"Query" name:"DryRun"`
	ZoneMappings   *[]UpdateLoadBalancerAddressTypeConfigZoneMappings `position:"Query" name:"ZoneMappings"  type:"Repeated"`
	LoadBalancerId string                                             `position:"Query" name:"LoadBalancerId"`
}

// UpdateLoadBalancerAddressTypeConfigZoneMappings is a repeated param struct in UpdateLoadBalancerAddressTypeConfigRequest
type UpdateLoadBalancerAddressTypeConfigZoneMappings struct {
	VSwitchId    string `name:"VSwitchId"`
	ZoneId       string `name:"ZoneId"`
	AllocationId string `name:"AllocationId"`
}

// UpdateLoadBalancerAddressTypeConfigResponse is the response struct for api UpdateLoadBalancerAddressTypeConfig
type UpdateLoadBalancerAddressTypeConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateUpdateLoadBalancerAddressTypeConfigRequest creates a request to invoke UpdateLoadBalancerAddressTypeConfig API
func CreateUpdateLoadBalancerAddressTypeConfigRequest() (request *UpdateLoadBalancerAddressTypeConfigRequest) {
	request = &UpdateLoadBalancerAddressTypeConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "UpdateLoadBalancerAddressTypeConfig", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLoadBalancerAddressTypeConfigResponse creates a response to parse from UpdateLoadBalancerAddressTypeConfig response
func CreateUpdateLoadBalancerAddressTypeConfigResponse() (response *UpdateLoadBalancerAddressTypeConfigResponse) {
	response = &UpdateLoadBalancerAddressTypeConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
