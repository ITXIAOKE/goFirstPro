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

// DescribeSynchronizationJobStatusList invokes the dts.DescribeSynchronizationJobStatusList API synchronously
func (client *Client) DescribeSynchronizationJobStatusList(request *DescribeSynchronizationJobStatusListRequest) (response *DescribeSynchronizationJobStatusListResponse, err error) {
	response = CreateDescribeSynchronizationJobStatusListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSynchronizationJobStatusListWithChan invokes the dts.DescribeSynchronizationJobStatusList API asynchronously
func (client *Client) DescribeSynchronizationJobStatusListWithChan(request *DescribeSynchronizationJobStatusListRequest) (<-chan *DescribeSynchronizationJobStatusListResponse, <-chan error) {
	responseChan := make(chan *DescribeSynchronizationJobStatusListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSynchronizationJobStatusList(request)
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

// DescribeSynchronizationJobStatusListWithCallback invokes the dts.DescribeSynchronizationJobStatusList API asynchronously
func (client *Client) DescribeSynchronizationJobStatusListWithCallback(request *DescribeSynchronizationJobStatusListRequest, callback func(response *DescribeSynchronizationJobStatusListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSynchronizationJobStatusListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSynchronizationJobStatusList(request)
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

// DescribeSynchronizationJobStatusListRequest is the request struct for api DescribeSynchronizationJobStatusList
type DescribeSynchronizationJobStatusListRequest struct {
	*requests.RpcRequest
	ClientToken                     string `position:"Query" name:"ClientToken"`
	OwnerId                         string `position:"Query" name:"OwnerId"`
	SynchronizationJobIdListJsonStr string `position:"Query" name:"SynchronizationJobIdListJsonStr"`
	AccountId                       string `position:"Query" name:"AccountId"`
}

// DescribeSynchronizationJobStatusListResponse is the response struct for api DescribeSynchronizationJobStatusList
type DescribeSynchronizationJobStatusListResponse struct {
	*responses.BaseResponse
	RequestId                        string                         `json:"RequestId" xml:"RequestId"`
	ErrCode                          string                         `json:"ErrCode" xml:"ErrCode"`
	PageRecordCount                  int                            `json:"PageRecordCount" xml:"PageRecordCount"`
	Success                          string                         `json:"Success" xml:"Success"`
	TotalRecordCount                 int64                          `json:"TotalRecordCount" xml:"TotalRecordCount"`
	ErrMessage                       string                         `json:"ErrMessage" xml:"ErrMessage"`
	PageNumber                       int                            `json:"PageNumber" xml:"PageNumber"`
	SynchronizationJobListStatusList []SynchronizationJobStatusInfo `json:"SynchronizationJobListStatusList" xml:"SynchronizationJobListStatusList"`
}

// CreateDescribeSynchronizationJobStatusListRequest creates a request to invoke DescribeSynchronizationJobStatusList API
func CreateDescribeSynchronizationJobStatusListRequest() (request *DescribeSynchronizationJobStatusListRequest) {
	request = &DescribeSynchronizationJobStatusListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "DescribeSynchronizationJobStatusList", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSynchronizationJobStatusListResponse creates a response to parse from DescribeSynchronizationJobStatusList response
func CreateDescribeSynchronizationJobStatusListResponse() (response *DescribeSynchronizationJobStatusListResponse) {
	response = &DescribeSynchronizationJobStatusListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}