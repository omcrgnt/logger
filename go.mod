module github.com/omcrgnt/logger

go 1.26.2

require (
	github.com/omcrgnt/builder v0.2.0
	github.com/omcrgnt/proto/gen/go v0.3.0
	github.com/omcrgnt/res v0.9.0
	github.com/omcrgnt/sdi v1.0.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.11-20260415201107-50325440f8f2.1 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace (
	github.com/omcrgnt/builder => /opt/github/builder
	github.com/omcrgnt/res => /opt/github/res
	github.com/omcrgnt/sdi => /opt/github/sdi
)
