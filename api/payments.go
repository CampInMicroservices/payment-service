package api

import (
	"net/http"
	"payment-service/db"

	"github.com/gin-gonic/gin"
)

type getPaymentRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getPaymentListRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
}

type createPaymentRequest struct {
	BookingID int64   `json:"booking_id" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Paid      *bool   `json:"paid" binding:"required"`
}

type updatePaymentRequest struct {
	Paid *bool `json:"paid" binding:"required"`
}

func (server *Server) GetPaymentByID(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getPaymentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.GetPaymentByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) GetAllPayments(ctx *gin.Context) {

	// Check if request has parameters offset and limit for pagination.
	var req getPaymentListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.ListPaymentParam{
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	// Execute query.
	result, err := server.store.GetAllPayments(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) CreatePayment(ctx *gin.Context) {

	// Check if request has all required fields in json body.
	var req createPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.CreatePaymentParam{
		BookingID: req.BookingID,
		Price:     req.Price,
		Paid:      *req.Paid,
	}

	// Execute query.
	result, err := server.store.CreatePayment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) UpdatePayment(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var uri getPaymentRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Check if request has all required fields in json body.
	var req updatePaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.UpdatePaymentParam{
		Paid: *req.Paid,
	}

	// Execute query.
	result, err := server.store.UpdatePayment(ctx, arg, uri.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}
