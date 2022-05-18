package inventory

import (
	"Shopify-Challenge/src/database"
	"Shopify-Challenge/src/helper"
	"Shopify-Challenge/src/logger"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// InventoryCtx is a middleware which extracts the inventoryID from the URL path
// and sets it in the context
func InventoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inventoryIDParam := chi.URLParam(r, "inventoryID")
		if inventoryIDParam == "" {
			logger.Sugar.Error("no inventoryID provided in route")
			helper.ServeResponse(w, r, http.StatusNotFound, nil)
			return
		}

		inventoryID, err := strconv.Atoi(inventoryIDParam)
		if err != nil {
			logger.Sugar.Errorf("failed to parse inventoryID from route: %+v", inventoryID)
			helper.ServeResponse(w, r, http.StatusNotFound, nil)
			return
		}

		inventory, err := database.GetInventory(inventoryID)
		if inventory == nil || err != nil {
			helper.ServeResponse(w, r, http.StatusNotFound, nil)
			return
		}

		ctx := context.WithValue(r.Context(), "inventoryID", inventoryID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetInventories(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for GetInventories")
	inventories, err := database.GetInventories()
	if err != nil {
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	getInventoriesResponse := GetInventoriesResponse{
		Count:       len(inventories),
		Inventories: inventories,
	}
	body, err := json.Marshal(getInventoriesResponse)
	if err != nil {
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusOK, body)
}

func CreateInventory(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for CreateInventory")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	var createInventoryRequest CreateInventoryRequest
	if err := json.Unmarshal(body, &createInventoryRequest); err != nil {
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	if err := database.CreateInventory(createInventoryRequest.Name); err != nil {
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusCreated, nil)
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for DeleteInventory")

	// shouldn't need to validate inventoryID since that's handled in middleware
	inventoryID := r.Context().Value("inventoryID").(int)

	if err := database.DeleteInventory(inventoryID); err == gorm.ErrRecordNotFound {
		helper.ServeResponse(w, r, http.StatusNotFound, nil)
		return
	} else if err != nil {
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusNoContent, nil)
}

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for UpdateInventory")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	var updateInventoryRequest UpdateInventoryRequest
	if err := json.Unmarshal(body, &updateInventoryRequest); err != nil {
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	// shouldn't need to validate inventoryID since that's handled in middleware
	inventoryID := r.Context().Value("inventoryID").(int)

	if err := database.UpdateInventory(inventoryID, updateInventoryRequest.Name); err == gorm.ErrRecordNotFound {
		helper.ServeResponse(w, r, http.StatusNotFound, nil)
		return
	} else if err != nil {
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusNoContent, nil)
}
