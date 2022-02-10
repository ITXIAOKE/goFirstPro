package aliyuncvc

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

// StartLive invokes the aliyuncvc.StartLive API synchronously
func (client *Client) StartLive(request *StartLiveRequest) (response *StartLiveResponse, err error) {
	response = CreateStartLiveResponse()
	err = client.DoAction(request, response)
	return
}

// StartLiveWithChan invokes the aliyuncvc.StartLive API asynchronously
func (client *Client) StartLiveWithChan(request *StartLiveRequest) (<-chan *StartLiveResponse, <-chan error) {
	responseChan := make(chan *StartLiveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartLive(request)
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

// StartLiveWithCallback invokes the aliyuncvc.StartLive API asynchronously
func (client *Client) StartLiveWithCallback(request *StartLiveRequest, callback func(response *StartLiveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartLiveResponse
		var err error
		defer close(result)
		response, err = client.StartLive(request)
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

// StartLiveRequest is the request struct for api StartLive
type StartLiveRequest struct {
	*requests.RpcRequest
	LiveUUID   string `position:"Body" name:"LiveUUID"`
	PushInfo   string `position:"Body" name:"PushInfo"`
	UserId     string `position:"Body" name:"UserId"`
	LayoutInfo string `position:"Body" name:"LayoutInfo"`
}

// StartLiveResponse is the response struct for api StartLive
type StartLiveResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartLiveRequest creates a request to invoke StartLive API
func CreateStartLiveRequest() (request *StartLiveRequest) {
	request = &StartLiveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "StartLive", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartLiveResponse creates a response to parse from StartLive response
func CreateStartLiveResponse() (response *StartLiveResponse) {
	response = &StartLiveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
