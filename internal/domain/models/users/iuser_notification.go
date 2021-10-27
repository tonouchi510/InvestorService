//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE -self_package=github.com/tonouchi510/InvestorService/internal/domain/models/$GOPACKAGE
package users

import "github.com/tonouchi510/InvestorService/internal/domain/models/investmentitem"

type IUserNotification interface {
	SetId(id UserId)
	SetName(name UserName)
	SetStatus(status BillingStatus)
	SetItemList(items []investmentitem.InvestmentItemId)
}
