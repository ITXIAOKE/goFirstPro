package quickbi_public

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

// AddWorkspaceUsers invokes the quickbi_public.AddWorkspaceUsers API synchronously
func (client *Client) AddWorkspaceUsers(request *AddWorkspaceUsersRequest) (response *AddWorkspaceUsersResponse, err error) {
	response = CreateAddWorkspaceUsersResponse()
	err = client.DoAction(request, response)
	return
}

// AddWorkspaceUsersWithChan invokes the quickbi_public.AddWorkspaceUsers API asynchronously
func (client *Client) AddWorkspaceUsersWithChan(request *AddWorkspaceUsersRequest) (<-chan *AddWorkspaceUsersResponse, <-chan error) {
	responseChan := make(chan *AddWorkspaceUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddWorkspaceUsers(request)
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

// AddWorkspaceUsersWithCallback invokes the quickbi_public.AddWorkspaceUsers API asynchronously
func (client *Client) AddWorkspaceUsersWithCallback(request *AddWorkspaceUsersRequest, callback func(response *AddWorkspaceUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddWorkspaceUsersResponse
		var err error
		defer close(result)
		response, err = client.AddWorkspaceUsers(request)
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

// AddWorkspaceUsersRequest is the request struct for api AddWorkspaceUsers
type AddWorkspaceUsersRequest struct {
	*requests.RpcRequest
	UserIds     string           `position:"Query" name:"UserIds"`
	RoleId      requests.Integer `position:"Query" name:"RoleId"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	SignType    string           `position:"Query" name:"SignType"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// AddWorkspaceUsersResponse is the response struct for api AddWorkspaceUsers
type AddWorkspaceUsersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateAddWorkspaceUsersRequest creates a request to invoke AddWorkspaceUsers API
func CreateAddWorkspaceUsersRequest() (request *AddWorkspaceUsersRequest) {
	request = &AddWorkspaceUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2020-08-03", "AddWorkspaceUsers", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddWorkspaceUsersResponse creates a response to parse from AddWorkspaceUsers response
func CreateAddWorkspaceUsersResponse() (response *AddWorkspaceUsersResponse) {
	response = &AddWorkspaceUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
