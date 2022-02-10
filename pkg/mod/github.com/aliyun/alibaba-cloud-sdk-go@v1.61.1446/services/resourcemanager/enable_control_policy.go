package resourcemanager

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

// EnableControlPolicy invokes the resourcemanager.EnableControlPolicy API synchronously
func (client *Client) EnableControlPolicy(request *EnableControlPolicyRequest) (response *EnableControlPolicyResponse, err error) {
	response = CreateEnableControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// EnableControlPolicyWithChan invokes the resourcemanager.EnableControlPolicy API asynchronously
func (client *Client) EnableControlPolicyWithChan(request *EnableControlPolicyRequest) (<-chan *EnableControlPolicyResponse, <-chan error) {
	responseChan := make(chan *EnableControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableControlPolicy(request)
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

// EnableControlPolicyWithCallback invokes the resourcemanager.EnableControlPolicy API asynchronously
func (client *Client) EnableControlPolicyWithCallback(request *EnableControlPolicyRequest, callback func(response *EnableControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.EnableControlPolicy(request)
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

// EnableControlPolicyRequest is the request struct for api EnableControlPolicy
type EnableControlPolicyRequest struct {
	*requests.RpcRequest
}

// EnableControlPolicyResponse is the response struct for api EnableControlPolicy
type EnableControlPolicyResponse struct {
	*responses.BaseResponse
	EnablementStatus string `json:"EnablementStatus" xml:"EnablementStatus"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableControlPolicyRequest creates a request to invoke EnableControlPolicy API
func CreateEnableControlPolicyRequest() (request *EnableControlPolicyRequest) {
	request = &EnableControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "EnableControlPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableControlPolicyResponse creates a response to parse from EnableControlPolicy response
func CreateEnableControlPolicyResponse() (response *EnableControlPolicyResponse) {
	response = &EnableControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
