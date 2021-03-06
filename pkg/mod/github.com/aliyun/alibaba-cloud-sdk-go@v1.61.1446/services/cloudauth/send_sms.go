package cloudauth

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

// SendSms invokes the cloudauth.SendSms API synchronously
func (client *Client) SendSms(request *SendSmsRequest) (response *SendSmsResponse, err error) {
	response = CreateSendSmsResponse()
	err = client.DoAction(request, response)
	return
}

// SendSmsWithChan invokes the cloudauth.SendSms API asynchronously
func (client *Client) SendSmsWithChan(request *SendSmsRequest) (<-chan *SendSmsResponse, <-chan error) {
	responseChan := make(chan *SendSmsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendSms(request)
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

// SendSmsWithCallback invokes the cloudauth.SendSms API asynchronously
func (client *Client) SendSmsWithCallback(request *SendSmsRequest, callback func(response *SendSmsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendSmsResponse
		var err error
		defer close(result)
		response, err = client.SendSms(request)
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

// SendSmsRequest is the request struct for api SendSms
type SendSmsRequest struct {
	*requests.RpcRequest
	Mobile        string `position:"Body" name:"Mobile"`
	SignName      string `position:"Body" name:"SignName"`
	OuterOrderNo  string `position:"Body" name:"OuterOrderNo"`
	TemplateCode  string `position:"Body" name:"TemplateCode"`
	TemplateParam string `position:"Body" name:"TemplateParam"`
}

// SendSmsResponse is the response struct for api SendSms
type SendSmsResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateSendSmsRequest creates a request to invoke SendSms API
func CreateSendSmsRequest() (request *SendSmsRequest) {
	request = &SendSmsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2020-06-18", "SendSms", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSendSmsResponse creates a response to parse from SendSms response
func CreateSendSmsResponse() (response *SendSmsResponse) {
	response = &SendSmsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
