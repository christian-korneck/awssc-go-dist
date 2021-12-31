package awssolutionsconstructsawskinesisstreamsgluejob

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-kinesisstreams-gluejob.KinesisstreamsToGluejob",
		reflect.TypeOf((*KinesisstreamsToGluejob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAlarms", GoGetter: "CloudwatchAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "glueJob", GoGetter: "GlueJob"},
			_jsii_.MemberProperty{JsiiProperty: "glueJobRole", GoGetter: "GlueJobRole"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStream", GoGetter: "KinesisStream"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputBucket", GoGetter: "OutputBucket"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisstreamsToGluejob{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-kinesisstreams-gluejob.KinesisstreamsToGluejobProps",
		reflect.TypeOf((*KinesisstreamsToGluejobProps)(nil)).Elem(),
	)
}
