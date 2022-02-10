package outboundbot

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

// ListJobGroupsAsync invokes the outboundbot.ListJobGroupsAsync API synchronously
func (client *Client) ListJobGroupsAsync(request *ListJobGroupsAsyncRequest) (response *ListJobGroupsAsyncResponse, err error) {
	response = CreateListJobGroupsAsyncResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobGroupsAsyncWithChan invokes the outboundbot.ListJobGroupsAsync API asynchronously
func (client *Client) ListJobGroupsAsyncWithChan(request *ListJobGroupsAsyncRequest) (<-chan *ListJobGroupsAsyncResponse, <-chan error) {
	responseChan := make(chan *ListJobGroupsAsyncResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobGroupsAsync(request)
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

// ListJobGroupsAsyncWithCallback invokes the outboundbot.ListJobGroupsAsync API asynchronously
func (client *Client) ListJobGroupsAsyncWithCallback(request *ListJobGroupsAsyncRequest, callback func(response *ListJobGroupsAsyncResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobGroupsAsyncResponse
		var err error
		defer close(result)
		response, err = client.ListJobGroupsAsync(request)
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

// ListJobGroupsAsyncRequest is the request struct for api ListJobGroupsAsync
type ListJobGroupsAsyncRequest struct {
	*requests.RpcRequest
	AsyncTaskId string `position:"Query" name:"AsyncTaskId"`
}

// ListJobGroupsAsyncResponse is the response struct for api ListJobGroupsAsync
type ListJobGroupsAsyncResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	TotalCount     int        `json:"TotalCount" xml:"TotalCount"`
	PageNumber     int        `json:"PageNumber" xml:"PageNumber"`
	PageSize       int        `json:"PageSize" xml:"PageSize"`
	Vaild          bool       `json:"Vaild" xml:"Vaild"`
	Timeout        bool       `json:"Timeout" xml:"Timeout"`
	JobGroups      []JobGroup `json:"JobGroups" xml:"JobGroups"`
}

// CreateListJobGroupsAsyncRequest creates a request to invoke ListJobGroupsAsync API
func CreateListJobGroupsAsyncRequest() (request *ListJobGroupsAsyncRequest) {
	request = &ListJobGroupsAsyncRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListJobGroupsAsync", "outboundbot", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListJobGroupsAsyncResponse creates a response to parse from ListJobGroupsAsync response
func CreateListJobGroupsAsyncResponse() (response *ListJobGroupsAsyncResponse) {
	response = &ListJobGroupsAsyncResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
