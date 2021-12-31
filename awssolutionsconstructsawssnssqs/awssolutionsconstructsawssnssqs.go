// CDK constructs for defining an interaction between an Amazon SNS topic and an Amazon SQS queue.
package awssolutionsconstructsawssnssqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawssnssqs/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawssnssqs/v2/internal"
)

type SnsToSqs interface {
	constructs.Construct
	DeadLetterQueue() *awssqs.DeadLetterQueue
	EncryptionKey() awskms.Key
	Node() constructs.Node
	SnsTopic() awssns.Topic
	SqsQueue() awssqs.Queue
	ToString() *string
}

// The jsii proxy struct for SnsToSqs
type jsiiProxy_SnsToSqs struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SnsToSqs) DeadLetterQueue() *awssqs.DeadLetterQueue {
	var returns *awssqs.DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsToSqs) EncryptionKey() awskms.Key {
	var returns awskms.Key
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsToSqs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsToSqs) SnsTopic() awssns.Topic {
	var returns awssns.Topic
	_jsii_.Get(
		j,
		"snsTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsToSqs) SqsQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}


func NewSnsToSqs(scope constructs.Construct, id *string, props *SnsToSqsProps) SnsToSqs {
	_init_.Initialize()

	j := jsiiProxy_SnsToSqs{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-sns-sqs.SnsToSqs",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSnsToSqs_Override(s SnsToSqs, scope constructs.Construct, id *string, props *SnsToSqsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-sns-sqs.SnsToSqs",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SnsToSqs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-sns-sqs.SnsToSqs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SnsToSqs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SnsToSqsProps struct {
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
	// Existing instance of SQS queue object, Providing both this and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// Existing instance of SNS topic object, providing both this and topicProps will cause an error..
	ExistingTopicObj awssns.Topic `json:"existingTopicObj"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead-letter queue.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// Optional user provided properties.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
	// Optional user provided properties to override the default properties for the SNS topic.
	TopicProps *awssns.TopicProps `json:"topicProps"`
}

