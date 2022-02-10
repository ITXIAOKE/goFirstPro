package vcs

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

// GetProfileList invokes the vcs.GetProfileList API synchronously
func (client *Client) GetProfileList(request *GetProfileListRequest) (response *GetProfileListResponse, err error) {
	response = CreateGetProfileListResponse()
	err = client.DoAction(request, response)
	return
}

// GetProfileListWithChan invokes the vcs.GetProfileList API asynchronously
func (client *Client) GetProfileListWithChan(request *GetProfileListRequest) (<-chan *GetProfileListResponse, <-chan error) {
	responseChan := make(chan *GetProfileListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProfileList(request)
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

// GetProfileListWithCallback invokes the vcs.GetProfileList API asynchronously
func (client *Client) GetProfileListWithCallback(request *GetProfileListRequest, callback func(response *GetProfileListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProfileListResponse
		var err error
		defer close(result)
		response, err = client.GetProfileList(request)
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

// GetProfileListRequest is the request struct for api GetProfileList
type GetProfileListRequest struct {
	*requests.RpcRequest
	ProfileIdList         map[string]interface{} `position:"Body" name:"ProfileIdList"`
	CorpId                string                 `position:"Body" name:"CorpId"`
	Gender                requests.Integer       `position:"Body" name:"Gender"`
	PlateNo               string                 `position:"Body" name:"PlateNo"`
	IdNumber              string                 `position:"Body" name:"IdNumber"`
	PageNumber            requests.Integer       `position:"Body" name:"PageNumber"`
	FaceImageId           string                 `position:"Body" name:"FaceImageId"`
	FaceUrl               string                 `position:"Body" name:"FaceUrl"`
	PageSize              requests.Integer       `position:"Body" name:"PageSize"`
	PersonIdList          map[string]interface{} `position:"Body" name:"PersonIdList"`
	LiveAddress           string                 `position:"Body" name:"LiveAddress"`
	IsvSubId              string                 `position:"Body" name:"IsvSubId"`
	SceneType             string                 `position:"Body" name:"SceneType"`
	PhoneNo               string                 `position:"Body" name:"PhoneNo"`
	CatalogId             requests.Integer       `position:"Body" name:"CatalogId"`
	Name                  string                 `position:"Body" name:"Name"`
	BizId                 string                 `position:"Body" name:"BizId"`
	MatchingRateThreshold string                 `position:"Body" name:"MatchingRateThreshold"`
}

// GetProfileListResponse is the response struct for api GetProfileList
type GetProfileListResponse struct {
	*responses.BaseResponse
	Code      string               `json:"Code" xml:"Code"`
	Message   string               `json:"Message" xml:"Message"`
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Data      DataInGetProfileList `json:"Data" xml:"Data"`
}

// CreateGetProfileListRequest creates a request to invoke GetProfileList API
func CreateGetProfileListRequest() (request *GetProfileListRequest) {
	request = &GetProfileListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetProfileList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetProfileListResponse creates a response to parse from GetProfileList response
func CreateGetProfileListResponse() (response *GetProfileListResponse) {
	response = &GetProfileListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}