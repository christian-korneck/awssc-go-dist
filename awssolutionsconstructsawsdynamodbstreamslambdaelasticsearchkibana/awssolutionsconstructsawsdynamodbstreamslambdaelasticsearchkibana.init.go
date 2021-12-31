package awssolutionsconstructsawsdynamodbstreamslambdaelasticsearchkibana

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda-elasticsearch-kibana.DynamoDBStreamsToLambdaToElasticSearchAndKibana",
		reflect.TypeOf((*DynamoDBStreamsToLambdaToElasticSearchAndKibana)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAlarms", GoGetter: "CloudwatchAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "dynamoTable", GoGetter: "DynamoTable"},
			_jsii_.MemberProperty{JsiiProperty: "dynamoTableInterface", GoGetter: "DynamoTableInterface"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchDomain", GoGetter: "ElasticsearchDomain"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchRole", GoGetter: "ElasticsearchRole"},
			_jsii_.MemberProperty{JsiiProperty: "identityPool", GoGetter: "IdentityPool"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userPool", GoGetter: "UserPool"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolClient", GoGetter: "UserPoolClient"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoDBStreamsToLambdaToElasticSearchAndKibana{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda-elasticsearch-kibana.DynamoDBStreamsToLambdaToElasticSearchAndKibanaProps",
		reflect.TypeOf((*DynamoDBStreamsToLambdaToElasticSearchAndKibanaProps)(nil)).Elem(),
	)
}
