// CDK constructs for defining an interaction between an AWS Lambda function and an Amazon S3 bucket.
package awssolutionsconstructsawsapigatewaysqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewaysqs/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewaysqs/v2/internal"
)

type ApiGatewayToSqs interface {
	constructs.Construct
	ApiGateway() awsapigateway.RestApi
	ApiGatewayCloudWatchRole() awsiam.Role
	ApiGatewayLogGroup() awslogs.LogGroup
	ApiGatewayRole() awsiam.Role
	DeadLetterQueue() *awssqs.DeadLetterQueue
	Node() constructs.Node
	SqsQueue() awssqs.Queue
	ToString() *string
}

// The jsii proxy struct for ApiGatewayToSqs
type jsiiProxy_ApiGatewayToSqs struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApiGatewayToSqs) ApiGateway() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSqs) ApiGatewayCloudWatchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSqs) ApiGatewayLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"apiGatewayLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSqs) ApiGatewayRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSqs) DeadLetterQueue() *awssqs.DeadLetterQueue {
	var returns *awssqs.DeadLetterQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSqs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToSqs) SqsQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"sqsQueue",
		&returns,
	)
	return returns
}


func NewApiGatewayToSqs(scope constructs.Construct, id *string, props *ApiGatewayToSqsProps) ApiGatewayToSqs {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayToSqs{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-sqs.ApiGatewayToSqs",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApiGatewayToSqs_Override(a ApiGatewayToSqs, scope constructs.Construct, id *string, props *ApiGatewayToSqsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-sqs.ApiGatewayToSqs",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayToSqs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-apigateway-sqs.ApiGatewayToSqs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayToSqs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiGatewayToSqsProps struct {
	// Whether to deploy an API Gateway Method for Create operations on the queue (i.e. sqs:SendMessage).
	AllowCreateOperation *bool `json:"allowCreateOperation"`
	// Whether to deploy an API Gateway Method for Delete operations on the queue (i.e. sqs:DeleteMessage).
	AllowDeleteOperation *bool `json:"allowDeleteOperation"`
	// Whether to deploy an API Gateway Method for Read operations on the queue (i.e. sqs:ReceiveMessage).
	AllowReadOperation *bool `json:"allowReadOperation"`
	// Optional user-provided props to override the default props for the API Gateway.
	ApiGatewayProps interface{} `json:"apiGatewayProps"`
	// API Gateway Request template for Create method, if allowCreateOperation set to true.
	CreateRequestTemplate *string `json:"createRequestTemplate"`
	// Optional user provided properties for the dead letter queue.
	DeadLetterQueueProps *awssqs.QueueProps `json:"deadLetterQueueProps"`
	// API Gateway Request template for Delete method, if allowDeleteOperation set to true.
	DeleteRequestTemplate *string `json:"deleteRequestTemplate"`
	// Whether to deploy a secondary queue to be used as a dead letter queue.
	DeployDeadLetterQueue *bool `json:"deployDeadLetterQueue"`
	// Existing instance of SQS queue object, providing both this  and queueProps will cause an error.
	ExistingQueueObj awssqs.Queue `json:"existingQueueObj"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
	// The number of times a message can be unsuccessfully dequeued before being moved to the dead-letter queue.
	MaxReceiveCount *float64 `json:"maxReceiveCount"`
	// User provided props to override the default props for the SQS queue.
	QueueProps *awssqs.QueueProps `json:"queueProps"`
	// API Gateway Request template for Get method, if allowReadOperation set to true.
	ReadRequestTemplate *string `json:"readRequestTemplate"`
}

