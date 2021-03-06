package vpc

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

// ListNatGatewayEcsMetric invokes the vpc.ListNatGatewayEcsMetric API synchronously
func (client *Client) ListNatGatewayEcsMetric(request *ListNatGatewayEcsMetricRequest) (response *ListNatGatewayEcsMetricResponse, err error) {
	response = CreateListNatGatewayEcsMetricResponse()
	err = client.DoAction(request, response)
	return
}

// ListNatGatewayEcsMetricWithChan invokes the vpc.ListNatGatewayEcsMetric API asynchronously
func (client *Client) ListNatGatewayEcsMetricWithChan(request *ListNatGatewayEcsMetricRequest) (<-chan *ListNatGatewayEcsMetricResponse, <-chan error) {
	responseChan := make(chan *ListNatGatewayEcsMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNatGatewayEcsMetric(request)
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

// ListNatGatewayEcsMetricWithCallback invokes the vpc.ListNatGatewayEcsMetric API asynchronously
func (client *Client) ListNatGatewayEcsMetricWithCallback(request *ListNatGatewayEcsMetricRequest, callback func(response *ListNatGatewayEcsMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNatGatewayEcsMetricResponse
		var err error
		defer close(result)
		response, err = client.ListNatGatewayEcsMetric(request)
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

// ListNatGatewayEcsMetricRequest is the request struct for api ListNatGatewayEcsMetric
type ListNatGatewayEcsMetricRequest struct {
	*requests.RpcRequest
	OrderKey         string           `position:"Query" name:"OrderKey"`
	NextToken        string           `position:"Query" name:"NextToken"`
	NatGatewayId     string           `position:"Query" name:"NatGatewayId"`
	DryRun           requests.Boolean `position:"Query" name:"DryRun"`
	PrivateIpAddress string           `position:"Query" name:"PrivateIpAddress"`
	MaxResults       string           `position:"Query" name:"MaxResults"`
	TimePoint        requests.Integer `position:"Query" name:"TimePoint"`
	OrderType        string           `position:"Query" name:"OrderType"`
}

// ListNatGatewayEcsMetricResponse is the response struct for api ListNatGatewayEcsMetric
type ListNatGatewayEcsMetricResponse struct {
	*responses.BaseResponse
	NextToken      string       `json:"NextToken" xml:"NextToken"`
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	MaxResults     int          `json:"MaxResults" xml:"MaxResults"`
	MetricDataList []MetricData `json:"MetricDataList" xml:"MetricDataList"`
}

// CreateListNatGatewayEcsMetricRequest creates a request to invoke ListNatGatewayEcsMetric API
func CreateListNatGatewayEcsMetricRequest() (request *ListNatGatewayEcsMetricRequest) {
	request = &ListNatGatewayEcsMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ListNatGatewayEcsMetric", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNatGatewayEcsMetricResponse creates a response to parse from ListNatGatewayEcsMetric response
func CreateListNatGatewayEcsMetricResponse() (response *ListNatGatewayEcsMetricResponse) {
	response = &ListNatGatewayEcsMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
