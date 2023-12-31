package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/One-lab-Homework-1/api"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
	"log"
	"net/http"
)

// createUser registration new User
//  @Summary      Create User
//  @Description  Create new User
//  @Tags         auth
//  @Accept       json
//  @Produce      json
//  @Param        req body api.RegisterRequest true "req body"
//  @Success      201
//  @Failure      400  {object} api.Error
//  @Failure      500  {object}  api.Error
//  @Router       /register [post]

func (h *Handler) createUser(ctx *gin.Context) {
	var req api.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	u := &entity.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}

	err = h.srvs.CreateUser(ctx, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	accessToken, err := h.srvs.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    accessToken,
	})
}

func (h *Handler) userItems(ctx *gin.Context) {
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}

	items, err := h.srvs.GetuserItems(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (h *Handler) buyItem(ctx *gin.Context) {
	id := ctx.Param("id")

	itemID := 0
	_, err := fmt.Sscanf(id, "%d", &itemID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	userID, ok := ctx.MustGet(authUserID).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    -1,
			Message: "can't get user id from auth",
		})
		return
	}
	err = h.srvs.BuyItem(ctx, int64(itemID), userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusAccepted)
}
