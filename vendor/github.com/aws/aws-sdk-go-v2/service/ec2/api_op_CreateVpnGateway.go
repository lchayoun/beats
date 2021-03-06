// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for CreateVpnGateway.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpnGatewayRequest
type CreateVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	// A private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	// If you're using a 16-bit ASN, it must be in the 64512 to 65534 range. If
	// you're using a 32-bit ASN, it must be in the 4200000000 to 4294967294 range.
	//
	// Default: 64512
	AmazonSideAsn *int64 `type:"long"`

	// The Availability Zone for the virtual private gateway.
	AvailabilityZone *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The type of VPN connection this virtual private gateway supports.
	//
	// Type is a required field
	Type GatewayType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateVpnGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpnGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVpnGatewayInput"}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of CreateVpnGateway.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpnGatewayResult
type CreateVpnGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Information about the virtual private gateway.
	VpnGateway *VpnGateway `locationName:"vpnGateway" type:"structure"`
}

// String returns the string representation
func (s CreateVpnGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVpnGateway = "CreateVpnGateway"

// CreateVpnGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a virtual private gateway. A virtual private gateway is the endpoint
// on the VPC side of your VPN connection. You can create a virtual private
// gateway before creating the VPC itself.
//
// For more information, see AWS Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html)
// in the AWS Site-to-Site VPN User Guide.
//
//    // Example sending a request using CreateVpnGatewayRequest.
//    req := client.CreateVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpnGateway
func (c *Client) CreateVpnGatewayRequest(input *CreateVpnGatewayInput) CreateVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &CreateVpnGatewayOutput{})
	return CreateVpnGatewayRequest{Request: req, Input: input, Copy: c.CreateVpnGatewayRequest}
}

// CreateVpnGatewayRequest is the request type for the
// CreateVpnGateway API operation.
type CreateVpnGatewayRequest struct {
	*aws.Request
	Input *CreateVpnGatewayInput
	Copy  func(*CreateVpnGatewayInput) CreateVpnGatewayRequest
}

// Send marshals and sends the CreateVpnGateway API request.
func (r CreateVpnGatewayRequest) Send(ctx context.Context) (*CreateVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpnGatewayResponse{
		CreateVpnGatewayOutput: r.Request.Data.(*CreateVpnGatewayOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpnGatewayResponse is the response type for the
// CreateVpnGateway API operation.
type CreateVpnGatewayResponse struct {
	*CreateVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpnGateway request.
func (r *CreateVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
