package idrsservice

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

// CreateStatisticsTask invokes the idrsservice.CreateStatisticsTask API synchronously
func (client *Client) CreateStatisticsTask(request *CreateStatisticsTaskRequest) (response *CreateStatisticsTaskResponse, err error) {
	response = CreateCreateStatisticsTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStatisticsTaskWithChan invokes the idrsservice.CreateStatisticsTask API asynchronously
func (client *Client) CreateStatisticsTaskWithChan(request *CreateStatisticsTaskRequest) (<-chan *CreateStatisticsTaskResponse, <-chan error) {
	responseChan := make(chan *CreateStatisticsTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStatisticsTask(request)
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

// CreateStatisticsTaskWithCallback invokes the idrsservice.CreateStatisticsTask API asynchronously
func (client *Client) CreateStatisticsTaskWithCallback(request *CreateStatisticsTaskRequest, callback func(response *CreateStatisticsTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStatisticsTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateStatisticsTask(request)
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

// CreateStatisticsTaskRequest is the request struct for api CreateStatisticsTask
type CreateStatisticsTaskRequest struct {
	*requests.RpcRequest
	ClientToken  string    `position:"Query" name:"ClientToken"`
	DepartmentId *[]string `position:"Query" name:"DepartmentId"  type:"Repeated"`
	DateTo       string    `position:"Query" name:"DateTo"`
	DateFrom     string    `position:"Query" name:"DateFrom"`
}

// CreateStatisticsTaskResponse is the response struct for api CreateStatisticsTask
type CreateStatisticsTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateStatisticsTaskRequest creates a request to invoke CreateStatisticsTask API
func CreateCreateStatisticsTaskRequest() (request *CreateStatisticsTaskRequest) {
	request = &CreateStatisticsTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "CreateStatisticsTask", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateStatisticsTaskResponse creates a response to parse from CreateStatisticsTask response
func CreateCreateStatisticsTaskResponse() (response *CreateStatisticsTaskResponse) {
	response = &CreateStatisticsTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
