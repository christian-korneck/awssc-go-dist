// CDK constructs for defining an interaction between an AWS Lambda function and an AWS Step Function.
package awssolutionsconstructsawslambdastepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdastepfunctions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdastepfunctions/v2/internal"
)

type LambdaToStepfunctions interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	StateMachine() awsstepfunctions.StateMachine
	StateMachineLogGroup() awslogs.ILogGroup
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToStepfunctions
type jsiiProxy_LambdaToStepfunctions struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToStepfunctions) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToStepfunctions) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToStepfunctions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToStepfunctions) StateMachine() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"stateMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToStepfunctions) StateMachineLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"stateMachineLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToStepfunctions) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToStepfunctions(scope constructs.Construct, id *string, props *LambdaToStepfunctionsProps) LambdaToStepfunctions {
	_init_.Initialize()

	j := jsiiProxy_LambdaToStepfunctions{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-stepfunctions.LambdaToStepfunctions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToStepfunctions_Override(l LambdaToStepfunctions, scope constructs.Construct, id *string, props *LambdaToStepfunctionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-stepfunctions.LambdaToStepfunctions",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToStepfunctions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-stepfunctions.LambdaToStepfunctions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToStepfunctions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToStepfunctionsProps struct {
	// User provided StateMachineProps to override the defaults.
	StateMachineProps *awsstepfunctions.StateMachineProps `json:"stateMachineProps"`
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
	// Optional Name for the Step Functions state machine environment variable set for the producer Lambda function.
	StateMachineEnvironmentVariableName *string `json:"stateMachineEnvironmentVariableName"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

