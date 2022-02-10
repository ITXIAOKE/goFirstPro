package oos

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

// ListTemplateVersions invokes the oos.ListTemplateVersions API synchronously
func (client *Client) ListTemplateVersions(request *ListTemplateVersionsRequest) (response *ListTemplateVersionsResponse, err error) {
	response = CreateListTemplateVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTemplateVersionsWithChan invokes the oos.ListTemplateVersions API asynchronously
func (client *Client) ListTemplateVersionsWithChan(request *ListTemplateVersionsRequest) (<-chan *ListTemplateVersionsResponse, <-chan error) {
	responseChan := make(chan *ListTemplateVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTemplateVersions(request)
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

// ListTemplateVersionsWithCallback invokes the oos.ListTemplateVersions API asynchronously
func (client *Client) ListTemplateVersionsWithCallback(request *ListTemplateVersionsRequest, callback func(response *ListTemplateVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTemplateVersionsResponse
		var err error
		defer close(result)
		response, err = client.ListTemplateVersions(request)
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

// ListTemplateVersionsRequest is the request struct for api ListTemplateVersions
type ListTemplateVersionsRequest struct {
	*requests.RpcRequest
	NextToken    string           `position:"Query" name:"NextToken"`
	MaxResults   requests.Integer `position:"Query" name:"MaxResults"`
	TemplateName string           `position:"Query" name:"TemplateName"`
	ShareType    string           `position:"Query" name:"ShareType"`
}

// ListTemplateVersionsResponse is the response struct for api ListTemplateVersions
type ListTemplateVersionsResponse struct {
	*responses.BaseResponse
	RequestId        string            `json:"RequestId" xml:"RequestId"`
	MaxResults       int               `json:"MaxResults" xml:"MaxResults"`
	NextToken        string            `json:"NextToken" xml:"NextToken"`
	TemplateVersions []TemplateVersion `json:"TemplateVersions" xml:"TemplateVersions"`
}

// CreateListTemplateVersionsRequest creates a request to invoke ListTemplateVersions API
func CreateListTemplateVersionsRequest() (request *ListTemplateVersionsRequest) {
	request = &ListTemplateVersionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ListTemplateVersions", "", "")
	request.Method = requests.POST
	return
}

// CreateListTemplateVersionsResponse creates a response to parse from ListTemplateVersions response
func CreateListTemplateVersionsResponse() (response *ListTemplateVersionsResponse) {
	response = &ListTemplateVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
