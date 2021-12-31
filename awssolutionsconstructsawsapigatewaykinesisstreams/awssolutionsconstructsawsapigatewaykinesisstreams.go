// CDK Constructs for AWS API Gateway and Amazon Kinesis Data Streams integration.
package awssolutionsconstructsawsapigatewaykinesisstreams

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewaykinesisstreams/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewaykinesisstreams/v2/internal"
)

type ApiGatewayToKinesisStreams interface {
	constructs.Construct
	ApiGateway() awsapigateway.RestApi
	ApiGatewayCloudWatchRole() awsiam.Role
	ApiGatewayLogGroup() awslogs.LogGroup
	ApiGatewayRole() awsiam.Role
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	KinesisStream() awskinesis.Stream
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for ApiGatewayToKinesisStreams
type jsiiProxy_ApiGatewayToKinesisStreams struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApiGatewayToKinesisStreams) ApiGateway() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToKinesisStreams) ApiGatewayCloudWatchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToKinesisStreams) ApiGatewayLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"apiGatewayLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToKinesisStreams) ApiGatewayRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToKinesisStreams) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToKinesisStreams) KinesisStream() awskinesis.Stream {
	var returns awskinesis.Stream
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToKinesisStreams) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewApiGatewayToKinesisStreams(scope constructs.Construct, id *string, props *ApiGatewayToKinesisStreamsProps) ApiGatewayToKinesisStreams {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayToKinesisStreams{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-kinesisstreams.ApiGatewayToKinesisStreams",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApiGatewayToKinesisStreams_Override(a ApiGatewayToKinesisStreams, scope constructs.Construct, id *string, props *ApiGatewayToKinesisStreamsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-kinesisstreams.ApiGatewayToKinesisStreams",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayToKinesisStreams_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-apigateway-kinesisstreams.ApiGatewayToKinesisStreams",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayToKinesisStreams) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiGatewayToKinesisStreamsProps struct {
	// Optional user-provided props to override the default props for the API Gateway.
	ApiGatewayProps *awsapigateway.RestApiProps `json:"apiGatewayProps"`
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// Existing instance of Kinesis Stream, providing both this and `kinesisStreamProps` will cause an error.
	ExistingStreamObj awskinesis.Stream `json:"existingStreamObj"`
	// Optional user-provided props to override the default props for the Kinesis Data Stream.
	KinesisStreamProps *awskinesis.StreamProps `json:"kinesisStreamProps"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
	// API Gateway request model for the PutRecord action.
	//
	// If not provided, a default one will be created.
	PutRecordRequestModel *awsapigateway.ModelOptions `json:"putRecordRequestModel"`
	// API Gateway request template for the PutRecord action.
	//
	// If not provided, a default one will be used.
	PutRecordRequestTemplate *string `json:"putRecordRequestTemplate"`
	// API Gateway request model for the PutRecords action.
	//
	// If not provided, a default one will be created.
	PutRecordsRequestModel *awsapigateway.ModelOptions `json:"putRecordsRequestModel"`
	// API Gateway request template for the PutRecords action.
	//
	// If not provided, a default one will be used.
	PutRecordsRequestTemplate *string `json:"putRecordsRequestTemplate"`
}

