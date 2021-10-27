package users

import (
	"fmt"

	"github.com/tonouchi510/InvestorService/internal/domain/models/investmentitem"
)

type User struct {
	Id                 UserId
	name               UserName
	status             BillingStatus
	investmentItemList []investmentitem.InvestmentItemId
}

func NewUser(id UserId, name UserName, status BillingStatus, items []investmentitem.InvestmentItemId) (*User, error) {
	newUser := User{
		Id:                 id,
		name:               name,
		status:             status,
		investmentItemList: items,
	}
	return &newUser, nil
}

func (u User) Equals(other User) (bool, error) {
	return (u.Id == other.Id), nil
}

func (u *User) ChangeName(name UserName) error {
	u.name = name
	return nil
}

func (u *User) ChangeStatus(status BillingStatus) error {
	u.status = status
	return nil
}

func (u *User) AddInvestmentItem(itemId investmentitem.InvestmentItemId) error {
	if u.IsInvestmentItemLimit() {
		return fmt.Errorf("InvestmentItem数が既に上限です。(%d件)", len(u.investmentItemList))
	}
	if u.IsDuplicatedItemList(itemId) {
		return fmt.Errorf("InvestmentItemId: %s は既に追加されています。", itemId)
	}
	u.investmentItemList = append(u.investmentItemList, itemId)
	return nil
}

func (u *User) RemoveInvestmentItem(itemId investmentitem.InvestmentItemId) error {
	results := []investmentitem.InvestmentItemId{}
	size := len(u.investmentItemList)
	for _, v := range u.investmentItemList {
		if v != itemId {
			results = append(results, v)
		}
	}
	if len(results) != (size - 1) {
		// 指定したitemIdが元々ない場合
		return fmt.Errorf("InvestmentItemId: %s が存在しません。", itemId)
	}
	u.investmentItemList = results
	return nil
}

func (u User) IsDuplicatedItemList(itemId investmentitem.InvestmentItemId) bool {
	for _, v := range u.investmentItemList {
		if v == itemId {
			return true
		}
	}
	return false
}

func (u User) IsInvestmentItemLimit() bool {
	c := len(u.investmentItemList)
	// 将来的にユーザランクの概念も導入し、ステータスと組み合わせて制限設定するのも良いかも
	switch u.status {
	case Free:
		return c <= 5
	case Premium:
		return c <= 30
	default:
		return c <= 5
	}
}

func (u User) Notify(note IUserNotification) error {
	note.SetId(u.Id)
	note.SetName(u.name)
	note.SetStatus(u.status)
	note.SetItemList(u.investmentItemList)
	return nil
}
