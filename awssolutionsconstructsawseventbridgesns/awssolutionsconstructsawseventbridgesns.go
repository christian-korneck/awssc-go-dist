// CDK Constructs for deploying AWS Events Rule that invokes AWS SNS
package awssolutionsconstructsawseventbridgesns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgesns/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawseventbridgesns/v2/internal"
)

type EventbridgeToSns interface {
	constructs.Construct
	EncryptionKey() awskms.Key
	EventBus() awsevents.IEventBus
	EventsRule() awsevents.Rule
	Node() constructs.Node
	SnsTopic() awssns.Topic
	ToString() *string
}

// The jsii proxy struct for EventbridgeToSns
type jsiiProxy_EventbridgeToSns struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EventbridgeToSns) EncryptionKey() awskms.Key {
	var returns awskms.Key
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSns) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSns) EventsRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSns) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventbridgeToSns) SnsTopic() awssns.Topic {
	var returns awssns.Topic
	_jsii_.Get(
		j,
		"snsTopic",
		&returns,
	)
	return returns
}


func NewEventbridgeToSns(scope constructs.Construct, id *string, props *EventbridgeToSnsProps) EventbridgeToSns {
	_init_.Initialize()

	j := jsiiProxy_EventbridgeToSns{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-sns.EventbridgeToSns",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventbridgeToSns_Override(e EventbridgeToSns, scope constructs.Construct, id *string, props *EventbridgeToSnsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-eventbridge-sns.EventbridgeToSns",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventbridgeToSns_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-eventbridge-sns.EventbridgeToSns",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventbridgeToSns) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventbridgeToSnsProps struct {
	// User provided eventRuleProps to override the defaults.
	EventRuleProps *awsevents.RuleProps `json:"eventRuleProps"`
	// Use a KMS Key, either managed by this CDK app, or imported.
	//
	// If importing an encryption key, it must be specified in
	// the encryptionKey property for this construct.
	EnableEncryptionWithCustomerManagedKey *bool `json:"enableEncryptionWithCustomerManagedKey"`
	// An optional, imported encryption key to encrypt the SQS queue, and SNS Topic.
	EncryptionKey awskms.Key `json:"encryptionKey"`
	// Optional user-provided props to override the default props for the encryption key.
	EncryptionKeyProps *awskms.KeyProps `json:"encryptionKeyProps"`
	// A new custom EventBus is created with provided props.
	EventBusProps *awsevents.EventBusProps `json:"eventBusProps"`
	// Existing instance of a custom EventBus.
	ExistingEventBusInterface awsevents.IEventBus `json:"existingEventBusInterface"`
	// Existing instance of SNS Topic object, providing both this and topicProps will cause an error..
	ExistingTopicObj awssns.Topic `json:"existingTopicObj"`
	// User provided props to override the default props for the SNS Topic.
	TopicProps *awssns.TopicProps `json:"topicProps"`
}

