package sgw

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

// OpenSgwService invokes the sgw.OpenSgwService API synchronously
func (client *Client) OpenSgwService(request *OpenSgwServiceRequest) (response *OpenSgwServiceResponse, err error) {
	response = CreateOpenSgwServiceResponse()
	err = client.DoAction(request, response)
	return
}

// OpenSgwServiceWithChan invokes the sgw.OpenSgwService API asynchronously
func (client *Client) OpenSgwServiceWithChan(request *OpenSgwServiceRequest) (<-chan *OpenSgwServiceResponse, <-chan error) {
	responseChan := make(chan *OpenSgwServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenSgwService(request)
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

// OpenSgwServiceWithCallback invokes the sgw.OpenSgwService API asynchronously
func (client *Client) OpenSgwServiceWithCallback(request *OpenSgwServiceRequest, callback func(response *OpenSgwServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenSgwServiceResponse
		var err error
		defer close(result)
		response, err = client.OpenSgwService(request)
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

// OpenSgwServiceRequest is the request struct for api OpenSgwService
type OpenSgwServiceRequest struct {
	*requests.RpcRequest
}

// OpenSgwServiceResponse is the response struct for api OpenSgwService
type OpenSgwServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateOpenSgwServiceRequest creates a request to invoke OpenSgwService API
func CreateOpenSgwServiceRequest() (request *OpenSgwServiceRequest) {
	request = &OpenSgwServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "OpenSgwService", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOpenSgwServiceResponse creates a response to parse from OpenSgwService response
func CreateOpenSgwServiceResponse() (response *OpenSgwServiceResponse) {
	response = &OpenSgwServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
