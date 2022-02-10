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

// GetNextResultToVerify invokes the qualitycheck.GetNextResultToVerify API synchronously
func (client *Client) GetNextResultToVerify(request *GetNextResultToVerifyRequest) (response *GetNextResultToVerifyResponse, err error) {
	response = CreateGetNextResultToVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// GetNextResultToVerifyWithChan invokes the qualitycheck.GetNextResultToVerify API asynchronously
func (client *Client) GetNextResultToVerifyWithChan(request *GetNextResultToVerifyRequest) (<-chan *GetNextResultToVerifyResponse, <-chan error) {
	responseChan := make(chan *GetNextResultToVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNextResultToVerify(request)
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

// GetNextResultToVerifyWithCallback invokes the qualitycheck.GetNextResultToVerify API asynchronously
func (client *Client) GetNextResultToVerifyWithCallback(request *GetNextResultToVerifyRequest, callback func(response *GetNextResultToVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNextResultToVerifyResponse
		var err error
		defer close(result)
		response, err = client.GetNextResultToVerify(request)
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

// GetNextResultToVerifyRequest is the request struct for api GetNextResultToVerify
type GetNextResultToVerifyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetNextResultToVerifyResponse is the response struct for api GetNextResultToVerify
type GetNextResultToVerifyResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetNextResultToVerifyRequest creates a request to invoke GetNextResultToVerify API
func CreateGetNextResultToVerifyRequest() (request *GetNextResultToVerifyRequest) {
	request = &GetNextResultToVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetNextResultToVerify", "", "")
	request.Method = requests.POST
	return
}

// CreateGetNextResultToVerifyResponse creates a response to parse from GetNextResultToVerify response
func CreateGetNextResultToVerifyResponse() (response *GetNextResultToVerifyResponse) {
	response = &GetNextResultToVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
