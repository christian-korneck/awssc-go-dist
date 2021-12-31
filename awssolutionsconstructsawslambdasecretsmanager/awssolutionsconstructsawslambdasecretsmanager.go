// CDK constructs for defining an interaction between an AWS Lambda function and AWS Secrets Manager.
package awssolutionsconstructsawslambdasecretsmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasecretsmanager/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdasecretsmanager/v2/internal"
)

type LambdaToSecretsmanager interface {
	constructs.Construct
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	Secret() awssecretsmanager.Secret
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToSecretsmanager
type jsiiProxy_LambdaToSecretsmanager struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToSecretsmanager) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSecretsmanager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSecretsmanager) Secret() awssecretsmanager.Secret {
	var returns awssecretsmanager.Secret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToSecretsmanager) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToSecretsmanager(scope constructs.Construct, id *string, props *LambdaToSecretsmanagerProps) LambdaToSecretsmanager {
	_init_.Initialize()

	j := jsiiProxy_LambdaToSecretsmanager{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-secretsmanager.LambdaToSecretsmanager",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToSecretsmanager_Override(l LambdaToSecretsmanager, scope constructs.Construct, id *string, props *LambdaToSecretsmanagerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-secretsmanager.LambdaToSecretsmanager",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToSecretsmanager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-secretsmanager.LambdaToSecretsmanager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToSecretsmanager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToSecretsmanagerProps struct {
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of Secret object, providing both this and secretProps will cause an error.
	ExistingSecretObj awssecretsmanager.Secret `json:"existingSecretObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// Optional secret permissions to grant to the Lambda function.
	//
	// One of the following may be specified: "Read" or "ReadWrite".
	GrantWriteAccess *string `json:"grantWriteAccess"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional Name for the Secret environment variable set for the Lambda function.
	SecretEnvironmentVariableName *string `json:"secretEnvironmentVariableName"`
	// Optional user-provided props to override the default props for the Secret.
	SecretProps *awssecretsmanager.SecretProps `json:"secretProps"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

