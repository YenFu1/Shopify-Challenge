package inventory

import "Shopify-Challenge/src/database"

// GetInventoriesResponse represents the response of the GetInventories endpoint
type GetInventoriesResponse struct {
	Count       int                  `json:"count"`
	Inventories []database.Inventory `json:"inventories"`
}

// CreateInventoryRequest represents the body needed for the CreateInventory endpoint
type CreateInventoryRequest struct {
	Name string `json:"name"`
}

// UpdateInventoryRequest represents the body needed for the UpdateInventory endpoint
type UpdateInventoryRequest struct {
	Name string `json:"name"`
}
