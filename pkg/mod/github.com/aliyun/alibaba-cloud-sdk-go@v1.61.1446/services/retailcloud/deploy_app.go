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

// DeployApp invokes the retailcloud.DeployApp API synchronously
func (client *Client) DeployApp(request *DeployAppRequest) (response *DeployAppResponse, err error) {
	response = CreateDeployAppResponse()
	err = client.DoAction(request, response)
	return
}

// DeployAppWithChan invokes the retailcloud.DeployApp API asynchronously
func (client *Client) DeployAppWithChan(request *DeployAppRequest) (<-chan *DeployAppResponse, <-chan error) {
	responseChan := make(chan *DeployAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployApp(request)
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

// DeployAppWithCallback invokes the retailcloud.DeployApp API asynchronously
func (client *Client) DeployAppWithCallback(request *DeployAppRequest, callback func(response *DeployAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployAppResponse
		var err error
		defer close(result)
		response, err = client.DeployApp(request)
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

// DeployAppRequest is the request struct for api DeployApp
type DeployAppRequest struct {
	*requests.RpcRequest
	DeployPacketUrl         string           `position:"Query" name:"DeployPacketUrl"`
	TotalPartitions         requests.Integer `position:"Query" name:"TotalPartitions"`
	Description             string           `position:"Query" name:"Description"`
	EnvId                   requests.Integer `position:"Query" name:"EnvId"`
	UpdateStrategyType      string           `position:"Query" name:"UpdateStrategyType"`
	PauseType               string           `position:"Query" name:"PauseType"`
	DeployPacketId          requests.Integer `position:"Query" name:"DeployPacketId"`
	ContainerImageList      *[]string        `position:"Query" name:"ContainerImageList"  type:"Repeated"`
	Name                    string           `position:"Query" name:"Name"`
	InitContainerImageList  *[]string        `position:"Query" name:"InitContainerImageList"  type:"Repeated"`
	DefaultPacketOfAppGroup string           `position:"Query" name:"DefaultPacketOfAppGroup"`
	ArmsFlag                requests.Boolean `position:"Query" name:"ArmsFlag"`
}

// DeployAppResponse is the response struct for api DeployApp
type DeployAppResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDeployAppRequest creates a request to invoke DeployApp API
func CreateDeployAppRequest() (request *DeployAppRequest) {
	request = &DeployAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DeployApp", "", "")
	request.Method = requests.POST
	return
}

// CreateDeployAppResponse creates a response to parse from DeployApp response
func CreateDeployAppResponse() (response *DeployAppResponse) {
	response = &DeployAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
