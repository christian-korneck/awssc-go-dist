// CDK Constructs for AWS IoT to AWS Kinesis Data Stream.
package awssolutionsconstructsawsiotkinesisstreams

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotkinesisstreams/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotkinesisstreams/v2/internal"
)

type IotToKinesisStreams interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	IotActionsRole() awsiam.Role
	IotTopicRule() awsiot.CfnTopicRule
	KinesisStream() awskinesis.Stream
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for IotToKinesisStreams
type jsiiProxy_IotToKinesisStreams struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IotToKinesisStreams) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisStreams) IotActionsRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"iotActionsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisStreams) IotTopicRule() awsiot.CfnTopicRule {
	var returns awsiot.CfnTopicRule
	_jsii_.Get(
		j,
		"iotTopicRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisStreams) KinesisStream() awskinesis.Stream {
	var returns awskinesis.Stream
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToKinesisStreams) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewIotToKinesisStreams(scope constructs.Construct, id *string, props *IotToKinesisStreamsProps) IotToKinesisStreams {
	_init_.Initialize()

	j := jsiiProxy_IotToKinesisStreams{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-kinesisstreams.IotToKinesisStreams",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIotToKinesisStreams_Override(i IotToKinesisStreams, scope constructs.Construct, id *string, props *IotToKinesisStreamsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-kinesisstreams.IotToKinesisStreams",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotToKinesisStreams_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-iot-kinesisstreams.IotToKinesisStreams",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotToKinesisStreams) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotToKinesisStreamsProps struct {
	// User provided CfnTopicRuleProps to override the defaults.
	IotTopicRuleProps *awsiot.CfnTopicRuleProps `json:"iotTopicRuleProps"`
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// Existing instance of Kinesis Stream object, providing both this and KinesisStreamProps will cause an error.
	ExistingStreamObj awskinesis.Stream `json:"existingStreamObj"`
	// User provided props to override the default props for the Kinesis Stream.
	KinesisStreamProps interface{} `json:"kinesisStreamProps"`
}

