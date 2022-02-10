package cloudesl

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

// AddUser invokes the cloudesl.AddUser API synchronously
func (client *Client) AddUser(request *AddUserRequest) (response *AddUserResponse, err error) {
	response = CreateAddUserResponse()
	err = client.DoAction(request, response)
	return
}

// AddUserWithChan invokes the cloudesl.AddUser API asynchronously
func (client *Client) AddUserWithChan(request *AddUserRequest) (<-chan *AddUserResponse, <-chan error) {
	responseChan := make(chan *AddUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddUser(request)
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

// AddUserWithCallback invokes the cloudesl.AddUser API asynchronously
func (client *Client) AddUserWithCallback(request *AddUserRequest, callback func(response *AddUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddUserResponse
		var err error
		defer close(result)
		response, err = client.AddUser(request)
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

// AddUserRequest is the request struct for api AddUser
type AddUserRequest struct {
	*requests.RpcRequest
	ExtraParams string `position:"Body" name:"ExtraParams"`
	ClientToken string `position:"Body" name:"ClientToken"`
	UserId      string `position:"Body" name:"UserId"`
}

// AddUserResponse is the response struct for api AddUser
type AddUserResponse struct {
	*responses.BaseResponse
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Message        string `json:"Message" xml:"Message"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	Code           string `json:"Code" xml:"Code"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateAddUserRequest creates a request to invoke AddUser API
func CreateAddUserRequest() (request *AddUserRequest) {
	request = &AddUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "AddUser", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddUserResponse creates a response to parse from AddUser response
func CreateAddUserResponse() (response *AddUserResponse) {
	response = &AddUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
