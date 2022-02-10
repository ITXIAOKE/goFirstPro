package companyreg

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

// ListUserSolutions invokes the companyreg.ListUserSolutions API synchronously
func (client *Client) ListUserSolutions(request *ListUserSolutionsRequest) (response *ListUserSolutionsResponse, err error) {
	response = CreateListUserSolutionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserSolutionsWithChan invokes the companyreg.ListUserSolutions API asynchronously
func (client *Client) ListUserSolutionsWithChan(request *ListUserSolutionsRequest) (<-chan *ListUserSolutionsResponse, <-chan error) {
	responseChan := make(chan *ListUserSolutionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserSolutions(request)
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

// ListUserSolutionsWithCallback invokes the companyreg.ListUserSolutions API asynchronously
func (client *Client) ListUserSolutionsWithCallback(request *ListUserSolutionsRequest, callback func(response *ListUserSolutionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserSolutionsResponse
		var err error
		defer close(result)
		response, err = client.ListUserSolutions(request)
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

// ListUserSolutionsRequest is the request struct for api ListUserSolutions
type ListUserSolutionsRequest struct {
	*requests.RpcRequest
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	IntentionBizId string           `position:"Query" name:"IntentionBizId"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
}

// ListUserSolutionsResponse is the response struct for api ListUserSolutions
type ListUserSolutionsResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int        `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int        `json:"CurrentPageNum" xml:"CurrentPageNum"`
	PageSize       int        `json:"PageSize" xml:"PageSize"`
	TotalPageNum   int        `json:"TotalPageNum" xml:"TotalPageNum"`
	Data           []Solution `json:"Data" xml:"Data"`
}

// CreateListUserSolutionsRequest creates a request to invoke ListUserSolutions API
func CreateListUserSolutionsRequest() (request *ListUserSolutionsRequest) {
	request = &ListUserSolutionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "ListUserSolutions", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListUserSolutionsResponse creates a response to parse from ListUserSolutions response
func CreateListUserSolutionsResponse() (response *ListUserSolutionsResponse) {
	response = &ListUserSolutionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
