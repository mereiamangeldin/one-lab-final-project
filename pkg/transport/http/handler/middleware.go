package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Manager) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(authorizationHeader)
		if header == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "empty auth header")
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth header")
		}
		userId, err := h.srv.Auth.ParseToken(headerParts[1])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		ctx := context.WithValue(c.Request().Context(), userCtx, userId)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func (h *Manager) UserWeakIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(authorizationHeader)
		if header == "" {
			return next(c)
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			return next(c)
		}
		userId, err := h.srv.Auth.ParseToken(headerParts[1])
		if err != nil {
			return next(c)
		}
		ctx := context.WithValue(c.Request().Context(), userCtx, userId)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
