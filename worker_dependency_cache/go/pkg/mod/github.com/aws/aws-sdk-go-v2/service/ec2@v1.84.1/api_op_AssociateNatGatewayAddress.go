// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates Elastic IP addresses (EIPs) and private IPv4 addresses with a public
// NAT gateway. For more information, see Work with NAT gateways
// (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-working-with)
// in the Amazon Virtual Private Cloud User Guide. By default, you can associate up
// to 2 Elastic IP addresses per public NAT gateway. You can increase the limit by
// requesting a quota adjustment. For more information, see Elastic IP address
// quotas
// (https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html#vpc-limits-eips)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) AssociateNatGatewayAddress(ctx context.Context, params *AssociateNatGatewayAddressInput, optFns ...func(*Options)) (*AssociateNatGatewayAddressOutput, error) {
	if params == nil {
		params = &AssociateNatGatewayAddressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateNatGatewayAddress", params, optFns, c.addOperationAssociateNatGatewayAddressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateNatGatewayAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateNatGatewayAddressInput struct {

	// The allocation IDs of EIPs that you want to associate with your NAT gateway.
	//
	// This member is required.
	AllocationIds []string

	// The NAT gateway ID.
	//
	// This member is required.
	NatGatewayId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The private IPv4 addresses that you want to assign to the NAT gateway.
	PrivateIpAddresses []string

	noSmithyDocumentSerde
}

type AssociateNatGatewayAddressOutput struct {

	// The IP addresses.
	NatGatewayAddresses []types.NatGatewayAddress

	// The NAT gateway ID.
	NatGatewayId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateNatGatewayAddressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAssociateNatGatewayAddress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAssociateNatGatewayAddress{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateNatGatewayAddressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateNatGatewayAddress(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateNatGatewayAddress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AssociateNatGatewayAddress",
	}
}