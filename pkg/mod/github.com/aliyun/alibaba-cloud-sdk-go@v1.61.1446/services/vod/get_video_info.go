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

// GetVideoInfo invokes the vod.GetVideoInfo API synchronously
func (client *Client) GetVideoInfo(request *GetVideoInfoRequest) (response *GetVideoInfoResponse, err error) {
	response = CreateGetVideoInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetVideoInfoWithChan invokes the vod.GetVideoInfo API asynchronously
func (client *Client) GetVideoInfoWithChan(request *GetVideoInfoRequest) (<-chan *GetVideoInfoResponse, <-chan error) {
	responseChan := make(chan *GetVideoInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVideoInfo(request)
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

// GetVideoInfoWithCallback invokes the vod.GetVideoInfo API asynchronously
func (client *Client) GetVideoInfoWithCallback(request *GetVideoInfoRequest, callback func(response *GetVideoInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVideoInfoResponse
		var err error
		defer close(result)
		response, err = client.GetVideoInfo(request)
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

// GetVideoInfoRequest is the request struct for api GetVideoInfo
type GetVideoInfoRequest struct {
	*requests.RpcRequest
	VideoId      string `position:"Query" name:"VideoId"`
	AdditionType string `position:"Query" name:"AdditionType"`
	ResultTypes  string `position:"Query" name:"ResultTypes"`
}

// GetVideoInfoResponse is the response struct for api GetVideoInfo
type GetVideoInfoResponse struct {
	*responses.BaseResponse
	RequestId string              `json:"RequestId" xml:"RequestId"`
	AI        string              `json:"AI" xml:"AI"`
	Video     VideoInGetVideoInfo `json:"Video" xml:"Video"`
}

// CreateGetVideoInfoRequest creates a request to invoke GetVideoInfo API
func CreateGetVideoInfoRequest() (request *GetVideoInfoRequest) {
	request = &GetVideoInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetVideoInfo", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVideoInfoResponse creates a response to parse from GetVideoInfo response
func CreateGetVideoInfoResponse() (response *GetVideoInfoResponse) {
	response = &GetVideoInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
