package users

import "fmt"

type BillingStatus int

const (
	Free BillingStatus = iota
	Premium
)

func NewBilingStatus(value int) (BillingStatus, error) {
	switch value {
	case 0:
		return Free, nil
	case 1:
		return Premium, nil
	default:
		return -1, fmt.Errorf("ValueError: ステータスが不正です。")
	}
}
