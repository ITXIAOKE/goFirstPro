package dcdn

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

// UntagDcdnResources invokes the dcdn.UntagDcdnResources API synchronously
func (client *Client) UntagDcdnResources(request *UntagDcdnResourcesRequest) (response *UntagDcdnResourcesResponse, err error) {
	response = CreateUntagDcdnResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// UntagDcdnResourcesWithChan invokes the dcdn.UntagDcdnResources API asynchronously
func (client *Client) UntagDcdnResourcesWithChan(request *UntagDcdnResourcesRequest) (<-chan *UntagDcdnResourcesResponse, <-chan error) {
	responseChan := make(chan *UntagDcdnResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UntagDcdnResources(request)
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

// UntagDcdnResourcesWithCallback invokes the dcdn.UntagDcdnResources API asynchronously
func (client *Client) UntagDcdnResourcesWithCallback(request *UntagDcdnResourcesRequest, callback func(response *UntagDcdnResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UntagDcdnResourcesResponse
		var err error
		defer close(result)
		response, err = client.UntagDcdnResources(request)
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

// UntagDcdnResourcesRequest is the request struct for api UntagDcdnResources
type UntagDcdnResourcesRequest struct {
	*requests.RpcRequest
	All          requests.Boolean `position:"Query" name:"All"`
	ResourceId   *[]string        `position:"Query" name:"ResourceId"  type:"Repeated"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	TagKey       *[]string        `position:"Query" name:"TagKey"  type:"Repeated"`
}

// UntagDcdnResourcesResponse is the response struct for api UntagDcdnResources
type UntagDcdnResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUntagDcdnResourcesRequest creates a request to invoke UntagDcdnResources API
func CreateUntagDcdnResourcesRequest() (request *UntagDcdnResourcesRequest) {
	request = &UntagDcdnResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "UntagDcdnResources", "", "")
	request.Method = requests.POST
	return
}

// CreateUntagDcdnResourcesResponse creates a response to parse from UntagDcdnResources response
func CreateUntagDcdnResourcesResponse() (response *UntagDcdnResourcesResponse) {
	response = &UntagDcdnResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
