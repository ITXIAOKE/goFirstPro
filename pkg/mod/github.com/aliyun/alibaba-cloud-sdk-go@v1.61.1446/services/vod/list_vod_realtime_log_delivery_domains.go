package vod

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

// ListVodRealtimeLogDeliveryDomains invokes the vod.ListVodRealtimeLogDeliveryDomains API synchronously
func (client *Client) ListVodRealtimeLogDeliveryDomains(request *ListVodRealtimeLogDeliveryDomainsRequest) (response *ListVodRealtimeLogDeliveryDomainsResponse, err error) {
	response = CreateListVodRealtimeLogDeliveryDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVodRealtimeLogDeliveryDomainsWithChan invokes the vod.ListVodRealtimeLogDeliveryDomains API asynchronously
func (client *Client) ListVodRealtimeLogDeliveryDomainsWithChan(request *ListVodRealtimeLogDeliveryDomainsRequest) (<-chan *ListVodRealtimeLogDeliveryDomainsResponse, <-chan error) {
	responseChan := make(chan *ListVodRealtimeLogDeliveryDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVodRealtimeLogDeliveryDomains(request)
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

// ListVodRealtimeLogDeliveryDomainsWithCallback invokes the vod.ListVodRealtimeLogDeliveryDomains API asynchronously
func (client *Client) ListVodRealtimeLogDeliveryDomainsWithCallback(request *ListVodRealtimeLogDeliveryDomainsRequest, callback func(response *ListVodRealtimeLogDeliveryDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVodRealtimeLogDeliveryDomainsResponse
		var err error
		defer close(result)
		response, err = client.ListVodRealtimeLogDeliveryDomains(request)
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

// ListVodRealtimeLogDeliveryDomainsRequest is the request struct for api ListVodRealtimeLogDeliveryDomains
type ListVodRealtimeLogDeliveryDomainsRequest struct {
	*requests.RpcRequest
	Project  string           `position:"Query" name:"Project"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	Region   string           `position:"Query" name:"Region"`
	Logstore string           `position:"Query" name:"Logstore"`
}

// ListVodRealtimeLogDeliveryDomainsResponse is the response struct for api ListVodRealtimeLogDeliveryDomains
type ListVodRealtimeLogDeliveryDomainsResponse struct {
	*responses.BaseResponse
	RequestId string                                     `json:"RequestId" xml:"RequestId"`
	Content   ContentInListVodRealtimeLogDeliveryDomains `json:"Content" xml:"Content"`
}

// CreateListVodRealtimeLogDeliveryDomainsRequest creates a request to invoke ListVodRealtimeLogDeliveryDomains API
func CreateListVodRealtimeLogDeliveryDomainsRequest() (request *ListVodRealtimeLogDeliveryDomainsRequest) {
	request = &ListVodRealtimeLogDeliveryDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListVodRealtimeLogDeliveryDomains", "vod", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListVodRealtimeLogDeliveryDomainsResponse creates a response to parse from ListVodRealtimeLogDeliveryDomains response
func CreateListVodRealtimeLogDeliveryDomainsResponse() (response *ListVodRealtimeLogDeliveryDomainsResponse) {
	response = &ListVodRealtimeLogDeliveryDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
