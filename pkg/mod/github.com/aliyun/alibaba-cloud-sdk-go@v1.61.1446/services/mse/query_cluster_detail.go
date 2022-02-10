package mse

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

// QueryClusterDetail invokes the mse.QueryClusterDetail API synchronously
func (client *Client) QueryClusterDetail(request *QueryClusterDetailRequest) (response *QueryClusterDetailResponse, err error) {
	response = CreateQueryClusterDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryClusterDetailWithChan invokes the mse.QueryClusterDetail API asynchronously
func (client *Client) QueryClusterDetailWithChan(request *QueryClusterDetailRequest) (<-chan *QueryClusterDetailResponse, <-chan error) {
	responseChan := make(chan *QueryClusterDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryClusterDetail(request)
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

// QueryClusterDetailWithCallback invokes the mse.QueryClusterDetail API asynchronously
func (client *Client) QueryClusterDetailWithCallback(request *QueryClusterDetailRequest, callback func(response *QueryClusterDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryClusterDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryClusterDetail(request)
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

// QueryClusterDetailRequest is the request struct for api QueryClusterDetail
type QueryClusterDetailRequest struct {
	*requests.RpcRequest
	OrderId    string `position:"Query" name:"OrderId"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// QueryClusterDetailResponse is the response struct for api QueryClusterDetail
type QueryClusterDetailResponse struct {
	*responses.BaseResponse
	Message   string                   `json:"Message" xml:"Message"`
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	ErrorCode string                   `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool                     `json:"Success" xml:"Success"`
	Data      DataInQueryClusterDetail `json:"Data" xml:"Data"`
}

// CreateQueryClusterDetailRequest creates a request to invoke QueryClusterDetail API
func CreateQueryClusterDetailRequest() (request *QueryClusterDetailRequest) {
	request = &QueryClusterDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "QueryClusterDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryClusterDetailResponse creates a response to parse from QueryClusterDetail response
func CreateQueryClusterDetailResponse() (response *QueryClusterDetailResponse) {
	response = &QueryClusterDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}