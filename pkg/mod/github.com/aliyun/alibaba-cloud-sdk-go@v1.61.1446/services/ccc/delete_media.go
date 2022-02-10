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

// DeleteMedia invokes the ccc.DeleteMedia API synchronously
func (client *Client) DeleteMedia(request *DeleteMediaRequest) (response *DeleteMediaResponse, err error) {
	response = CreateDeleteMediaResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMediaWithChan invokes the ccc.DeleteMedia API asynchronously
func (client *Client) DeleteMediaWithChan(request *DeleteMediaRequest) (<-chan *DeleteMediaResponse, <-chan error) {
	responseChan := make(chan *DeleteMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMedia(request)
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

// DeleteMediaWithCallback invokes the ccc.DeleteMedia API asynchronously
func (client *Client) DeleteMediaWithCallback(request *DeleteMediaRequest, callback func(response *DeleteMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMediaResponse
		var err error
		defer close(result)
		response, err = client.DeleteMedia(request)
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

// DeleteMediaRequest is the request struct for api DeleteMedia
type DeleteMediaRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
}

// DeleteMediaResponse is the response struct for api DeleteMedia
type DeleteMediaResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeleteMediaRequest creates a request to invoke DeleteMedia API
func CreateDeleteMediaRequest() (request *DeleteMediaRequest) {
	request = &DeleteMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "DeleteMedia", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMediaResponse creates a response to parse from DeleteMedia response
func CreateDeleteMediaResponse() (response *DeleteMediaResponse) {
	response = &DeleteMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
