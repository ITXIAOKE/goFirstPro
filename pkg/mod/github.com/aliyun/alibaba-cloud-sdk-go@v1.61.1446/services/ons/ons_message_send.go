package ons

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

// OnsMessageSend invokes the ons.OnsMessageSend API synchronously
func (client *Client) OnsMessageSend(request *OnsMessageSendRequest) (response *OnsMessageSendResponse, err error) {
	response = CreateOnsMessageSendResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMessageSendWithChan invokes the ons.OnsMessageSend API asynchronously
func (client *Client) OnsMessageSendWithChan(request *OnsMessageSendRequest) (<-chan *OnsMessageSendResponse, <-chan error) {
	responseChan := make(chan *OnsMessageSendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMessageSend(request)
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

// OnsMessageSendWithCallback invokes the ons.OnsMessageSend API asynchronously
func (client *Client) OnsMessageSendWithCallback(request *OnsMessageSendRequest, callback func(response *OnsMessageSendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMessageSendResponse
		var err error
		defer close(result)
		response, err = client.OnsMessageSend(request)
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

// OnsMessageSendRequest is the request struct for api OnsMessageSend
type OnsMessageSendRequest struct {
	*requests.RpcRequest
	Message    string `position:"Query" name:"Message"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
	Tag        string `position:"Query" name:"Tag"`
	Key        string `position:"Query" name:"Key"`
}

// OnsMessageSendResponse is the response struct for api OnsMessageSend
type OnsMessageSendResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsMessageSendRequest creates a request to invoke OnsMessageSend API
func CreateOnsMessageSendRequest() (request *OnsMessageSendRequest) {
	request = &OnsMessageSendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMessageSend", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsMessageSendResponse creates a response to parse from OnsMessageSend response
func CreateOnsMessageSendResponse() (response *OnsMessageSendResponse) {
	response = &OnsMessageSendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
