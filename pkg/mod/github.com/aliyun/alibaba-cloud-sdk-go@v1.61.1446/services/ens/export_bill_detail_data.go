package ens

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

// ExportBillDetailData invokes the ens.ExportBillDetailData API synchronously
func (client *Client) ExportBillDetailData(request *ExportBillDetailDataRequest) (response *ExportBillDetailDataResponse, err error) {
	response = CreateExportBillDetailDataResponse()
	err = client.DoAction(request, response)
	return
}

// ExportBillDetailDataWithChan invokes the ens.ExportBillDetailData API asynchronously
func (client *Client) ExportBillDetailDataWithChan(request *ExportBillDetailDataRequest) (<-chan *ExportBillDetailDataResponse, <-chan error) {
	responseChan := make(chan *ExportBillDetailDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportBillDetailData(request)
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

// ExportBillDetailDataWithCallback invokes the ens.ExportBillDetailData API asynchronously
func (client *Client) ExportBillDetailDataWithCallback(request *ExportBillDetailDataRequest, callback func(response *ExportBillDetailDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportBillDetailDataResponse
		var err error
		defer close(result)
		response, err = client.ExportBillDetailData(request)
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

// ExportBillDetailDataRequest is the request struct for api ExportBillDetailData
type ExportBillDetailDataRequest struct {
	*requests.RpcRequest
	StartDate string `position:"Query" name:"StartDate"`
	EndDate   string `position:"Query" name:"EndDate"`
}

// ExportBillDetailDataResponse is the response struct for api ExportBillDetailData
type ExportBillDetailDataResponse struct {
	*responses.BaseResponse
	FilePath  string `json:"FilePath" xml:"FilePath"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateExportBillDetailDataRequest creates a request to invoke ExportBillDetailData API
func CreateExportBillDetailDataRequest() (request *ExportBillDetailDataRequest) {
	request = &ExportBillDetailDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ExportBillDetailData", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportBillDetailDataResponse creates a response to parse from ExportBillDetailData response
func CreateExportBillDetailDataResponse() (response *ExportBillDetailDataResponse) {
	response = &ExportBillDetailDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
