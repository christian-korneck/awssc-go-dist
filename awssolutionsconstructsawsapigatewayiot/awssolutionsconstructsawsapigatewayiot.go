// CDK constructs to proxy communication to IotCore using a APIGateway(REST).
package awssolutionsconstructsawsapigatewayiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewayiot/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewayiot/v2/internal"
)

type ApiGatewayToIot interface {
	constructs.Construct
	ApiGateway() awsapigateway.RestApi
	ApiGatewayCloudWatchRole() awsiam.Role
	ApiGatewayLogGroup() awslogs.LogGroup
	ApiGatewayRole() awsiam.IRole
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for ApiGatewayToIot
type jsiiProxy_ApiGatewayToIot struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApiGatewayToIot) ApiGateway() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToIot) ApiGatewayCloudWatchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToIot) ApiGatewayLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"apiGatewayLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToIot) ApiGatewayRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"apiGatewayRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToIot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewApiGatewayToIot(scope constructs.Construct, id *string, props *ApiGatewayToIotProps) ApiGatewayToIot {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayToIot{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-iot.ApiGatewayToIot",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApiGatewayToIot_Override(a ApiGatewayToIot, scope constructs.Construct, id *string, props *ApiGatewayToIotProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-iot.ApiGatewayToIot",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayToIot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-apigateway-iot.ApiGatewayToIot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayToIot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The properties for the ApiGatewayIot class.
type ApiGatewayToIotProps struct {
	// The AWS IoT endpoint subdomain to integrate the API Gateway with (e.g ab123cdefghij4l-ats). Added as AWS Subdomain to the Integration Request.
	IotEndpoint *string `json:"iotEndpoint"`
	// Creates an api key and associates to usage plan if set to true.
	ApiGatewayCreateApiKey *bool `json:"apiGatewayCreateApiKey"`
	// The IAM role that is used by API Gateway to publish messages to IoT topics and Thing shadows.
	ApiGatewayExecutionRole awsiam.IRole `json:"apiGatewayExecutionRole"`
	// Optional user-provided props to override the default props for the API.
	ApiGatewayProps *awsapigateway.RestApiProps `json:"apiGatewayProps"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
}

