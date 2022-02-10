package foas

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

// UpdateProject invokes the foas.UpdateProject API synchronously
func (client *Client) UpdateProject(request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
	response = CreateUpdateProjectResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateProjectWithChan invokes the foas.UpdateProject API asynchronously
func (client *Client) UpdateProjectWithChan(request *UpdateProjectRequest) (<-chan *UpdateProjectResponse, <-chan error) {
	responseChan := make(chan *UpdateProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateProject(request)
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

// UpdateProjectWithCallback invokes the foas.UpdateProject API asynchronously
func (client *Client) UpdateProjectWithCallback(request *UpdateProjectRequest, callback func(response *UpdateProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateProjectResponse
		var err error
		defer close(result)
		response, err = client.UpdateProject(request)
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

// UpdateProjectRequest is the request struct for api UpdateProject
type UpdateProjectRequest struct {
	*requests.RoaRequest
	ProjectName     string `position:"Path" name:"projectName"`
	GlobalJobConfig string `position:"Body" name:"globalJobConfig"`
}

// UpdateProjectResponse is the response struct for api UpdateProject
type UpdateProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateProjectRequest creates a request to invoke UpdateProject API
func CreateUpdateProjectRequest() (request *UpdateProjectRequest) {
	request = &UpdateProjectRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "UpdateProject", "/api/v2/projects/[projectName]", "foas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateProjectResponse creates a response to parse from UpdateProject response
func CreateUpdateProjectResponse() (response *UpdateProjectResponse) {
	response = &UpdateProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
