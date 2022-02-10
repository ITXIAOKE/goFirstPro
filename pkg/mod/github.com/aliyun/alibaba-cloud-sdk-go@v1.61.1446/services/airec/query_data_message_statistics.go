package airec

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

// QueryDataMessageStatistics invokes the airec.QueryDataMessageStatistics API synchronously
func (client *Client) QueryDataMessageStatistics(request *QueryDataMessageStatisticsRequest) (response *QueryDataMessageStatisticsResponse, err error) {
	response = CreateQueryDataMessageStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDataMessageStatisticsWithChan invokes the airec.QueryDataMessageStatistics API asynchronously
func (client *Client) QueryDataMessageStatisticsWithChan(request *QueryDataMessageStatisticsRequest) (<-chan *QueryDataMessageStatisticsResponse, <-chan error) {
	responseChan := make(chan *QueryDataMessageStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDataMessageStatistics(request)
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

// QueryDataMessageStatisticsWithCallback invokes the airec.QueryDataMessageStatistics API asynchronously
func (client *Client) QueryDataMessageStatisticsWithCallback(request *QueryDataMessageStatisticsRequest, callback func(response *QueryDataMessageStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDataMessageStatisticsResponse
		var err error
		defer close(result)
		response, err = client.QueryDataMessageStatistics(request)
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

// QueryDataMessageStatisticsRequest is the request struct for api QueryDataMessageStatistics
type QueryDataMessageStatisticsRequest struct {
	*requests.RoaRequest
	TraceId       string           `position:"Query" name:"traceId"`
	MessageSource string           `position:"Query" name:"messageSource"`
	EndTime       requests.Integer `position:"Query" name:"endTime"`
	UserType      string           `position:"Query" name:"userType"`
	StartTime     requests.Integer `position:"Query" name:"startTime"`
	UserId        string           `position:"Query" name:"userId"`
	ItemId        string           `position:"Query" name:"itemId"`
	InstanceId    string           `position:"Path" name:"instanceId"`
	ItemType      string           `position:"Query" name:"itemType"`
	CmdType       string           `position:"Query" name:"cmdType"`
	SceneId       string           `position:"Query" name:"sceneId"`
	BhvType       string           `position:"Query" name:"bhvType"`
	Table         string           `position:"Path" name:"table"`
}

// QueryDataMessageStatisticsResponse is the response struct for api QueryDataMessageStatistics
type QueryDataMessageStatisticsResponse struct {
	*responses.BaseResponse
	Code      string                 `json:"code" xml:"code"`
	Message   string                 `json:"message" xml:"message"`
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateQueryDataMessageStatisticsRequest creates a request to invoke QueryDataMessageStatistics API
func CreateQueryDataMessageStatisticsRequest() (request *QueryDataMessageStatisticsRequest) {
	request = &QueryDataMessageStatisticsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "QueryDataMessageStatistics", "/v2/openapi/instances/[instanceId]/tables/[table]/data-message-statistics", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryDataMessageStatisticsResponse creates a response to parse from QueryDataMessageStatistics response
func CreateQueryDataMessageStatisticsResponse() (response *QueryDataMessageStatisticsResponse) {
	response = &QueryDataMessageStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}