package dyplsapi

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

// QueryCallStatus invokes the dyplsapi.QueryCallStatus API synchronously
func (client *Client) QueryCallStatus(request *QueryCallStatusRequest) (response *QueryCallStatusResponse, err error) {
	response = CreateQueryCallStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCallStatusWithChan invokes the dyplsapi.QueryCallStatus API asynchronously
func (client *Client) QueryCallStatusWithChan(request *QueryCallStatusRequest) (<-chan *QueryCallStatusResponse, <-chan error) {
	responseChan := make(chan *QueryCallStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCallStatus(request)
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

// QueryCallStatusWithCallback invokes the dyplsapi.QueryCallStatus API asynchronously
func (client *Client) QueryCallStatusWithCallback(request *QueryCallStatusRequest, callback func(response *QueryCallStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCallStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryCallStatus(request)
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

// QueryCallStatusRequest is the request struct for api QueryCallStatus
type QueryCallStatusRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SubsId               string           `position:"Query" name:"SubsId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CallNo               string           `position:"Query" name:"CallNo"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
}

// QueryCallStatusResponse is the response struct for api QueryCallStatus
type QueryCallStatusResponse struct {
	*responses.BaseResponse
	Code                string              `json:"Code" xml:"Code"`
	Message             string              `json:"Message" xml:"Message"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	SecretCallStatusDTO SecretCallStatusDTO `json:"SecretCallStatusDTO" xml:"SecretCallStatusDTO"`
}

// CreateQueryCallStatusRequest creates a request to invoke QueryCallStatus API
func CreateQueryCallStatusRequest() (request *QueryCallStatusRequest) {
	request = &QueryCallStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "QueryCallStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryCallStatusResponse creates a response to parse from QueryCallStatus response
func CreateQueryCallStatusResponse() (response *QueryCallStatusResponse) {
	response = &QueryCallStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
