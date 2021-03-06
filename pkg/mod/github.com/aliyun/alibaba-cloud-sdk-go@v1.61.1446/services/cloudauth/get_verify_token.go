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

// GetVerifyToken invokes the cloudauth.GetVerifyToken API synchronously
func (client *Client) GetVerifyToken(request *GetVerifyTokenRequest) (response *GetVerifyTokenResponse, err error) {
	response = CreateGetVerifyTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetVerifyTokenWithChan invokes the cloudauth.GetVerifyToken API asynchronously
func (client *Client) GetVerifyTokenWithChan(request *GetVerifyTokenRequest) (<-chan *GetVerifyTokenResponse, <-chan error) {
	responseChan := make(chan *GetVerifyTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVerifyToken(request)
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

// GetVerifyTokenWithCallback invokes the cloudauth.GetVerifyToken API asynchronously
func (client *Client) GetVerifyTokenWithCallback(request *GetVerifyTokenRequest, callback func(response *GetVerifyTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVerifyTokenResponse
		var err error
		defer close(result)
		response, err = client.GetVerifyToken(request)
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

// GetVerifyTokenRequest is the request struct for api GetVerifyToken
type GetVerifyTokenRequest struct {
	*requests.RpcRequest
	Binding       string `position:"Body" name:"Binding"`
	VerifyConfigs string `position:"Query" name:"VerifyConfigs"`
	UserData      string `position:"Query" name:"UserData"`
	Biz           string `position:"Query" name:"Biz"`
	TicketId      string `position:"Query" name:"TicketId"`
}

// GetVerifyTokenResponse is the response struct for api GetVerifyToken
type GetVerifyTokenResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetVerifyTokenRequest creates a request to invoke GetVerifyToken API
func CreateGetVerifyTokenRequest() (request *GetVerifyTokenRequest) {
	request = &GetVerifyTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2018-09-16", "GetVerifyToken", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVerifyTokenResponse creates a response to parse from GetVerifyToken response
func CreateGetVerifyTokenResponse() (response *GetVerifyTokenResponse) {
	response = &GetVerifyTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
