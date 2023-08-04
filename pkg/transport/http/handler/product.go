package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"net/http"
	"strconv"
)

func (h *Manager) GetProducts(c echo.Context) error {
	auth := false
	userID := c.Request().Context().Value(userCtx)
	if userID != nil {
		auth = true
	}
	title := c.QueryParam("title")
	var products []model.Product
	var err error
	if title != "" {
		products, err = h.srv.Product.GetProductsByTitle(title)
	} else {
		products, err = h.srv.Product.GetProducts()
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if auth {
		products, err = h.srv.User.GetLikedUserProducts(products, userID.(uint))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, products)
}

func (h *Manager) GetCategoryProducts(c echo.Context) error {
	auth := false
	userID := c.Request().Context().Value(userCtx)
	if userID != nil {
		auth = true
	}
	id_ := c.Param("category_id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	products, err := h.srv.Product.GetCategoryProducts(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if auth {
		products, err = h.srv.User.GetLikedUserProducts(products, userID.(uint))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusOK, products)
}

func (h *Manager) GetCategories(c echo.Context) error {
	categories, err := h.srv.Product.GetCategories()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, categories)
}

func (h *Manager) CreateProduct(c echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	id, err := h.srv.Product.CreateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Manager) UpdateProduct(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	res := h.srv.Product.UpdateProduct(uint(id), product)
	if res != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Product updated successfully",
	})
}
func (h *Manager) DeleteProduct(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.srv.Product.DeleteProduct(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Product deleted successfully",
	})
}

func (h *Manager) GetProductById(c echo.Context) error {
	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	product, err := h.srv.Product.GetProductById(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "No such product",
		})
	}
	return c.JSON(http.StatusOK, product)
}

func (h *Manager) LikeAction(c echo.Context) error {
	userID := c.Request().Context().Value(userCtx).(uint)
	id_ := c.Param("product_id")
	productID, err := strconv.Atoi(id_)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = h.srv.Product.LikeAction(userID, uint(productID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"error": "like error",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Product liked successfully",
	})
}
