package live

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

// AddCasterEpisodeGroupContent invokes the live.AddCasterEpisodeGroupContent API synchronously
func (client *Client) AddCasterEpisodeGroupContent(request *AddCasterEpisodeGroupContentRequest) (response *AddCasterEpisodeGroupContentResponse, err error) {
	response = CreateAddCasterEpisodeGroupContentResponse()
	err = client.DoAction(request, response)
	return
}

// AddCasterEpisodeGroupContentWithChan invokes the live.AddCasterEpisodeGroupContent API asynchronously
func (client *Client) AddCasterEpisodeGroupContentWithChan(request *AddCasterEpisodeGroupContentRequest) (<-chan *AddCasterEpisodeGroupContentResponse, <-chan error) {
	responseChan := make(chan *AddCasterEpisodeGroupContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCasterEpisodeGroupContent(request)
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

// AddCasterEpisodeGroupContentWithCallback invokes the live.AddCasterEpisodeGroupContent API asynchronously
func (client *Client) AddCasterEpisodeGroupContentWithCallback(request *AddCasterEpisodeGroupContentRequest, callback func(response *AddCasterEpisodeGroupContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCasterEpisodeGroupContentResponse
		var err error
		defer close(result)
		response, err = client.AddCasterEpisodeGroupContent(request)
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

// AddCasterEpisodeGroupContentRequest is the request struct for api AddCasterEpisodeGroupContent
type AddCasterEpisodeGroupContentRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Query" name:"ClientToken"`
	Content     string           `position:"Query" name:"Content"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// AddCasterEpisodeGroupContentResponse is the response struct for api AddCasterEpisodeGroupContent
type AddCasterEpisodeGroupContentResponse struct {
	*responses.BaseResponse
	RequestId string                                `json:"RequestId" xml:"RequestId"`
	ProgramId string                                `json:"ProgramId" xml:"ProgramId"`
	ItemIds   ItemIdsInAddCasterEpisodeGroupContent `json:"ItemIds" xml:"ItemIds"`
}

// CreateAddCasterEpisodeGroupContentRequest creates a request to invoke AddCasterEpisodeGroupContent API
func CreateAddCasterEpisodeGroupContentRequest() (request *AddCasterEpisodeGroupContentRequest) {
	request = &AddCasterEpisodeGroupContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddCasterEpisodeGroupContent", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddCasterEpisodeGroupContentResponse creates a response to parse from AddCasterEpisodeGroupContent response
func CreateAddCasterEpisodeGroupContentResponse() (response *AddCasterEpisodeGroupContentResponse) {
	response = &AddCasterEpisodeGroupContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
