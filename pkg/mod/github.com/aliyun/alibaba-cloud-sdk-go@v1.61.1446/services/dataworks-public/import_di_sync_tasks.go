package dataworks_public

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

// ImportDISyncTasks invokes the dataworks_public.ImportDISyncTasks API synchronously
func (client *Client) ImportDISyncTasks(request *ImportDISyncTasksRequest) (response *ImportDISyncTasksResponse, err error) {
	response = CreateImportDISyncTasksResponse()
	err = client.DoAction(request, response)
	return
}

// ImportDISyncTasksWithChan invokes the dataworks_public.ImportDISyncTasks API asynchronously
func (client *Client) ImportDISyncTasksWithChan(request *ImportDISyncTasksRequest) (<-chan *ImportDISyncTasksResponse, <-chan error) {
	responseChan := make(chan *ImportDISyncTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportDISyncTasks(request)
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

// ImportDISyncTasksWithCallback invokes the dataworks_public.ImportDISyncTasks API asynchronously
func (client *Client) ImportDISyncTasksWithCallback(request *ImportDISyncTasksRequest, callback func(response *ImportDISyncTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportDISyncTasksResponse
		var err error
		defer close(result)
		response, err = client.ImportDISyncTasks(request)
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

// ImportDISyncTasksRequest is the request struct for api ImportDISyncTasks
type ImportDISyncTasksRequest struct {
	*requests.RpcRequest
	TaskType  string           `position:"Query" name:"TaskType"`
	TaskParam string           `position:"Query" name:"TaskParam"`
	Body      string           `position:"Body" name:"body"`
	ProjectId requests.Integer `position:"Query" name:"ProjectId"`
}

// ImportDISyncTasksResponse is the response struct for api ImportDISyncTasks
type ImportDISyncTasksResponse struct {
	*responses.BaseResponse
	Success   bool                        `json:"Success" xml:"Success"`
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	TaskInfo  TaskInfoInImportDISyncTasks `json:"TaskInfo" xml:"TaskInfo"`
}

// CreateImportDISyncTasksRequest creates a request to invoke ImportDISyncTasks API
func CreateImportDISyncTasksRequest() (request *ImportDISyncTasksRequest) {
	request = &ImportDISyncTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ImportDISyncTasks", "", "")
	request.Method = requests.POST
	return
}

// CreateImportDISyncTasksResponse creates a response to parse from ImportDISyncTasks response
func CreateImportDISyncTasksResponse() (response *ImportDISyncTasksResponse) {
	response = &ImportDISyncTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
