package handler

import (
	"github.com/labstack/echo/v4"
	model2 "github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"net/http"
	"strconv"
)

func (h *Manager) UpdateUser(c echo.Context) error {
	id := c.Request().Context().Value(userCtx).(uint)
	var user model2.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res := h.srv.User.Update(uint(id), user)
	if res != nil {
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "User updated successfully",
	})
}

func (h *Manager) DepositBalance(c echo.Context) error {
	id := c.Request().Context().Value(userCtx).(uint)
	res := h.srv.User.DepositBalance(uint(id))
	if res != nil {
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Balance updated successfully",
	})
}

func (h *Manager) UpdatePassword(c echo.Context) error {
	id := c.Request().Context().Value(userCtx).(uint)
	var pass model2.ChangePassword
	if err := c.Bind(&pass); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res := h.srv.Auth.UpdatePassword(uint(id), pass)
	if res != nil {
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Password changed successfully",
	})
}

func (h *Manager) DeleteUser(c echo.Context) error {
	id := c.Request().Context().Value(userCtx).(uint)
	err := h.srv.User.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "User deleted successfully",
	})
}

func (h *Manager) GetUserById(c echo.Context) error {
	id := c.Request().Context().Value(userCtx).(uint)
	user, err := h.srv.User.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "No such user",
		})
	}
	return c.JSON(http.StatusOK, user)
}
func (h *Manager) GetBalance(c echo.Context) error {
	id_ := c.Param("id")
	_, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{})
}

func (h *Manager) GetUserProducts(c echo.Context) error {
	userID := c.Request().Context().Value(userCtx).(uint)
	var products []model2.Product
	var err error
	products, err = h.srv.User.GetUserProducts(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (h *Manager) GetUserFavoriteProducts(c echo.Context) error {
	userID := c.Request().Context().Value(userCtx).(uint)
	var products []model2.Product
	var err error
	products, err = h.srv.User.GetUserFavoriteProducts(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	products, err = h.srv.User.GetLikedUserProducts(products, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (h *Manager) BuyProduct(c echo.Context) error {
	userID := c.Request().Context().Value(userCtx).(uint)
	id_ := c.Param("product_id")
	productID, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.srv.User.BuyProduct(userID, uint(productID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "purchase error",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Product purchased successfully",
	})
}
