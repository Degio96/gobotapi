// Code AutoGenerated; DO NOT EDIT.

package filters

import (
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/types"
)

func filterableData(client *gobotapi.Client, filterable any) *DataFilter {
	dataResult := &DataFilter{}
	switch filterable.(type) {
	case types.CallbackQuery:
		x := filterable.(types.CallbackQuery)
		dataResult.From = &x.From
		dataResult.Message = x.Message
	case types.ChatJoinRequest:
		x := filterable.(types.ChatJoinRequest)
		dataResult.Chat = &x.Chat
		dataResult.From = &x.From
		dataResult.Date = &x.Date
	case types.ChatMemberUpdated:
		x := filterable.(types.ChatMemberUpdated)
		dataResult.Chat = &x.Chat
		dataResult.From = &x.From
		dataResult.Date = &x.Date
	case types.ChosenInlineResult:
		x := filterable.(types.ChosenInlineResult)
		dataResult.From = &x.From
	case types.InlineQuery:
		x := filterable.(types.InlineQuery)
		dataResult.From = &x.From
	case types.Message:
		x := filterable.(types.Message)
		dataResult.From = x.From
		dataResult.Date = &x.Date
		dataResult.Chat = &x.Chat
	case types.PreCheckoutQuery:
		x := filterable.(types.PreCheckoutQuery)
		dataResult.From = &x.From
	case types.ShippingQuery:
		x := filterable.(types.ShippingQuery)
		dataResult.From = &x.From
	}
	dataResult.Client = client
	dataResult.RawUpdate = filterable
	return dataResult
}
