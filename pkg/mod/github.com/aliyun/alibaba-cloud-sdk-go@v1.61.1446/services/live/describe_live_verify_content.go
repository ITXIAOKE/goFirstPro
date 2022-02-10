package live

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

// DescribeLiveVerifyContent invokes the live.DescribeLiveVerifyContent API synchronously
func (client *Client) DescribeLiveVerifyContent(request *DescribeLiveVerifyContentRequest) (response *DescribeLiveVerifyContentResponse, err error) {
	response = CreateDescribeLiveVerifyContentResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveVerifyContentWithChan invokes the live.DescribeLiveVerifyContent API asynchronously
func (client *Client) DescribeLiveVerifyContentWithChan(request *DescribeLiveVerifyContentRequest) (<-chan *DescribeLiveVerifyContentResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveVerifyContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveVerifyContent(request)
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

// DescribeLiveVerifyContentWithCallback invokes the live.DescribeLiveVerifyContent API asynchronously
func (client *Client) DescribeLiveVerifyContentWithCallback(request *DescribeLiveVerifyContentRequest, callback func(response *DescribeLiveVerifyContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveVerifyContentResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveVerifyContent(request)
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

// DescribeLiveVerifyContentRequest is the request struct for api DescribeLiveVerifyContent
type DescribeLiveVerifyContentRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveVerifyContentResponse is the response struct for api DescribeLiveVerifyContent
type DescribeLiveVerifyContentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Content   string `json:"Content" xml:"Content"`
}

// CreateDescribeLiveVerifyContentRequest creates a request to invoke DescribeLiveVerifyContent API
func CreateDescribeLiveVerifyContentRequest() (request *DescribeLiveVerifyContentRequest) {
	request = &DescribeLiveVerifyContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveVerifyContent", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveVerifyContentResponse creates a response to parse from DescribeLiveVerifyContent response
func CreateDescribeLiveVerifyContentResponse() (response *DescribeLiveVerifyContentResponse) {
	response = &DescribeLiveVerifyContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
