// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeDBLogFilesInput struct {
	_ struct{} `type:"structure"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	FileLastWritten *int64 `locationName:"FileLastWritten" type:"integer"`

	FileSize *int64 `locationName:"FileSize" type:"integer"`

	FilenameContains *string `locationName:"FilenameContains" type:"string"`

	Marker *string `locationName:"Marker" type:"string"`

	MaxRecords *int64 `locationName:"MaxRecords" type:"integer"`
}

// String returns the string representation
func (s DescribeDBLogFilesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeDBLogFilesOutput struct {
	_ struct{} `type:"structure"`

	DescribeDBLogFiles []DescribeDBLogFilesDetails `locationNameList:"DescribeDBLogFilesDetails" type:"list"`

	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBLogFilesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeDBLogFiles = "DescribeDBLogFiles"

// DescribeDBLogFilesRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DescribeDBLogFilesRequest.
//    req := client.DescribeDBLogFilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/DescribeDBLogFiles
func (c *Client) DescribeDBLogFilesRequest(input *DescribeDBLogFilesInput) DescribeDBLogFilesRequest {
	op := &aws.Operation{
		Name:       opDescribeDBLogFiles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBLogFilesInput{}
	}

	req := c.newRequest(op, input, &DescribeDBLogFilesOutput{})
	return DescribeDBLogFilesRequest{Request: req, Input: input, Copy: c.DescribeDBLogFilesRequest}
}

// DescribeDBLogFilesRequest is the request type for the
// DescribeDBLogFiles API operation.
type DescribeDBLogFilesRequest struct {
	*aws.Request
	Input *DescribeDBLogFilesInput
	Copy  func(*DescribeDBLogFilesInput) DescribeDBLogFilesRequest
}

// Send marshals and sends the DescribeDBLogFiles API request.
func (r DescribeDBLogFilesRequest) Send(ctx context.Context) (*DescribeDBLogFilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBLogFilesResponse{
		DescribeDBLogFilesOutput: r.Request.Data.(*DescribeDBLogFilesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDBLogFilesResponse is the response type for the
// DescribeDBLogFiles API operation.
type DescribeDBLogFilesResponse struct {
	*DescribeDBLogFilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBLogFiles request.
func (r *DescribeDBLogFilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
