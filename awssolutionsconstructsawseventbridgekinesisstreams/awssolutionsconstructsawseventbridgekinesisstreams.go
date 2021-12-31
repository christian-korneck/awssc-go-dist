// CDK Constructs for deploying Amazon CloudWatch Events Rule that invokes Amazon Kinesis Data Stream
package awssolutionsconstructsawseventbridgekinesisstreams

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgekinesisstreams/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgekinesisstreams/v2/internal"
)

type EventbridgeToKinesisStreams interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	EventBus() awsevents.IEventBus
	EventsRole() awsiam.Role
	EventsRule() awsevents.Rule
	KinesisStream() awskinesis.Stream
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for EventbridgeToKinesisStreams
type jsiiProxy_EventbridgeToKinesisStreams struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EventbridgeToKinesisStreams) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisStreams) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisStreams) EventsRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"eventsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisStreams) EventsRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisStreams) KinesisStream() awskinesis.Stream {
	var returns awskinesis.Stream
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToKinesisStreams) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewEventbridgeToKinesisStreams(scope constructs.Construct, id *string, props *EventbridgeToKinesisStreamsProps) EventbridgeToKinesisStreams {
	_init_.Initialize()

	j := jsiiProxy_EventbridgeToKinesisStreams{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-kinesisstreams.EventbridgeToKinesisStreams",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventbridgeToKinesisStreams_Override(e EventbridgeToKinesisStreams, scope constructs.Construct, id *string, props *EventbridgeToKinesisStreamsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-kinesisstreams.EventbridgeToKinesisStreams",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventbridgeToKinesisStreams_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-eventbridge-kinesisstreams.EventbridgeToKinesisStreams",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventbridgeToKinesisStreams) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventbridgeToKinesisStreamsProps struct {
	// User provided eventRuleProps to override the defaults.
	EventRuleProps *awsevents.RuleProps `json:"eventRuleProps"`
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// A new custom EventBus is created with provided props.
	EventBusProps *awsevents.EventBusProps `json:"eventBusProps"`
	// Existing instance of a custom EventBus.
	ExistingEventBusInterface awsevents.IEventBus `json:"existingEventBusInterface"`
	// Existing instance of Kinesis Stream object, providing both this and KinesisStreamProps will cause an error.
	ExistingStreamObj awskinesis.Stream `json:"existingStreamObj"`
	// User provided props to override the default props for the Kinesis Stream.
	KinesisStreamProps interface{} `json:"kinesisStreamProps"`
}

