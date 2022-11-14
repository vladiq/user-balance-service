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

var DepositFundsValid = `
{
	"id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": 5
}
`

var DepositFundsInvalidJSON = `
{
	id: "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": 5
}
`

var DepositFundsNegativeAmount = `
{
	"id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": -5
}
`

var DepositFundsInvalidUUID = `
	"id": "ooooooooooooooooooooooooooooooo",
	"amount": 5
`

var WithdrawFundsValid = `
{
	"id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": 5
}
`

var WithdrawFundsInvalidJSON = `
{
	id: "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": 5
}
`

var WithdrawFundsNegativeAmount = `
{
	"id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": -5
}
`

var WithdrawFundsInvalidUUID = `
	"id": "ooooooooooooooooooooooooooooooo",
	"amount": 5
`
