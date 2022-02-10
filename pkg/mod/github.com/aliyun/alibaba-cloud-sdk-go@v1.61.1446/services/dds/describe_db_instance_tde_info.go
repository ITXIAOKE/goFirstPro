package dds

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

// DescribeDBInstanceTDEInfo invokes the dds.DescribeDBInstanceTDEInfo API synchronously
func (client *Client) DescribeDBInstanceTDEInfo(request *DescribeDBInstanceTDEInfoRequest) (response *DescribeDBInstanceTDEInfoResponse, err error) {
	response = CreateDescribeDBInstanceTDEInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceTDEInfoWithChan invokes the dds.DescribeDBInstanceTDEInfo API asynchronously
func (client *Client) DescribeDBInstanceTDEInfoWithChan(request *DescribeDBInstanceTDEInfoRequest) (<-chan *DescribeDBInstanceTDEInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceTDEInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceTDEInfo(request)
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

// DescribeDBInstanceTDEInfoWithCallback invokes the dds.DescribeDBInstanceTDEInfo API asynchronously
func (client *Client) DescribeDBInstanceTDEInfoWithCallback(request *DescribeDBInstanceTDEInfoRequest, callback func(response *DescribeDBInstanceTDEInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceTDEInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceTDEInfo(request)
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

// DescribeDBInstanceTDEInfoRequest is the request struct for api DescribeDBInstanceTDEInfo
type DescribeDBInstanceTDEInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBInstanceTDEInfoResponse is the response struct for api DescribeDBInstanceTDEInfo
type DescribeDBInstanceTDEInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TDEStatus string `json:"TDEStatus" xml:"TDEStatus"`
}

// CreateDescribeDBInstanceTDEInfoRequest creates a request to invoke DescribeDBInstanceTDEInfo API
func CreateDescribeDBInstanceTDEInfoRequest() (request *DescribeDBInstanceTDEInfoRequest) {
	request = &DescribeDBInstanceTDEInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeDBInstanceTDEInfo", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceTDEInfoResponse creates a response to parse from DescribeDBInstanceTDEInfo response
func CreateDescribeDBInstanceTDEInfoResponse() (response *DescribeDBInstanceTDEInfoResponse) {
	response = &DescribeDBInstanceTDEInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
