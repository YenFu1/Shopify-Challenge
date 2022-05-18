package router

import (
	"Shopify-Challenge/src/inventory"
	"Shopify-Challenge/src/item"
	"Shopify-Challenge/src/shipment"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Route("/inventories", func(r chi.Router) {
			r.Post("/", inventory.CreateInventory)
			r.Get("/", inventory.GetInventories)

			r.Route("/{inventoryID}", func(r chi.Router) {
				r.Use(inventory.InventoryCtx)
				r.Put("/", inventory.UpdateInventory)
				r.Delete("/", inventory.DeleteInventory)

				r.Route("/items", func(r chi.Router) {
					r.Post("/", item.CreateItem)
					r.Get("/", item.GetItems)

					r.Route("/{itemID}", func(r chi.Router) {
						r.Use(item.ItemCtx)
						r.Put("/", item.UpdateItem)
						r.Delete("/", item.DeleteItem)
					})
				})
			})
		})

		r.Route("/shipments", func(r chi.Router) {
			r.Get("/", shipment.GetShipments)
			r.Post("/", shipment.CreateShipment)
		})
	})

	return r
}
