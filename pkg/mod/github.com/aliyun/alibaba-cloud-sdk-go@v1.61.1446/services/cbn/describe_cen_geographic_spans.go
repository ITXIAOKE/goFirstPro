package cbn

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

// DescribeCenGeographicSpans invokes the cbn.DescribeCenGeographicSpans API synchronously
func (client *Client) DescribeCenGeographicSpans(request *DescribeCenGeographicSpansRequest) (response *DescribeCenGeographicSpansResponse, err error) {
	response = CreateDescribeCenGeographicSpansResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCenGeographicSpansWithChan invokes the cbn.DescribeCenGeographicSpans API asynchronously
func (client *Client) DescribeCenGeographicSpansWithChan(request *DescribeCenGeographicSpansRequest) (<-chan *DescribeCenGeographicSpansResponse, <-chan error) {
	responseChan := make(chan *DescribeCenGeographicSpansResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCenGeographicSpans(request)
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

// DescribeCenGeographicSpansWithCallback invokes the cbn.DescribeCenGeographicSpans API asynchronously
func (client *Client) DescribeCenGeographicSpansWithCallback(request *DescribeCenGeographicSpansRequest, callback func(response *DescribeCenGeographicSpansResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCenGeographicSpansResponse
		var err error
		defer close(result)
		response, err = client.DescribeCenGeographicSpans(request)
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

// DescribeCenGeographicSpansRequest is the request struct for api DescribeCenGeographicSpans
type DescribeCenGeographicSpansRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	GeographicSpanId     string           `position:"Query" name:"GeographicSpanId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCenGeographicSpansResponse is the response struct for api DescribeCenGeographicSpans
type DescribeCenGeographicSpansResponse struct {
	*responses.BaseResponse
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	PageNumber           int                  `json:"PageNumber" xml:"PageNumber"`
	TotalCount           int                  `json:"TotalCount" xml:"TotalCount"`
	GeographicSpanModels GeographicSpanModels `json:"GeographicSpanModels" xml:"GeographicSpanModels"`
}

// CreateDescribeCenGeographicSpansRequest creates a request to invoke DescribeCenGeographicSpans API
func CreateDescribeCenGeographicSpansRequest() (request *DescribeCenGeographicSpansRequest) {
	request = &DescribeCenGeographicSpansRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenGeographicSpans", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCenGeographicSpansResponse creates a response to parse from DescribeCenGeographicSpans response
func CreateDescribeCenGeographicSpansResponse() (response *DescribeCenGeographicSpansResponse) {
	response = &DescribeCenGeographicSpansResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}