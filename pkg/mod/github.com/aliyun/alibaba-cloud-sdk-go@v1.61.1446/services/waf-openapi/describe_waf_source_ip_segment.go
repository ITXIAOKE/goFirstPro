package waf_openapi

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

// DescribeWafSourceIpSegment invokes the waf_openapi.DescribeWafSourceIpSegment API synchronously
func (client *Client) DescribeWafSourceIpSegment(request *DescribeWafSourceIpSegmentRequest) (response *DescribeWafSourceIpSegmentResponse, err error) {
	response = CreateDescribeWafSourceIpSegmentResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWafSourceIpSegmentWithChan invokes the waf_openapi.DescribeWafSourceIpSegment API asynchronously
func (client *Client) DescribeWafSourceIpSegmentWithChan(request *DescribeWafSourceIpSegmentRequest) (<-chan *DescribeWafSourceIpSegmentResponse, <-chan error) {
	responseChan := make(chan *DescribeWafSourceIpSegmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWafSourceIpSegment(request)
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

// DescribeWafSourceIpSegmentWithCallback invokes the waf_openapi.DescribeWafSourceIpSegment API asynchronously
func (client *Client) DescribeWafSourceIpSegmentWithCallback(request *DescribeWafSourceIpSegmentRequest, callback func(response *DescribeWafSourceIpSegmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWafSourceIpSegmentResponse
		var err error
		defer close(result)
		response, err = client.DescribeWafSourceIpSegment(request)
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

// DescribeWafSourceIpSegmentRequest is the request struct for api DescribeWafSourceIpSegment
type DescribeWafSourceIpSegmentRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeWafSourceIpSegmentResponse is the response struct for api DescribeWafSourceIpSegment
type DescribeWafSourceIpSegmentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Ips       string `json:"Ips" xml:"Ips"`
	IpV6s     string `json:"IpV6s" xml:"IpV6s"`
}

// CreateDescribeWafSourceIpSegmentRequest creates a request to invoke DescribeWafSourceIpSegment API
func CreateDescribeWafSourceIpSegmentRequest() (request *DescribeWafSourceIpSegmentRequest) {
	request = &DescribeWafSourceIpSegmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeWafSourceIpSegment", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWafSourceIpSegmentResponse creates a response to parse from DescribeWafSourceIpSegment response
func CreateDescribeWafSourceIpSegmentResponse() (response *DescribeWafSourceIpSegmentResponse) {
	response = &DescribeWafSourceIpSegmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
