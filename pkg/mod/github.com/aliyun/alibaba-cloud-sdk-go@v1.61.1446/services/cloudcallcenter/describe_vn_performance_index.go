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

// DescribeVnPerformanceIndex invokes the cloudcallcenter.DescribeVnPerformanceIndex API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnperformanceindex.html
func (client *Client) DescribeVnPerformanceIndex(request *DescribeVnPerformanceIndexRequest) (response *DescribeVnPerformanceIndexResponse, err error) {
	response = CreateDescribeVnPerformanceIndexResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVnPerformanceIndexWithChan invokes the cloudcallcenter.DescribeVnPerformanceIndex API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnperformanceindex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVnPerformanceIndexWithChan(request *DescribeVnPerformanceIndexRequest) (<-chan *DescribeVnPerformanceIndexResponse, <-chan error) {
	responseChan := make(chan *DescribeVnPerformanceIndexResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVnPerformanceIndex(request)
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

// DescribeVnPerformanceIndexWithCallback invokes the cloudcallcenter.DescribeVnPerformanceIndex API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/describevnperformanceindex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVnPerformanceIndexWithCallback(request *DescribeVnPerformanceIndexRequest, callback func(response *DescribeVnPerformanceIndexResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVnPerformanceIndexResponse
		var err error
		defer close(result)
		response, err = client.DescribeVnPerformanceIndex(request)
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

// DescribeVnPerformanceIndexRequest is the request struct for api DescribeVnPerformanceIndex
type DescribeVnPerformanceIndexRequest struct {
	*requests.RpcRequest
	TimeUnit   string `position:"Query" name:"TimeUnit"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeVnPerformanceIndexResponse is the response struct for api DescribeVnPerformanceIndex
type DescribeVnPerformanceIndexResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	ResolutionRate   ResolutionRate   `json:"ResolutionRate" xml:"ResolutionRate"`
	ValidAnswerRate  ValidAnswerRate  `json:"ValidAnswerRate" xml:"ValidAnswerRate"`
	DialoguePassRate DialoguePassRate `json:"DialoguePassRate" xml:"DialoguePassRate"`
	KnowledgeHitRate KnowledgeHitRate `json:"KnowledgeHitRate" xml:"KnowledgeHitRate"`
}

// CreateDescribeVnPerformanceIndexRequest creates a request to invoke DescribeVnPerformanceIndex API
func CreateDescribeVnPerformanceIndexRequest() (request *DescribeVnPerformanceIndexRequest) {
	request = &DescribeVnPerformanceIndexRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DescribeVnPerformanceIndex", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeVnPerformanceIndexResponse creates a response to parse from DescribeVnPerformanceIndex response
func CreateDescribeVnPerformanceIndexResponse() (response *DescribeVnPerformanceIndexResponse) {
	response = &DescribeVnPerformanceIndexResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
