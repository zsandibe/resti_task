package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/pkg"
)

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
