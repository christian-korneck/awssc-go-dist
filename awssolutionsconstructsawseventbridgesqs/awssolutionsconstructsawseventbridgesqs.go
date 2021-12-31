// CDK Constructs for deploying AWS Eventbridge that invokes AWS SQS
package awssolutionsconstructsawseventbridgesqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgesqs/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgesqs/v2/internal"
)

type EventbridgeToSqs interface {
	constructs.Construct
	DeadLetterQueue() *awssqs.DeadLetterQueue
	EncryptionKey() awskms.IKey
	EventBus() awsevents.IEventBus
	EventsRule() awsevents.Rule
	Node() constructs.Node
	SqsQueue() awssqs.Queue
	ToString() *string
}

// The jsii proxy struct for EventbridgeToSqs
type jsiiProxy_EventbridgeToSqs struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EventbridgeToSqs) DeadLetterQueue() *awssqs.DeadLetterQueue {
	var returns *awssqs.DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSqs) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSqs) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSqs) EventsRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSqs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSqs) SqsQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}


func NewEventbridgeToSqs(scope constructs.Construct, id *string, props *EventbridgeToSqsProps) EventbridgeToSqs {
	_init_.Initialize()

	j := jsiiProxy_EventbridgeToSqs{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-sqs.EventbridgeToSqs",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventbridgeToSqs_Override(e EventbridgeToSqs, scope constructs.Construct, id *string, props *EventbridgeToSqsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-sqs.EventbridgeToSqs",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventbridgeToSqs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-eventbridge-sqs.EventbridgeToSqs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventbridgeToSqs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventbridgeToSqsProps struct {
	// User provided eventRuleProps to override the defaults.
	EventRuleProps *awsevents.RuleProps `json:"eventRuleProps"`
	// Optional user provided properties for the dead letter queue.
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	// Whether to deploy a secondary queue to be used as a dead letter queue.
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	// Use a KMS Key, either managed by this CDK app, or imported.
	//
	// If importing an encryption key, it must be specified in
	// the encryptionKey property for this construct.
	EnableEncryptionWithCustomerManagedKey *bool `json:"enableEncryptionWithCustomerManagedKey"`
	// Whether to grant additional permissions to the Lambda function enabling it to purge the SQS queue.
	EnableQueuePurging *bool `json:"enableQueuePurging"`
	// An optional, imported encryption key to encrypt the SQS queue, and SNS Topic.
	EncryptionKey awskms.Key `json:"encryptionKey"`
	// Optional user-provided props to override the default props for the encryption key.
	EncryptionKeyProps *awskms.KeyProps `json:"encryptionKeyProps"`
	// A new custom EventBus is created with provided props.
	EventBusProps *awsevents.EventBusProps `json:"eventBusProps"`
	// Existing instance of a custom EventBus.
	ExistingEventBusInterface awsevents.IEventBus `json:"existingEventBusInterface"`
	// Existing instance of SQS queue object, providing both this and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead-letter queue.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// User provided props to override the default props for the SQS queue.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
}

