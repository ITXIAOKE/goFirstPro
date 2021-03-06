package edas

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

// UpdateEnvAppGroup invokes the edas.UpdateEnvAppGroup API synchronously
func (client *Client) UpdateEnvAppGroup(request *UpdateEnvAppGroupRequest) (response *UpdateEnvAppGroupResponse, err error) {
	response = CreateUpdateEnvAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEnvAppGroupWithChan invokes the edas.UpdateEnvAppGroup API asynchronously
func (client *Client) UpdateEnvAppGroupWithChan(request *UpdateEnvAppGroupRequest) (<-chan *UpdateEnvAppGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateEnvAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEnvAppGroup(request)
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

// UpdateEnvAppGroupWithCallback invokes the edas.UpdateEnvAppGroup API asynchronously
func (client *Client) UpdateEnvAppGroupWithCallback(request *UpdateEnvAppGroupRequest, callback func(response *UpdateEnvAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEnvAppGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateEnvAppGroup(request)
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

// UpdateEnvAppGroupRequest is the request struct for api UpdateEnvAppGroup
type UpdateEnvAppGroupRequest struct {
	*requests.RoaRequest
	PointcutName string `position:"Body" name:"PointcutName"`
	Scopes       string `position:"Body" name:"Scopes"`
}

// UpdateEnvAppGroupResponse is the response struct for api UpdateEnvAppGroup
type UpdateEnvAppGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateEnvAppGroupRequest creates a request to invoke UpdateEnvAppGroup API
func CreateUpdateEnvAppGroupRequest() (request *UpdateEnvAppGroupRequest) {
	request = &UpdateEnvAppGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateEnvAppGroup", "/pop/v5/gray/env_app_groups", "edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateEnvAppGroupResponse creates a response to parse from UpdateEnvAppGroup response
func CreateUpdateEnvAppGroupResponse() (response *UpdateEnvAppGroupResponse) {
	response = &UpdateEnvAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
