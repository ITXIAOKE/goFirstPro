package ccc

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

// PublishContactFlowVersion invokes the ccc.PublishContactFlowVersion API synchronously
func (client *Client) PublishContactFlowVersion(request *PublishContactFlowVersionRequest) (response *PublishContactFlowVersionResponse, err error) {
	response = CreatePublishContactFlowVersionResponse()
	err = client.DoAction(request, response)
	return
}

// PublishContactFlowVersionWithChan invokes the ccc.PublishContactFlowVersion API asynchronously
func (client *Client) PublishContactFlowVersionWithChan(request *PublishContactFlowVersionRequest) (<-chan *PublishContactFlowVersionResponse, <-chan error) {
	responseChan := make(chan *PublishContactFlowVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishContactFlowVersion(request)
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

// PublishContactFlowVersionWithCallback invokes the ccc.PublishContactFlowVersion API asynchronously
func (client *Client) PublishContactFlowVersionWithCallback(request *PublishContactFlowVersionRequest, callback func(response *PublishContactFlowVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishContactFlowVersionResponse
		var err error
		defer close(result)
		response, err = client.PublishContactFlowVersion(request)
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

// PublishContactFlowVersionRequest is the request struct for api PublishContactFlowVersion
type PublishContactFlowVersionRequest struct {
	*requests.RpcRequest
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ContactFlowVersionId string           `position:"Query" name:"ContactFlowVersionId"`
	UseTianGong          requests.Boolean `position:"Query" name:"UseTianGong"`
}

// PublishContactFlowVersionResponse is the response struct for api PublishContactFlowVersion
type PublishContactFlowVersionResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreatePublishContactFlowVersionRequest creates a request to invoke PublishContactFlowVersion API
func CreatePublishContactFlowVersionRequest() (request *PublishContactFlowVersionRequest) {
	request = &PublishContactFlowVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "PublishContactFlowVersion", "", "")
	request.Method = requests.POST
	return
}

// CreatePublishContactFlowVersionResponse creates a response to parse from PublishContactFlowVersion response
func CreatePublishContactFlowVersionResponse() (response *PublishContactFlowVersionResponse) {
	response = &PublishContactFlowVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
