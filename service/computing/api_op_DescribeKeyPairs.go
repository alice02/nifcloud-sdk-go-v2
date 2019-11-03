// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeKeyPairsInput struct {
	_ struct{} `type:"structure"`

	KeyName []string `locationName:"KeyName" type:"list"`
}

// String returns the string representation
func (s DescribeKeyPairsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeKeyPairsOutput struct {
	_ struct{} `type:"structure"`

	KeySet []KeySetItem `locationName:"keySet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s DescribeKeyPairsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeKeyPairs = "DescribeKeyPairs"

// DescribeKeyPairsRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeKeyPairsRequest.
//    req := client.DescribeKeyPairsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeKeyPairs
func (c *Client) DescribeKeyPairsRequest(input *DescribeKeyPairsInput) DescribeKeyPairsRequest {
	op := &aws.Operation{
		Name:       opDescribeKeyPairs,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeKeyPairsInput{}
	}

	req := c.newRequest(op, input, &DescribeKeyPairsOutput{})
	return DescribeKeyPairsRequest{Request: req, Input: input, Copy: c.DescribeKeyPairsRequest}
}

// DescribeKeyPairsRequest is the request type for the
// DescribeKeyPairs API operation.
type DescribeKeyPairsRequest struct {
	*aws.Request
	Input *DescribeKeyPairsInput
	Copy  func(*DescribeKeyPairsInput) DescribeKeyPairsRequest
}

// Send marshals and sends the DescribeKeyPairs API request.
func (r DescribeKeyPairsRequest) Send(ctx context.Context) (*DescribeKeyPairsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeKeyPairsResponse{
		DescribeKeyPairsOutput: r.Request.Data.(*DescribeKeyPairsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeKeyPairsResponse is the response type for the
// DescribeKeyPairs API operation.
type DescribeKeyPairsResponse struct {
	*DescribeKeyPairsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeKeyPairs request.
func (r *DescribeKeyPairsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
