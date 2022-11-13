package response

import "time"

type GetUserMonthlyReport struct {
	Timestamp time.Time
	IsAccrual bool
	Info      string
	Amount    float64
}
