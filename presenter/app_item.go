package presenter

import "git.orion.home/oxhead/casa/model"

type AppItem struct {
	ID          uint
	HomeItemID  *uint
	Title       string
	Description string
	ImageURL    string
	URL         string
}

func AppItemsFromModelList(catalogItems []model.CatalogItem, homeItems []model.HomeItem) []AppItem {
	homeItemMapByAppID := map[uint]model.HomeItem{}

	for _, homeItem := range homeItems {
		homeItemMapByAppID[homeItem.CatalogItem.ID] = homeItem
	}

	appItems := []AppItem{}

	for _, catalogItem := range catalogItems {
		appItem := AppItem{
			ID:          catalogItem.ID,
			Title:       catalogItem.Title,
			Description: catalogItem.Description,
			ImageURL:    catalogItem.ImageURL,
			URL:         catalogItem.URL,
		}

		if homeItem, ok := homeItemMapByAppID[catalogItem.ID]; ok {
			homeItemID := homeItem.ID
			appItem.HomeItemID = &homeItemID
		}

		appItems = append(appItems, appItem)
	}

	return appItems
}
