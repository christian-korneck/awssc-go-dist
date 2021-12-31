// CDK Constructs for deploying AWS Events Rule that invokes AWS Step Functions
package awssolutionsconstructsawseventbridgestepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgestepfunctions/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgestepfunctions/v2/internal"
)

type EventbridgeToStepfunctions interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	EventBus() awsevents.IEventBus
	EventsRule() awsevents.Rule
	Node() constructs.Node
	StateMachine() awsstepfunctions.StateMachine
	StateMachineLogGroup() awslogs.ILogGroup
	ToString() *string
}

// The jsii proxy struct for EventbridgeToStepfunctions
type jsiiProxy_EventbridgeToStepfunctions struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EventbridgeToStepfunctions) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToStepfunctions) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToStepfunctions) EventsRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToStepfunctions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToStepfunctions) StateMachine() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"stateMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToStepfunctions) StateMachineLogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"stateMachineLogGroup",
		&returns,
	)
	return returns
}


func NewEventbridgeToStepfunctions(scope constructs.Construct, id *string, props *EventbridgeToStepfunctionsProps) EventbridgeToStepfunctions {
	_init_.Initialize()

	j := jsiiProxy_EventbridgeToStepfunctions{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-stepfunctions.EventbridgeToStepfunctions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventbridgeToStepfunctions_Override(e EventbridgeToStepfunctions, scope constructs.Construct, id *string, props *EventbridgeToStepfunctionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-stepfunctions.EventbridgeToStepfunctions",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventbridgeToStepfunctions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-eventbridge-stepfunctions.EventbridgeToStepfunctions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventbridgeToStepfunctions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventbridgeToStepfunctionsProps struct {
	// User provided eventRuleProps to override the defaults.
	EventRuleProps *awsevents.RuleProps `json:"eventRuleProps"`
	// User provided StateMachineProps to override the defaults.
	StateMachineProps *awsstepfunctions.StateMachineProps `json:"stateMachineProps"`
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// A new custom EventBus is created with provided props.
	EventBusProps *awsevents.EventBusProps `json:"eventBusProps"`
	// Existing instance of a custom EventBus.
	ExistingEventBusInterface awsevents.IEventBus `json:"existingEventBusInterface"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
}

