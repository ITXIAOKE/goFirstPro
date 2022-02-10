package domain

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

// SaveBatchTaskForCreatingOrderActivate invokes the domain.SaveBatchTaskForCreatingOrderActivate API synchronously
func (client *Client) SaveBatchTaskForCreatingOrderActivate(request *SaveBatchTaskForCreatingOrderActivateRequest) (response *SaveBatchTaskForCreatingOrderActivateResponse, err error) {
	response = CreateSaveBatchTaskForCreatingOrderActivateResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForCreatingOrderActivateWithChan invokes the domain.SaveBatchTaskForCreatingOrderActivate API asynchronously
func (client *Client) SaveBatchTaskForCreatingOrderActivateWithChan(request *SaveBatchTaskForCreatingOrderActivateRequest) (<-chan *SaveBatchTaskForCreatingOrderActivateResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForCreatingOrderActivateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForCreatingOrderActivate(request)
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

// SaveBatchTaskForCreatingOrderActivateWithCallback invokes the domain.SaveBatchTaskForCreatingOrderActivate API asynchronously
func (client *Client) SaveBatchTaskForCreatingOrderActivateWithCallback(request *SaveBatchTaskForCreatingOrderActivateRequest, callback func(response *SaveBatchTaskForCreatingOrderActivateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForCreatingOrderActivateResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForCreatingOrderActivate(request)
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

// SaveBatchTaskForCreatingOrderActivateRequest is the request struct for api SaveBatchTaskForCreatingOrderActivate
type SaveBatchTaskForCreatingOrderActivateRequest struct {
	*requests.RpcRequest
	OrderActivateParam *[]SaveBatchTaskForCreatingOrderActivateOrderActivateParam `position:"Query" name:"OrderActivateParam"  type:"Repeated"`
	CouponNo           string                                                     `position:"Query" name:"CouponNo"`
	UseCoupon          requests.Boolean                                           `position:"Query" name:"UseCoupon"`
	PromotionNo        string                                                     `position:"Query" name:"PromotionNo"`
	UserClientIp       string                                                     `position:"Query" name:"UserClientIp"`
	Lang               string                                                     `position:"Query" name:"Lang"`
	UsePromotion       requests.Boolean                                           `position:"Query" name:"UsePromotion"`
}

// SaveBatchTaskForCreatingOrderActivateOrderActivateParam is a repeated param struct in SaveBatchTaskForCreatingOrderActivateRequest
type SaveBatchTaskForCreatingOrderActivateOrderActivateParam struct {
	Country                   string `name:"Country"`
	SubscriptionDuration      string `name:"SubscriptionDuration"`
	PermitPremiumActivation   string `name:"PermitPremiumActivation"`
	City                      string `name:"City"`
	Dns2                      string `name:"Dns2"`
	Dns1                      string `name:"Dns1"`
	RegistrantProfileId       string `name:"RegistrantProfileId"`
	AliyunDns                 string `name:"AliyunDns"`
	ZhCity                    string `name:"ZhCity"`
	TelExt                    string `name:"TelExt"`
	ZhRegistrantName          string `name:"ZhRegistrantName"`
	Province                  string `name:"Province"`
	PostalCode                string `name:"PostalCode"`
	Email                     string `name:"Email"`
	ZhRegistrantOrganization  string `name:"ZhRegistrantOrganization"`
	Address                   string `name:"Address"`
	TelArea                   string `name:"TelArea"`
	DomainName                string `name:"DomainName"`
	ZhAddress                 string `name:"ZhAddress"`
	RegistrantType            string `name:"RegistrantType"`
	Telephone                 string `name:"Telephone"`
	TrademarkDomainActivation string `name:"TrademarkDomainActivation"`
	ZhProvince                string `name:"ZhProvince"`
	RegistrantOrganization    string `name:"RegistrantOrganization"`
	EnableDomainProxy         string `name:"EnableDomainProxy"`
	RegistrantName            string `name:"RegistrantName"`
}

// SaveBatchTaskForCreatingOrderActivateResponse is the response struct for api SaveBatchTaskForCreatingOrderActivate
type SaveBatchTaskForCreatingOrderActivateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForCreatingOrderActivateRequest creates a request to invoke SaveBatchTaskForCreatingOrderActivate API
func CreateSaveBatchTaskForCreatingOrderActivateRequest() (request *SaveBatchTaskForCreatingOrderActivateRequest) {
	request = &SaveBatchTaskForCreatingOrderActivateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForCreatingOrderActivate", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveBatchTaskForCreatingOrderActivateResponse creates a response to parse from SaveBatchTaskForCreatingOrderActivate response
func CreateSaveBatchTaskForCreatingOrderActivateResponse() (response *SaveBatchTaskForCreatingOrderActivateResponse) {
	response = &SaveBatchTaskForCreatingOrderActivateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
