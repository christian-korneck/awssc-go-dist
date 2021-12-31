// CDK constructs for defining an interaction between an AWS Lambda function and AWS Systems Manager Parameter Store String parameter
package awssolutionsconstructsawslambdassmstringparameter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdassmstringparameter/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdassmstringparameter/v2/internal"
)

type LambdaToSsmstringparameter interface {
	constructs.Construct
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	StringParameter() awsssm.StringParameter
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToSsmstringparameter
type jsiiProxy_LambdaToSsmstringparameter struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToSsmstringparameter) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSsmstringparameter) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSsmstringparameter) StringParameter() awsssm.StringParameter {
	var returns awsssm.StringParameter
	_jsii_.Get(
		j,
		"stringParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSsmstringparameter) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToSsmstringparameter(scope constructs.Construct, id *string, props *LambdaToSsmstringparameterProps) LambdaToSsmstringparameter {
	_init_.Initialize()

	j := jsiiProxy_LambdaToSsmstringparameter{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-ssmstringparameter.LambdaToSsmstringparameter",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToSsmstringparameter_Override(l LambdaToSsmstringparameter, scope constructs.Construct, id *string, props *LambdaToSsmstringparameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-ssmstringparameter.LambdaToSsmstringparameter",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToSsmstringparameter_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-ssmstringparameter.LambdaToSsmstringparameter",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToSsmstringparameter) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToSsmstringparameterProps struct {
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// Existing instance of Lambda Function object, if this is set then the lambdaFunctionProps is ignored.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of SSM String parameter object, If this is set then the stringParameterProps is ignored.
	ExistingStringParameterObj awsssm.StringParameter `json:"existingStringParameterObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional Name for the SSM String parameter environment variable set for the Lambda function.
	StringParameterEnvironmentVariableName *string `json:"stringParameterEnvironmentVariableName"`
	// Optional SSM String parameter permissions to grant to the Lambda function.
	//
	// One of the following may be specified: "Read", "ReadWrite".
	StringParameterPermissions *string `json:"stringParameterPermissions"`
	// Optional user provided props to override the default props for SSM String parameter.
	//
	// If existingStringParameterObj
	// is not set stringParameterProps is required. The only supported string parameter type is ParameterType.STRING.
	StringParameterProps *awsssm.StringParameterProps `json:"stringParameterProps"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

