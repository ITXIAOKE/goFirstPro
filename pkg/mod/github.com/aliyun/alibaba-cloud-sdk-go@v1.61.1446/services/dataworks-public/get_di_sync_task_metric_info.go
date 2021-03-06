package dataworks_public

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

// GetDISyncTaskMetricInfo invokes the dataworks_public.GetDISyncTaskMetricInfo API synchronously
func (client *Client) GetDISyncTaskMetricInfo(request *GetDISyncTaskMetricInfoRequest) (response *GetDISyncTaskMetricInfoResponse, err error) {
	response = CreateGetDISyncTaskMetricInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetDISyncTaskMetricInfoWithChan invokes the dataworks_public.GetDISyncTaskMetricInfo API asynchronously
func (client *Client) GetDISyncTaskMetricInfoWithChan(request *GetDISyncTaskMetricInfoRequest) (<-chan *GetDISyncTaskMetricInfoResponse, <-chan error) {
	responseChan := make(chan *GetDISyncTaskMetricInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDISyncTaskMetricInfo(request)
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

// GetDISyncTaskMetricInfoWithCallback invokes the dataworks_public.GetDISyncTaskMetricInfo API asynchronously
func (client *Client) GetDISyncTaskMetricInfoWithCallback(request *GetDISyncTaskMetricInfoRequest, callback func(response *GetDISyncTaskMetricInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDISyncTaskMetricInfoResponse
		var err error
		defer close(result)
		response, err = client.GetDISyncTaskMetricInfo(request)
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

// GetDISyncTaskMetricInfoRequest is the request struct for api GetDISyncTaskMetricInfo
type GetDISyncTaskMetricInfoRequest struct {
	*requests.RpcRequest
	EndDate   requests.Integer `position:"Query" name:"EndDate"`
	StartDate requests.Integer `position:"Query" name:"StartDate"`
	ProjectId requests.Integer `position:"Query" name:"ProjectId"`
	FileId    requests.Integer `position:"Query" name:"FileId"`
}

// GetDISyncTaskMetricInfoResponse is the response struct for api GetDISyncTaskMetricInfo
type GetDISyncTaskMetricInfoResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Success    bool       `json:"Success" xml:"Success"`
	MetricInfo MetricInfo `json:"MetricInfo" xml:"MetricInfo"`
}

// CreateGetDISyncTaskMetricInfoRequest creates a request to invoke GetDISyncTaskMetricInfo API
func CreateGetDISyncTaskMetricInfoRequest() (request *GetDISyncTaskMetricInfoRequest) {
	request = &GetDISyncTaskMetricInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetDISyncTaskMetricInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateGetDISyncTaskMetricInfoResponse creates a response to parse from GetDISyncTaskMetricInfo response
func CreateGetDISyncTaskMetricInfoResponse() (response *GetDISyncTaskMetricInfoResponse) {
	response = &GetDISyncTaskMetricInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
