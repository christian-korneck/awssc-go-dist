// CDK constructs for defining an interaction between an Amazon S3 bucket and an Amazon SQS queue.
package awssolutionsconstructsawss3sqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawss3sqs/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawss3sqs/v2/internal"
)

type S3ToSqs interface {
	constructs.Construct
	DeadLetterQueue() *awssqs.DeadLetterQueue
	EncryptionKey() awskms.IKey
	Node() constructs.Node
	S3Bucket() awss3.Bucket
	S3BucketInterface() awss3.IBucket
	S3LoggingBucket() awss3.Bucket
	SqsQueue() awssqs.Queue
	ToString() *string
}

// The jsii proxy struct for S3ToSqs
type jsiiProxy_S3ToSqs struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_S3ToSqs) DeadLetterQueue() *awssqs.DeadLetterQueue {
	var returns *awssqs.DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToSqs) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToSqs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToSqs) S3Bucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToSqs) S3BucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3BucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToSqs) S3LoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3LoggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ToSqs) SqsQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}


func NewS3ToSqs(scope constructs.Construct, id *string, props *S3ToSqsProps) S3ToSqs {
	_init_.Initialize()

	j := jsiiProxy_S3ToSqs{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-s3-sqs.S3ToSqs",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewS3ToSqs_Override(s S3ToSqs, scope constructs.Construct, id *string, props *S3ToSqsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-s3-sqs.S3ToSqs",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func S3ToSqs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-s3-sqs.S3ToSqs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_S3ToSqs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type S3ToSqsProps struct {
	// Optional user provided props to override the default props for the S3 Bucket.
	BucketProps *awss3.BucketProps `json:"bucketProps"`
	// Optional user provided properties for the dead letter queue.
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	// Whether to deploy a secondary queue to be used as a dead letter queue.
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	// Use a KMS Key, either managed by this CDK app, or imported.
	//
	// If importing an encryption key, it must be specified in
	// the encryptionKey property for this construct.
	EnableEncryptionWithCustomerManagedKey *bool `json:"enableEncryptionWithCustomerManagedKey"`
	// Optional imported encryption key to encrypt the SQS queue.
	EncryptionKey awskms.Key `json:"encryptionKey"`
	// Optional user provided props to override the default props for the encryption key.
	EncryptionKeyProps *awskms.KeyProps `json:"encryptionKeyProps"`
	// Existing instance of S3 Bucket object, providing both this and `bucketProps` will cause an error.
	ExistingBucketObj awss3.Bucket `json:"existingBucketObj"`
	// Existing instance of SQS queue object, Providing both this and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// Optional user provided props to override the default props for the S3 Logging Bucket.
	LoggingBucketProps *awss3.BucketProps `json:"loggingBucketProps"`
	// Whether to turn on Access Logs for the S3 bucket with the associated storage costs.
	//
	// Enabling Access Logging is a best practice.
	LogS3AccessLogs *bool `json:"logS3AccessLogs"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead-letter queue.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// Optional user provided properties.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
	// S3 object key filter rules to determine which objects trigger this event.
	S3EventFilters *[]*awss3.NotificationKeyFilter `json:"s3EventFilters"`
	// The S3 event types that will trigger the notification.
	S3EventTypes *[]awss3.EventType `json:"s3EventTypes"`
}

