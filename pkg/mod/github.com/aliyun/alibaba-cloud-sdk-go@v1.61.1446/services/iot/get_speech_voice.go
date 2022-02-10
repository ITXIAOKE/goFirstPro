package iot

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

// GetSpeechVoice invokes the iot.GetSpeechVoice API synchronously
func (client *Client) GetSpeechVoice(request *GetSpeechVoiceRequest) (response *GetSpeechVoiceResponse, err error) {
	response = CreateGetSpeechVoiceResponse()
	err = client.DoAction(request, response)
	return
}

// GetSpeechVoiceWithChan invokes the iot.GetSpeechVoice API asynchronously
func (client *Client) GetSpeechVoiceWithChan(request *GetSpeechVoiceRequest) (<-chan *GetSpeechVoiceResponse, <-chan error) {
	responseChan := make(chan *GetSpeechVoiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSpeechVoice(request)
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

// GetSpeechVoiceWithCallback invokes the iot.GetSpeechVoice API asynchronously
func (client *Client) GetSpeechVoiceWithCallback(request *GetSpeechVoiceRequest, callback func(response *GetSpeechVoiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSpeechVoiceResponse
		var err error
		defer close(result)
		response, err = client.GetSpeechVoice(request)
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

// GetSpeechVoiceRequest is the request struct for api GetSpeechVoice
type GetSpeechVoiceRequest struct {
	*requests.RpcRequest
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// GetSpeechVoiceResponse is the response struct for api GetSpeechVoice
type GetSpeechVoiceResponse struct {
	*responses.BaseResponse
	RequestId    string               `json:"RequestId" xml:"RequestId"`
	Success      bool                 `json:"Success" xml:"Success"`
	Code         string               `json:"Code" xml:"Code"`
	ErrorMessage string               `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInGetSpeechVoice `json:"Data" xml:"Data"`
}

// CreateGetSpeechVoiceRequest creates a request to invoke GetSpeechVoice API
func CreateGetSpeechVoiceRequest() (request *GetSpeechVoiceRequest) {
	request = &GetSpeechVoiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetSpeechVoice", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSpeechVoiceResponse creates a response to parse from GetSpeechVoice response
func CreateGetSpeechVoiceResponse() (response *GetSpeechVoiceResponse) {
	response = &GetSpeechVoiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
