module github.com/uvalib/librabus-pub

go 1.22

toolchain go1.24.2

require github.com/uvalib/librabus-sdk/uvalibrabus v0.0.0-20250722185634-5799610d7eb7

// for local development
replace github.com/uvalib/librabus-sdk/uvalibrabus => ../../uvalibrabus

require (
	github.com/aws/aws-sdk-go-v2 v1.37.0 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.30.0 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.18.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.17.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.29.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.31.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.35.0 // indirect
	github.com/aws/smithy-go v1.22.5 // indirect
)
