package ehpc

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

// ListCloudMetricProfilings invokes the ehpc.ListCloudMetricProfilings API synchronously
func (client *Client) ListCloudMetricProfilings(request *ListCloudMetricProfilingsRequest) (response *ListCloudMetricProfilingsResponse, err error) {
	response = CreateListCloudMetricProfilingsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCloudMetricProfilingsWithChan invokes the ehpc.ListCloudMetricProfilings API asynchronously
func (client *Client) ListCloudMetricProfilingsWithChan(request *ListCloudMetricProfilingsRequest) (<-chan *ListCloudMetricProfilingsResponse, <-chan error) {
	responseChan := make(chan *ListCloudMetricProfilingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCloudMetricProfilings(request)
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

// ListCloudMetricProfilingsWithCallback invokes the ehpc.ListCloudMetricProfilings API asynchronously
func (client *Client) ListCloudMetricProfilingsWithCallback(request *ListCloudMetricProfilingsRequest, callback func(response *ListCloudMetricProfilingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCloudMetricProfilingsResponse
		var err error
		defer close(result)
		response, err = client.ListCloudMetricProfilings(request)
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

// ListCloudMetricProfilingsRequest is the request struct for api ListCloudMetricProfilings
type ListCloudMetricProfilingsRequest struct {
	*requests.RpcRequest
	ClusterId  string           `position:"Query" name:"ClusterId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListCloudMetricProfilingsResponse is the response struct for api ListCloudMetricProfilings
type ListCloudMetricProfilingsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	Profilings Profilings `json:"Profilings" xml:"Profilings"`
}

// CreateListCloudMetricProfilingsRequest creates a request to invoke ListCloudMetricProfilings API
func CreateListCloudMetricProfilingsRequest() (request *ListCloudMetricProfilingsRequest) {
	request = &ListCloudMetricProfilingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListCloudMetricProfilings", "", "")
	request.Method = requests.GET
	return
}

// CreateListCloudMetricProfilingsResponse creates a response to parse from ListCloudMetricProfilings response
func CreateListCloudMetricProfilingsResponse() (response *ListCloudMetricProfilingsResponse) {
	response = &ListCloudMetricProfilingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}