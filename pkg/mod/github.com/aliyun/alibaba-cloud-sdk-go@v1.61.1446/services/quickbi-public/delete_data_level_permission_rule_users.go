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

// DeleteDataLevelPermissionRuleUsers invokes the quickbi_public.DeleteDataLevelPermissionRuleUsers API synchronously
func (client *Client) DeleteDataLevelPermissionRuleUsers(request *DeleteDataLevelPermissionRuleUsersRequest) (response *DeleteDataLevelPermissionRuleUsersResponse, err error) {
	response = CreateDeleteDataLevelPermissionRuleUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDataLevelPermissionRuleUsersWithChan invokes the quickbi_public.DeleteDataLevelPermissionRuleUsers API asynchronously
func (client *Client) DeleteDataLevelPermissionRuleUsersWithChan(request *DeleteDataLevelPermissionRuleUsersRequest) (<-chan *DeleteDataLevelPermissionRuleUsersResponse, <-chan error) {
	responseChan := make(chan *DeleteDataLevelPermissionRuleUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDataLevelPermissionRuleUsers(request)
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

// DeleteDataLevelPermissionRuleUsersWithCallback invokes the quickbi_public.DeleteDataLevelPermissionRuleUsers API asynchronously
func (client *Client) DeleteDataLevelPermissionRuleUsersWithCallback(request *DeleteDataLevelPermissionRuleUsersRequest, callback func(response *DeleteDataLevelPermissionRuleUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDataLevelPermissionRuleUsersResponse
		var err error
		defer close(result)
		response, err = client.DeleteDataLevelPermissionRuleUsers(request)
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

// DeleteDataLevelPermissionRuleUsersRequest is the request struct for api DeleteDataLevelPermissionRuleUsers
type DeleteDataLevelPermissionRuleUsersRequest struct {
	*requests.RpcRequest
	AccessPoint     string `position:"Query" name:"AccessPoint"`
	SignType        string `position:"Query" name:"SignType"`
	DeleteUserModel string `position:"Query" name:"DeleteUserModel"`
}

// DeleteDataLevelPermissionRuleUsersResponse is the response struct for api DeleteDataLevelPermissionRuleUsers
type DeleteDataLevelPermissionRuleUsersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteDataLevelPermissionRuleUsersRequest creates a request to invoke DeleteDataLevelPermissionRuleUsers API
func CreateDeleteDataLevelPermissionRuleUsersRequest() (request *DeleteDataLevelPermissionRuleUsersRequest) {
	request = &DeleteDataLevelPermissionRuleUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2020-11-11", "DeleteDataLevelPermissionRuleUsers", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDataLevelPermissionRuleUsersResponse creates a response to parse from DeleteDataLevelPermissionRuleUsers response
func CreateDeleteDataLevelPermissionRuleUsersResponse() (response *DeleteDataLevelPermissionRuleUsersResponse) {
	response = &DeleteDataLevelPermissionRuleUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
