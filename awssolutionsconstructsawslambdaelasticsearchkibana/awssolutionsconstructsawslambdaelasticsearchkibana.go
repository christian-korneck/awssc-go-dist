// CDK Constructs for AWS Lambda to AWS Elasticsearch with Kibana integration
package awssolutionsconstructsawslambdaelasticsearchkibana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdaelasticsearchkibana/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawslambdaelasticsearchkibana/v2/internal"
)

type LambdaToElasticSearchAndKibana interface {
	constructs.Construct
	CloudwatchAlarms() *[]awscloudwatch.Alarm
	ElasticsearchDomain() awselasticsearch.CfnDomain
	ElasticsearchRole() awsiam.Role
	IdentityPool() awscognito.CfnIdentityPool
	LambdaFunction() awslambda.Function
	Node() constructs.Node
	UserPool() awscognito.UserPool
	UserPoolClient() awscognito.UserPoolClient
	ToString() *string
}

// The jsii proxy struct for LambdaToElasticSearchAndKibana
type jsiiProxy_LambdaToElasticSearchAndKibana struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LambdaToElasticSearchAndKibana) CloudwatchAlarms() *[]awscloudwatch.Alarm {
	var returns *[]awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"cloudwatchAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToElasticSearchAndKibana) ElasticsearchDomain() awselasticsearch.CfnDomain {
	var returns awselasticsearch.CfnDomain
	_jsii_.Get(
		j,
		"elasticsearchDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToElasticSearchAndKibana) ElasticsearchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"elasticsearchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToElasticSearchAndKibana) IdentityPool() awscognito.CfnIdentityPool {
	var returns awscognito.CfnIdentityPool
	_jsii_.Get(
		j,
		"identityPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToElasticSearchAndKibana) LambdaFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToElasticSearchAndKibana) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToElasticSearchAndKibana) UserPool() awscognito.UserPool {
	var returns awscognito.UserPool
	_jsii_.Get(
		j,
		"userPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaToElasticSearchAndKibana) UserPoolClient() awscognito.UserPoolClient {
	var returns awscognito.UserPoolClient
	_jsii_.Get(
		j,
		"userPoolClient",
		&returns,
	)
	return returns
}


func NewLambdaToElasticSearchAndKibana(scope constructs.Construct, id *string, props *LambdaToElasticSearchAndKibanaProps) LambdaToElasticSearchAndKibana {
	_init_.Initialize()

	j := jsiiProxy_LambdaToElasticSearchAndKibana{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-elasticsearch-kibana.LambdaToElasticSearchAndKibana",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaToElasticSearchAndKibana_Override(l LambdaToElasticSearchAndKibana, scope constructs.Construct, id *string, props *LambdaToElasticSearchAndKibanaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-lambda-elasticsearch-kibana.LambdaToElasticSearchAndKibana",
		[]interface{}{scope, id, props},
		l,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LambdaToElasticSearchAndKibana_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-lambda-elasticsearch-kibana.LambdaToElasticSearchAndKibana",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LambdaToElasticSearchAndKibana) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LambdaToElasticSearchAndKibanaProps struct {
	// Cognito & ES Domain Name.
	DomainName *string `json:"domainName"`
	// Optional Cognito Domain Name, if provided it will be used for Cognito Domain, and domainName will be used for the Elasticsearch Domain.
	CognitoDomainName *string `json:"cognitoDomainName"`
	// Whether to create recommended CloudWatch alarms.
	CreateCloudWatchAlarms *bool `json:"createCloudWatchAlarms"`
	// Optional Name for the ElasticSearch domain endpoint environment variable set for the Lambda function.
	DomainEndpointEnvironmentVariableName *string `json:"domainEndpointEnvironmentVariableName"`
	// Optional user provided props to override the default props for the Elasticsearch Service.
	EsDomainProps *awselasticsearch.CfnDomainProps `json:"esDomainProps"`
	// Existing instance of Lambda Function object, providing both this and `lambdaFunctionProps` will cause an error.
	ExistingLambdaObj awslambda.Function `json:"existingLambdaObj"`
	// User provided props to override the default props for the Lambda function.
	LambdaFunctionProps *awslambda.FunctionProps `json:"lambdaFunctionProps"`
}

