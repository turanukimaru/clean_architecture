module github.com/turanukimaru/ca/usecase

go 1.17

replace github.com/turanukimaru/ca/domain => ../domain

require (
	github.com/stretchr/testify v1.7.0
	github.com/turanukimaru/ca/domain v0.0.0-00010101000000-000000000000
	github.com/turanukimaru/goastart v0.0.0-20211228053956-e54a448e3088
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	goa.design/goa/v3 v3.5.3 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)
