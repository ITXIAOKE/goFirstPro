package servicemesh

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

// UpdateIstioInjectionConfig invokes the servicemesh.UpdateIstioInjectionConfig API synchronously
func (client *Client) UpdateIstioInjectionConfig(request *UpdateIstioInjectionConfigRequest) (response *UpdateIstioInjectionConfigResponse, err error) {
	response = CreateUpdateIstioInjectionConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateIstioInjectionConfigWithChan invokes the servicemesh.UpdateIstioInjectionConfig API asynchronously
func (client *Client) UpdateIstioInjectionConfigWithChan(request *UpdateIstioInjectionConfigRequest) (<-chan *UpdateIstioInjectionConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateIstioInjectionConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateIstioInjectionConfig(request)
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

// UpdateIstioInjectionConfigWithCallback invokes the servicemesh.UpdateIstioInjectionConfig API asynchronously
func (client *Client) UpdateIstioInjectionConfigWithCallback(request *UpdateIstioInjectionConfigRequest, callback func(response *UpdateIstioInjectionConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateIstioInjectionConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateIstioInjectionConfig(request)
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

// UpdateIstioInjectionConfigRequest is the request struct for api UpdateIstioInjectionConfig
type UpdateIstioInjectionConfigRequest struct {
	*requests.RpcRequest
	EnableIstioInjection requests.Boolean `position:"Body" name:"EnableIstioInjection"`
	Namespace            string           `position:"Body" name:"Namespace"`
	ServiceMeshId        string           `position:"Body" name:"ServiceMeshId"`
}

// UpdateIstioInjectionConfigResponse is the response struct for api UpdateIstioInjectionConfig
type UpdateIstioInjectionConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateIstioInjectionConfigRequest creates a request to invoke UpdateIstioInjectionConfig API
func CreateUpdateIstioInjectionConfigRequest() (request *UpdateIstioInjectionConfigRequest) {
	request = &UpdateIstioInjectionConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("servicemesh", "2020-01-11", "UpdateIstioInjectionConfig", "servicemesh", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateIstioInjectionConfigResponse creates a response to parse from UpdateIstioInjectionConfig response
func CreateUpdateIstioInjectionConfigResponse() (response *UpdateIstioInjectionConfigResponse) {
	response = &UpdateIstioInjectionConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}