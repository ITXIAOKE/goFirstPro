package vs

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

// DescribeVsDomainRegionData invokes the vs.DescribeVsDomainRegionData API synchronously
func (client *Client) DescribeVsDomainRegionData(request *DescribeVsDomainRegionDataRequest) (response *DescribeVsDomainRegionDataResponse, err error) {
	response = CreateDescribeVsDomainRegionDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVsDomainRegionDataWithChan invokes the vs.DescribeVsDomainRegionData API asynchronously
func (client *Client) DescribeVsDomainRegionDataWithChan(request *DescribeVsDomainRegionDataRequest) (<-chan *DescribeVsDomainRegionDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVsDomainRegionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVsDomainRegionData(request)
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

// DescribeVsDomainRegionDataWithCallback invokes the vs.DescribeVsDomainRegionData API asynchronously
func (client *Client) DescribeVsDomainRegionDataWithCallback(request *DescribeVsDomainRegionDataRequest, callback func(response *DescribeVsDomainRegionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVsDomainRegionDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVsDomainRegionData(request)
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

// DescribeVsDomainRegionDataRequest is the request struct for api DescribeVsDomainRegionData
type DescribeVsDomainRegionDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVsDomainRegionDataResponse is the response struct for api DescribeVsDomainRegionData
type DescribeVsDomainRegionDataResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DomainName   string `json:"DomainName" xml:"DomainName"`
	DataInterval string `json:"DataInterval" xml:"DataInterval"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	EndTime      string `json:"EndTime" xml:"EndTime"`
	Value        Value  `json:"Value" xml:"Value"`
}

// CreateDescribeVsDomainRegionDataRequest creates a request to invoke DescribeVsDomainRegionData API
func CreateDescribeVsDomainRegionDataRequest() (request *DescribeVsDomainRegionDataRequest) {
	request = &DescribeVsDomainRegionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeVsDomainRegionData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVsDomainRegionDataResponse creates a response to parse from DescribeVsDomainRegionData response
func CreateDescribeVsDomainRegionDataResponse() (response *DescribeVsDomainRegionDataResponse) {
	response = &DescribeVsDomainRegionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
