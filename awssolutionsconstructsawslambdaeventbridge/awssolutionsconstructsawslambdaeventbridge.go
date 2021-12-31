// CDK constructs for defining an interaction between an AWS Lambda function and an Amazon EventBridge.
package awssolutionsconstructsawslambdaeventbridge

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdaeventbridge/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdaeventbridge/v2/internal"
)

type LambdaToEventbridge interface {
	constructs.Construct
	EventBus() awsevents.IEventBus
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	Vpc() awsec2.IVpc
	ToString() *string
}

// The jsii proxy struct for LambdaToEventbridge
type jsiiProxy_LambdaToEventbridge struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToEventbridge) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToEventbridge) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToEventbridge) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToEventbridge) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewLambdaToEventbridge(scope constructs.Construct, id *string, props *LambdaToEventbridgeProps) LambdaToEventbridge {
	_init_.Initialize()

	j := jsiiProxy_LambdaToEventbridge{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-eventbridge.LambdaToEventbridge",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToEventbridge_Override(l LambdaToEventbridge, scope constructs.Construct, id *string, props *LambdaToEventbridgeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-eventbridge.LambdaToEventbridge",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToEventbridge_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-eventbridge.LambdaToEventbridge",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToEventbridge) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToEventbridgeProps struct {
	// Whether to deploy a new VPC.
	DeployVpc *bool `json:"deployVpc"`
	// Optional Name for the EventBus's name environment variable set for the Lambda function.
	EventBusEnvironmentVariableName *string `json:"eventBusEnvironmentVariableName"`
	// A new custom EventBus is created with provided props.
	EventBusProps *awsevents.EventBusProps `json:"eventBusProps"`
	// Existing instance of a custom EventBus.
	ExistingEventBusInterface awsevents.IEventBus `json:"existingEventBusInterface"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// An existing VPC for the construct to use (construct will NOT create a new VPC in this case).
	ExistingVpc awsec2.IVpc `json:"existingVpc"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Properties to override default properties if deployVpc is true.
	VpcProps *awsec2.VpcProps `json:"vpcProps"`
}

