package rtc

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

// CreateMPULayout invokes the rtc.CreateMPULayout API synchronously
func (client *Client) CreateMPULayout(request *CreateMPULayoutRequest) (response *CreateMPULayoutResponse, err error) {
	response = CreateCreateMPULayoutResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMPULayoutWithChan invokes the rtc.CreateMPULayout API asynchronously
func (client *Client) CreateMPULayoutWithChan(request *CreateMPULayoutRequest) (<-chan *CreateMPULayoutResponse, <-chan error) {
	responseChan := make(chan *CreateMPULayoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMPULayout(request)
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

// CreateMPULayoutWithCallback invokes the rtc.CreateMPULayout API asynchronously
func (client *Client) CreateMPULayoutWithCallback(request *CreateMPULayoutRequest, callback func(response *CreateMPULayoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMPULayoutResponse
		var err error
		defer close(result)
		response, err = client.CreateMPULayout(request)
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

// CreateMPULayoutRequest is the request struct for api CreateMPULayout
type CreateMPULayoutRequest struct {
	*requests.RpcRequest
	Panes         *[]CreateMPULayoutPanes `position:"Query" name:"Panes"  type:"Repeated"`
	ShowLog       string                  `position:"Query" name:"ShowLog"`
	OwnerId       requests.Integer        `position:"Query" name:"OwnerId"`
	AppId         string                  `position:"Query" name:"AppId"`
	Name          string                  `position:"Query" name:"Name"`
	AudioMixCount requests.Integer        `position:"Query" name:"AudioMixCount"`
}

// CreateMPULayoutPanes is a repeated param struct in CreateMPULayoutRequest
type CreateMPULayoutPanes struct {
	MajorPane string `name:"MajorPane"`
	Width     string `name:"Width"`
	Height    string `name:"Height"`
	Y         string `name:"Y"`
	PaneId    string `name:"PaneId"`
	ZOrder    string `name:"ZOrder"`
	X         string `name:"X"`
}

// CreateMPULayoutResponse is the response struct for api CreateMPULayout
type CreateMPULayoutResponse struct {
	*responses.BaseResponse
	LayoutId  int64  `json:"LayoutId" xml:"LayoutId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateMPULayoutRequest creates a request to invoke CreateMPULayout API
func CreateCreateMPULayoutRequest() (request *CreateMPULayoutRequest) {
	request = &CreateMPULayoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "CreateMPULayout", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMPULayoutResponse creates a response to parse from CreateMPULayout response
func CreateCreateMPULayoutResponse() (response *CreateMPULayoutResponse) {
	response = &CreateMPULayoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
