// CDK Constructs for AWS Cloudfront to AWS API Gateway integration.
package awssolutionsconstructsawscloudfrontapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfrontapigateway/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfrontapigateway/v2/internal"
)

type CloudFrontToApiGateway interface {
	constructs.Construct
	ApiGateway() awsapigateway.RestApi
	CloudFrontFunction() awscloudfront.Function
	CloudFrontLoggingBucket() awss3.Bucket
	CloudFrontWebDistribution() awscloudfront.Distribution
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for CloudFrontToApiGateway
type jsiiProxy_CloudFrontToApiGateway struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CloudFrontToApiGateway) ApiGateway() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGateway) CloudFrontFunction() awscloudfront.Function {
	var returns awscloudfront.Function
	_jsii_.Get(
		j,
		"cloudFrontFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGateway) CloudFrontLoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"cloudFrontLoggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGateway) CloudFrontWebDistribution() awscloudfront.Distribution {
	var returns awscloudfront.Distribution
	_jsii_.Get(
		j,
		"cloudFrontWebDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCloudFrontToApiGateway(scope constructs.Construct, id *string, props *CloudFrontToApiGatewayProps) CloudFrontToApiGateway {
	_init_.Initialize()

	j := jsiiProxy_CloudFrontToApiGateway{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cloudfront-apigateway.CloudFrontToApiGateway",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCloudFrontToApiGateway_Override(c CloudFrontToApiGateway, scope constructs.Construct, id *string, props *CloudFrontToApiGatewayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cloudfront-apigateway.CloudFrontToApiGateway",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudFrontToApiGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-cloudfront-apigateway.CloudFrontToApiGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CloudFrontToApiGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudFrontToApiGatewayProps struct {
	// Existing instance of api.RestApi object.
	ExistingApiGatewayObj awsapigateway.RestApi `json:"existingApiGatewayObj"`
	// Optional user provided props to override the default props.
	CloudFrontDistributionProps interface{} `json:"cloudFrontDistributionProps"`
	// Optional user provided props to override the default props for the CloudFront Logging Bucket.
	CloudFrontLoggingBucketProps *awss3.BucketProps `json:"cloudFrontLoggingBucketProps"`
	// Optional user provided props to turn on/off the automatic injection of best practice HTTP security headers in all responses from cloudfront.
	InsertHttpSecurityHeaders *bool `json:"insertHttpSecurityHeaders"`
}

