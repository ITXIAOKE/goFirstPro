package scsp

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

// ListOutboundPhoneNumber invokes the scsp.ListOutboundPhoneNumber API synchronously
func (client *Client) ListOutboundPhoneNumber(request *ListOutboundPhoneNumberRequest) (response *ListOutboundPhoneNumberResponse, err error) {
	response = CreateListOutboundPhoneNumberResponse()
	err = client.DoAction(request, response)
	return
}

// ListOutboundPhoneNumberWithChan invokes the scsp.ListOutboundPhoneNumber API asynchronously
func (client *Client) ListOutboundPhoneNumberWithChan(request *ListOutboundPhoneNumberRequest) (<-chan *ListOutboundPhoneNumberResponse, <-chan error) {
	responseChan := make(chan *ListOutboundPhoneNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOutboundPhoneNumber(request)
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

// ListOutboundPhoneNumberWithCallback invokes the scsp.ListOutboundPhoneNumber API asynchronously
func (client *Client) ListOutboundPhoneNumberWithCallback(request *ListOutboundPhoneNumberRequest, callback func(response *ListOutboundPhoneNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOutboundPhoneNumberResponse
		var err error
		defer close(result)
		response, err = client.ListOutboundPhoneNumber(request)
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

// ListOutboundPhoneNumberRequest is the request struct for api ListOutboundPhoneNumber
type ListOutboundPhoneNumberRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query"`
	InstanceId  string `position:"Query"`
	AccountName string `position:"Query"`
}

// ListOutboundPhoneNumberResponse is the response struct for api ListOutboundPhoneNumber
type ListOutboundPhoneNumberResponse struct {
	*responses.BaseResponse
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Code           string   `json:"Code" xml:"Code"`
	Success        bool     `json:"Success" xml:"Success"`
	HttpStatusCode int64    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           []string `json:"Data" xml:"Data"`
}

// CreateListOutboundPhoneNumberRequest creates a request to invoke ListOutboundPhoneNumber API
func CreateListOutboundPhoneNumberRequest() (request *ListOutboundPhoneNumberRequest) {
	request = &ListOutboundPhoneNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "ListOutboundPhoneNumber", "", "")
	request.Method = requests.GET
	return
}

// CreateListOutboundPhoneNumberResponse creates a response to parse from ListOutboundPhoneNumber response
func CreateListOutboundPhoneNumberResponse() (response *ListOutboundPhoneNumberResponse) {
	response = &ListOutboundPhoneNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
