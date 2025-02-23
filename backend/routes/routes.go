package routes

import (
	"voucher-app/handler"
	"voucher-app/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Router defines the router structure
type Router struct {
	brandService service.BrandService
	brandHandler handler.BrandHandler

	voucherService service.VoucherService
	voucherHandler handler.VoucherHandler

	transactionHandler handler.TransactionHandler
	transactionService service.TransactionService
}

// MakeRouter creates a new router
func MakeRouter(db *gorm.DB) Router {
	brandService := service.NewBrandService(db)
	brandHandler := handler.NewBrandHandler(brandService)

	voucherService := service.NewVoucherService(db)
	voucherHandler := handler.NewVoucherHandler(voucherService)

	transactionService := service.NewTransactionService()
	transactionHandler := handler.NewTransactionHandler(transactionService)

	return Router{
		brandService:       *brandService,
		brandHandler:       *brandHandler,
		voucherService:     *voucherService,
		voucherHandler:     *voucherHandler,
		transactionService: *transactionService,
		transactionHandler: *transactionHandler,
	}
}

// InitRoutes initializes routes
func (r Router) InitRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")

	api.POST("/brand", r.brandHandler.CreateBrand)

	api.POST("/voucher", r.voucherHandler.CreateVoucher)
	api.GET("/voucher", r.voucherHandler.GetVoucher)
	api.GET("/voucher/brand", r.voucherHandler.GetVoucherByBrand)

	api.POST("/transaction/redemption", r.transactionHandler.CreateRedemptionTransaction)
	api.GET("/transaction/redemption", r.transactionHandler.GetRedemptionTransactionDetail)
}
