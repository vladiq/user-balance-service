package testdata

var MakeTransferValid = `
{
	"from_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"to_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": 5.5
}
`

var MakeTransferInvalidJSON = `
{
	"from_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"to_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": 5.5
`

var MakeTransferNegativeAmount = `
{
	"from_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"to_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": -5.5
}
`

var MakeTransferInvalidUUID = `
{
	"from_id": "foo",
	"to_id": "bar",
	"amount": 5.5
}
`
