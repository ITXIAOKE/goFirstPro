package qualitycheck

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

// GetAsrVocab invokes the qualitycheck.GetAsrVocab API synchronously
func (client *Client) GetAsrVocab(request *GetAsrVocabRequest) (response *GetAsrVocabResponse, err error) {
	response = CreateGetAsrVocabResponse()
	err = client.DoAction(request, response)
	return
}

// GetAsrVocabWithChan invokes the qualitycheck.GetAsrVocab API asynchronously
func (client *Client) GetAsrVocabWithChan(request *GetAsrVocabRequest) (<-chan *GetAsrVocabResponse, <-chan error) {
	responseChan := make(chan *GetAsrVocabResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAsrVocab(request)
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

// GetAsrVocabWithCallback invokes the qualitycheck.GetAsrVocab API asynchronously
func (client *Client) GetAsrVocabWithCallback(request *GetAsrVocabRequest, callback func(response *GetAsrVocabResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAsrVocabResponse
		var err error
		defer close(result)
		response, err = client.GetAsrVocab(request)
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

// GetAsrVocabRequest is the request struct for api GetAsrVocab
type GetAsrVocabRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetAsrVocabResponse is the response struct for api GetAsrVocab
type GetAsrVocabResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAsrVocabRequest creates a request to invoke GetAsrVocab API
func CreateGetAsrVocabRequest() (request *GetAsrVocabRequest) {
	request = &GetAsrVocabRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetAsrVocab", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAsrVocabResponse creates a response to parse from GetAsrVocab response
func CreateGetAsrVocabResponse() (response *GetAsrVocabResponse) {
	response = &GetAsrVocabResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}