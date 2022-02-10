package edas

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

// AuthorizeApplication invokes the edas.AuthorizeApplication API synchronously
func (client *Client) AuthorizeApplication(request *AuthorizeApplicationRequest) (response *AuthorizeApplicationResponse, err error) {
	response = CreateAuthorizeApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// AuthorizeApplicationWithChan invokes the edas.AuthorizeApplication API asynchronously
func (client *Client) AuthorizeApplicationWithChan(request *AuthorizeApplicationRequest) (<-chan *AuthorizeApplicationResponse, <-chan error) {
	responseChan := make(chan *AuthorizeApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthorizeApplication(request)
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

// AuthorizeApplicationWithCallback invokes the edas.AuthorizeApplication API asynchronously
func (client *Client) AuthorizeApplicationWithCallback(request *AuthorizeApplicationRequest, callback func(response *AuthorizeApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthorizeApplicationResponse
		var err error
		defer close(result)
		response, err = client.AuthorizeApplication(request)
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

// AuthorizeApplicationRequest is the request struct for api AuthorizeApplication
type AuthorizeApplicationRequest struct {
	*requests.RoaRequest
	AppIds       string `position:"Query" name:"AppIds"`
	TargetUserId string `position:"Query" name:"TargetUserId"`
}

// AuthorizeApplicationResponse is the response struct for api AuthorizeApplication
type AuthorizeApplicationResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAuthorizeApplicationRequest creates a request to invoke AuthorizeApplication API
func CreateAuthorizeApplicationRequest() (request *AuthorizeApplicationRequest) {
	request = &AuthorizeApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "AuthorizeApplication", "/pop/v5/account/authorize_app", "edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAuthorizeApplicationResponse creates a response to parse from AuthorizeApplication response
func CreateAuthorizeApplicationResponse() (response *AuthorizeApplicationResponse) {
	response = &AuthorizeApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}