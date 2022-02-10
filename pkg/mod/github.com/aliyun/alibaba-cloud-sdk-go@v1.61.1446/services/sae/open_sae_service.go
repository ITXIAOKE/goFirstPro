package sae

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

// OpenSaeService invokes the sae.OpenSaeService API synchronously
func (client *Client) OpenSaeService(request *OpenSaeServiceRequest) (response *OpenSaeServiceResponse, err error) {
	response = CreateOpenSaeServiceResponse()
	err = client.DoAction(request, response)
	return
}

// OpenSaeServiceWithChan invokes the sae.OpenSaeService API asynchronously
func (client *Client) OpenSaeServiceWithChan(request *OpenSaeServiceRequest) (<-chan *OpenSaeServiceResponse, <-chan error) {
	responseChan := make(chan *OpenSaeServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenSaeService(request)
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

// OpenSaeServiceWithCallback invokes the sae.OpenSaeService API asynchronously
func (client *Client) OpenSaeServiceWithCallback(request *OpenSaeServiceRequest, callback func(response *OpenSaeServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenSaeServiceResponse
		var err error
		defer close(result)
		response, err = client.OpenSaeService(request)
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

// OpenSaeServiceRequest is the request struct for api OpenSaeService
type OpenSaeServiceRequest struct {
	*requests.RoaRequest
}

// OpenSaeServiceResponse is the response struct for api OpenSaeService
type OpenSaeServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateOpenSaeServiceRequest creates a request to invoke OpenSaeService API
func CreateOpenSaeServiceRequest() (request *OpenSaeServiceRequest) {
	request = &OpenSaeServiceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "OpenSaeService", "/service/open", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOpenSaeServiceResponse creates a response to parse from OpenSaeService response
func CreateOpenSaeServiceResponse() (response *OpenSaeServiceResponse) {
	response = &OpenSaeServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}