package cloudcallcenter

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

// DeletePrivacyNumberPool invokes the cloudcallcenter.DeletePrivacyNumberPool API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deleteprivacynumberpool.html
func (client *Client) DeletePrivacyNumberPool(request *DeletePrivacyNumberPoolRequest) (response *DeletePrivacyNumberPoolResponse, err error) {
	response = CreateDeletePrivacyNumberPoolResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePrivacyNumberPoolWithChan invokes the cloudcallcenter.DeletePrivacyNumberPool API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deleteprivacynumberpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePrivacyNumberPoolWithChan(request *DeletePrivacyNumberPoolRequest) (<-chan *DeletePrivacyNumberPoolResponse, <-chan error) {
	responseChan := make(chan *DeletePrivacyNumberPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePrivacyNumberPool(request)
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

// DeletePrivacyNumberPoolWithCallback invokes the cloudcallcenter.DeletePrivacyNumberPool API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deleteprivacynumberpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePrivacyNumberPoolWithCallback(request *DeletePrivacyNumberPoolRequest, callback func(response *DeletePrivacyNumberPoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePrivacyNumberPoolResponse
		var err error
		defer close(result)
		response, err = client.DeletePrivacyNumberPool(request)
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

// DeletePrivacyNumberPoolRequest is the request struct for api DeletePrivacyNumberPool
type DeletePrivacyNumberPoolRequest struct {
	*requests.RpcRequest
	PoolId string `position:"Query" name:"PoolId"`
}

// DeletePrivacyNumberPoolResponse is the response struct for api DeletePrivacyNumberPool
type DeletePrivacyNumberPoolResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeletePrivacyNumberPoolRequest creates a request to invoke DeletePrivacyNumberPool API
func CreateDeletePrivacyNumberPoolRequest() (request *DeletePrivacyNumberPoolRequest) {
	request = &DeletePrivacyNumberPoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeletePrivacyNumberPool", "", "")
	request.Method = requests.POST
	return
}

// CreateDeletePrivacyNumberPoolResponse creates a response to parse from DeletePrivacyNumberPool response
func CreateDeletePrivacyNumberPoolResponse() (response *DeletePrivacyNumberPoolResponse) {
	response = &DeletePrivacyNumberPoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
