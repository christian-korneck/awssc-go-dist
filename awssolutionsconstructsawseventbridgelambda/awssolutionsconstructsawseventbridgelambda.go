// CDK Constructs for deploying AWS Events Rule that inveokes AWS Lambda
package awssolutionsconstructsawseventbridgelambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgelambda/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgelambda/v2/internal"
)

type EventbridgeToLambda interface {
	constructs.Construct
	EventBus() awsevents.IEventBus
	EventsRule() awsevents.Rule
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for EventbridgeToLambda
type jsiiProxy_EventbridgeToLambda struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EventbridgeToLambda) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToLambda) EventsRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToLambda) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEventbridgeToLambda(scope constructs.Construct, id *string, props *EventbridgeToLambdaProps) EventbridgeToLambda {
	_init_.Initialize()

	j := jsiiProxy_EventbridgeToLambda{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-lambda.EventbridgeToLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventbridgeToLambda_Override(e EventbridgeToLambda, scope constructs.Construct, id *string, props *EventbridgeToLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-lambda.EventbridgeToLambda",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventbridgeToLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-eventbridge-lambda.EventbridgeToLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventbridgeToLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventbridgeToLambdaProps struct {
	// User provided eventRuleProps to override the defaults.
	EventRuleProps *awsevents.RuleProps `json:"eventRuleProps"`
	// A new custom EventBus is created with provided props.
	EventBusProps *awsevents.EventBusProps `json:"eventBusProps"`
	// Existing instance of a custom EventBus.
	ExistingEventBusInterface awsevents.IEventBus `json:"existingEventBusInterface"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
}

