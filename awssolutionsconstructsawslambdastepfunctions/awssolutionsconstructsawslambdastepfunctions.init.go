package awssolutionsconstructsawslambdastepfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-lambda-stepfunctions.LambdaToStepfunctions",
		reflect.TypeOf((*LambdaToStepfunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAlarms", GoGetter: "CloudwatchAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachine", GoGetter: "StateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachineLogGroup", GoGetter: "StateMachineLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaToStepfunctions{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-lambda-stepfunctions.LambdaToStepfunctionsProps",
		reflect.TypeOf((*LambdaToStepfunctionsProps)(nil)).Elem(),
	)
}
