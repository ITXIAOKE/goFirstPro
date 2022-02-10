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

// UpdateDepartment invokes the idrsservice.UpdateDepartment API synchronously
func (client *Client) UpdateDepartment(request *UpdateDepartmentRequest) (response *UpdateDepartmentResponse, err error) {
	response = CreateUpdateDepartmentResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDepartmentWithChan invokes the idrsservice.UpdateDepartment API asynchronously
func (client *Client) UpdateDepartmentWithChan(request *UpdateDepartmentRequest) (<-chan *UpdateDepartmentResponse, <-chan error) {
	responseChan := make(chan *UpdateDepartmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDepartment(request)
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

// UpdateDepartmentWithCallback invokes the idrsservice.UpdateDepartment API asynchronously
func (client *Client) UpdateDepartmentWithCallback(request *UpdateDepartmentRequest, callback func(response *UpdateDepartmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDepartmentResponse
		var err error
		defer close(result)
		response, err = client.UpdateDepartment(request)
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

// UpdateDepartmentRequest is the request struct for api UpdateDepartment
type UpdateDepartmentRequest struct {
	*requests.RpcRequest
	Description string `position:"Body" name:"Description"`
	Label       string `position:"Body" name:"Label"`
	Name        string `position:"Body" name:"Name"`
	Id          string `position:"Query" name:"Id"`
}

// UpdateDepartmentResponse is the response struct for api UpdateDepartment
type UpdateDepartmentResponse struct {
	*responses.BaseResponse
	Code      string                 `json:"Code" xml:"Code"`
	Data      map[string]interface{} `json:"Data" xml:"Data"`
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDepartmentRequest creates a request to invoke UpdateDepartment API
func CreateUpdateDepartmentRequest() (request *UpdateDepartmentRequest) {
	request = &UpdateDepartmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "UpdateDepartment", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDepartmentResponse creates a response to parse from UpdateDepartment response
func CreateUpdateDepartmentResponse() (response *UpdateDepartmentResponse) {
	response = &UpdateDepartmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
