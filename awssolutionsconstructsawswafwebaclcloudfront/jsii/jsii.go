// Package jsii contains the functionaility needed for jsii packages to
// initialize their dependencies and themselves. Users should never need to use this package
// directly. If you find you need to - please report a bug at
// https://github.com/aws/jsii/issues/new/choose
package jsii

import (
	_                                                   "embed"

	_jsii_                                              "github.com/aws/jsii-runtime-go/runtime"

	awscdk                                              "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	constructs                                          "github.com/aws/constructs-go/constructs/v10/jsii"
	awssolutionsconstructsawscloudfrontapigatewaylambda "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfrontapigatewaylambda/v2/jsii"
	awssolutionsconstructsawscloudfrontmediastore       "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfrontmediastore/v2/jsii"
	awssolutionsconstructsawscloudfronts3               "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfronts3/v2/jsii"
	awssolutionsconstructscore                          "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructscore/v2/jsii"
)

//go:embed aws-solutions-constructs-aws-wafwebacl-cloudfront-2.2.0.tgz
var tarball []byte

// Initialize loads the necessary packages in the @jsii/kernel to support the enclosing module.
// The implementation is idempotent (and hence safe to be called over and over).
func Initialize() {
	// Ensure all dependencies are initialized
	awssolutionsconstructsawscloudfrontapigatewaylambda.Initialize()
	awssolutionsconstructsawscloudfrontmediastore.Initialize()
	awssolutionsconstructsawscloudfronts3.Initialize()
	awssolutionsconstructscore.Initialize()
	awscdk.Initialize()
	constructs.Initialize()

	// Load this library into the kernel
	_jsii_.Load("@aws-solutions-constructs/aws-wafwebacl-cloudfront", "2.2.0", tarball)
}
