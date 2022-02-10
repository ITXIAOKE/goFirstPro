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

// SetLoadBalancerHTTPListenerAttribute invokes the ens.SetLoadBalancerHTTPListenerAttribute API synchronously
func (client *Client) SetLoadBalancerHTTPListenerAttribute(request *SetLoadBalancerHTTPListenerAttributeRequest) (response *SetLoadBalancerHTTPListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerHTTPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// SetLoadBalancerHTTPListenerAttributeWithChan invokes the ens.SetLoadBalancerHTTPListenerAttribute API asynchronously
func (client *Client) SetLoadBalancerHTTPListenerAttributeWithChan(request *SetLoadBalancerHTTPListenerAttributeRequest) (<-chan *SetLoadBalancerHTTPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerHTTPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerHTTPListenerAttribute(request)
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

// SetLoadBalancerHTTPListenerAttributeWithCallback invokes the ens.SetLoadBalancerHTTPListenerAttribute API asynchronously
func (client *Client) SetLoadBalancerHTTPListenerAttributeWithCallback(request *SetLoadBalancerHTTPListenerAttributeRequest, callback func(response *SetLoadBalancerHTTPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerHTTPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerHTTPListenerAttribute(request)
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

// SetLoadBalancerHTTPListenerAttributeRequest is the request struct for api SetLoadBalancerHTTPListenerAttribute
type SetLoadBalancerHTTPListenerAttributeRequest struct {
	*requests.RpcRequest
	HealthCheckTimeout     requests.Integer `position:"Query" name:"HealthCheckTimeout"`
	HealthCheckURI         string           `position:"Query" name:"HealthCheckURI"`
	HealthCheck            string           `position:"Query" name:"HealthCheck"`
	Protocol               string           `position:"Query" name:"Protocol"`
	Cookie                 string           `position:"Query" name:"Cookie"`
	HealthCheckMethod      string           `position:"Query" name:"HealthCheckMethod"`
	HealthCheckDomain      string           `position:"Query" name:"HealthCheckDomain"`
	RequestTimeout         requests.Integer `position:"Query" name:"RequestTimeout"`
	LoadBalancerId         string           `position:"Query" name:"LoadBalancerId"`
	HealthCheckInterval    requests.Integer `position:"Query" name:"HealthCheckInterval"`
	Description            string           `position:"Query" name:"Description"`
	UnhealthyThreshold     requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold       requests.Integer `position:"Query" name:"HealthyThreshold"`
	Scheduler              string           `position:"Query" name:"Scheduler"`
	CookieTimeout          requests.Integer `position:"Query" name:"CookieTimeout"`
	StickySessionType      string           `position:"Query" name:"StickySessionType"`
	ListenerPort           requests.Integer `position:"Query" name:"ListenerPort"`
	StickySession          string           `position:"Query" name:"StickySession"`
	IdleTimeout            requests.Integer `position:"Query" name:"IdleTimeout"`
	HealthCheckConnectPort requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string           `position:"Query" name:"HealthCheckHttpCode"`
}

// SetLoadBalancerHTTPListenerAttributeResponse is the response struct for api SetLoadBalancerHTTPListenerAttribute
type SetLoadBalancerHTTPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLoadBalancerHTTPListenerAttributeRequest creates a request to invoke SetLoadBalancerHTTPListenerAttribute API
func CreateSetLoadBalancerHTTPListenerAttributeRequest() (request *SetLoadBalancerHTTPListenerAttributeRequest) {
	request = &SetLoadBalancerHTTPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "SetLoadBalancerHTTPListenerAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetLoadBalancerHTTPListenerAttributeResponse creates a response to parse from SetLoadBalancerHTTPListenerAttribute response
func CreateSetLoadBalancerHTTPListenerAttributeResponse() (response *SetLoadBalancerHTTPListenerAttributeResponse) {
	response = &SetLoadBalancerHTTPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
