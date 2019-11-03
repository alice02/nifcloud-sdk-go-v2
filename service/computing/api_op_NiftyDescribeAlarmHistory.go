// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDescribeAlarmHistoryInput struct {
	_ struct{} `type:"structure"`

	Rule []RequestRuleStruct `locationName:"Rule" type:"list"`
}

// String returns the string representation
func (s NiftyDescribeAlarmHistoryInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDescribeAlarmHistoryOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	ReservationSet []ReservationSetItem `locationName:"reservationSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s NiftyDescribeAlarmHistoryOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDescribeAlarmHistory = "NiftyDescribeAlarmHistory"

// NiftyDescribeAlarmHistoryRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDescribeAlarmHistoryRequest.
//    req := client.NiftyDescribeAlarmHistoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDescribeAlarmHistory
func (c *Client) NiftyDescribeAlarmHistoryRequest(input *NiftyDescribeAlarmHistoryInput) NiftyDescribeAlarmHistoryRequest {
	op := &aws.Operation{
		Name:       opNiftyDescribeAlarmHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDescribeAlarmHistoryInput{}
	}

	req := c.newRequest(op, input, &NiftyDescribeAlarmHistoryOutput{})
	return NiftyDescribeAlarmHistoryRequest{Request: req, Input: input, Copy: c.NiftyDescribeAlarmHistoryRequest}
}

// NiftyDescribeAlarmHistoryRequest is the request type for the
// NiftyDescribeAlarmHistory API operation.
type NiftyDescribeAlarmHistoryRequest struct {
	*aws.Request
	Input *NiftyDescribeAlarmHistoryInput
	Copy  func(*NiftyDescribeAlarmHistoryInput) NiftyDescribeAlarmHistoryRequest
}

// Send marshals and sends the NiftyDescribeAlarmHistory API request.
func (r NiftyDescribeAlarmHistoryRequest) Send(ctx context.Context) (*NiftyDescribeAlarmHistoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDescribeAlarmHistoryResponse{
		NiftyDescribeAlarmHistoryOutput: r.Request.Data.(*NiftyDescribeAlarmHistoryOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDescribeAlarmHistoryResponse is the response type for the
// NiftyDescribeAlarmHistory API operation.
type NiftyDescribeAlarmHistoryResponse struct {
	*NiftyDescribeAlarmHistoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDescribeAlarmHistory request.
func (r *NiftyDescribeAlarmHistoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
