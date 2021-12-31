// CDK Constructs for Amazon Dynamodb streams to AWS Lambda to AWS Elasticsearch with Kibana integration
package awssolutionsconstructsawsdynamodbstreamslambdaelasticsearchkibana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsdynamodbstreamslambdaelasticsearchkibana/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsdynamodbstreamslambdaelasticsearchkibana/v2/internal"
)

type DynamoDBStreamsToLambdaToElasticSearchAndKibana interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	DynamoTable() awsdynamodb.Table
	DynamoTableInterface() awsdynamodb.ITable
	ElasticsearchDomain() awselasticsearch.CfnDomain
	ElasticsearchRole() awsiam.Role
	IdentityPool() awscognito.CfnIdentityPool
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	UserPool() awscognito.UserPool
	UserPoolClient() awscognito.UserPoolClient
	ToString() *string
}

// The jsii proxy struct for DynamoDBStreamsToLambdaToElasticSearchAndKibana
type jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) DynamoTable() awsdynamodb.Table {
	var returns awsdynamodb.Table
	_jsii_.Get(
		j,
		"dynamoTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) DynamoTableInterface() awsdynamodb.ITable {
	var returns awsdynamodb.ITable
	_jsii_.Get(
		j,
		"dynamoTableInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) ElasticsearchDomain() awselasticsearch.CfnDomain {
	var returns awselasticsearch.CfnDomain
	_jsii_.Get(
		j,
		"elasticsearchDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) ElasticsearchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"elasticsearchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) IdentityPool() awscognito.CfnIdentityPool {
	var returns awscognito.CfnIdentityPool
	_jsii_.Get(
		j,
		"identityPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) UserPool() awscognito.UserPool {
	var returns awscognito.UserPool
	_jsii_.Get(
		j,
		"userPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) UserPoolClient() awscognito.UserPoolClient {
	var returns awscognito.UserPoolClient
	_jsii_.Get(
		j,
		"userPoolClient",
		&returns,
	)
	return returns
}


func NewDynamoDBStreamsToLambdaToElasticSearchAndKibana(scope constructs.Construct, id *string, props *DynamoDBStreamsToLambdaToElasticSearchAndKibanaProps) DynamoDBStreamsToLambdaToElasticSearchAndKibana {
	_init_.Initialize()

	j := jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda-elasticsearch-kibana.DynamoDBStreamsToLambdaToElasticSearchAndKibana",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDynamoDBStreamsToLambdaToElasticSearchAndKibana_Override(d DynamoDBStreamsToLambdaToElasticSearchAndKibana, scope constructs.Construct, id *string, props *DynamoDBStreamsToLambdaToElasticSearchAndKibanaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda-elasticsearch-kibana.DynamoDBStreamsToLambdaToElasticSearchAndKibana",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DynamoDBStreamsToLambdaToElasticSearchAndKibana_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda-elasticsearch-kibana.DynamoDBStreamsToLambdaToElasticSearchAndKibana",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DynamoDBStreamsToLambdaToElasticSearchAndKibanaProps struct {
	// Cognito & ES Domain Name.
	DomainName *string `json:"domainName"`
	// Optional Cognito Domain Name, if provided it will be used for Cognito Domain, and domainName will be used for the Elasticsearch Domain.
	CognitoDomainName *string `json:"cognitoDomainName"`
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// Whether to deploy a SQS dead letter queue when a data record reaches the Maximum Retry Attempts or Maximum Record Age, its metadata like shard ID and stream ARN will be sent to an SQS queue.
	DeploySqsDlqQueue *bool `json:"deploySqsDlqQueue"`
	// Optional user provided props to override the default props.
	DynamoEventSourceProps *awslambdaeventsources.DynamoEventSourceProps `json:"dynamoEventSourceProps"`
	// Optional user provided props to override the default props.
	DynamoTableProps *awsdynamodb.TableProps `json:"dynamoTableProps"`
	// Optional user provided props to override the default props for the API Gateway.
	EsDomainProps *awselasticsearch.CfnDomainProps `json:"esDomainProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// Existing instance of DynamoDB table object, providing both this and `dynamoTableProps` will cause an error.
	ExistingTableInterface awsdynamodb.ITable `json:"existingTableInterface"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
	// Optional user provided properties for the SQS dead letter queue.
	SqsDlqQueueProps *awssqs.QueueProps `json:"sqsDlqQueueProps"`
}

