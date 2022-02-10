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

// SubmitPrecisionTask invokes the qualitycheck.SubmitPrecisionTask API synchronously
func (client *Client) SubmitPrecisionTask(request *SubmitPrecisionTaskRequest) (response *SubmitPrecisionTaskResponse, err error) {
	response = CreateSubmitPrecisionTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitPrecisionTaskWithChan invokes the qualitycheck.SubmitPrecisionTask API asynchronously
func (client *Client) SubmitPrecisionTaskWithChan(request *SubmitPrecisionTaskRequest) (<-chan *SubmitPrecisionTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitPrecisionTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitPrecisionTask(request)
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

// SubmitPrecisionTaskWithCallback invokes the qualitycheck.SubmitPrecisionTask API asynchronously
func (client *Client) SubmitPrecisionTaskWithCallback(request *SubmitPrecisionTaskRequest, callback func(response *SubmitPrecisionTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitPrecisionTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitPrecisionTask(request)
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

// SubmitPrecisionTaskRequest is the request struct for api SubmitPrecisionTask
type SubmitPrecisionTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// SubmitPrecisionTaskResponse is the response struct for api SubmitPrecisionTask
type SubmitPrecisionTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSubmitPrecisionTaskRequest creates a request to invoke SubmitPrecisionTask API
func CreateSubmitPrecisionTaskRequest() (request *SubmitPrecisionTaskRequest) {
	request = &SubmitPrecisionTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "SubmitPrecisionTask", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitPrecisionTaskResponse creates a response to parse from SubmitPrecisionTask response
func CreateSubmitPrecisionTaskResponse() (response *SubmitPrecisionTaskResponse) {
	response = &SubmitPrecisionTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
