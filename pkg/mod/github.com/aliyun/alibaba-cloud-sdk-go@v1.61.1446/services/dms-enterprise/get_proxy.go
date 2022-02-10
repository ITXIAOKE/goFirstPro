package dms_enterprise

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

// GetProxy invokes the dms_enterprise.GetProxy API synchronously
func (client *Client) GetProxy(request *GetProxyRequest) (response *GetProxyResponse, err error) {
	response = CreateGetProxyResponse()
	err = client.DoAction(request, response)
	return
}

// GetProxyWithChan invokes the dms_enterprise.GetProxy API asynchronously
func (client *Client) GetProxyWithChan(request *GetProxyRequest) (<-chan *GetProxyResponse, <-chan error) {
	responseChan := make(chan *GetProxyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProxy(request)
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

// GetProxyWithCallback invokes the dms_enterprise.GetProxy API asynchronously
func (client *Client) GetProxyWithCallback(request *GetProxyRequest, callback func(response *GetProxyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProxyResponse
		var err error
		defer close(result)
		response, err = client.GetProxy(request)
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

// GetProxyRequest is the request struct for api GetProxy
type GetProxyRequest struct {
	*requests.RpcRequest
	Tid        requests.Integer `position:"Query" name:"Tid"`
	InstanceId requests.Integer `position:"Query" name:"InstanceId"`
	ProxyId    requests.Integer `position:"Query" name:"ProxyId"`
}

// GetProxyResponse is the response struct for api GetProxy
type GetProxyResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Success       bool   `json:"Success" xml:"Success"`
	ErrorMessage  string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode     string `json:"ErrorCode" xml:"ErrorCode"`
	ProxyId       int64  `json:"ProxyId" xml:"ProxyId"`
	CreatorId     int64  `json:"CreatorId" xml:"CreatorId"`
	CreatorName   string `json:"CreatorName" xml:"CreatorName"`
	InstanceId    int64  `json:"InstanceId" xml:"InstanceId"`
	PrivateEnable bool   `json:"PrivateEnable" xml:"PrivateEnable"`
	PrivateHost   string `json:"PrivateHost" xml:"PrivateHost"`
	PublicEnable  bool   `json:"PublicEnable" xml:"PublicEnable"`
	PublicHost    string `json:"PublicHost" xml:"PublicHost"`
	MysqlPort     int    `json:"MysqlPort" xml:"MysqlPort"`
	HttpsPort     int    `json:"HttpsPort" xml:"HttpsPort"`
}

// CreateGetProxyRequest creates a request to invoke GetProxy API
func CreateGetProxyRequest() (request *GetProxyRequest) {
	request = &GetProxyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetProxy", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetProxyResponse creates a response to parse from GetProxy response
func CreateGetProxyResponse() (response *GetProxyResponse) {
	response = &GetProxyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
