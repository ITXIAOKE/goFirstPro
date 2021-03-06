package dts

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

// ResetDtsJob invokes the dts.ResetDtsJob API synchronously
func (client *Client) ResetDtsJob(request *ResetDtsJobRequest) (response *ResetDtsJobResponse, err error) {
	response = CreateResetDtsJobResponse()
	err = client.DoAction(request, response)
	return
}

// ResetDtsJobWithChan invokes the dts.ResetDtsJob API asynchronously
func (client *Client) ResetDtsJobWithChan(request *ResetDtsJobRequest) (<-chan *ResetDtsJobResponse, <-chan error) {
	responseChan := make(chan *ResetDtsJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetDtsJob(request)
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

// ResetDtsJobWithCallback invokes the dts.ResetDtsJob API asynchronously
func (client *Client) ResetDtsJobWithCallback(request *ResetDtsJobRequest, callback func(response *ResetDtsJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetDtsJobResponse
		var err error
		defer close(result)
		response, err = client.ResetDtsJob(request)
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

// ResetDtsJobRequest is the request struct for api ResetDtsJob
type ResetDtsJobRequest struct {
	*requests.RpcRequest
	DtsJobId                 string `position:"Query" name:"DtsJobId"`
	DtsInstanceId            string `position:"Query" name:"DtsInstanceId"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
}

// ResetDtsJobResponse is the response struct for api ResetDtsJob
type ResetDtsJobResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
}

// CreateResetDtsJobRequest creates a request to invoke ResetDtsJob API
func CreateResetDtsJobRequest() (request *ResetDtsJobRequest) {
	request = &ResetDtsJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ResetDtsJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetDtsJobResponse creates a response to parse from ResetDtsJob response
func CreateResetDtsJobResponse() (response *ResetDtsJobResponse) {
	response = &ResetDtsJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
