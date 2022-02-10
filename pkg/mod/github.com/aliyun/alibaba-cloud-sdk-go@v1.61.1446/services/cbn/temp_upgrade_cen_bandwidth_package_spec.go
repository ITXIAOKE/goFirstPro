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

// TempUpgradeCenBandwidthPackageSpec invokes the cbn.TempUpgradeCenBandwidthPackageSpec API synchronously
func (client *Client) TempUpgradeCenBandwidthPackageSpec(request *TempUpgradeCenBandwidthPackageSpecRequest) (response *TempUpgradeCenBandwidthPackageSpecResponse, err error) {
	response = CreateTempUpgradeCenBandwidthPackageSpecResponse()
	err = client.DoAction(request, response)
	return
}

// TempUpgradeCenBandwidthPackageSpecWithChan invokes the cbn.TempUpgradeCenBandwidthPackageSpec API asynchronously
func (client *Client) TempUpgradeCenBandwidthPackageSpecWithChan(request *TempUpgradeCenBandwidthPackageSpecRequest) (<-chan *TempUpgradeCenBandwidthPackageSpecResponse, <-chan error) {
	responseChan := make(chan *TempUpgradeCenBandwidthPackageSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TempUpgradeCenBandwidthPackageSpec(request)
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

// TempUpgradeCenBandwidthPackageSpecWithCallback invokes the cbn.TempUpgradeCenBandwidthPackageSpec API asynchronously
func (client *Client) TempUpgradeCenBandwidthPackageSpecWithCallback(request *TempUpgradeCenBandwidthPackageSpecRequest, callback func(response *TempUpgradeCenBandwidthPackageSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TempUpgradeCenBandwidthPackageSpecResponse
		var err error
		defer close(result)
		response, err = client.TempUpgradeCenBandwidthPackageSpec(request)
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

// TempUpgradeCenBandwidthPackageSpecRequest is the request struct for api TempUpgradeCenBandwidthPackageSpec
type TempUpgradeCenBandwidthPackageSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth             requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	EndTime               string           `position:"Query" name:"EndTime"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	CenBandwidthPackageId string           `position:"Query" name:"CenBandwidthPackageId"`
}

// TempUpgradeCenBandwidthPackageSpecResponse is the response struct for api TempUpgradeCenBandwidthPackageSpec
type TempUpgradeCenBandwidthPackageSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTempUpgradeCenBandwidthPackageSpecRequest creates a request to invoke TempUpgradeCenBandwidthPackageSpec API
func CreateTempUpgradeCenBandwidthPackageSpecRequest() (request *TempUpgradeCenBandwidthPackageSpecRequest) {
	request = &TempUpgradeCenBandwidthPackageSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "TempUpgradeCenBandwidthPackageSpec", "", "")
	request.Method = requests.POST
	return
}

// CreateTempUpgradeCenBandwidthPackageSpecResponse creates a response to parse from TempUpgradeCenBandwidthPackageSpec response
func CreateTempUpgradeCenBandwidthPackageSpecResponse() (response *TempUpgradeCenBandwidthPackageSpecResponse) {
	response = &TempUpgradeCenBandwidthPackageSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
