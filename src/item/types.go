package item

import "Shopify-Challenge/src/database"

type GetItemsResponse struct {
	Count int             `json:"count"`
	Items []database.Item `json:"items"`
}

// CreateItemRequest represents the request body for the CreateItem POST endpoint
type CreateItemRequest struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// UpdateItemRequest represents the request body for the UpdateItem PUT endpoint
type UpdateItemRequest struct {
	Name       string `json:"name"`
	Count      int    `json:"count"`
	ShipmentID *int   `json:"shipmentId"`
}
