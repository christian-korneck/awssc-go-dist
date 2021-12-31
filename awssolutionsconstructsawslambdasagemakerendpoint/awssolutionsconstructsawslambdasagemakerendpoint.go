// CDK constructs for defining an interaction between an AWS Lambda function and an Amazon SageMaker inference endpoint.
package awssolutionsconstructsawslambdasagemakerendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasagemakerendpoint/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasagemakerendpoint/v2/internal"
)

type LambdaToSagemakerEndpoint interface {
	constructs.Construct
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	SagemakerEndpoint() awssagemaker.CfnEndpoint
	SagemakerEndpointConfig() awssagemaker.CfnEndpointConfig
	SagemakerModel() awssagemaker.CfnModel
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToSagemakerEndpoint
type jsiiProxy_LambdaToSagemakerEndpoint struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToSagemakerEndpoint) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSagemakerEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSagemakerEndpoint) SagemakerEndpoint() awssagemaker.CfnEndpoint {
	var returns awssagemaker.CfnEndpoint
	_jsii_.Get(
		j,
		"sagemakerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSagemakerEndpoint) SagemakerEndpointConfig() awssagemaker.CfnEndpointConfig {
	var returns awssagemaker.CfnEndpointConfig
	_jsii_.Get(
		j,
		"sagemakerEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSagemakerEndpoint) SagemakerModel() awssagemaker.CfnModel {
	var returns awssagemaker.CfnModel
	_jsii_.Get(
		j,
		"sagemakerModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSagemakerEndpoint) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToSagemakerEndpoint(scope constructs.Construct, id *string, props *LambdaToSagemakerEndpointProps) LambdaToSagemakerEndpoint {
	_init_.Initialize()

	j := jsiiProxy_LambdaToSagemakerEndpoint{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-sagemakerendpoint.LambdaToSagemakerEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToSagemakerEndpoint_Override(l LambdaToSagemakerEndpoint, scope constructs.Construct, id *string, props *LambdaToSagemakerEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-sagemakerendpoint.LambdaToSagemakerEndpoint",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToSagemakerEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-sagemakerendpoint.LambdaToSagemakerEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToSagemakerEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToSagemakerEndpointProps struct {
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// User provided props to create Sagemaker Endpoint Configuration.
	EndpointConfigProps *awssagemaker.CfnEndpointConfigProps `json:"endpointConfigProps"`
	// User provided props to create Sagemaker Endpoint.
	EndpointProps *awssagemaker.CfnEndpointProps `json:"endpointProps"`
	// Existing instance of Lambda Function object, Providing both this and lambdaFunctionProps will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing Sagemaker Enpoint object, providing both this and endpointProps will cause an error.
	ExistingSagemakerEndpointObj awssagemaker.CfnEndpoint `json:"existingSagemakerEndpointObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// User provided props to create Sagemaker Model.
	ModelProps interface{} `json:"modelProps"`
	// Optional Name for the SageMaker endpoint environment variable set for the Lambda function.
	SagemakerEnvironmentVariableName *string `json:"sagemakerEnvironmentVariableName"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

