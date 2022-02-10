package cloudauth

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

// CompareFaces invokes the cloudauth.CompareFaces API synchronously
func (client *Client) CompareFaces(request *CompareFacesRequest) (response *CompareFacesResponse, err error) {
	response = CreateCompareFacesResponse()
	err = client.DoAction(request, response)
	return
}

// CompareFacesWithChan invokes the cloudauth.CompareFaces API asynchronously
func (client *Client) CompareFacesWithChan(request *CompareFacesRequest) (<-chan *CompareFacesResponse, <-chan error) {
	responseChan := make(chan *CompareFacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CompareFaces(request)
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

// CompareFacesWithCallback invokes the cloudauth.CompareFaces API asynchronously
func (client *Client) CompareFacesWithCallback(request *CompareFacesRequest, callback func(response *CompareFacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CompareFacesResponse
		var err error
		defer close(result)
		response, err = client.CompareFaces(request)
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

// CompareFacesRequest is the request struct for api CompareFaces
type CompareFacesRequest struct {
	*requests.RpcRequest
	TargetImageUrl    string `position:"Body" name:"TargetImageUrl"`
	SourceImageUrl    string `position:"Body" name:"SourceImageUrl"`
	TargetImageBase64 string `position:"Body" name:"TargetImageBase64"`
	BizType           string `position:"Body" name:"BizType"`
	BizId             string `position:"Body" name:"BizId"`
	SourceImageBase64 string `position:"Body" name:"SourceImageBase64"`
}

// CompareFacesResponse is the response struct for api CompareFaces
type CompareFacesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	Success      bool         `json:"Success" xml:"Success"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateCompareFacesRequest creates a request to invoke CompareFaces API
func CreateCompareFacesRequest() (request *CompareFacesRequest) {
	request = &CompareFacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2020-11-12", "CompareFaces", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCompareFacesResponse creates a response to parse from CompareFaces response
func CreateCompareFacesResponse() (response *CompareFacesResponse) {
	response = &CompareFacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}