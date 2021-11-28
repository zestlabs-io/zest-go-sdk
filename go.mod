module github.com/zestlabs-io/zest-go-sdk

go 1.15

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/go-openapi/analysis v0.21.1 // indirect
	github.com/go-openapi/errors v0.20.1
	github.com/go-openapi/runtime v0.21.0
	github.com/go-openapi/strfmt v0.21.1
	github.com/go-openapi/swag v0.19.15
	github.com/go-openapi/validate v0.20.3
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golangci/golangci-lint v1.26.0 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	go.mongodb.org/mongo-driver v1.8.0 // indirect
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

// So far we don't support backward compatability - as soon as we do, we will jump to v2.0.0 with v2 package
retract (
	v1.0.1
	v1.2.0
	v1.2.1
)
