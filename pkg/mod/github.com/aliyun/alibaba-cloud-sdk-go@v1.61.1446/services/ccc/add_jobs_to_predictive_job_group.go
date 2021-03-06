package ccc

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

// AddJobsToPredictiveJobGroup invokes the ccc.AddJobsToPredictiveJobGroup API synchronously
func (client *Client) AddJobsToPredictiveJobGroup(request *AddJobsToPredictiveJobGroupRequest) (response *AddJobsToPredictiveJobGroupResponse, err error) {
	response = CreateAddJobsToPredictiveJobGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AddJobsToPredictiveJobGroupWithChan invokes the ccc.AddJobsToPredictiveJobGroup API asynchronously
func (client *Client) AddJobsToPredictiveJobGroupWithChan(request *AddJobsToPredictiveJobGroupRequest) (<-chan *AddJobsToPredictiveJobGroupResponse, <-chan error) {
	responseChan := make(chan *AddJobsToPredictiveJobGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddJobsToPredictiveJobGroup(request)
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

// AddJobsToPredictiveJobGroupWithCallback invokes the ccc.AddJobsToPredictiveJobGroup API asynchronously
func (client *Client) AddJobsToPredictiveJobGroupWithCallback(request *AddJobsToPredictiveJobGroupRequest, callback func(response *AddJobsToPredictiveJobGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddJobsToPredictiveJobGroupResponse
		var err error
		defer close(result)
		response, err = client.AddJobsToPredictiveJobGroup(request)
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

// AddJobsToPredictiveJobGroupRequest is the request struct for api AddJobsToPredictiveJobGroup
type AddJobsToPredictiveJobGroupRequest struct {
	*requests.RpcRequest
	ClientToken  string    `position:"Query" name:"ClientToken"`
	JobsJson     *[]string `position:"Body" name:"JobsJson"  type:"Repeated"`
	InstanceId   string    `position:"Query" name:"InstanceId"`
	SkillGroupId string    `position:"Query" name:"SkillGroupId"`
	JobGroupId   string    `position:"Query" name:"JobGroupId"`
}

// AddJobsToPredictiveJobGroupResponse is the response struct for api AddJobsToPredictiveJobGroup
type AddJobsToPredictiveJobGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateAddJobsToPredictiveJobGroupRequest creates a request to invoke AddJobsToPredictiveJobGroup API
func CreateAddJobsToPredictiveJobGroupRequest() (request *AddJobsToPredictiveJobGroupRequest) {
	request = &AddJobsToPredictiveJobGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "AddJobsToPredictiveJobGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateAddJobsToPredictiveJobGroupResponse creates a response to parse from AddJobsToPredictiveJobGroup response
func CreateAddJobsToPredictiveJobGroupResponse() (response *AddJobsToPredictiveJobGroupResponse) {
	response = &AddJobsToPredictiveJobGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
