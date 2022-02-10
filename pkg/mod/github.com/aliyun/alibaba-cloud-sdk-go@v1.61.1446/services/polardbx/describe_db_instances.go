package polardbx

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

// DescribeDBInstances invokes the polardbx.DescribeDBInstances API synchronously
func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
	response = CreateDescribeDBInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstancesWithChan invokes the polardbx.DescribeDBInstances API asynchronously
func (client *Client) DescribeDBInstancesWithChan(request *DescribeDBInstancesRequest) (<-chan *DescribeDBInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstances(request)
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

// DescribeDBInstancesWithCallback invokes the polardbx.DescribeDBInstances API asynchronously
func (client *Client) DescribeDBInstancesWithCallback(request *DescribeDBInstancesRequest, callback func(response *DescribeDBInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstances(request)
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

// DescribeDBInstancesRequest is the request struct for api DescribeDBInstances
type DescribeDBInstancesRequest struct {
	*requests.RpcRequest
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeDBInstancesResponse is the response struct for api DescribeDBInstances
type DescribeDBInstancesResponse struct {
	*responses.BaseResponse
	RequestId   string                            `json:"RequestId" xml:"RequestId"`
	PageNumber  int                               `json:"PageNumber" xml:"PageNumber"`
	PageSize    int                               `json:"PageSize" xml:"PageSize"`
	TotalNumber int                               `json:"TotalNumber" xml:"TotalNumber"`
	DBInstances []DBInstanceInDescribeDBInstances `json:"DBInstances" xml:"DBInstances"`
}

// CreateDescribeDBInstancesRequest creates a request to invoke DescribeDBInstances API
func CreateDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
	request = &DescribeDBInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "DescribeDBInstances", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstancesResponse creates a response to parse from DescribeDBInstances response
func CreateDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
	response = &DescribeDBInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}