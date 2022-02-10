package mse

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

// ListAlarmItems invokes the mse.ListAlarmItems API synchronously
func (client *Client) ListAlarmItems(request *ListAlarmItemsRequest) (response *ListAlarmItemsResponse, err error) {
	response = CreateListAlarmItemsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAlarmItemsWithChan invokes the mse.ListAlarmItems API asynchronously
func (client *Client) ListAlarmItemsWithChan(request *ListAlarmItemsRequest) (<-chan *ListAlarmItemsResponse, <-chan error) {
	responseChan := make(chan *ListAlarmItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAlarmItems(request)
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

// ListAlarmItemsWithCallback invokes the mse.ListAlarmItems API asynchronously
func (client *Client) ListAlarmItemsWithCallback(request *ListAlarmItemsRequest, callback func(response *ListAlarmItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAlarmItemsResponse
		var err error
		defer close(result)
		response, err = client.ListAlarmItems(request)
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

// ListAlarmItemsRequest is the request struct for api ListAlarmItems
type ListAlarmItemsRequest struct {
	*requests.RpcRequest
	RequestPars string `position:"Query" name:"RequestPars"`
}

// ListAlarmItemsResponse is the response struct for api ListAlarmItems
type ListAlarmItemsResponse struct {
	*responses.BaseResponse
	HttpCode   string      `json:"HttpCode" xml:"HttpCode"`
	TotalCount int         `json:"TotalCount" xml:"TotalCount"`
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Message    string      `json:"Message" xml:"Message"`
	PageSize   int         `json:"PageSize" xml:"PageSize"`
	PageNumber int         `json:"PageNumber" xml:"PageNumber"`
	ErrorCode  string      `json:"ErrorCode" xml:"ErrorCode"`
	Success    bool        `json:"Success" xml:"Success"`
	Data       []AlarmItem `json:"Data" xml:"Data"`
}

// CreateListAlarmItemsRequest creates a request to invoke ListAlarmItems API
func CreateListAlarmItemsRequest() (request *ListAlarmItemsRequest) {
	request = &ListAlarmItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListAlarmItems", "", "")
	request.Method = requests.GET
	return
}

// CreateListAlarmItemsResponse creates a response to parse from ListAlarmItems response
func CreateListAlarmItemsResponse() (response *ListAlarmItemsResponse) {
	response = &ListAlarmItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}