package iot

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

// ListProductTags invokes the iot.ListProductTags API synchronously
func (client *Client) ListProductTags(request *ListProductTagsRequest) (response *ListProductTagsResponse, err error) {
	response = CreateListProductTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ListProductTagsWithChan invokes the iot.ListProductTags API asynchronously
func (client *Client) ListProductTagsWithChan(request *ListProductTagsRequest) (<-chan *ListProductTagsResponse, <-chan error) {
	responseChan := make(chan *ListProductTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProductTags(request)
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

// ListProductTagsWithCallback invokes the iot.ListProductTags API asynchronously
func (client *Client) ListProductTagsWithCallback(request *ListProductTagsRequest, callback func(response *ListProductTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProductTagsResponse
		var err error
		defer close(result)
		response, err = client.ListProductTags(request)
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

// ListProductTagsRequest is the request struct for api ListProductTags
type ListProductTagsRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
}

// ListProductTagsResponse is the response struct for api ListProductTags
type ListProductTagsResponse struct {
	*responses.BaseResponse
	RequestId    string                `json:"RequestId" xml:"RequestId"`
	Success      bool                  `json:"Success" xml:"Success"`
	ErrorMessage string                `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string                `json:"Code" xml:"Code"`
	Data         DataInListProductTags `json:"Data" xml:"Data"`
}

// CreateListProductTagsRequest creates a request to invoke ListProductTags API
func CreateListProductTagsRequest() (request *ListProductTagsRequest) {
	request = &ListProductTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListProductTags", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListProductTagsResponse creates a response to parse from ListProductTags response
func CreateListProductTagsResponse() (response *ListProductTagsResponse) {
	response = &ListProductTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}