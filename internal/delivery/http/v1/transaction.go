package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/pkg"
)

// @Summary Create a new transaction
// @Description Creates a new transaction by taking a value and account id,group.Account2_id optional
// @Tags transaction
// @Accept  json
// @Produce  json
// @Param   input  body      domain.CreateTransactionInput  true  "Transaction Creation Data"
// @Success 201  {string} string "Successfully created"
// @Failure 400  {object}  Response
// @Failure 500 {object} Response
// @Router /transactions/create [post]
func (h *Handler) CreateTransaction(c *gin.Context) {
	fmt.Println("OK")
	var inp domain.CreateTransactionInput

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if err := pkg.ValidateTransaction(inp.Value, inp.Group, inp.Account_id, inp.Account2_id); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid input: %v", err))
		return
	}

	if err := h.service.CreateTransaction(c, inp); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("error with creating transaction: %v", err))
		return
	}

	c.JSON(http.StatusCreated, "Successfully created")
}

// @Summary Get transactions
// @Description Getting transactions by  account id
// @Tags transaction
// @Produce json
// @Success 200 {object} []entity.Transaction
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Router /{id} [get]
func (h *Handler) GetAllTransactionsByAccountId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid id: %v", err))
		return
	}

	transactions, err := h.service.GetAllTransactionsByAccountId(c, id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("error with gettung transactions by id: %v", err))
		return
	}
	if len(transactions) == 0 {
		errorResponse(c, http.StatusNotFound, fmt.Errorf("there`s no such transactions by this id"))
		return
	}
	c.JSON(http.StatusOK, transactions)
}
