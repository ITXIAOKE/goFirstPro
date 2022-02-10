package sas

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

// ModifyPushAllTask invokes the sas.ModifyPushAllTask API synchronously
func (client *Client) ModifyPushAllTask(request *ModifyPushAllTaskRequest) (response *ModifyPushAllTaskResponse, err error) {
	response = CreateModifyPushAllTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPushAllTaskWithChan invokes the sas.ModifyPushAllTask API asynchronously
func (client *Client) ModifyPushAllTaskWithChan(request *ModifyPushAllTaskRequest) (<-chan *ModifyPushAllTaskResponse, <-chan error) {
	responseChan := make(chan *ModifyPushAllTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPushAllTask(request)
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

// ModifyPushAllTaskWithCallback invokes the sas.ModifyPushAllTask API asynchronously
func (client *Client) ModifyPushAllTaskWithCallback(request *ModifyPushAllTaskRequest, callback func(response *ModifyPushAllTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPushAllTaskResponse
		var err error
		defer close(result)
		response, err = client.ModifyPushAllTask(request)
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

// ModifyPushAllTaskRequest is the request struct for api ModifyPushAllTask
type ModifyPushAllTaskRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Tasks    string `position:"Query" name:"Tasks"`
	Uuids    string `position:"Query" name:"Uuids"`
}

// ModifyPushAllTaskResponse is the response struct for api ModifyPushAllTask
type ModifyPushAllTaskResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PushTaskRsp PushTaskRsp `json:"PushTaskRsp" xml:"PushTaskRsp"`
}

// CreateModifyPushAllTaskRequest creates a request to invoke ModifyPushAllTask API
func CreateModifyPushAllTaskRequest() (request *ModifyPushAllTaskRequest) {
	request = &ModifyPushAllTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyPushAllTask", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyPushAllTaskResponse creates a response to parse from ModifyPushAllTask response
func CreateModifyPushAllTaskResponse() (response *ModifyPushAllTaskResponse) {
	response = &ModifyPushAllTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}