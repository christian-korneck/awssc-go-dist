// CDK Constructs for AWS API Gateway and Amazon DynamoDB integration.
package awssolutionsconstructsawsapigatewaydynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewaydynamodb/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawsapigatewaydynamodb/v2/internal"
)

type ApiGatewayToDynamoDB interface {
	constructs.Construct
	ApiGateway() awsapigateway.RestApi
	ApiGatewayCloudWatchRole() awsiam.Role
	ApiGatewayLogGroup() awslogs.LogGroup
	ApiGatewayRole() awsiam.Role
	DynamoTable() awsdynamodb.Table
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for ApiGatewayToDynamoDB
type jsiiProxy_ApiGatewayToDynamoDB struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApiGatewayToDynamoDB) ApiGateway() awsapigateway.RestApi {
	var returns awsapigateway.RestApi
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToDynamoDB) ApiGatewayCloudWatchRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayCloudWatchRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToDynamoDB) ApiGatewayLogGroup() awslogs.LogGroup {
	var returns awslogs.LogGroup
	_jsii_.Get(
		j,
		"apiGatewayLogGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToDynamoDB) ApiGatewayRole() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"apiGatewayRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToDynamoDB) DynamoTable() awsdynamodb.Table {
	var returns awsdynamodb.Table
	_jsii_.Get(
		j,
		"dynamoTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayToDynamoDB) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewApiGatewayToDynamoDB(scope constructs.Construct, id *string, props *ApiGatewayToDynamoDBProps) ApiGatewayToDynamoDB {
	_init_.Initialize()

	j := jsiiProxy_ApiGatewayToDynamoDB{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-dynamodb.ApiGatewayToDynamoDB",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApiGatewayToDynamoDB_Override(a ApiGatewayToDynamoDB, scope constructs.Construct, id *string, props *ApiGatewayToDynamoDBProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-apigateway-dynamodb.ApiGatewayToDynamoDB",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiGatewayToDynamoDB_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-apigateway-dynamodb.ApiGatewayToDynamoDB",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiGatewayToDynamoDB) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApiGatewayToDynamoDBProps struct {
	// Whether to deploy API Gateway Method for Create operation on DynamoDB table.
	AllowCreateOperation *bool `json:"allowCreateOperation"`
	// Whether to deploy API Gateway Method for Delete operation on DynamoDB table.
	AllowDeleteOperation *bool `json:"allowDeleteOperation"`
	// Whether to deploy API Gateway Method for Read operation on DynamoDB table.
	AllowReadOperation *bool `json:"allowReadOperation"`
	// Whether to deploy API Gateway Method for Update operation on DynamoDB table.
	AllowUpdateOperation *bool `json:"allowUpdateOperation"`
	// Optional user-provided props to override the default props for the API Gateway.
	ApiGatewayProps *awsapigateway.RestApiProps `json:"apiGatewayProps"`
	// API Gateway Request template for Create method, required if allowCreateOperation set to true.
	CreateRequestTemplate *string `json:"createRequestTemplate"`
	// Optional API Gateway Request template for Delete method, it will use the default template if allowDeleteOperation is true and deleteRequestTemplate is not provided.
	//
	// The default template only supports a partition key and not partition + sort keys.
	DeleteRequestTemplate *string `json:"deleteRequestTemplate"`
	// Optional user provided props to override the default props.
	DynamoTableProps *awsdynamodb.TableProps `json:"dynamoTableProps"`
	// Existing instance of DynamoDB table object, providing both this and `dynamoTableProps` will cause an error.
	ExistingTableObj awsdynamodb.Table `json:"existingTableObj"`
	// User provided props to override the default props for the CloudWatchLogs LogGroup.
	LogGroupProps *awslogs.LogGroupProps `json:"logGroupProps"`
	// Optional API Gateway Request template for Read method, it will use the default template if allowReadOperation is true and readRequestTemplate is not provided.
	//
	// The default template only supports a partition key and not partition + sort keys.
	ReadRequestTemplate *string `json:"readRequestTemplate"`
	// API Gateway Request template for Update method, required if allowUpdateOperation set to true.
	UpdateRequestTemplate *string `json:"updateRequestTemplate"`
}

