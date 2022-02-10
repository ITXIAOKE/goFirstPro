package idrsservice

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

// UpdateServiceConfiguration invokes the idrsservice.UpdateServiceConfiguration API synchronously
func (client *Client) UpdateServiceConfiguration(request *UpdateServiceConfigurationRequest) (response *UpdateServiceConfigurationResponse, err error) {
	response = CreateUpdateServiceConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateServiceConfigurationWithChan invokes the idrsservice.UpdateServiceConfiguration API asynchronously
func (client *Client) UpdateServiceConfigurationWithChan(request *UpdateServiceConfigurationRequest) (<-chan *UpdateServiceConfigurationResponse, <-chan error) {
	responseChan := make(chan *UpdateServiceConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateServiceConfiguration(request)
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

// UpdateServiceConfigurationWithCallback invokes the idrsservice.UpdateServiceConfiguration API asynchronously
func (client *Client) UpdateServiceConfigurationWithCallback(request *UpdateServiceConfigurationRequest, callback func(response *UpdateServiceConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateServiceConfigurationResponse
		var err error
		defer close(result)
		response, err = client.UpdateServiceConfiguration(request)
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

// UpdateServiceConfigurationRequest is the request struct for api UpdateServiceConfiguration
type UpdateServiceConfigurationRequest struct {
	*requests.RpcRequest
	LiveRecordMaxClient       requests.Integer `position:"Query" name:"LiveRecordMaxClient"`
	LiveRecordVideoResolution requests.Integer `position:"Query" name:"LiveRecordVideoResolution"`
	TaskItemQueueSize         requests.Integer `position:"Query" name:"TaskItemQueueSize"`
	LiveRecordLayout          requests.Integer `position:"Query" name:"LiveRecordLayout"`
	ClientQueueSize           requests.Integer `position:"Query" name:"ClientQueueSize"`
	LiveRecordTaskProfile     string           `position:"Query" name:"LiveRecordTaskProfile"`
	LiveRecordAll             requests.Boolean `position:"Query" name:"LiveRecordAll"`
	LiveRecordEveryOne        requests.Boolean `position:"Query" name:"LiveRecordEveryOne"`
}

// UpdateServiceConfigurationResponse is the response struct for api UpdateServiceConfiguration
type UpdateServiceConfigurationResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateServiceConfigurationRequest creates a request to invoke UpdateServiceConfiguration API
func CreateUpdateServiceConfigurationRequest() (request *UpdateServiceConfigurationRequest) {
	request = &UpdateServiceConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "UpdateServiceConfiguration", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateServiceConfigurationResponse creates a response to parse from UpdateServiceConfiguration response
func CreateUpdateServiceConfigurationResponse() (response *UpdateServiceConfigurationResponse) {
	response = &UpdateServiceConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
