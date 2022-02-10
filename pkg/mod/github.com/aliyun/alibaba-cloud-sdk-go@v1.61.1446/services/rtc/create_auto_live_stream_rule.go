package rtc

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

// CreateAutoLiveStreamRule invokes the rtc.CreateAutoLiveStreamRule API synchronously
func (client *Client) CreateAutoLiveStreamRule(request *CreateAutoLiveStreamRuleRequest) (response *CreateAutoLiveStreamRuleResponse, err error) {
	response = CreateCreateAutoLiveStreamRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAutoLiveStreamRuleWithChan invokes the rtc.CreateAutoLiveStreamRule API asynchronously
func (client *Client) CreateAutoLiveStreamRuleWithChan(request *CreateAutoLiveStreamRuleRequest) (<-chan *CreateAutoLiveStreamRuleResponse, <-chan error) {
	responseChan := make(chan *CreateAutoLiveStreamRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAutoLiveStreamRule(request)
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

// CreateAutoLiveStreamRuleWithCallback invokes the rtc.CreateAutoLiveStreamRule API asynchronously
func (client *Client) CreateAutoLiveStreamRuleWithCallback(request *CreateAutoLiveStreamRuleRequest, callback func(response *CreateAutoLiveStreamRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAutoLiveStreamRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateAutoLiveStreamRule(request)
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

// CreateAutoLiveStreamRuleRequest is the request struct for api CreateAutoLiveStreamRule
type CreateAutoLiveStreamRuleRequest struct {
	*requests.RpcRequest
	RuleName          string           `position:"Query" name:"RuleName"`
	ChannelIdPrefixes *[]string        `position:"Query" name:"ChannelIdPrefixes"  type:"Repeated"`
	ShowLog           string           `position:"Query" name:"ShowLog"`
	PlayDomain        string           `position:"Query" name:"PlayDomain"`
	OwnerId           requests.Integer `position:"Query" name:"OwnerId"`
	AppId             string           `position:"Query" name:"AppId"`
	CallBack          string           `position:"Query" name:"CallBack"`
	MediaEncode       requests.Integer `position:"Query" name:"MediaEncode"`
	ChannelIds        *[]string        `position:"Query" name:"ChannelIds"  type:"Repeated"`
}

// CreateAutoLiveStreamRuleResponse is the response struct for api CreateAutoLiveStreamRule
type CreateAutoLiveStreamRuleResponse struct {
	*responses.BaseResponse
	RuleId    int64  `json:"RuleId" xml:"RuleId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateAutoLiveStreamRuleRequest creates a request to invoke CreateAutoLiveStreamRule API
func CreateCreateAutoLiveStreamRuleRequest() (request *CreateAutoLiveStreamRuleRequest) {
	request = &CreateAutoLiveStreamRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "CreateAutoLiveStreamRule", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateAutoLiveStreamRuleResponse creates a response to parse from CreateAutoLiveStreamRule response
func CreateCreateAutoLiveStreamRuleResponse() (response *CreateAutoLiveStreamRuleResponse) {
	response = &CreateAutoLiveStreamRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
