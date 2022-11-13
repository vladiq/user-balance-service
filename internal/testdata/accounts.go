package testdata

var CreateAccountValid = `
{
	"amount": 5.0
}
`

var CreateAccountInvalidJson = `
{
	"amount": 5.0
`

var CreateAccountNegativeAmount = `
{
	"amount": -5.0
}
`
