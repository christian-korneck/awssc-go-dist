// CDK Constructs for AWS Cloudfront to AWS API Gateway to AWS Lambda integration.
package awssolutionsconstructsawscloudfrontapigatewaylambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfrontapigatewaylambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfrontapigatewaylambda/v2/internal"
)

type CloudFrontToApiGatewayToLambda interface {
	constructs.Construct
	ApiGateway() awsapigateway.RestApi
	ApiGatewayCloudWatchRole() awsiam.Role
	ApiGatewayLogGroup() awslogs.LogGroup
	CloudFrontFunction() awscloudfront.Function
	CloudFrontLoggingBucket() awss3.Bucket
	CloudFrontWebDistribution() awscloudfront.Distribution
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for CloudFrontToApiGatewayToLambda
type jsiiProxy_CloudFrontToApiGatewayToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CloudFrontToApiGatewayToLambda) ApiGateway() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGatewayToLambda) ApiGatewayCloudWatchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGatewayToLambda) ApiGatewayLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"apiGatewayLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGatewayToLambda) CloudFrontFunction() awscloudfront.Function {
	var returns awscloudfront.Function
	_jsii_.Get(
		j,
		"cloudFrontFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGatewayToLambda) CloudFrontLoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"cloudFrontLoggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGatewayToLambda) CloudFrontWebDistribution() awscloudfront.Distribution {
	var returns awscloudfront.Distribution
	_jsii_.Get(
		j,
		"cloudFrontWebDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGatewayToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToApiGatewayToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCloudFrontToApiGatewayToLambda(scope constructs.Construct, id *string, props *CloudFrontToApiGatewayToLambdaProps) CloudFrontToApiGatewayToLambda {
	_init_.Initialize()

	j := jsiiProxy_CloudFrontToApiGatewayToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cloudfront-apigateway-lambda.CloudFrontToApiGatewayToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCloudFrontToApiGatewayToLambda_Override(c CloudFrontToApiGatewayToLambda, scope constructs.Construct, id *string, props *CloudFrontToApiGatewayToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cloudfront-apigateway-lambda.CloudFrontToApiGatewayToLambda",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudFrontToApiGatewayToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-cloudfront-apigateway-lambda.CloudFrontToApiGatewayToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CloudFrontToApiGatewayToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudFrontToApiGatewayToLambdaProps struct {
	// Optional user provided props to override the default props for the API Gateway.
	ApiGatewayProps interface{} `json:"apiGatewayProps"`
	// Optional user provided props to override the default props.
	CloudFrontDistributionProps interface{} `json:"cloudFrontDistributionProps"`
	// Optional user provided props to override the default props for the CloudFront Logging Bucket.
	CloudFrontLoggingBucketProps *awss3.BucketProps `json:"cloudFrontLoggingBucketProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Optional user provided props to turn on/off the automatic injection of best practice HTTP security headers in all responses from cloudfront.
	InsertHttpSecurityHeaders *bool `json:"insertHttpSecurityHeaders"`
	// Optional user provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional user provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
}

