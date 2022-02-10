package vs

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

// UpdateRenderingDeviceSpec invokes the vs.UpdateRenderingDeviceSpec API synchronously
func (client *Client) UpdateRenderingDeviceSpec(request *UpdateRenderingDeviceSpecRequest) (response *UpdateRenderingDeviceSpecResponse, err error) {
	response = CreateUpdateRenderingDeviceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRenderingDeviceSpecWithChan invokes the vs.UpdateRenderingDeviceSpec API asynchronously
func (client *Client) UpdateRenderingDeviceSpecWithChan(request *UpdateRenderingDeviceSpecRequest) (<-chan *UpdateRenderingDeviceSpecResponse, <-chan error) {
	responseChan := make(chan *UpdateRenderingDeviceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRenderingDeviceSpec(request)
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

// UpdateRenderingDeviceSpecWithCallback invokes the vs.UpdateRenderingDeviceSpec API asynchronously
func (client *Client) UpdateRenderingDeviceSpecWithCallback(request *UpdateRenderingDeviceSpecRequest, callback func(response *UpdateRenderingDeviceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRenderingDeviceSpecResponse
		var err error
		defer close(result)
		response, err = client.UpdateRenderingDeviceSpec(request)
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

// UpdateRenderingDeviceSpecRequest is the request struct for api UpdateRenderingDeviceSpec
type UpdateRenderingDeviceSpecRequest struct {
	*requests.RpcRequest
	Description     string           `position:"Query" name:"Description"`
	EffectiveTime   requests.Boolean `position:"Query" name:"EffectiveTime"`
	ShowLog         string           `position:"Query" name:"ShowLog"`
	AutoRenewPeriod requests.Integer `position:"Query" name:"AutoRenewPeriod"`
	Period          requests.Integer `position:"Query" name:"Period"`
	Specification   string           `position:"Query" name:"Specification"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	PeriodUnit      string           `position:"Query" name:"PeriodUnit"`
	AutoRenew       requests.Boolean `position:"Query" name:"AutoRenew"`
	InstanceIds     string           `position:"Query" name:"InstanceIds"`
}

// UpdateRenderingDeviceSpecResponse is the response struct for api UpdateRenderingDeviceSpec
type UpdateRenderingDeviceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateRenderingDeviceSpecRequest creates a request to invoke UpdateRenderingDeviceSpec API
func CreateUpdateRenderingDeviceSpecRequest() (request *UpdateRenderingDeviceSpecRequest) {
	request = &UpdateRenderingDeviceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "UpdateRenderingDeviceSpec", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateRenderingDeviceSpecResponse creates a response to parse from UpdateRenderingDeviceSpec response
func CreateUpdateRenderingDeviceSpecResponse() (response *UpdateRenderingDeviceSpecResponse) {
	response = &UpdateRenderingDeviceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
