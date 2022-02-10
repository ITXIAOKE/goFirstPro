package ehpc

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

// InitializeEHPC invokes the ehpc.InitializeEHPC API synchronously
func (client *Client) InitializeEHPC(request *InitializeEHPCRequest) (response *InitializeEHPCResponse, err error) {
	response = CreateInitializeEHPCResponse()
	err = client.DoAction(request, response)
	return
}

// InitializeEHPCWithChan invokes the ehpc.InitializeEHPC API asynchronously
func (client *Client) InitializeEHPCWithChan(request *InitializeEHPCRequest) (<-chan *InitializeEHPCResponse, <-chan error) {
	responseChan := make(chan *InitializeEHPCResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitializeEHPC(request)
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

// InitializeEHPCWithCallback invokes the ehpc.InitializeEHPC API asynchronously
func (client *Client) InitializeEHPCWithCallback(request *InitializeEHPCRequest, callback func(response *InitializeEHPCResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitializeEHPCResponse
		var err error
		defer close(result)
		response, err = client.InitializeEHPC(request)
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

// InitializeEHPCRequest is the request struct for api InitializeEHPC
type InitializeEHPCRequest struct {
	*requests.RpcRequest
}

// InitializeEHPCResponse is the response struct for api InitializeEHPC
type InitializeEHPCResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInitializeEHPCRequest creates a request to invoke InitializeEHPC API
func CreateInitializeEHPCRequest() (request *InitializeEHPCRequest) {
	request = &InitializeEHPCRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "InitializeEHPC", "", "")
	request.Method = requests.GET
	return
}

// CreateInitializeEHPCResponse creates a response to parse from InitializeEHPC response
func CreateInitializeEHPCResponse() (response *InitializeEHPCResponse) {
	response = &InitializeEHPCResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
