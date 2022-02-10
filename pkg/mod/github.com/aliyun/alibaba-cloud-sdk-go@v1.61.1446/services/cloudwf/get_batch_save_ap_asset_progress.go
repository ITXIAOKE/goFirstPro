package cloudwf

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

// GetBatchSaveApAssetProgress invokes the cloudwf.GetBatchSaveApAssetProgress API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getbatchsaveapassetprogress.html
func (client *Client) GetBatchSaveApAssetProgress(request *GetBatchSaveApAssetProgressRequest) (response *GetBatchSaveApAssetProgressResponse, err error) {
	response = CreateGetBatchSaveApAssetProgressResponse()
	err = client.DoAction(request, response)
	return
}

// GetBatchSaveApAssetProgressWithChan invokes the cloudwf.GetBatchSaveApAssetProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getbatchsaveapassetprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBatchSaveApAssetProgressWithChan(request *GetBatchSaveApAssetProgressRequest) (<-chan *GetBatchSaveApAssetProgressResponse, <-chan error) {
	responseChan := make(chan *GetBatchSaveApAssetProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBatchSaveApAssetProgress(request)
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

// GetBatchSaveApAssetProgressWithCallback invokes the cloudwf.GetBatchSaveApAssetProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getbatchsaveapassetprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBatchSaveApAssetProgressWithCallback(request *GetBatchSaveApAssetProgressRequest, callback func(response *GetBatchSaveApAssetProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBatchSaveApAssetProgressResponse
		var err error
		defer close(result)
		response, err = client.GetBatchSaveApAssetProgress(request)
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

// GetBatchSaveApAssetProgressRequest is the request struct for api GetBatchSaveApAssetProgress
type GetBatchSaveApAssetProgressRequest struct {
	*requests.RpcRequest
}

// GetBatchSaveApAssetProgressResponse is the response struct for api GetBatchSaveApAssetProgress
type GetBatchSaveApAssetProgressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetBatchSaveApAssetProgressRequest creates a request to invoke GetBatchSaveApAssetProgress API
func CreateGetBatchSaveApAssetProgressRequest() (request *GetBatchSaveApAssetProgressRequest) {
	request = &GetBatchSaveApAssetProgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetBatchSaveApAssetProgress", "cloudwf", "openAPI")
	return
}

// CreateGetBatchSaveApAssetProgressResponse creates a response to parse from GetBatchSaveApAssetProgress response
func CreateGetBatchSaveApAssetProgressResponse() (response *GetBatchSaveApAssetProgressResponse) {
	response = &GetBatchSaveApAssetProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
