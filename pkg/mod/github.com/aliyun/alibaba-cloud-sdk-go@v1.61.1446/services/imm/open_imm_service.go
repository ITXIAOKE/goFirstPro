package imm

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

// OpenImmService invokes the imm.OpenImmService API synchronously
func (client *Client) OpenImmService(request *OpenImmServiceRequest) (response *OpenImmServiceResponse, err error) {
	response = CreateOpenImmServiceResponse()
	err = client.DoAction(request, response)
	return
}

// OpenImmServiceWithChan invokes the imm.OpenImmService API asynchronously
func (client *Client) OpenImmServiceWithChan(request *OpenImmServiceRequest) (<-chan *OpenImmServiceResponse, <-chan error) {
	responseChan := make(chan *OpenImmServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenImmService(request)
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

// OpenImmServiceWithCallback invokes the imm.OpenImmService API asynchronously
func (client *Client) OpenImmServiceWithCallback(request *OpenImmServiceRequest, callback func(response *OpenImmServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenImmServiceResponse
		var err error
		defer close(result)
		response, err = client.OpenImmService(request)
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

// OpenImmServiceRequest is the request struct for api OpenImmService
type OpenImmServiceRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// OpenImmServiceResponse is the response struct for api OpenImmService
type OpenImmServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateOpenImmServiceRequest creates a request to invoke OpenImmService API
func CreateOpenImmServiceRequest() (request *OpenImmServiceRequest) {
	request = &OpenImmServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "OpenImmService", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOpenImmServiceResponse creates a response to parse from OpenImmService response
func CreateOpenImmServiceResponse() (response *OpenImmServiceResponse) {
	response = &OpenImmServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}