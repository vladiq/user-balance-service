package testdata

const ValidUUIDString = "f896ad24-b44d-4fca-b382-c2e4beb2c280"

var CreateAccountValid = `
{
	"amount": 5.0
}
`

var CreateAccountInvalidJSON = `
{
	"amount": 5.0
`

var CreateAccountNegativeAmount = `
{
	"amount": -5.0
}
`
