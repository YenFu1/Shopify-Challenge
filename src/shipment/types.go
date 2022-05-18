package shipment

import "Shopify-Challenge/src/database"

type Shipment struct {
	database.Shipment
	Items []database.Item `json:"items"`
}

// GetShipmentsResponse represents the response body of the GetShipments endpoint
type GetShipmentsResponse struct {
	Count     int        `json:"count"`
	Shipments []Shipment `json:"shipments"`
}

// CreateShipmentRequest represents the request body of the CreateShipment endpoint
type CreateShipmentRequest struct {
	Name        string          `json:"name"`
	InventoryID int             `json:"inventoryId"`
	Items       []database.Item `json:"items"`
}
