package csb

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

// FindBrokerSLOHisList invokes the csb.FindBrokerSLOHisList API synchronously
// api document: https://help.aliyun.com/api/csb/findbrokerslohislist.html
func (client *Client) FindBrokerSLOHisList(request *FindBrokerSLOHisListRequest) (response *FindBrokerSLOHisListResponse, err error) {
	response = CreateFindBrokerSLOHisListResponse()
	err = client.DoAction(request, response)
	return
}

// FindBrokerSLOHisListWithChan invokes the csb.FindBrokerSLOHisList API asynchronously
// api document: https://help.aliyun.com/api/csb/findbrokerslohislist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindBrokerSLOHisListWithChan(request *FindBrokerSLOHisListRequest) (<-chan *FindBrokerSLOHisListResponse, <-chan error) {
	responseChan := make(chan *FindBrokerSLOHisListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindBrokerSLOHisList(request)
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

// FindBrokerSLOHisListWithCallback invokes the csb.FindBrokerSLOHisList API asynchronously
// api document: https://help.aliyun.com/api/csb/findbrokerslohislist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindBrokerSLOHisListWithCallback(request *FindBrokerSLOHisListRequest, callback func(response *FindBrokerSLOHisListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindBrokerSLOHisListResponse
		var err error
		defer close(result)
		response, err = client.FindBrokerSLOHisList(request)
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

// FindBrokerSLOHisListRequest is the request struct for api FindBrokerSLOHisList
type FindBrokerSLOHisListRequest struct {
	*requests.RpcRequest
	CsbId       requests.Integer `position:"Query" name:"CsbId"`
	BeginDdHHmm string           `position:"Query" name:"BeginDdHHmm"`
	EndDdHHmm   string           `position:"Query" name:"EndDdHHmm"`
}

// FindBrokerSLOHisListResponse is the response struct for api FindBrokerSLOHisList
type FindBrokerSLOHisListResponse struct {
	*responses.BaseResponse
	Code      int                    `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateFindBrokerSLOHisListRequest creates a request to invoke FindBrokerSLOHisList API
func CreateFindBrokerSLOHisListRequest() (request *FindBrokerSLOHisListRequest) {
	request = &FindBrokerSLOHisListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindBrokerSLOHisList", "", "")
	request.Method = requests.GET
	return
}

// CreateFindBrokerSLOHisListResponse creates a response to parse from FindBrokerSLOHisList response
func CreateFindBrokerSLOHisListResponse() (response *FindBrokerSLOHisListResponse) {
	response = &FindBrokerSLOHisListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}