package account_handlers

import (
	"log"
	"net/http"

	service "finance/internal/accounts/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AccountsHandlers struct {
	service service.AccountsService
}

func NewAccountsHandlers(svc service.AccountsService) *AccountsHandlers {
	return &AccountsHandlers{service: svc}
}

func (h *AccountsHandlers) GetAll(c echo.Context) error {
	accounts, err := h.service.GetAllAccounts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, accounts)
}

func (h *AccountsHandlers) GetOne(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	acc, err := h.service.GetAccount(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "account not found"})
	}
	return c.JSON(http.StatusOK, acc)
}

func (h *AccountsHandlers) Create(c echo.Context) error {
	log.Println("Account Service Create 1")

	var req service.AccountCreateDTO
	if err := c.Bind(&req); err != nil {
		log.Println("Account Service Create 2", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	log.Println("Account Service Create 3", req)

	created, err := h.service.CreateAccount(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, created)
}

func (h *AccountsHandlers) Update(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	var req service.AccountUpdateDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	updated, err := h.service.UpdateAccount(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

func (h *AccountsHandlers) Delete(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid UUID"})
	}

	if err := h.service.DeleteAccount(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
