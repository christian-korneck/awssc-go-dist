package awssolutionsconstructsawswafwebaclcloudfront

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-wafwebacl-cloudfront.WafwebaclToCloudFront",
		reflect.TypeOf((*WafwebaclToCloudFront)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontWebDistribution", GoGetter: "CloudFrontWebDistribution"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webacl", GoGetter: "Webacl"},
		},
		func() interface{} {
			j := jsiiProxy_WafwebaclToCloudFront{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-wafwebacl-cloudfront.WafwebaclToCloudFrontProps",
		reflect.TypeOf((*WafwebaclToCloudFrontProps)(nil)).Elem(),
	)
}
