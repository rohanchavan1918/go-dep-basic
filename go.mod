module go-dep-inj

go 1.21.1

replace go-dep-inj/libs v1.0.0 => ./libs

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/gomodule/redigo v1.9.2
)
