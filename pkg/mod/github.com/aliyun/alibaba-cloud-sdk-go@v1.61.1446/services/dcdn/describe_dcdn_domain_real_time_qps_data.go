package dcdn

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

// DescribeDcdnDomainRealTimeQpsData invokes the dcdn.DescribeDcdnDomainRealTimeQpsData API synchronously
func (client *Client) DescribeDcdnDomainRealTimeQpsData(request *DescribeDcdnDomainRealTimeQpsDataRequest) (response *DescribeDcdnDomainRealTimeQpsDataResponse, err error) {
	response = CreateDescribeDcdnDomainRealTimeQpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainRealTimeQpsDataWithChan invokes the dcdn.DescribeDcdnDomainRealTimeQpsData API asynchronously
func (client *Client) DescribeDcdnDomainRealTimeQpsDataWithChan(request *DescribeDcdnDomainRealTimeQpsDataRequest) (<-chan *DescribeDcdnDomainRealTimeQpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainRealTimeQpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainRealTimeQpsData(request)
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

// DescribeDcdnDomainRealTimeQpsDataWithCallback invokes the dcdn.DescribeDcdnDomainRealTimeQpsData API asynchronously
func (client *Client) DescribeDcdnDomainRealTimeQpsDataWithCallback(request *DescribeDcdnDomainRealTimeQpsDataRequest, callback func(response *DescribeDcdnDomainRealTimeQpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainRealTimeQpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainRealTimeQpsData(request)
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

// DescribeDcdnDomainRealTimeQpsDataRequest is the request struct for api DescribeDcdnDomainRealTimeQpsData
type DescribeDcdnDomainRealTimeQpsDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnDomainRealTimeQpsDataResponse is the response struct for api DescribeDcdnDomainRealTimeQpsData
type DescribeDcdnDomainRealTimeQpsDataResponse struct {
	*responses.BaseResponse
	RequestId string                                  `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDcdnDomainRealTimeQpsData `json:"Data" xml:"Data"`
}

// CreateDescribeDcdnDomainRealTimeQpsDataRequest creates a request to invoke DescribeDcdnDomainRealTimeQpsData API
func CreateDescribeDcdnDomainRealTimeQpsDataRequest() (request *DescribeDcdnDomainRealTimeQpsDataRequest) {
	request = &DescribeDcdnDomainRealTimeQpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainRealTimeQpsData", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeDcdnDomainRealTimeQpsDataResponse creates a response to parse from DescribeDcdnDomainRealTimeQpsData response
func CreateDescribeDcdnDomainRealTimeQpsDataResponse() (response *DescribeDcdnDomainRealTimeQpsDataResponse) {
	response = &DescribeDcdnDomainRealTimeQpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
