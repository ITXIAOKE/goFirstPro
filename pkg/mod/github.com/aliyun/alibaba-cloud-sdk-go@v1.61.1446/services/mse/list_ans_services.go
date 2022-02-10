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

// ListAnsServices invokes the mse.ListAnsServices API synchronously
func (client *Client) ListAnsServices(request *ListAnsServicesRequest) (response *ListAnsServicesResponse, err error) {
	response = CreateListAnsServicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAnsServicesWithChan invokes the mse.ListAnsServices API asynchronously
func (client *Client) ListAnsServicesWithChan(request *ListAnsServicesRequest) (<-chan *ListAnsServicesResponse, <-chan error) {
	responseChan := make(chan *ListAnsServicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAnsServices(request)
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

// ListAnsServicesWithCallback invokes the mse.ListAnsServices API asynchronously
func (client *Client) ListAnsServicesWithCallback(request *ListAnsServicesRequest, callback func(response *ListAnsServicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAnsServicesResponse
		var err error
		defer close(result)
		response, err = client.ListAnsServices(request)
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

// ListAnsServicesRequest is the request struct for api ListAnsServices
type ListAnsServicesRequest struct {
	*requests.RpcRequest
	ClusterId   string           `position:"Query" name:"ClusterId"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	GroupName   string           `position:"Query" name:"GroupName"`
	HasIpCount  string           `position:"Query" name:"HasIpCount"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	NamespaceId string           `position:"Query" name:"NamespaceId"`
	RequestPars string           `position:"Query" name:"RequestPars"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	ServiceName string           `position:"Query" name:"ServiceName"`
}

// ListAnsServicesResponse is the response struct for api ListAnsServices
type ListAnsServicesResponse struct {
	*responses.BaseResponse
	HttpCode   string                  `json:"HttpCode" xml:"HttpCode"`
	TotalCount int                     `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                  `json:"RequestId" xml:"RequestId"`
	Message    string                  `json:"Message" xml:"Message"`
	PageSize   int                     `json:"PageSize" xml:"PageSize"`
	PageNumber int                     `json:"PageNumber" xml:"PageNumber"`
	ErrorCode  string                  `json:"ErrorCode" xml:"ErrorCode"`
	Success    bool                    `json:"Success" xml:"Success"`
	Data       []SimpleNacosAnsService `json:"Data" xml:"Data"`
}

// CreateListAnsServicesRequest creates a request to invoke ListAnsServices API
func CreateListAnsServicesRequest() (request *ListAnsServicesRequest) {
	request = &ListAnsServicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListAnsServices", "", "")
	request.Method = requests.GET
	return
}

// CreateListAnsServicesResponse creates a response to parse from ListAnsServices response
func CreateListAnsServicesResponse() (response *ListAnsServicesResponse) {
	response = &ListAnsServicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
