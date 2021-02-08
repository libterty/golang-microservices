package domains

import (
	"../utils"
	"fmt"
	"net/http"
)

type itemDao struct {}

var (
	items = map[string]*Item{
		"lib11": &Item{
			Id: "lib23",
			Content: "libTest",
		},
	}
	ItemDao itemDao
)

func (i *itemDao) GetItem(itemId string) (*Item, *utils.ApplicationError)  {
	item := items[itemId]
	if item == nil {
		return nil, &utils.ApplicationError{
			Message: fmt.Sprintf("Item not found"),
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}
	}
	return item, nil
}