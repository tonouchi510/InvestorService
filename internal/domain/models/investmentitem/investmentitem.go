package investmentitem

import "fmt"

type InvestmentItemId string

func NewInvestmentItemId(value string) (InvestmentItemId, error) {
	if value == "" {
		return "", fmt.Errorf("ValueError: InvestmentItemIdが空です。")
	} else if len(value) != 32 {
		return "", fmt.Errorf("ValueError: InvestmentItemIdが不正です。")
	}
	id := InvestmentItemId(value)
	return id, nil
}
