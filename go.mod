module github.com/omcrgnt/logger

go 1.26.2

retract (
	[v1.0.0, v1.20.0]
	[v0.1.0, v0.20.0]
)

retract (
	[v1.0.0, v1.20.0]
	[v0.1.0, v0.20.0]
)

require (
	github.com/omcrgnt/builder v0.20.1
	github.com/omcrgnt/builder v0.20.1
	github.com/omcrgnt/proto/gen/go v0.3.0
	github.com/omcrgnt/res v0.20.1
	github.com/omcrgnt/sdi v0.20.1
	github.com/omcrgnt/res v0.20.1
	github.com/omcrgnt/sdi v0.20.1
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
