package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/pkg"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var inp domain.CreateAccountInput
	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if err := pkg.ValidateAccount(inp.Name, inp.Balance); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid input: %v", err))
		return
	}

	id, err := h.service.CreateAccount(c, &inp)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("error with creating account: %v", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) GetAllAccountsWithTransactions(c *gin.Context) {
	accounts, err := h.service.GetAllAccountsWithTransactions(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("errors with finding accounts: %v", err))
		return
	}

	c.JSON(http.StatusOK, accounts)
}
