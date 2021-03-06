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

// DescribeSQLLogByQueryId invokes the gpdb.DescribeSQLLogByQueryId API synchronously
func (client *Client) DescribeSQLLogByQueryId(request *DescribeSQLLogByQueryIdRequest) (response *DescribeSQLLogByQueryIdResponse, err error) {
	response = CreateDescribeSQLLogByQueryIdResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLLogByQueryIdWithChan invokes the gpdb.DescribeSQLLogByQueryId API asynchronously
func (client *Client) DescribeSQLLogByQueryIdWithChan(request *DescribeSQLLogByQueryIdRequest) (<-chan *DescribeSQLLogByQueryIdResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLLogByQueryIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLLogByQueryId(request)
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

// DescribeSQLLogByQueryIdWithCallback invokes the gpdb.DescribeSQLLogByQueryId API asynchronously
func (client *Client) DescribeSQLLogByQueryIdWithCallback(request *DescribeSQLLogByQueryIdRequest, callback func(response *DescribeSQLLogByQueryIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLLogByQueryIdResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLLogByQueryId(request)
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

// DescribeSQLLogByQueryIdRequest is the request struct for api DescribeSQLLogByQueryId
type DescribeSQLLogByQueryIdRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	QueryId      string `position:"Query" name:"QueryId"`
}

// DescribeSQLLogByQueryIdResponse is the response struct for api DescribeSQLLogByQueryId
type DescribeSQLLogByQueryIdResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Items     []SQLLog `json:"Items" xml:"Items"`
}

// CreateDescribeSQLLogByQueryIdRequest creates a request to invoke DescribeSQLLogByQueryId API
func CreateDescribeSQLLogByQueryIdRequest() (request *DescribeSQLLogByQueryIdRequest) {
	request = &DescribeSQLLogByQueryIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeSQLLogByQueryId", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSQLLogByQueryIdResponse creates a response to parse from DescribeSQLLogByQueryId response
func CreateDescribeSQLLogByQueryIdResponse() (response *DescribeSQLLogByQueryIdResponse) {
	response = &DescribeSQLLogByQueryIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
