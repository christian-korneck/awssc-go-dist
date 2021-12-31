// CDK Constructs for AWS Cognito to AWS API Gateway to AWS Lambda integration
package awssolutionsconstructsawscognitoapigatewaylambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscognitoapigatewaylambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscognitoapigatewaylambda/v2/internal"
)

type CognitoToApiGatewayToLambda interface {
	constructs.Construct
	ApiGateway() awsapigateway.RestApi
	ApiGatewayAuthorizer() awsapigateway.CfnAuthorizer
	ApiGatewayCloudWatchRole() awsiam.Role
	ApiGatewayLogGroup() awslogs.LogGroup
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	UserPool() awscognito.UserPool
	UserPoolClient() awscognito.UserPoolClient
	AddAuthorizers()
	ToString() *string
}

// The jsii proxy struct for CognitoToApiGatewayToLambda
type jsiiProxy_CognitoToApiGatewayToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CognitoToApiGatewayToLambda) ApiGateway() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoToApiGatewayToLambda) ApiGatewayAuthorizer() awsapigateway.CfnAuthorizer {
	var returns awsapigateway.CfnAuthorizer
	_jsii_.Get(
		j,
		"apiGatewayAuthorizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoToApiGatewayToLambda) ApiGatewayCloudWatchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoToApiGatewayToLambda) ApiGatewayLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"apiGatewayLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoToApiGatewayToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoToApiGatewayToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoToApiGatewayToLambda) UserPool() awscognito.UserPool {
	var returns awscognito.UserPool
	_jsii_.Get(
		j,
		"userPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoToApiGatewayToLambda) UserPoolClient() awscognito.UserPoolClient {
	var returns awscognito.UserPoolClient
	_jsii_.Get(
		j,
		"userPoolClient",
		&returns,
	)
	return returns
}


func NewCognitoToApiGatewayToLambda(scope constructs.Construct, id *string, props *CognitoToApiGatewayToLambdaProps) CognitoToApiGatewayToLambda {
	_init_.Initialize()

	j := jsiiProxy_CognitoToApiGatewayToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cognito-apigateway-lambda.CognitoToApiGatewayToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCognitoToApiGatewayToLambda_Override(c CognitoToApiGatewayToLambda, scope constructs.Construct, id *string, props *CognitoToApiGatewayToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cognito-apigateway-lambda.CognitoToApiGatewayToLambda",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CognitoToApiGatewayToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-cognito-apigateway-lambda.CognitoToApiGatewayToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoToApiGatewayToLambda) AddAuthorizers() {
	_jsii_.InvokeVoid(
		c,
		"addAuthorizers",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CognitoToApiGatewayToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CognitoToApiGatewayToLambdaProps struct {
	// Optional user provided props to override the default props for the API Gateway.
	ApiGatewayProps interface{} `json:"apiGatewayProps"`
	// Optional user provided props to override the default props.
	CognitoUserPoolClientProps interface{} `json:"cognitoUserPoolClientProps"`
	// Optional user provided props to override the default props.
	CognitoUserPoolProps *awscognito.UserPoolProps `json:"cognitoUserPoolProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
}

