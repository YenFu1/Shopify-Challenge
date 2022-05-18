package database

import (
	"errors"

	"gorm.io/gorm"
)

// GetShipments returns all the shipments in the database
func GetShipments() ([]Shipment, error) {
	var shipments []Shipment
	err := DB.Find(&shipments).Error
	return shipments, err
}

// CreateShipment creates a shipment entry in the database and
// updates the inventory to reflect the moved items
func CreateShipment(inventoryID int, name string, items []Item) error {
	// run as transaction
	return DB.Transaction(func(tx *gorm.DB) error {
		shipment := Shipment{Name: name}
		if err := tx.Create(&shipment).Error; err != nil {
			return err
		}

		for _, i := range items {
			var item Item
			if err := tx.Model(&Item{}).Where("id = ?", i.ID).Find(&item).Error; err != nil {
				return err
			}

			newCount := item.Count - i.Count
			if newCount < 0 {
				return errors.New("not enough items in stock for shipment")
			}

			// update item count in inventory
			result := tx.Model(&Item{}).Where("id = ?", i.ID).Update("count", newCount)
			if result.Error == nil && result.RowsAffected == 0 {
				return gorm.ErrRecordNotFound
			} else if result.Error != nil {
				return result.Error
			}

			// create new entry in item table representing the item in shipment
			shipmentItem := Item{
				Name:       item.Name,
				Count:      i.Count,
				ShipmentID: &shipment.ID,
			}

			if err := tx.Create(&shipmentItem).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
