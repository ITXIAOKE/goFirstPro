package vod

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

// SubmitLiveEditing invokes the vod.SubmitLiveEditing API synchronously
func (client *Client) SubmitLiveEditing(request *SubmitLiveEditingRequest) (response *SubmitLiveEditingResponse, err error) {
	response = CreateSubmitLiveEditingResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitLiveEditingWithChan invokes the vod.SubmitLiveEditing API asynchronously
func (client *Client) SubmitLiveEditingWithChan(request *SubmitLiveEditingRequest) (<-chan *SubmitLiveEditingResponse, <-chan error) {
	responseChan := make(chan *SubmitLiveEditingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitLiveEditing(request)
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

// SubmitLiveEditingWithCallback invokes the vod.SubmitLiveEditing API asynchronously
func (client *Client) SubmitLiveEditingWithCallback(request *SubmitLiveEditingRequest, callback func(response *SubmitLiveEditingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitLiveEditingResponse
		var err error
		defer close(result)
		response, err = client.SubmitLiveEditing(request)
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

// SubmitLiveEditingRequest is the request struct for api SubmitLiveEditing
type SubmitLiveEditingRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Clips                string           `position:"Query" name:"Clips"`
	Description          string           `position:"Query" name:"Description"`
	Title                string           `position:"Query" name:"Title"`
	CoverURL             string           `position:"Query" name:"CoverURL"`
	UserData             string           `position:"Query" name:"UserData"`
	AppName              string           `position:"Query" name:"AppName"`
	ProduceConfig        string           `position:"Query" name:"ProduceConfig"`
	StreamName           string           `position:"Query" name:"StreamName"`
	MediaMetadata        string           `position:"Query" name:"MediaMetadata"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DomainName           string           `position:"Query" name:"DomainName"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// SubmitLiveEditingResponse is the response struct for api SubmitLiveEditing
type SubmitLiveEditingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MediaId   string `json:"MediaId" xml:"MediaId"`
	ProjectId string `json:"ProjectId" xml:"ProjectId"`
}

// CreateSubmitLiveEditingRequest creates a request to invoke SubmitLiveEditing API
func CreateSubmitLiveEditingRequest() (request *SubmitLiveEditingRequest) {
	request = &SubmitLiveEditingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitLiveEditing", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitLiveEditingResponse creates a response to parse from SubmitLiveEditing response
func CreateSubmitLiveEditingResponse() (response *SubmitLiveEditingResponse) {
	response = &SubmitLiveEditingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
