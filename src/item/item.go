package item

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
)

// ItemCtx is a middleware which extracts the itemID from the URL path
// and sets it in the context
func ItemCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemIDParam := chi.URLParam(r, "itemID")
		if itemIDParam == "" {
			logger.Sugar.Error("no itemID provided in route")
			helper.ServeResponse(w, r, http.StatusNotFound, nil)
			return
		}

		itemID, err := strconv.Atoi(itemIDParam)
		if err != nil {
			logger.Sugar.Errorf("failed to parse itemID from route: %+v", itemID)
			helper.ServeResponse(w, r, http.StatusNotFound, nil)
			return
		}

		item, err := database.GetItem(itemID)
		if item == nil || err != nil {
			logger.Sugar.Errorf("failed to find item entry in database for itemID: %d", itemID)
			helper.ServeResponse(w, r, http.StatusNotFound, nil)
			return
		}

		ctx := context.WithValue(r.Context(), "itemID", itemID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for GetItems")

	// shouldn't need to validate inventoryID since that's handled in middleware
	inventoryID := r.Context().Value("inventoryID").(int)

	items, err := database.GetItemsByInventoryID(inventoryID)
	if err != nil {
		logger.Sugar.Errorf("failed to get items: %+v", err)
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	getItemsResponse := GetItemsResponse{
		Count: len(items),
		Items: items,
	}

	body, err := json.Marshal(getItemsResponse)
	if err != nil {
		logger.Sugar.Errorf("failed to marshal response: %+v", getItemsResponse)
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusOK, body)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for CreateItem")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Sugar.Errorf("failed to read request body: %+v, err: %+v", r.Body, err)
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	var createItemRequest CreateItemRequest
	if err := json.Unmarshal(body, &createItemRequest); err != nil {
		logger.Sugar.Errorf("failed to unmarshal body: %+v, err: %+v", createItemRequest, err)
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	// shouldn't need to validate inventoryID since that's handled in middleware
	inventoryID := r.Context().Value("inventoryID").(int)
	if err := database.CreateItem(&inventoryID, nil, createItemRequest.Count, createItemRequest.Name); err != nil {
		logger.Sugar.Errorf("failed to create item: %+v", err)
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusCreated, nil)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("request received for DeleteItem")

	// shouldn't need to validate itemID since that's handled in middleware
	itemID := r.Context().Value("itemID").(int)

	if err := database.DeleteItem(itemID); err != nil {
		logger.Sugar.Errorf("failed to delete item: %+v", err)
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusNoContent, nil)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	logger.Sugar.Info("reequest received for UpdateItem")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Sugar.Errorf("failed to read request body: %+v, err: %+v", r.Body, err)
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	var updateItemRequest UpdateItemRequest
	if err := json.Unmarshal(body, &updateItemRequest); err != nil {
		logger.Sugar.Errorf("failed to unmarshal request body: %+v, err: %+v", updateItemRequest, err)
		helper.ServeResponse(w, r, http.StatusBadRequest, []byte(helper.INVALID_BODY))
		return
	}

	// shouldn't need to validate itemID since that's handled in middleware
	itemID := r.Context().Value("itemID").(int)
	if err := database.UpdateItem(
		itemID,
		updateItemRequest.Count,
		updateItemRequest.Name,
	); err != nil {
		logger.Sugar.Errorf("failed to update item: %+v", err)
		helper.ServeResponse(w, r, http.StatusInternalServerError, []byte(helper.UNKNOWN_ERROR))
		return
	}

	helper.ServeResponse(w, r, http.StatusNoContent, nil)
}
