package handlers

import (
	spendsService "finance/internal/spends/service"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SpendsHandlers interface {
	GetSpend(id int) (spendsService.Spend, int)
	GetAllSpends() ([]spendsService.Spend, error)
	PostSpend(spendsService.Spend) (spendsService.Spend, error)
	PatchSpend(id int, spend spendsService.Spend) (oldSpend spendsService.Spend, newSpend spendsService.Spend, err error)
	DeleteSpend(id int) error
}

type spendsHandlers struct {
	service spendsService.SpendsService
}

func NewSpendsHandlers(s spendsService.SpendsService) spendsHandlers {
	return spendsHandlers{service: s}
}

func (h *spendsHandlers) GetSpend(c echo.Context) error {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid UUID",
		})
	}

	spend, err := h.service.GetSpend(id)

	if err != nil {
		return c.String(500, "Ошибка GET")
	}

	return c.JSON(201, spend)
}

func (h *spendsHandlers) GetAllSpends(c echo.Context) error {
	spends, err := h.service.GetAllSpends()

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(201, spends)
}

func (h *spendsHandlers) PostSpend(c echo.Context) error {
	reqSpend := spendsService.Spend{}

	if err := c.Bind(&reqSpend); err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]string{"error": "invalid JSON"})
	}

	spend, err := h.service.CreateSpend(reqSpend)

	if err != nil {
		return c.JSON(400, map[string]string{"error": "post handler was not worked"})
	}

	return c.JSON(201, spend)
}

func (h *spendsHandlers) PatchSpend(c echo.Context) error {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid UUID",
		})
	}

	reqSpend := spendsService.Spend{}

	if err := c.Bind(&reqSpend); err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]string{"error": "invalid JSON"})
	}

	spend, err := h.service.UpdateSpend(id, reqSpend)

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(201, spend)
}

func (h *spendsHandlers) DeleteSpend(c echo.Context) error {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid UUID",
		})
	}

	err = h.service.DeleteSpend(id)

	if err != nil {
		return err
	}

	return c.JSON(201, "ok")
}
