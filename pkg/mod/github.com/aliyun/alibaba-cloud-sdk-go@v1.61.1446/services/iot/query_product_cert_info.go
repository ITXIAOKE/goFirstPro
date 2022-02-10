package iot

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

// QueryProductCertInfo invokes the iot.QueryProductCertInfo API synchronously
func (client *Client) QueryProductCertInfo(request *QueryProductCertInfoRequest) (response *QueryProductCertInfoResponse, err error) {
	response = CreateQueryProductCertInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryProductCertInfoWithChan invokes the iot.QueryProductCertInfo API asynchronously
func (client *Client) QueryProductCertInfoWithChan(request *QueryProductCertInfoRequest) (<-chan *QueryProductCertInfoResponse, <-chan error) {
	responseChan := make(chan *QueryProductCertInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryProductCertInfo(request)
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

// QueryProductCertInfoWithCallback invokes the iot.QueryProductCertInfo API asynchronously
func (client *Client) QueryProductCertInfoWithCallback(request *QueryProductCertInfoRequest, callback func(response *QueryProductCertInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryProductCertInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryProductCertInfo(request)
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

// QueryProductCertInfoRequest is the request struct for api QueryProductCertInfo
type QueryProductCertInfoRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// QueryProductCertInfoResponse is the response struct for api QueryProductCertInfo
type QueryProductCertInfoResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Success         bool            `json:"Success" xml:"Success"`
	Code            string          `json:"Code" xml:"Code"`
	ErrorMessage    string          `json:"ErrorMessage" xml:"ErrorMessage"`
	ProductCertInfo ProductCertInfo `json:"ProductCertInfo" xml:"ProductCertInfo"`
}

// CreateQueryProductCertInfoRequest creates a request to invoke QueryProductCertInfo API
func CreateQueryProductCertInfoRequest() (request *QueryProductCertInfoRequest) {
	request = &QueryProductCertInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryProductCertInfo", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryProductCertInfoResponse creates a response to parse from QueryProductCertInfo response
func CreateQueryProductCertInfoResponse() (response *QueryProductCertInfoResponse) {
	response = &QueryProductCertInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
