package testdata

var CreateReservationValid = `
{
	"user_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"service_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"order_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": 5.5
}
`

var CreateReservationInvalidJSON = `
	"user_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"service_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"order_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": 5.5
}
`

var CreateReservationNegativeAmount = `
{
	"user_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"service_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"order_id": "f896ad24-b44d-4fca-b382-c2e4beb2c280",
	"amount": -5.5
}
`

var CreateReservationInvalidUUID = `
{
	"user_id": "foo",
	"service_id": "bar",
	"order_id": "foobar",
	"amount": 5.5
}
`
