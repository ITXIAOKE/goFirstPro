package codeup

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

// ListRepositoryMemberWithInherited invokes the codeup.ListRepositoryMemberWithInherited API synchronously
func (client *Client) ListRepositoryMemberWithInherited(request *ListRepositoryMemberWithInheritedRequest) (response *ListRepositoryMemberWithInheritedResponse, err error) {
	response = CreateListRepositoryMemberWithInheritedResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepositoryMemberWithInheritedWithChan invokes the codeup.ListRepositoryMemberWithInherited API asynchronously
func (client *Client) ListRepositoryMemberWithInheritedWithChan(request *ListRepositoryMemberWithInheritedRequest) (<-chan *ListRepositoryMemberWithInheritedResponse, <-chan error) {
	responseChan := make(chan *ListRepositoryMemberWithInheritedResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepositoryMemberWithInherited(request)
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

// ListRepositoryMemberWithInheritedWithCallback invokes the codeup.ListRepositoryMemberWithInherited API asynchronously
func (client *Client) ListRepositoryMemberWithInheritedWithCallback(request *ListRepositoryMemberWithInheritedRequest, callback func(response *ListRepositoryMemberWithInheritedResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepositoryMemberWithInheritedResponse
		var err error
		defer close(result)
		response, err = client.ListRepositoryMemberWithInherited(request)
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

// ListRepositoryMemberWithInheritedRequest is the request struct for api ListRepositoryMemberWithInherited
type ListRepositoryMemberWithInheritedRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// ListRepositoryMemberWithInheritedResponse is the response struct for api ListRepositoryMemberWithInherited
type ListRepositoryMemberWithInheritedResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ErrorCode    string       `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool         `json:"Success" xml:"Success"`
	ErrorMessage string       `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       []ResultItem `json:"Result" xml:"Result"`
}

// CreateListRepositoryMemberWithInheritedRequest creates a request to invoke ListRepositoryMemberWithInherited API
func CreateListRepositoryMemberWithInheritedRequest() (request *ListRepositoryMemberWithInheritedRequest) {
	request = &ListRepositoryMemberWithInheritedRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "ListRepositoryMemberWithInherited", "/api/v4/projects/[ProjectId]/all_members", "", "")
	request.Method = requests.GET
	return
}

// CreateListRepositoryMemberWithInheritedResponse creates a response to parse from ListRepositoryMemberWithInherited response
func CreateListRepositoryMemberWithInheritedResponse() (response *ListRepositoryMemberWithInheritedResponse) {
	response = &ListRepositoryMemberWithInheritedResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
