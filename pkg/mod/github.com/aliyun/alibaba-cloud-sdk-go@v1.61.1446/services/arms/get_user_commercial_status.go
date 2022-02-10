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

// GetUserCommercialStatus invokes the arms.GetUserCommercialStatus API synchronously
func (client *Client) GetUserCommercialStatus(request *GetUserCommercialStatusRequest) (response *GetUserCommercialStatusResponse, err error) {
	response = CreateGetUserCommercialStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserCommercialStatusWithChan invokes the arms.GetUserCommercialStatus API asynchronously
func (client *Client) GetUserCommercialStatusWithChan(request *GetUserCommercialStatusRequest) (<-chan *GetUserCommercialStatusResponse, <-chan error) {
	responseChan := make(chan *GetUserCommercialStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserCommercialStatus(request)
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

// GetUserCommercialStatusWithCallback invokes the arms.GetUserCommercialStatus API asynchronously
func (client *Client) GetUserCommercialStatusWithCallback(request *GetUserCommercialStatusRequest, callback func(response *GetUserCommercialStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserCommercialStatusResponse
		var err error
		defer close(result)
		response, err = client.GetUserCommercialStatus(request)
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

// GetUserCommercialStatusRequest is the request struct for api GetUserCommercialStatus
type GetUserCommercialStatusRequest struct {
	*requests.RpcRequest
	UserId       string `position:"Query" name:"UserId"`
	ParentId     string `position:"Query" name:"ParentId"`
	TargetUserId string `position:"Query" name:"TargetUserId"`
}

// GetUserCommercialStatusResponse is the response struct for api GetUserCommercialStatus
type GetUserCommercialStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetUserCommercialStatusRequest creates a request to invoke GetUserCommercialStatus API
func CreateGetUserCommercialStatusRequest() (request *GetUserCommercialStatusRequest) {
	request = &GetUserCommercialStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetUserCommercialStatus", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserCommercialStatusResponse creates a response to parse from GetUserCommercialStatus response
func CreateGetUserCommercialStatusResponse() (response *GetUserCommercialStatusResponse) {
	response = &GetUserCommercialStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
