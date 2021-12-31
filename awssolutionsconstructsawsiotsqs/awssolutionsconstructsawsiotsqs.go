// CDK Constructs for AWS IoT to AWS SQS integration
package awssolutionsconstructsawsiotsqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotsqs/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsiotsqs/v2/internal"
)

type IotToSqs interface {
	constructs.Construct
	DeadLetterQueue() *awssqs.DeadLetterQueue
	EncryptionKey() awskms.IKey
	IotActionsRole() awsiam.Role
	IotTopicRule() awsiot.CfnTopicRule
	Node() constructs.Node
	SqsQueue() awssqs.Queue
	ToString() *string
}

// The jsii proxy struct for IotToSqs
type jsiiProxy_IotToSqs struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IotToSqs) DeadLetterQueue() *awssqs.DeadLetterQueue {
	var returns *awssqs.DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToSqs) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToSqs) IotActionsRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"iotActionsRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToSqs) IotTopicRule() awsiot.CfnTopicRule {
	var returns awsiot.CfnTopicRule
	_jsii_.Get(
		j,
		"iotTopicRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToSqs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotToSqs) SqsQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}


func NewIotToSqs(scope constructs.Construct, id *string, props *IotToSqsProps) IotToSqs {
	_init_.Initialize()

	j := jsiiProxy_IotToSqs{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-sqs.IotToSqs",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIotToSqs_Override(i IotToSqs, scope constructs.Construct, id *string, props *IotToSqsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-iot-sqs.IotToSqs",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func IotToSqs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-iot-sqs.IotToSqs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (i *jsiiProxy_IotToSqs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type IotToSqsProps struct {
	// User provided CfnTopicRuleProps to override the defaults.
	IotTopicRuleProps *awsiot.CfnTopicRuleProps `json:"iotTopicRuleProps"`
	// Optional user provided properties for the dead letter queue.
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	// Whether to deploy a secondary queue to be used as a dead letter queue.
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	// Use a KMS Key, either managed by this CDK app, or imported.
	//
	// If importing an encryption key, it must be specified in
	// the encryptionKey property for this construct.
	EnableEncryptionWithCustomerManagedKey *bool `json:"enableEncryptionWithCustomerManagedKey"`
	// An optional, imported encryption key to encrypt the SQS queue, and SNS Topic.
	EncryptionKey awskms.Key `json:"encryptionKey"`
	// Optional user-provided props to override the default props for the encryption key.
	EncryptionKeyProps *awskms.KeyProps `json:"encryptionKeyProps"`
	// Existing instance of SQS queue object, providing both this and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead-letter queue.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// User provided props to override the default props for the SQS queue.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
}

