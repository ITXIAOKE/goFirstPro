package cbn

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

// PublishRouteEntries invokes the cbn.PublishRouteEntries API synchronously
func (client *Client) PublishRouteEntries(request *PublishRouteEntriesRequest) (response *PublishRouteEntriesResponse, err error) {
	response = CreatePublishRouteEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// PublishRouteEntriesWithChan invokes the cbn.PublishRouteEntries API asynchronously
func (client *Client) PublishRouteEntriesWithChan(request *PublishRouteEntriesRequest) (<-chan *PublishRouteEntriesResponse, <-chan error) {
	responseChan := make(chan *PublishRouteEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishRouteEntries(request)
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

// PublishRouteEntriesWithCallback invokes the cbn.PublishRouteEntries API asynchronously
func (client *Client) PublishRouteEntriesWithCallback(request *PublishRouteEntriesRequest, callback func(response *PublishRouteEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishRouteEntriesResponse
		var err error
		defer close(result)
		response, err = client.PublishRouteEntries(request)
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

// PublishRouteEntriesRequest is the request struct for api PublishRouteEntries
type PublishRouteEntriesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                     string           `position:"Query" name:"CenId"`
	ChildInstanceRegionId     string           `position:"Query" name:"ChildInstanceRegionId"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	DestinationCidrBlock      string           `position:"Query" name:"DestinationCidrBlock"`
	ChildInstanceType         string           `position:"Query" name:"ChildInstanceType"`
	ChildInstanceId           string           `position:"Query" name:"ChildInstanceId"`
	ChildInstanceRouteTableId string           `position:"Query" name:"ChildInstanceRouteTableId"`
}

// PublishRouteEntriesResponse is the response struct for api PublishRouteEntries
type PublishRouteEntriesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePublishRouteEntriesRequest creates a request to invoke PublishRouteEntries API
func CreatePublishRouteEntriesRequest() (request *PublishRouteEntriesRequest) {
	request = &PublishRouteEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "PublishRouteEntries", "", "")
	request.Method = requests.POST
	return
}

// CreatePublishRouteEntriesResponse creates a response to parse from PublishRouteEntries response
func CreatePublishRouteEntriesResponse() (response *PublishRouteEntriesResponse) {
	response = &PublishRouteEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
