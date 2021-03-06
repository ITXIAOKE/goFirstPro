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

// UnTagVodResources invokes the vod.UnTagVodResources API synchronously
func (client *Client) UnTagVodResources(request *UnTagVodResourcesRequest) (response *UnTagVodResourcesResponse, err error) {
	response = CreateUnTagVodResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// UnTagVodResourcesWithChan invokes the vod.UnTagVodResources API asynchronously
func (client *Client) UnTagVodResourcesWithChan(request *UnTagVodResourcesRequest) (<-chan *UnTagVodResourcesResponse, <-chan error) {
	responseChan := make(chan *UnTagVodResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnTagVodResources(request)
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

// UnTagVodResourcesWithCallback invokes the vod.UnTagVodResources API asynchronously
func (client *Client) UnTagVodResourcesWithCallback(request *UnTagVodResourcesRequest, callback func(response *UnTagVodResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnTagVodResourcesResponse
		var err error
		defer close(result)
		response, err = client.UnTagVodResources(request)
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

// UnTagVodResourcesRequest is the request struct for api UnTagVodResources
type UnTagVodResourcesRequest struct {
	*requests.RpcRequest
	All          requests.Boolean `position:"Query" name:"All"`
	ResourceId   *[]string        `position:"Query" name:"ResourceId"  type:"Repeated"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	TagKey       *[]string        `position:"Query" name:"TagKey"  type:"Repeated"`
}

// UnTagVodResourcesResponse is the response struct for api UnTagVodResources
type UnTagVodResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnTagVodResourcesRequest creates a request to invoke UnTagVodResources API
func CreateUnTagVodResourcesRequest() (request *UnTagVodResourcesRequest) {
	request = &UnTagVodResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "UnTagVodResources", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnTagVodResourcesResponse creates a response to parse from UnTagVodResources response
func CreateUnTagVodResourcesResponse() (response *UnTagVodResourcesResponse) {
	response = &UnTagVodResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
