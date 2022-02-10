package airec

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

// DescribeUserMetrics invokes the airec.DescribeUserMetrics API synchronously
func (client *Client) DescribeUserMetrics(request *DescribeUserMetricsRequest) (response *DescribeUserMetricsResponse, err error) {
	response = CreateDescribeUserMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserMetricsWithChan invokes the airec.DescribeUserMetrics API asynchronously
func (client *Client) DescribeUserMetricsWithChan(request *DescribeUserMetricsRequest) (<-chan *DescribeUserMetricsResponse, <-chan error) {
	responseChan := make(chan *DescribeUserMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserMetrics(request)
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

// DescribeUserMetricsWithCallback invokes the airec.DescribeUserMetrics API asynchronously
func (client *Client) DescribeUserMetricsWithCallback(request *DescribeUserMetricsRequest, callback func(response *DescribeUserMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserMetricsResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserMetrics(request)
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

// DescribeUserMetricsRequest is the request struct for api DescribeUserMetrics
type DescribeUserMetricsRequest struct {
	*requests.RoaRequest
	MetricType string           `position:"Query" name:"metricType"`
	InstanceId string           `position:"Path" name:"instanceId"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	StartTime  requests.Integer `position:"Query" name:"startTime"`
}

// DescribeUserMetricsResponse is the response struct for api DescribeUserMetrics
type DescribeUserMetricsResponse struct {
	*responses.BaseResponse
	Code      string   `json:"code" xml:"code"`
	Message   string   `json:"message" xml:"message"`
	RequestId string   `json:"requestId" xml:"requestId"`
	Result    []Result `json:"result" xml:"result"`
}

// CreateDescribeUserMetricsRequest creates a request to invoke DescribeUserMetrics API
func CreateDescribeUserMetricsRequest() (request *DescribeUserMetricsRequest) {
	request = &DescribeUserMetricsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "DescribeUserMetrics", "/v2/openapi/instances/[instanceId]/metrics", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeUserMetricsResponse creates a response to parse from DescribeUserMetrics response
func CreateDescribeUserMetricsResponse() (response *DescribeUserMetricsResponse) {
	response = &DescribeUserMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}