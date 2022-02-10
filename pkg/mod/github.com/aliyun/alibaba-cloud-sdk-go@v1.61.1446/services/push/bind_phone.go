package push

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

// BindPhone invokes the push.BindPhone API synchronously
func (client *Client) BindPhone(request *BindPhoneRequest) (response *BindPhoneResponse, err error) {
	response = CreateBindPhoneResponse()
	err = client.DoAction(request, response)
	return
}

// BindPhoneWithChan invokes the push.BindPhone API asynchronously
func (client *Client) BindPhoneWithChan(request *BindPhoneRequest) (<-chan *BindPhoneResponse, <-chan error) {
	responseChan := make(chan *BindPhoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindPhone(request)
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

// BindPhoneWithCallback invokes the push.BindPhone API asynchronously
func (client *Client) BindPhoneWithCallback(request *BindPhoneRequest, callback func(response *BindPhoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindPhoneResponse
		var err error
		defer close(result)
		response, err = client.BindPhone(request)
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

// BindPhoneRequest is the request struct for api BindPhone
type BindPhoneRequest struct {
	*requests.RpcRequest
	PhoneNumber string           `position:"Query" name:"PhoneNumber"`
	DeviceId    string           `position:"Query" name:"DeviceId"`
	AppKey      requests.Integer `position:"Query" name:"AppKey"`
}

// BindPhoneResponse is the response struct for api BindPhone
type BindPhoneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBindPhoneRequest creates a request to invoke BindPhone API
func CreateBindPhoneRequest() (request *BindPhoneRequest) {
	request = &BindPhoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "BindPhone", "", "")
	request.Method = requests.POST
	return
}

// CreateBindPhoneResponse creates a response to parse from BindPhone response
func CreateBindPhoneResponse() (response *BindPhoneResponse) {
	response = &BindPhoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
