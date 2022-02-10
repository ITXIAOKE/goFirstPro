package sgw

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

// DescribeMqttConfig invokes the sgw.DescribeMqttConfig API synchronously
func (client *Client) DescribeMqttConfig(request *DescribeMqttConfigRequest) (response *DescribeMqttConfigResponse, err error) {
	response = CreateDescribeMqttConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMqttConfigWithChan invokes the sgw.DescribeMqttConfig API asynchronously
func (client *Client) DescribeMqttConfigWithChan(request *DescribeMqttConfigRequest) (<-chan *DescribeMqttConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeMqttConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMqttConfig(request)
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

// DescribeMqttConfigWithCallback invokes the sgw.DescribeMqttConfig API asynchronously
func (client *Client) DescribeMqttConfigWithCallback(request *DescribeMqttConfigRequest, callback func(response *DescribeMqttConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMqttConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeMqttConfig(request)
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

// DescribeMqttConfigRequest is the request struct for api DescribeMqttConfig
type DescribeMqttConfigRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// DescribeMqttConfigResponse is the response struct for api DescribeMqttConfig
type DescribeMqttConfigResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	Success           bool   `json:"Success" xml:"Success"`
	Code              string `json:"Code" xml:"Code"`
	Message           string `json:"Message" xml:"Message"`
	IsEnabled         bool   `json:"IsEnabled" xml:"IsEnabled"`
	BrokerUrl         string `json:"BrokerUrl" xml:"BrokerUrl"`
	InternalBrokerUrl string `json:"InternalBrokerUrl" xml:"InternalBrokerUrl"`
	PublishTopic      string `json:"PublishTopic" xml:"PublishTopic"`
	SubscribeTopic    string `json:"SubscribeTopic" xml:"SubscribeTopic"`
	GroupId           string `json:"GroupId" xml:"GroupId"`
	MqttInstanceId    string `json:"MqttInstanceId" xml:"MqttInstanceId"`
	AuthType          string `json:"AuthType" xml:"AuthType"`
	Username          string `json:"Username" xml:"Username"`
	Password          string `json:"Password" xml:"Password"`
}

// CreateDescribeMqttConfigRequest creates a request to invoke DescribeMqttConfig API
func CreateDescribeMqttConfigRequest() (request *DescribeMqttConfigRequest) {
	request = &DescribeMqttConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeMqttConfig", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMqttConfigResponse creates a response to parse from DescribeMqttConfig response
func CreateDescribeMqttConfigResponse() (response *DescribeMqttConfigResponse) {
	response = &DescribeMqttConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
