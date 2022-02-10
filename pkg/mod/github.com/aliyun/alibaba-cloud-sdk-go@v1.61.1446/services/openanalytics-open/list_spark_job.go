package openanalytics_open

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

// ListSparkJob invokes the openanalytics_open.ListSparkJob API synchronously
func (client *Client) ListSparkJob(request *ListSparkJobRequest) (response *ListSparkJobResponse, err error) {
	response = CreateListSparkJobResponse()
	err = client.DoAction(request, response)
	return
}

// ListSparkJobWithChan invokes the openanalytics_open.ListSparkJob API asynchronously
func (client *Client) ListSparkJobWithChan(request *ListSparkJobRequest) (<-chan *ListSparkJobResponse, <-chan error) {
	responseChan := make(chan *ListSparkJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSparkJob(request)
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

// ListSparkJobWithCallback invokes the openanalytics_open.ListSparkJob API asynchronously
func (client *Client) ListSparkJobWithCallback(request *ListSparkJobRequest, callback func(response *ListSparkJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSparkJobResponse
		var err error
		defer close(result)
		response, err = client.ListSparkJob(request)
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

// ListSparkJobRequest is the request struct for api ListSparkJob
type ListSparkJobRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Condition  string           `position:"Query" name:"Condition"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	VcName     string           `position:"Query" name:"VcName"`
}

// ListSparkJobResponse is the response struct for api ListSparkJob
type ListSparkJobResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	DataResult DataResult `json:"DataResult" xml:"DataResult"`
}

// CreateListSparkJobRequest creates a request to invoke ListSparkJob API
func CreateListSparkJobRequest() (request *ListSparkJobRequest) {
	request = &ListSparkJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "ListSparkJob", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSparkJobResponse creates a response to parse from ListSparkJob response
func CreateListSparkJobResponse() (response *ListSparkJobResponse) {
	response = &ListSparkJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}