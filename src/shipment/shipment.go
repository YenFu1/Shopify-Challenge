package shipment

import (
	"Shopify-Challenge/src/database"
	"Shopify-Challenge/src/helper"
	"Shopify-Challenge/src/logger"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetShipments(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for GetInventories")
	shipments, err := database.GetShipments()
	if err != nil {
		logger.Sugar.Errorf("failed to get shipments: %+v", err)
	}

	var getShipmentsResponse GetShipmentsResponse
	for _, s := range shipments {
		items, err := database.GetItemsByShipmentID(s.ID)
		if err != nil {
			logger.Sugar.Errorf("failed to get item from database: %+v", err)
			helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
			return
		}
		shipment := Shipment{
			Items:    items,
			Shipment: s,
		}
		getShipmentsResponse.Shipments = append(getShipmentsResponse.Shipments, shipment)
	}

	body, err := json.Marshal(getShipmentsResponse)
	if err != nil {
		logger.Sugar.Errorf("failed to marshal response body")
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusOK, body)
}

func CreateShipment(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for CreateShipment")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Sugar.Errorf("failed to read request body: %+v, err: %+v", r.Body, err)
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	var createShipmentRequest CreateShipmentRequest
	if err := json.Unmarshal(body, &createShipmentRequest); err != nil {
		logger.Sugar.Errorf("failed to unmarshal request body: %+v, err: %+v", createShipmentRequest, err)
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	if err := database.CreateShipment(
		createShipmentRequest.InventoryID,
		createShipmentRequest.Name,
		createShipmentRequest.Items,
	); err != nil {
		logger.Sugar.Errorf("failed to create shipment: %+v", err)
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusCreated, nil)
}
