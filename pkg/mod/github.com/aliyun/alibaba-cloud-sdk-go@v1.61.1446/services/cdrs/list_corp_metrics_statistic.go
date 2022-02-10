package cdrs

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

// ListCorpMetricsStatistic invokes the cdrs.ListCorpMetricsStatistic API synchronously
func (client *Client) ListCorpMetricsStatistic(request *ListCorpMetricsStatisticRequest) (response *ListCorpMetricsStatisticResponse, err error) {
	response = CreateListCorpMetricsStatisticResponse()
	err = client.DoAction(request, response)
	return
}

// ListCorpMetricsStatisticWithChan invokes the cdrs.ListCorpMetricsStatistic API asynchronously
func (client *Client) ListCorpMetricsStatisticWithChan(request *ListCorpMetricsStatisticRequest) (<-chan *ListCorpMetricsStatisticResponse, <-chan error) {
	responseChan := make(chan *ListCorpMetricsStatisticResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCorpMetricsStatistic(request)
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

// ListCorpMetricsStatisticWithCallback invokes the cdrs.ListCorpMetricsStatistic API asynchronously
func (client *Client) ListCorpMetricsStatisticWithCallback(request *ListCorpMetricsStatisticRequest, callback func(response *ListCorpMetricsStatisticResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCorpMetricsStatisticResponse
		var err error
		defer close(result)
		response, err = client.ListCorpMetricsStatistic(request)
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

// ListCorpMetricsStatisticRequest is the request struct for api ListCorpMetricsStatistic
type ListCorpMetricsStatisticRequest struct {
	*requests.RpcRequest
	Schema          string           `position:"Body" name:"Schema"`
	CorpId          string           `position:"Body" name:"CorpId"`
	EndTime         string           `position:"Body" name:"EndTime"`
	StartTime       string           `position:"Body" name:"StartTime"`
	PageNumber      requests.Integer `position:"Body" name:"PageNumber"`
	DeviceGroupList string           `position:"Body" name:"DeviceGroupList"`
	TagCode         string           `position:"Body" name:"TagCode"`
	UserGroupList   string           `position:"Body" name:"UserGroupList"`
	PageSize        requests.Integer `position:"Body" name:"PageSize"`
	DeviceIdList    string           `position:"Body" name:"DeviceIdList"`
}

// ListCorpMetricsStatisticResponse is the response struct for api ListCorpMetricsStatistic
type ListCorpMetricsStatisticResponse struct {
	*responses.BaseResponse
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	Message    string     `json:"Message" xml:"Message"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	Code       string     `json:"Code" xml:"Code"`
	Success    string     `json:"Success" xml:"Success"`
	Data       []DataItem `json:"Data" xml:"Data"`
}

// CreateListCorpMetricsStatisticRequest creates a request to invoke ListCorpMetricsStatistic API
func CreateListCorpMetricsStatisticRequest() (request *ListCorpMetricsStatisticRequest) {
	request = &ListCorpMetricsStatisticRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListCorpMetricsStatistic", "", "")
	request.Method = requests.POST
	return
}

// CreateListCorpMetricsStatisticResponse creates a response to parse from ListCorpMetricsStatistic response
func CreateListCorpMetricsStatisticResponse() (response *ListCorpMetricsStatisticResponse) {
	response = &ListCorpMetricsStatisticResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
