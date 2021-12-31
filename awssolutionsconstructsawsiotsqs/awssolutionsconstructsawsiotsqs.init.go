package awssolutionsconstructsawsiotsqs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-iot-sqs.IotToSqs",
		reflect.TypeOf((*IotToSqs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "iotActionsRole", GoGetter: "IotActionsRole"},
			_jsii_.MemberProperty{JsiiProperty: "iotTopicRule", GoGetter: "IotTopicRule"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "sqsQueue", GoGetter: "SqsQueue"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IotToSqs{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-iot-sqs.IotToSqsProps",
		reflect.TypeOf((*IotToSqsProps)(nil)).Elem(),
	)
}
