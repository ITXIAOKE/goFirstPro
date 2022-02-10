package retailcloud

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

// ScaleApp invokes the retailcloud.ScaleApp API synchronously
func (client *Client) ScaleApp(request *ScaleAppRequest) (response *ScaleAppResponse, err error) {
	response = CreateScaleAppResponse()
	err = client.DoAction(request, response)
	return
}

// ScaleAppWithChan invokes the retailcloud.ScaleApp API asynchronously
func (client *Client) ScaleAppWithChan(request *ScaleAppRequest) (<-chan *ScaleAppResponse, <-chan error) {
	responseChan := make(chan *ScaleAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScaleApp(request)
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

// ScaleAppWithCallback invokes the retailcloud.ScaleApp API asynchronously
func (client *Client) ScaleAppWithCallback(request *ScaleAppRequest, callback func(response *ScaleAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScaleAppResponse
		var err error
		defer close(result)
		response, err = client.ScaleApp(request)
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

// ScaleAppRequest is the request struct for api ScaleApp
type ScaleAppRequest struct {
	*requests.RpcRequest
	TotalPartitions requests.Integer `position:"Query" name:"TotalPartitions"`
	Replicas        requests.Integer `position:"Query" name:"Replicas"`
	EnvId           requests.Integer `position:"Query" name:"EnvId"`
}

// ScaleAppResponse is the response struct for api ScaleApp
type ScaleAppResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateScaleAppRequest creates a request to invoke ScaleApp API
func CreateScaleAppRequest() (request *ScaleAppRequest) {
	request = &ScaleAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ScaleApp", "", "")
	request.Method = requests.POST
	return
}

// CreateScaleAppResponse creates a response to parse from ScaleApp response
func CreateScaleAppResponse() (response *ScaleAppResponse) {
	response = &ScaleAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
