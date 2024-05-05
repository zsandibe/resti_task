package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/pkg"
)

type AccountResponse struct {
	Id int `json:"id"`
}

// @Summary Create a new account
// @Description Creates a new account by taking a name and initial balance
// @Tags account
// @Accept  json
// @Produce  json
// @Param   input  body      domain.CreateAccountInput  true  "Account Creation Data"
// @Success 201  {object} AccountResponse
// @Failure 400  {object}  Response
// @Failure 500 {object} Response
// @Router /accounts/create [post]
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
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("error with creating account: %v", err))
		return
	}

	resp := &AccountResponse{
		Id: id,
	}
	c.JSON(http.StatusCreated, resp)
}

// @Summary Get accounts
// @Description Getting account with transactions
// @Tags account
// @Produce json
// @Success 200 {object} []domain.GetAccountOutput
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Router / [get]
func (h *Handler) GetAllAccountsWithTransactions(c *gin.Context) {
	accounts, err := h.service.GetAllAccountsWithTransactions(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("errors with finding accounts: %v", err))
		return
	}

	c.JSON(http.StatusOK, accounts)
}
