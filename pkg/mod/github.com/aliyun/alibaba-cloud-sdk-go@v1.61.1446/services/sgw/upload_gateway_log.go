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

// UploadGatewayLog invokes the sgw.UploadGatewayLog API synchronously
func (client *Client) UploadGatewayLog(request *UploadGatewayLogRequest) (response *UploadGatewayLogResponse, err error) {
	response = CreateUploadGatewayLogResponse()
	err = client.DoAction(request, response)
	return
}

// UploadGatewayLogWithChan invokes the sgw.UploadGatewayLog API asynchronously
func (client *Client) UploadGatewayLogWithChan(request *UploadGatewayLogRequest) (<-chan *UploadGatewayLogResponse, <-chan error) {
	responseChan := make(chan *UploadGatewayLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadGatewayLog(request)
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

// UploadGatewayLogWithCallback invokes the sgw.UploadGatewayLog API asynchronously
func (client *Client) UploadGatewayLogWithCallback(request *UploadGatewayLogRequest, callback func(response *UploadGatewayLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadGatewayLogResponse
		var err error
		defer close(result)
		response, err = client.UploadGatewayLog(request)
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

// UploadGatewayLogRequest is the request struct for api UploadGatewayLog
type UploadGatewayLogRequest struct {
	*requests.RpcRequest
	GatewayId string `position:"Query" name:"GatewayId"`
}

// UploadGatewayLogResponse is the response struct for api UploadGatewayLog
type UploadGatewayLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateUploadGatewayLogRequest creates a request to invoke UploadGatewayLog API
func CreateUploadGatewayLogRequest() (request *UploadGatewayLogRequest) {
	request = &UploadGatewayLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "UploadGatewayLog", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUploadGatewayLogResponse creates a response to parse from UploadGatewayLog response
func CreateUploadGatewayLogResponse() (response *UploadGatewayLogResponse) {
	response = &UploadGatewayLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
