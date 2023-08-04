package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/One-lab-Homework-1/api"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
	"log"
	"net/http"
)

func (h *Handler) createItem(ctx *gin.Context) {
	var item entity.Item

	err := ctx.ShouldBindJSON(&item)

	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateItem(ctx, &item)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) getItems(ctx *gin.Context) {
	items, err := h.srvs.GetItems(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, items)
}
func (h *Handler) getItemByID(ctx *gin.Context) {
	id := ctx.Param("id")

	itemID := 0
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	item, err := h.srvs.GetItemByID(ctx, int64(itemID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, item)

}

func (h *Handler) deleteItem(ctx *gin.Context) {
	id := ctx.Param("id")

	itemID := 0
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.srvs.DeleteItem(ctx, int64(itemID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (h *Handler) updateItem(ctx *gin.Context) {
	id := ctx.Param("id")

	itemID := 0
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var item entity.Item

	err = ctx.ShouldBindJSON(&item)

	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.UpdateItem(ctx, int64(itemID), item)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusOK)

}
