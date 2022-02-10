package dytnsapi

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

// DescribeEmptyNumberDetect invokes the dytnsapi.DescribeEmptyNumberDetect API synchronously
func (client *Client) DescribeEmptyNumberDetect(request *DescribeEmptyNumberDetectRequest) (response *DescribeEmptyNumberDetectResponse, err error) {
	response = CreateDescribeEmptyNumberDetectResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEmptyNumberDetectWithChan invokes the dytnsapi.DescribeEmptyNumberDetect API asynchronously
func (client *Client) DescribeEmptyNumberDetectWithChan(request *DescribeEmptyNumberDetectRequest) (<-chan *DescribeEmptyNumberDetectResponse, <-chan error) {
	responseChan := make(chan *DescribeEmptyNumberDetectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEmptyNumberDetect(request)
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

// DescribeEmptyNumberDetectWithCallback invokes the dytnsapi.DescribeEmptyNumberDetect API asynchronously
func (client *Client) DescribeEmptyNumberDetectWithCallback(request *DescribeEmptyNumberDetectRequest, callback func(response *DescribeEmptyNumberDetectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEmptyNumberDetectResponse
		var err error
		defer close(result)
		response, err = client.DescribeEmptyNumberDetect(request)
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

// DescribeEmptyNumberDetectRequest is the request struct for api DescribeEmptyNumberDetect
type DescribeEmptyNumberDetectRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EncryptType          string           `position:"Query" name:"EncryptType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Phone                string           `position:"Query" name:"Phone"`
}

// DescribeEmptyNumberDetectResponse is the response struct for api DescribeEmptyNumberDetect
type DescribeEmptyNumberDetectResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	Data      []DataList `json:"Data" xml:"Data"`
}

// CreateDescribeEmptyNumberDetectRequest creates a request to invoke DescribeEmptyNumberDetect API
func CreateDescribeEmptyNumberDetectRequest() (request *DescribeEmptyNumberDetectRequest) {
	request = &DescribeEmptyNumberDetectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dytnsapi", "2020-02-17", "DescribeEmptyNumberDetect", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeEmptyNumberDetectResponse creates a response to parse from DescribeEmptyNumberDetect response
func CreateDescribeEmptyNumberDetectResponse() (response *DescribeEmptyNumberDetectResponse) {
	response = &DescribeEmptyNumberDetectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
