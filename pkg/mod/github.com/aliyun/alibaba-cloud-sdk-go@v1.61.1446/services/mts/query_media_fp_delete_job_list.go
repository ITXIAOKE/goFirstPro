package mts

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

// QueryMediaFpDeleteJobList invokes the mts.QueryMediaFpDeleteJobList API synchronously
func (client *Client) QueryMediaFpDeleteJobList(request *QueryMediaFpDeleteJobListRequest) (response *QueryMediaFpDeleteJobListResponse, err error) {
	response = CreateQueryMediaFpDeleteJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMediaFpDeleteJobListWithChan invokes the mts.QueryMediaFpDeleteJobList API asynchronously
func (client *Client) QueryMediaFpDeleteJobListWithChan(request *QueryMediaFpDeleteJobListRequest) (<-chan *QueryMediaFpDeleteJobListResponse, <-chan error) {
	responseChan := make(chan *QueryMediaFpDeleteJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaFpDeleteJobList(request)
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

// QueryMediaFpDeleteJobListWithCallback invokes the mts.QueryMediaFpDeleteJobList API asynchronously
func (client *Client) QueryMediaFpDeleteJobListWithCallback(request *QueryMediaFpDeleteJobListRequest, callback func(response *QueryMediaFpDeleteJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaFpDeleteJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaFpDeleteJobList(request)
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

// QueryMediaFpDeleteJobListRequest is the request struct for api QueryMediaFpDeleteJobList
type QueryMediaFpDeleteJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
}

// QueryMediaFpDeleteJobListResponse is the response struct for api QueryMediaFpDeleteJobList
type QueryMediaFpDeleteJobListResponse struct {
	*responses.BaseResponse
	RequestId            string                                 `json:"RequestId" xml:"RequestId"`
	NonExistIds          NonExistIdsInQueryMediaFpDeleteJobList `json:"NonExistIds" xml:"NonExistIds"`
	MediaFpDeleteJobList MediaFpDeleteJobList                   `json:"MediaFpDeleteJobList" xml:"MediaFpDeleteJobList"`
}

// CreateQueryMediaFpDeleteJobListRequest creates a request to invoke QueryMediaFpDeleteJobList API
func CreateQueryMediaFpDeleteJobListRequest() (request *QueryMediaFpDeleteJobListRequest) {
	request = &QueryMediaFpDeleteJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaFpDeleteJobList", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryMediaFpDeleteJobListResponse creates a response to parse from QueryMediaFpDeleteJobList response
func CreateQueryMediaFpDeleteJobListResponse() (response *QueryMediaFpDeleteJobListResponse) {
	response = &QueryMediaFpDeleteJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}