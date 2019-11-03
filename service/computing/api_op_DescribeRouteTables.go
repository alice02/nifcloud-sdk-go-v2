// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeRouteTablesInput struct {
	_ struct{} `type:"structure"`

	Filter []RequestFilterStruct `locationName:"Filter" type:"list"`

	RouteTableId []string `locationName:"RouteTableId" type:"list"`
}

// String returns the string representation
func (s DescribeRouteTablesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeRouteTablesOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	RouteTableSet []RouteTableSetItem `locationName:"routeTableSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeRouteTablesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeRouteTables = "DescribeRouteTables"

// DescribeRouteTablesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeRouteTablesRequest.
//    req := client.DescribeRouteTablesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeRouteTables
func (c *Client) DescribeRouteTablesRequest(input *DescribeRouteTablesInput) DescribeRouteTablesRequest {
	op := &aws.Operation{
		Name:       opDescribeRouteTables,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeRouteTablesInput{}
	}

	req := c.newRequest(op, input, &DescribeRouteTablesOutput{})
	return DescribeRouteTablesRequest{Request: req, Input: input, Copy: c.DescribeRouteTablesRequest}
}

// DescribeRouteTablesRequest is the request type for the
// DescribeRouteTables API operation.
type DescribeRouteTablesRequest struct {
	*aws.Request
	Input *DescribeRouteTablesInput
	Copy  func(*DescribeRouteTablesInput) DescribeRouteTablesRequest
}

// Send marshals and sends the DescribeRouteTables API request.
func (r DescribeRouteTablesRequest) Send(ctx context.Context) (*DescribeRouteTablesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRouteTablesResponse{
		DescribeRouteTablesOutput: r.Request.Data.(*DescribeRouteTablesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRouteTablesResponse is the response type for the
// DescribeRouteTables API operation.
type DescribeRouteTablesResponse struct {
	*DescribeRouteTablesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRouteTables request.
func (r *DescribeRouteTablesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
