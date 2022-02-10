package eipanycast

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

// ModifyAnycastEipAddressAttribute invokes the eipanycast.ModifyAnycastEipAddressAttribute API synchronously
func (client *Client) ModifyAnycastEipAddressAttribute(request *ModifyAnycastEipAddressAttributeRequest) (response *ModifyAnycastEipAddressAttributeResponse, err error) {
	response = CreateModifyAnycastEipAddressAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAnycastEipAddressAttributeWithChan invokes the eipanycast.ModifyAnycastEipAddressAttribute API asynchronously
func (client *Client) ModifyAnycastEipAddressAttributeWithChan(request *ModifyAnycastEipAddressAttributeRequest) (<-chan *ModifyAnycastEipAddressAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyAnycastEipAddressAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAnycastEipAddressAttribute(request)
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

// ModifyAnycastEipAddressAttributeWithCallback invokes the eipanycast.ModifyAnycastEipAddressAttribute API asynchronously
func (client *Client) ModifyAnycastEipAddressAttributeWithCallback(request *ModifyAnycastEipAddressAttributeRequest, callback func(response *ModifyAnycastEipAddressAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAnycastEipAddressAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyAnycastEipAddressAttribute(request)
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

// ModifyAnycastEipAddressAttributeRequest is the request struct for api ModifyAnycastEipAddressAttribute
type ModifyAnycastEipAddressAttributeRequest struct {
	*requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	AnycastId   string `position:"Query" name:"AnycastId"`
	Name        string `position:"Query" name:"Name"`
}

// ModifyAnycastEipAddressAttributeResponse is the response struct for api ModifyAnycastEipAddressAttribute
type ModifyAnycastEipAddressAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAnycastEipAddressAttributeRequest creates a request to invoke ModifyAnycastEipAddressAttribute API
func CreateModifyAnycastEipAddressAttributeRequest() (request *ModifyAnycastEipAddressAttributeRequest) {
	request = &ModifyAnycastEipAddressAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eipanycast", "2020-03-09", "ModifyAnycastEipAddressAttribute", "eipanycast", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyAnycastEipAddressAttributeResponse creates a response to parse from ModifyAnycastEipAddressAttribute response
func CreateModifyAnycastEipAddressAttributeResponse() (response *ModifyAnycastEipAddressAttributeResponse) {
	response = &ModifyAnycastEipAddressAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
