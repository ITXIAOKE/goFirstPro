package live

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

// StopPlaylist invokes the live.StopPlaylist API synchronously
func (client *Client) StopPlaylist(request *StopPlaylistRequest) (response *StopPlaylistResponse, err error) {
	response = CreateStopPlaylistResponse()
	err = client.DoAction(request, response)
	return
}

// StopPlaylistWithChan invokes the live.StopPlaylist API asynchronously
func (client *Client) StopPlaylistWithChan(request *StopPlaylistRequest) (<-chan *StopPlaylistResponse, <-chan error) {
	responseChan := make(chan *StopPlaylistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopPlaylist(request)
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

// StopPlaylistWithCallback invokes the live.StopPlaylist API asynchronously
func (client *Client) StopPlaylistWithCallback(request *StopPlaylistRequest, callback func(response *StopPlaylistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopPlaylistResponse
		var err error
		defer close(result)
		response, err = client.StopPlaylist(request)
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

// StopPlaylistRequest is the request struct for api StopPlaylist
type StopPlaylistRequest struct {
	*requests.RpcRequest
	ProgramId string           `position:"Query" name:"ProgramId"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// StopPlaylistResponse is the response struct for api StopPlaylist
type StopPlaylistResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ProgramId string `json:"ProgramId" xml:"ProgramId"`
}

// CreateStopPlaylistRequest creates a request to invoke StopPlaylist API
func CreateStopPlaylistRequest() (request *StopPlaylistRequest) {
	request = &StopPlaylistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "StopPlaylist", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopPlaylistResponse creates a response to parse from StopPlaylist response
func CreateStopPlaylistResponse() (response *StopPlaylistResponse) {
	response = &StopPlaylistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
