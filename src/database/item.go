package database

import (
	"gorm.io/gorm"
)

// GetItems returns a slice containing all items in an inventory
func GetItemsByInventoryID(inventoryID int) ([]Item, error) {
	var items []Item
	err := DB.Where("inventory_id = ?", inventoryID).Find(&items).Error
	return items, err
}

func GetItemsByShipmentID(shipmentID int) ([]Item, error) {
	var items []Item
	err := DB.Where("shipment_id = ?", shipmentID).Find(&items).Error
	return items, err
}

// GetItem returns an inventory whose ID matches itemID
func GetItem(itemID int) (*Item, error) {
	var item Item
	result := DB.First(&item, itemID)
	if result.Error == nil && result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &item, result.Error
}

// CreateItem creates an item
func CreateItem(inventoryID, shipmentID *int, count int, name string) error {
	item := Item{
		Name:        name,
		Count:       count,
		InventoryID: inventoryID,
		ShipmentID:  shipmentID,
	}
	return DB.Create(&item).Error
}

// DeleteItem deletes an item based on itemID
func DeleteItem(itemID int) error {
	result := DB.Delete(&Item{}, itemID)
	if result.Error == nil && result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

// UpdateItem updates an item based on itemID
// Throws an error if shipmentID doesn't exist in DB
func UpdateItem(itemID int, count int, shipmentID *int, name string) error {
	var shipment Shipment
	result := DB.Where("shipment_id = ?", shipmentID).Find(&shipment)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	item := Item{
		Name:       name,
		ShipmentID: shipmentID,
		Count:      count,
	}

	result = DB.Model(&Item{}).Where("id = ?", itemID).Updates(item)
	if result.Error == nil && result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
