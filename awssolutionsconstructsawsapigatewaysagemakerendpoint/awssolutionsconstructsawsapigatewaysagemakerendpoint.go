// CDK Constructs for AWS API Gateway and Amazon SageMaker Endpoint integration.
package awssolutionsconstructsawsapigatewaysagemakerendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewaysagemakerendpoint/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewaysagemakerendpoint/v2/internal"
)

type ApiGatewayToSageMakerEndpoint interface {
	constructs.Construct
	ApiGateway() awsapigateway.RestApi
	ApiGatewayCloudWatchRole() awsiam.Role
	ApiGatewayLogGroup() awslogs.LogGroup
	ApiGatewayRole() awsiam.Role
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for ApiGatewayToSageMakerEndpoint
type jsiiProxy_ApiGatewayToSageMakerEndpoint struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApiGatewayToSageMakerEndpoint) ApiGateway() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSageMakerEndpoint) ApiGatewayCloudWatchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSageMakerEndpoint) ApiGatewayLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"apiGatewayLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSageMakerEndpoint) ApiGatewayRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSageMakerEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewApiGatewayToSageMakerEndpoint(scope constructs.Construct, id *string, props *ApiGatewayToSageMakerEndpointProps) ApiGatewayToSageMakerEndpoint {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayToSageMakerEndpoint{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-sagemakerendpoint.ApiGatewayToSageMakerEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApiGatewayToSageMakerEndpoint_Override(a ApiGatewayToSageMakerEndpoint, scope constructs.Construct, id *string, props *ApiGatewayToSageMakerEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-sagemakerendpoint.ApiGatewayToSageMakerEndpoint",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayToSageMakerEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-apigateway-sagemakerendpoint.ApiGatewayToSageMakerEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayToSageMakerEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiGatewayToSageMakerEndpointProps struct {
	// Name of the deployed SageMaker inference endpoint.
	EndpointName *string `json:"endpointName"`
	// Mapping template to convert GET requests received on the REST API to POST requests expected by the SageMaker endpoint.
	RequestMappingTemplate *string `json:"requestMappingTemplate"`
	// Resource path for the GET method.
	//
	// The variable defined here can be referenced in `requestMappingTemplate`.
	ResourcePath *string `json:"resourcePath"`
	// Optional IAM role that is used by API Gateway to invoke the SageMaker endpoint.
	ApiGatewayExecutionRole awsiam.Role `json:"apiGatewayExecutionRole"`
	// Optional user-provided props to override the default props for the API Gateway.
	ApiGatewayProps *awsapigateway.RestApiProps `json:"apiGatewayProps"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
	// Optional resource name where the GET method will be available.
	ResourceName *string `json:"resourceName"`
	// Optional mapping template to convert responses received from the SageMaker endpoint.
	ResponseMappingTemplate *string `json:"responseMappingTemplate"`
}

