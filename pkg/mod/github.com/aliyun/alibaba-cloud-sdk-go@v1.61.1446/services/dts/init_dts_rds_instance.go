package dts

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

// InitDtsRdsInstance invokes the dts.InitDtsRdsInstance API synchronously
func (client *Client) InitDtsRdsInstance(request *InitDtsRdsInstanceRequest) (response *InitDtsRdsInstanceResponse, err error) {
	response = CreateInitDtsRdsInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// InitDtsRdsInstanceWithChan invokes the dts.InitDtsRdsInstance API asynchronously
func (client *Client) InitDtsRdsInstanceWithChan(request *InitDtsRdsInstanceRequest) (<-chan *InitDtsRdsInstanceResponse, <-chan error) {
	responseChan := make(chan *InitDtsRdsInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitDtsRdsInstance(request)
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

// InitDtsRdsInstanceWithCallback invokes the dts.InitDtsRdsInstance API asynchronously
func (client *Client) InitDtsRdsInstanceWithCallback(request *InitDtsRdsInstanceRequest, callback func(response *InitDtsRdsInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitDtsRdsInstanceResponse
		var err error
		defer close(result)
		response, err = client.InitDtsRdsInstance(request)
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

// InitDtsRdsInstanceRequest is the request struct for api InitDtsRdsInstance
type InitDtsRdsInstanceRequest struct {
	*requests.RpcRequest
	EndpointInstanceId   string `position:"Query" name:"EndpointInstanceId"`
	EndpointRegion       string `position:"Query" name:"EndpointRegion"`
	EndpointCenId        string `position:"Query" name:"EndpointCenId"`
	EndpointInstanceType string `position:"Query" name:"EndpointInstanceType"`
	DtsInstanceId        string `position:"Query" name:"DtsInstanceId"`
}

// InitDtsRdsInstanceResponse is the response struct for api InitDtsRdsInstance
type InitDtsRdsInstanceResponse struct {
	*responses.BaseResponse
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        string `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	AdminAccount   string `json:"AdminAccount" xml:"AdminAccount"`
	AdminPassword  string `json:"AdminPassword" xml:"AdminPassword"`
}

// CreateInitDtsRdsInstanceRequest creates a request to invoke InitDtsRdsInstance API
func CreateInitDtsRdsInstanceRequest() (request *InitDtsRdsInstanceRequest) {
	request = &InitDtsRdsInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "InitDtsRdsInstance", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInitDtsRdsInstanceResponse creates a response to parse from InitDtsRdsInstance response
func CreateInitDtsRdsInstanceResponse() (response *InitDtsRdsInstanceResponse) {
	response = &InitDtsRdsInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
