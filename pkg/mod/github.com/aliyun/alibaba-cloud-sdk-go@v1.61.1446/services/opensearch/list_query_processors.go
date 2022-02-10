package opensearch

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

// ListQueryProcessors invokes the opensearch.ListQueryProcessors API synchronously
func (client *Client) ListQueryProcessors(request *ListQueryProcessorsRequest) (response *ListQueryProcessorsResponse, err error) {
	response = CreateListQueryProcessorsResponse()
	err = client.DoAction(request, response)
	return
}

// ListQueryProcessorsWithChan invokes the opensearch.ListQueryProcessors API asynchronously
func (client *Client) ListQueryProcessorsWithChan(request *ListQueryProcessorsRequest) (<-chan *ListQueryProcessorsResponse, <-chan error) {
	responseChan := make(chan *ListQueryProcessorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQueryProcessors(request)
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

// ListQueryProcessorsWithCallback invokes the opensearch.ListQueryProcessors API asynchronously
func (client *Client) ListQueryProcessorsWithCallback(request *ListQueryProcessorsRequest, callback func(response *ListQueryProcessorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQueryProcessorsResponse
		var err error
		defer close(result)
		response, err = client.ListQueryProcessors(request)
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

// ListQueryProcessorsRequest is the request struct for api ListQueryProcessors
type ListQueryProcessorsRequest struct {
	*requests.RoaRequest
	AppId            requests.Integer `position:"Path" name:"appId"`
	IsActive         requests.Integer `position:"Query" name:"isActive"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// ListQueryProcessorsResponse is the response struct for api ListQueryProcessors
type ListQueryProcessorsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"requestId" xml:"requestId"`
	Result    []ResultItem `json:"result" xml:"result"`
}

// CreateListQueryProcessorsRequest creates a request to invoke ListQueryProcessors API
func CreateListQueryProcessorsRequest() (request *ListQueryProcessorsRequest) {
	request = &ListQueryProcessorsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListQueryProcessors", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/query-processors", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListQueryProcessorsResponse creates a response to parse from ListQueryProcessors response
func CreateListQueryProcessorsResponse() (response *ListQueryProcessorsResponse) {
	response = &ListQueryProcessorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
