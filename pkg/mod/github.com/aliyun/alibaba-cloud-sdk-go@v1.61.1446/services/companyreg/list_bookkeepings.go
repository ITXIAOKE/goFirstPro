package companyreg

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

// ListBookkeepings invokes the companyreg.ListBookkeepings API synchronously
func (client *Client) ListBookkeepings(request *ListBookkeepingsRequest) (response *ListBookkeepingsResponse, err error) {
	response = CreateListBookkeepingsResponse()
	err = client.DoAction(request, response)
	return
}

// ListBookkeepingsWithChan invokes the companyreg.ListBookkeepings API asynchronously
func (client *Client) ListBookkeepingsWithChan(request *ListBookkeepingsRequest) (<-chan *ListBookkeepingsResponse, <-chan error) {
	responseChan := make(chan *ListBookkeepingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBookkeepings(request)
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

// ListBookkeepingsWithCallback invokes the companyreg.ListBookkeepings API asynchronously
func (client *Client) ListBookkeepingsWithCallback(request *ListBookkeepingsRequest, callback func(response *ListBookkeepingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBookkeepingsResponse
		var err error
		defer close(result)
		response, err = client.ListBookkeepings(request)
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

// ListBookkeepingsRequest is the request struct for api ListBookkeepings
type ListBookkeepingsRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListBookkeepingsResponse is the response struct for api ListBookkeepings
type ListBookkeepingsResponse struct {
	*responses.BaseResponse
	RequestId         string        `json:"RequestId" xml:"RequestId"`
	TotalItemNumber   int           `json:"TotalItemNumber" xml:"TotalItemNumber"`
	CurrentPageNumber int           `json:"CurrentPageNumber" xml:"CurrentPageNumber"`
	PageSize          int           `json:"PageSize" xml:"PageSize"`
	TotalPageNumber   int           `json:"TotalPageNumber" xml:"TotalPageNumber"`
	Bookkeepings      []Bookkeeping `json:"Bookkeepings" xml:"Bookkeepings"`
}

// CreateListBookkeepingsRequest creates a request to invoke ListBookkeepings API
func CreateListBookkeepingsRequest() (request *ListBookkeepingsRequest) {
	request = &ListBookkeepingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ListBookkeepings", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListBookkeepingsResponse creates a response to parse from ListBookkeepings response
func CreateListBookkeepingsResponse() (response *ListBookkeepingsResponse) {
	response = &ListBookkeepingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}