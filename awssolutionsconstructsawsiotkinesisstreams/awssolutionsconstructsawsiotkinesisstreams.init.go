package awssolutionsconstructsawsiotkinesisstreams

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-iot-kinesisstreams.IotToKinesisStreams",
		reflect.TypeOf((*IotToKinesisStreams)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAlarms", GoGetter: "CloudwatchAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "iotActionsRole", GoGetter: "IotActionsRole"},
			_jsii_.MemberProperty{JsiiProperty: "iotTopicRule", GoGetter: "IotTopicRule"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStream", GoGetter: "KinesisStream"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IotToKinesisStreams{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-iot-kinesisstreams.IotToKinesisStreamsProps",
		reflect.TypeOf((*IotToKinesisStreamsProps)(nil)).Elem(),
	)
}
