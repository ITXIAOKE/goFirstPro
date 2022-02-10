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

// ModifyMPULayout invokes the rtc.ModifyMPULayout API synchronously
func (client *Client) ModifyMPULayout(request *ModifyMPULayoutRequest) (response *ModifyMPULayoutResponse, err error) {
	response = CreateModifyMPULayoutResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMPULayoutWithChan invokes the rtc.ModifyMPULayout API asynchronously
func (client *Client) ModifyMPULayoutWithChan(request *ModifyMPULayoutRequest) (<-chan *ModifyMPULayoutResponse, <-chan error) {
	responseChan := make(chan *ModifyMPULayoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMPULayout(request)
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

// ModifyMPULayoutWithCallback invokes the rtc.ModifyMPULayout API asynchronously
func (client *Client) ModifyMPULayoutWithCallback(request *ModifyMPULayoutRequest, callback func(response *ModifyMPULayoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMPULayoutResponse
		var err error
		defer close(result)
		response, err = client.ModifyMPULayout(request)
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

// ModifyMPULayoutRequest is the request struct for api ModifyMPULayout
type ModifyMPULayoutRequest struct {
	*requests.RpcRequest
	LayoutId      requests.Integer        `position:"Query" name:"LayoutId"`
	Panes         *[]ModifyMPULayoutPanes `position:"Query" name:"Panes"  type:"Repeated"`
	ShowLog       string                  `position:"Query" name:"ShowLog"`
	OwnerId       requests.Integer        `position:"Query" name:"OwnerId"`
	AppId         string                  `position:"Query" name:"AppId"`
	Name          string                  `position:"Query" name:"Name"`
	AudioMixCount requests.Integer        `position:"Query" name:"AudioMixCount"`
}

// ModifyMPULayoutPanes is a repeated param struct in ModifyMPULayoutRequest
type ModifyMPULayoutPanes struct {
	MajorPane string `name:"MajorPane"`
	Width     string `name:"Width"`
	Height    string `name:"Height"`
	Y         string `name:"Y"`
	PaneId    string `name:"PaneId"`
	ZOrder    string `name:"ZOrder"`
	X         string `name:"X"`
}

// ModifyMPULayoutResponse is the response struct for api ModifyMPULayout
type ModifyMPULayoutResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyMPULayoutRequest creates a request to invoke ModifyMPULayout API
func CreateModifyMPULayoutRequest() (request *ModifyMPULayoutRequest) {
	request = &ModifyMPULayoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "ModifyMPULayout", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyMPULayoutResponse creates a response to parse from ModifyMPULayout response
func CreateModifyMPULayoutResponse() (response *ModifyMPULayoutResponse) {
	response = &ModifyMPULayoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
