package edas

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

// DescribeLocalitySetting invokes the edas.DescribeLocalitySetting API synchronously
func (client *Client) DescribeLocalitySetting(request *DescribeLocalitySettingRequest) (response *DescribeLocalitySettingResponse, err error) {
	response = CreateDescribeLocalitySettingResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLocalitySettingWithChan invokes the edas.DescribeLocalitySetting API asynchronously
func (client *Client) DescribeLocalitySettingWithChan(request *DescribeLocalitySettingRequest) (<-chan *DescribeLocalitySettingResponse, <-chan error) {
	responseChan := make(chan *DescribeLocalitySettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLocalitySetting(request)
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

// DescribeLocalitySettingWithCallback invokes the edas.DescribeLocalitySetting API asynchronously
func (client *Client) DescribeLocalitySettingWithCallback(request *DescribeLocalitySettingRequest, callback func(response *DescribeLocalitySettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLocalitySettingResponse
		var err error
		defer close(result)
		response, err = client.DescribeLocalitySetting(request)
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

// DescribeLocalitySettingRequest is the request struct for api DescribeLocalitySetting
type DescribeLocalitySettingRequest struct {
	*requests.RoaRequest
	NamespaceId string `position:"Query" name:"NamespaceId"`
	AppId       string `position:"Query" name:"AppId"`
	Region      string `position:"Query" name:"Region"`
}

// DescribeLocalitySettingResponse is the response struct for api DescribeLocalitySetting
type DescribeLocalitySettingResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateDescribeLocalitySettingRequest creates a request to invoke DescribeLocalitySetting API
func CreateDescribeLocalitySettingRequest() (request *DescribeLocalitySettingRequest) {
	request = &DescribeLocalitySettingRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DescribeLocalitySetting", "/pop/sp/applications/locality/setting", "edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeLocalitySettingResponse creates a response to parse from DescribeLocalitySetting response
func CreateDescribeLocalitySettingResponse() (response *DescribeLocalitySettingResponse) {
	response = &DescribeLocalitySettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}