package arms

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

// StartAlert invokes the arms.StartAlert API synchronously
func (client *Client) StartAlert(request *StartAlertRequest) (response *StartAlertResponse, err error) {
	response = CreateStartAlertResponse()
	err = client.DoAction(request, response)
	return
}

// StartAlertWithChan invokes the arms.StartAlert API asynchronously
func (client *Client) StartAlertWithChan(request *StartAlertRequest) (<-chan *StartAlertResponse, <-chan error) {
	responseChan := make(chan *StartAlertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartAlert(request)
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

// StartAlertWithCallback invokes the arms.StartAlert API asynchronously
func (client *Client) StartAlertWithCallback(request *StartAlertRequest, callback func(response *StartAlertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartAlertResponse
		var err error
		defer close(result)
		response, err = client.StartAlert(request)
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

// StartAlertRequest is the request struct for api StartAlert
type StartAlertRequest struct {
	*requests.RpcRequest
	AlertId     string `position:"Query" name:"AlertId"`
	ProxyUserId string `position:"Query" name:"ProxyUserId"`
}

// StartAlertResponse is the response struct for api StartAlert
type StartAlertResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	StartAlertIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
}

// CreateStartAlertRequest creates a request to invoke StartAlert API
func CreateStartAlertRequest() (request *StartAlertRequest) {
	request = &StartAlertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "StartAlert", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartAlertResponse creates a response to parse from StartAlert response
func CreateStartAlertResponse() (response *StartAlertResponse) {
	response = &StartAlertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}