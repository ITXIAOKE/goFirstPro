package gpdb

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

// DescribeSQLLogsOnSlice invokes the gpdb.DescribeSQLLogsOnSlice API synchronously
func (client *Client) DescribeSQLLogsOnSlice(request *DescribeSQLLogsOnSliceRequest) (response *DescribeSQLLogsOnSliceResponse, err error) {
	response = CreateDescribeSQLLogsOnSliceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLLogsOnSliceWithChan invokes the gpdb.DescribeSQLLogsOnSlice API asynchronously
func (client *Client) DescribeSQLLogsOnSliceWithChan(request *DescribeSQLLogsOnSliceRequest) (<-chan *DescribeSQLLogsOnSliceResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLLogsOnSliceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLLogsOnSlice(request)
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

// DescribeSQLLogsOnSliceWithCallback invokes the gpdb.DescribeSQLLogsOnSlice API asynchronously
func (client *Client) DescribeSQLLogsOnSliceWithCallback(request *DescribeSQLLogsOnSliceRequest, callback func(response *DescribeSQLLogsOnSliceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLLogsOnSliceResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLLogsOnSlice(request)
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

// DescribeSQLLogsOnSliceRequest is the request struct for api DescribeSQLLogsOnSlice
type DescribeSQLLogsOnSliceRequest struct {
	*requests.RpcRequest
	SliceId        string           `position:"Query" name:"SliceId"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	MinExecuteCost string           `position:"Query" name:"MinExecuteCost"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId   string           `position:"Query" name:"DBInstanceId"`
	MaxExecuteCost string           `position:"Query" name:"MaxExecuteCost"`
	ExecuteState   string           `position:"Query" name:"ExecuteState"`
	QueryId        string           `position:"Query" name:"QueryId"`
}

// DescribeSQLLogsOnSliceResponse is the response struct for api DescribeSQLLogsOnSlice
type DescribeSQLLogsOnSliceResponse struct {
	*responses.BaseResponse
	RequestId       string             `json:"RequestId" xml:"RequestId"`
	PageRecordCount int                `json:"PageRecordCount" xml:"PageRecordCount"`
	PageNumber      int                `json:"PageNumber" xml:"PageNumber"`
	SliceLogItems   []SQLLogsSliceItem `json:"SliceLogItems" xml:"SliceLogItems"`
}

// CreateDescribeSQLLogsOnSliceRequest creates a request to invoke DescribeSQLLogsOnSlice API
func CreateDescribeSQLLogsOnSliceRequest() (request *DescribeSQLLogsOnSliceRequest) {
	request = &DescribeSQLLogsOnSliceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeSQLLogsOnSlice", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSQLLogsOnSliceResponse creates a response to parse from DescribeSQLLogsOnSlice response
func CreateDescribeSQLLogsOnSliceResponse() (response *DescribeSQLLogsOnSliceResponse) {
	response = &DescribeSQLLogsOnSliceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
