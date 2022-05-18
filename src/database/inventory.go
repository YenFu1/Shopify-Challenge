package database

import (
	"gorm.io/gorm"
)

// GetInventories returns a slice containing all inventories
func GetInventories() ([]Inventory, error) {
	var inventories []Inventory
	err := DB.Find(&inventories).Error
	return inventories, err
}

// GetInventory returns an inventory whose ID matches inventoryID
func GetInventory(inventoryID int) (*Inventory, error) {
	var inventory Inventory
	result := DB.First(&inventory, inventoryID)
	if result.Error == nil && result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &inventory, result.Error
}

// CreateInventory creates an inventory
func CreateInventory(name string) error {
	return DB.Create(&Inventory{Name: name}).Error
}

// DeleteInventory deletes an inventory based on inventoryID
func DeleteInventory(inventoryID int) error {
	result := DB.Delete(&Inventory{}, inventoryID)
	if result.Error == nil && result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

// UpdateInventory updates an inventory based on inventoryID
func UpdateInventory(inventoryID int, name string) error {
	result := DB.Model(&Inventory{}).Where("id = ?", inventoryID).Update("Name", name)
	if result.Error == nil && result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
