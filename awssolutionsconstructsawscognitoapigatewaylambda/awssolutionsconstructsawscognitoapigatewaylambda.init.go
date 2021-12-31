package awssolutionsconstructsawscognitoapigatewaylambda

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-cognito-apigateway-lambda.CognitoToApiGatewayToLambda",
		reflect.TypeOf((*CognitoToApiGatewayToLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAuthorizers", GoMethod: "AddAuthorizers"},
			_jsii_.MemberProperty{JsiiProperty: "apiGateway", GoGetter: "ApiGateway"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayAuthorizer", GoGetter: "ApiGatewayAuthorizer"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayCloudWatchRole", GoGetter: "ApiGatewayCloudWatchRole"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayLogGroup", GoGetter: "ApiGatewayLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userPool", GoGetter: "UserPool"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolClient", GoGetter: "UserPoolClient"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoToApiGatewayToLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-cognito-apigateway-lambda.CognitoToApiGatewayToLambdaProps",
		reflect.TypeOf((*CognitoToApiGatewayToLambdaProps)(nil)).Elem(),
	)
}
