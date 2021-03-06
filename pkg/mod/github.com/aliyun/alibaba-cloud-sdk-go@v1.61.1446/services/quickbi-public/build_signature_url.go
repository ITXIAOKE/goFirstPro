package quickbi_public

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

// BuildSignatureUrl invokes the quickbi_public.BuildSignatureUrl API synchronously
func (client *Client) BuildSignatureUrl(request *BuildSignatureUrlRequest) (response *BuildSignatureUrlResponse, err error) {
	response = CreateBuildSignatureUrlResponse()
	err = client.DoAction(request, response)
	return
}

// BuildSignatureUrlWithChan invokes the quickbi_public.BuildSignatureUrl API asynchronously
func (client *Client) BuildSignatureUrlWithChan(request *BuildSignatureUrlRequest) (<-chan *BuildSignatureUrlResponse, <-chan error) {
	responseChan := make(chan *BuildSignatureUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BuildSignatureUrl(request)
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

// BuildSignatureUrlWithCallback invokes the quickbi_public.BuildSignatureUrl API asynchronously
func (client *Client) BuildSignatureUrlWithCallback(request *BuildSignatureUrlRequest, callback func(response *BuildSignatureUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BuildSignatureUrlResponse
		var err error
		defer close(result)
		response, err = client.BuildSignatureUrl(request)
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

// BuildSignatureUrlRequest is the request struct for api BuildSignatureUrl
type BuildSignatureUrlRequest struct {
	*requests.RpcRequest
	Watermark   string           `position:"Query" name:"Watermark"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	AccountType requests.Integer `position:"Query" name:"AccountType"`
	SignType    string           `position:"Query" name:"SignType"`
	UserId      string           `position:"Query" name:"UserId"`
	AccountName string           `position:"Query" name:"AccountName"`
	WorksId     string           `position:"Query" name:"WorksId"`
	WithHost    requests.Boolean `position:"Query" name:"WithHost"`
	ExpireMin   requests.Integer `position:"Query" name:"ExpireMin"`
}

// BuildSignatureUrlResponse is the response struct for api BuildSignatureUrl
type BuildSignatureUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateBuildSignatureUrlRequest creates a request to invoke BuildSignatureUrl API
func CreateBuildSignatureUrlRequest() (request *BuildSignatureUrlRequest) {
	request = &BuildSignatureUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2021-03-25", "BuildSignatureUrl", "quickbi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBuildSignatureUrlResponse creates a response to parse from BuildSignatureUrl response
func CreateBuildSignatureUrlResponse() (response *BuildSignatureUrlResponse) {
	response = &BuildSignatureUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
