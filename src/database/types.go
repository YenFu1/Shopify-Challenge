package database

type Inventory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	InventoryID *int   `json:"inventoryId"`
	ShipmentID  *int   `json:"shipmentId"`
	Count       int    `json:"count"`
}

type Shipment struct {
	ID          int    `json:"id"`
	InventoryID int    `json:"inventoryId"`
	Name        string `json:"name"`
}
