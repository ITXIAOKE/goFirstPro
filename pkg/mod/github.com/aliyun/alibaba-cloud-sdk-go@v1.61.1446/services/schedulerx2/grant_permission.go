package schedulerx2

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

// GrantPermission invokes the schedulerx2.GrantPermission API synchronously
func (client *Client) GrantPermission(request *GrantPermissionRequest) (response *GrantPermissionResponse, err error) {
	response = CreateGrantPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// GrantPermissionWithChan invokes the schedulerx2.GrantPermission API asynchronously
func (client *Client) GrantPermissionWithChan(request *GrantPermissionRequest) (<-chan *GrantPermissionResponse, <-chan error) {
	responseChan := make(chan *GrantPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrantPermission(request)
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

// GrantPermissionWithCallback invokes the schedulerx2.GrantPermission API asynchronously
func (client *Client) GrantPermissionWithCallback(request *GrantPermissionRequest, callback func(response *GrantPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrantPermissionResponse
		var err error
		defer close(result)
		response, err = client.GrantPermission(request)
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

// GrantPermissionRequest is the request struct for api GrantPermission
type GrantPermissionRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	GroupId         string           `position:"Query" name:"GroupId"`
	UserId          string           `position:"Query" name:"UserId"`
	GrantOption     requests.Boolean `position:"Query" name:"GrantOption"`
	Namespace       string           `position:"Query" name:"Namespace"`
	UserName        string           `position:"Query" name:"UserName"`
}

// GrantPermissionResponse is the response struct for api GrantPermission
type GrantPermissionResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGrantPermissionRequest creates a request to invoke GrantPermission API
func CreateGrantPermissionRequest() (request *GrantPermissionRequest) {
	request = &GrantPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "GrantPermission", "", "")
	request.Method = requests.POST
	return
}

// CreateGrantPermissionResponse creates a response to parse from GrantPermission response
func CreateGrantPermissionResponse() (response *GrantPermissionResponse) {
	response = &GrantPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}