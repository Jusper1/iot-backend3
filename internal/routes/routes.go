package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"iot-backend/internal/handlers"
	masterHandlers "iot-backend/internal/handlers/master"
	"iot-backend/internal/middleware"
	"iot-backend/internal/repositories"
	masterRepo "iot-backend/internal/repositories/master"
	"iot-backend/internal/services"
	masterService "iot-backend/internal/services/master"
)


func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	api := r.Group("/api")

	// PUBLIC ROUTES


	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	// REPOSITORY

	instansiRepository := masterRepo.NewInstansiRepository(db)
	picRepository := masterRepo.NewPICRepository(db)
	produkRepository := masterRepo.NewProdukRepository(db)
	wilayahRepository := masterRepo.NewWilayahRepository(db)
	ekspedisiRepository := masterRepo.NewEkspedisiRepository(db)

	orderRepository := repositories.NewOrderRepository(db)
	orderItemRepository := repositories.NewOrderItemRepository(db)
	inaprocRepository := repositories.NewOrderInaprocRepository(db)
	manualRepository := repositories.NewOrderManualRepository(db)
	procurementRepository := repositories.NewOrderProcurementRepository(db)
	pricingRepository := repositories.NewOrderPricingRepository(db)
	orderDocumentRepository := repositories.NewOrderDocumentRepository(db)
	paymentRepository := repositories.NewPaymentRepository(db)
	spjRepository := repositories.NewSPJRepository(db)
	documentRepository := repositories.NewDocumentRepository(db)
	shipmentRepository := repositories.NewShipmentRepository(db)
	iotInaprocRepository := repositories.NewIOTInaprocRepository(db)

	// SERVICE

	instansiServiceInst := masterService.NewInstansiService(instansiRepository)
	picServiceInst := masterService.NewPICService(picRepository)
	produkServiceInst := masterService.NewProdukService(produkRepository)
	wilayahServiceInst := masterService.NewWilayahService(wilayahRepository)
	ekspedisiServiceInst := masterService.NewEkspedisiService(ekspedisiRepository)

	orderServiceInst := services.NewOrderService(orderRepository)
	orderItemServiceInst := services.NewOrderItemService(orderItemRepository)
	inaprocServiceInst := services.NewOrderInaprocService(inaprocRepository)
	manualServiceInst := services.NewOrderManualService(manualRepository)
	procurementServiceInst := services.NewOrderProcurementService(procurementRepository)
	pricingServiceInst := services.NewOrderPricingService(pricingRepository)
	orderDocumentServiceInst := services.NewOrderDocumentService(orderDocumentRepository)
	paymentServiceInst := services.NewPaymentService(paymentRepository)
	spjServiceInst := services.NewSPJService(spjRepository)
	documentServiceInst := services.NewDocumentService(documentRepository)
	shipmentServiceInst := services.NewShipmentService(shipmentRepository)
	iotInaprocService := services.NewIOTInaprocService(iotInaprocRepository)

	// HANDLER

	instansiHandler := masterHandlers.NewInstansiHandler(instansiServiceInst)
	picHandler := masterHandlers.NewPICHandler(picServiceInst)
	produkHandler := masterHandlers.NewProdukHandler(produkServiceInst)
	wilayahHandler := masterHandlers.NewWilayahHandler(wilayahServiceInst)
	ekspedisiHandler := masterHandlers.NewEkspedisiHandler(ekspedisiServiceInst)

	orderHandler := handlers.NewOrderHandler(orderServiceInst)
	orderItemHandler := handlers.NewOrderItemHandler(orderItemServiceInst)
	inaprocHandler := handlers.NewOrderInaprocHandler(inaprocServiceInst)
	manualHandler := handlers.NewOrderManualHandler(manualServiceInst)
	procurementHandler := handlers.NewOrderProcurementHandler(procurementServiceInst)
	pricingHandler := handlers.NewOrderPricingHandler(pricingServiceInst)
	orderDocumentHandler := handlers.NewOrderDocumentHandler(orderDocumentServiceInst)
	paymentHandler := handlers.NewPaymentHandler(paymentServiceInst)
	spjHandler := handlers.NewSPJHandler(spjServiceInst)
	documentHandler := handlers.NewDocumentHandler(documentServiceInst)
	shipmentHandler := handlers.NewShipmentHandler(shipmentServiceInst)
	iotInaprocHandler := handlers.NewIOTInaprocHandler(iotInaprocService)

	// PROTECTED ROUTES

	protected := api.Group("/")
	protected.Use(middleware.JWTAuth())

	// PROFILE

	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Token valid",
			"user": gin.H{
				"id":    c.MustGet("user_id"),
				"name":  c.MustGet("name"),
				"email": c.MustGet("email"),
				"role":  c.MustGet("role"),
			},
		})
	})

	// MASTER INSTANSI

	protected.GET("/instansi", instansiHandler.FindAll)
	protected.GET("/instansi/search", instansiHandler.Search) // static route, taruh sebelum /:id
	protected.GET("/instansi/:id", instansiHandler.FindByID)
	protected.POST("/instansi", instansiHandler.Create)
	protected.PUT("/instansi/:id", instansiHandler.Update)
	protected.DELETE("/instansi/:id", instansiHandler.Delete)

	// MASTER PIC

	protected.GET("/pic", picHandler.FindAll)
	protected.GET("/pic/instansi/:instansi_id", picHandler.FindByInstansiID)
	protected.GET("/pic/:id", picHandler.FindByID)
	protected.POST("/pic", picHandler.Create)
	protected.PUT("/pic/:id", picHandler.Update)
	protected.DELETE("/pic/:id", picHandler.Delete)

	// MASTER PRODUK

	protected.GET("/produk", produkHandler.FindAll)
	protected.GET("/produk/active", produkHandler.FindActive)
	protected.GET("/produk/kode/:kode", produkHandler.FindByKode)
	protected.GET("/produk/:id", produkHandler.FindByID)
	protected.POST("/produk", produkHandler.Create)
	protected.PUT("/produk/:id", produkHandler.Update)
	protected.DELETE("/produk/:id", produkHandler.Delete)

	// MASTER WILAYAH

	protected.GET("/wilayah", wilayahHandler.FindAll)
	protected.GET("/wilayah/provinsi/:provinsi", wilayahHandler.FindByProvinsi)
	protected.GET("/wilayah/:id", wilayahHandler.FindByID)
	protected.POST("/wilayah", wilayahHandler.Create)
	protected.PUT("/wilayah/:id", wilayahHandler.Update)
	protected.DELETE("/wilayah/:id", wilayahHandler.Delete)

	// MASTER EKSPEDISI

	protected.GET("/ekspedisi", ekspedisiHandler.FindAll)
	protected.GET("/ekspedisi/active", ekspedisiHandler.FindActive)
	protected.GET("/ekspedisi/:id", ekspedisiHandler.FindByID)
	protected.POST("/ekspedisi", ekspedisiHandler.Create)
	protected.PUT("/ekspedisi/:id", ekspedisiHandler.Update)
	protected.DELETE("/ekspedisi/:id", ekspedisiHandler.Delete)

	// ORDERS

	protected.GET("/orders", orderHandler.FindAll)
	protected.GET("/orders/kode/:kode", orderHandler.FindByKode)
	protected.GET("/orders/:id", orderHandler.FindByID)
	protected.POST("/orders", orderHandler.Create)
	protected.PUT("/orders/:id", orderHandler.Update)
	protected.DELETE("/orders/:id", orderHandler.Delete)

	// ORDER ITEMS

	// Catatan: OrderItemHandler tidak punya method FindAll, hanya FindByOrderID/FindByID.
	protected.GET("/order-items/order/:order_id", orderItemHandler.FindByOrderID)
	protected.GET("/order-items/:id", orderItemHandler.FindByID)
	protected.POST("/order-items", orderItemHandler.Create)
	protected.POST("/order-items/bulk", orderItemHandler.CreateMany)
	protected.PUT("/order-items/:id", orderItemHandler.Update)
	protected.DELETE("/order-items/:id", orderItemHandler.Delete)

	// INAPROC
	// (handler hanya punya: Create, FindByID, FindByOrderID, Update, DeleteByOrderID
	//  -- TIDAK ada FindAll / Delete by id)

	protected.GET("/inaproc/order/:order_id", inaprocHandler.FindByOrderID)
	protected.GET("/inaproc/:id", inaprocHandler.FindByID)
	protected.POST("/inaproc", inaprocHandler.Create)
	protected.PUT("/inaproc/:id", inaprocHandler.Update)
	protected.DELETE("/inaproc/order/:order_id", inaprocHandler.DeleteByOrderID)

	// MANUAL (order_manual)
	// (handler hanya punya: Create, FindByID, FindByOrderID, Update, DeleteByOrderID
	//  -- TIDAK ada FindAll / Delete by id)

	protected.GET("/manual/order/:order_id", manualHandler.FindByOrderID)
	protected.GET("/manual/:id", manualHandler.FindByID)
	protected.POST("/manual", manualHandler.Create)
	protected.PUT("/manual/:id", manualHandler.Update)
	protected.DELETE("/manual/order/:order_id", manualHandler.DeleteByOrderID)

	// ORDER PROCUREMENT (sebelumnya belum ada di router)
	// (Create, FindByID, FindByOrderID, Update, DeleteByOrderID)

	protected.GET("/order-procurement/order/:order_id", procurementHandler.FindByOrderID)
	protected.GET("/order-procurement/:id", procurementHandler.FindByID)
	protected.POST("/order-procurement", procurementHandler.Create)
	protected.PUT("/order-procurement/:id", procurementHandler.Update)
	protected.DELETE("/order-procurement/order/:order_id", procurementHandler.DeleteByOrderID)

	// ORDER PRICING (sebelumnya belum ada di router)
	// (Create, FindByID, FindByOrderID, Update, DeleteByOrderID)

	protected.GET("/order-pricing/order/:order_id", pricingHandler.FindByOrderID)
	protected.GET("/order-pricing/:id", pricingHandler.FindByID)
	protected.POST("/order-pricing", pricingHandler.Create)
	protected.PUT("/order-pricing/:id", pricingHandler.Update)
	protected.DELETE("/order-pricing/order/:order_id", pricingHandler.DeleteByOrderID)

	// ORDER DOCUMENT (sebelumnya belum ada di router)
	// beda dengan DOCUMENTS di bawah -- ini dokumen bawaan order
	// (Create, FindByID, FindByOrderID, Update, Delete)

	protected.GET("/order-documents/order/:order_id", orderDocumentHandler.FindByOrderID)
	protected.GET("/order-documents/:id", orderDocumentHandler.FindByID)
	protected.POST("/order-documents", orderDocumentHandler.Create)
	protected.PUT("/order-documents/:id", orderDocumentHandler.Update)
	protected.DELETE("/order-documents/:id", orderDocumentHandler.Delete)

	// PAYMENTS

	protected.GET("/payments", paymentHandler.FindAll)
	protected.GET("/payments/order/:order_id", paymentHandler.FindByOrderID)
	protected.GET("/payments/:id", paymentHandler.FindByID)
	protected.POST("/payments", paymentHandler.Create)
	protected.PUT("/payments/:id", paymentHandler.Update)
	protected.DELETE("/payments/:id", paymentHandler.Delete)

	// SHIPMENT / PENGIRIMAN (sebelumnya belum ada di router)
	// beda dengan master EKSPEDISI di atas

	protected.GET("/shipments", shipmentHandler.FindAll)
	protected.GET("/shipments/order/:order_id", shipmentHandler.FindByOrderID)
	protected.GET("/shipments/:id", shipmentHandler.FindByID)
	protected.POST("/shipments", shipmentHandler.Create)
	protected.PUT("/shipments/:id", shipmentHandler.Update)
	protected.DELETE("/shipments/:id", shipmentHandler.Delete)

	// SPJ

	protected.GET("/spj", spjHandler.FindAll)
	protected.GET("/spj/order/:order_id", spjHandler.FindByOrderID)
	protected.GET("/spj/:id", spjHandler.FindByID)
	protected.POST("/spj", spjHandler.Create)
	protected.PUT("/spj/:id", spjHandler.Update)
	protected.DELETE("/spj/:id", spjHandler.Delete)

	// DOCUMENTS (dokumen umum, bisa terkait order ATAU spj)

	protected.GET("/documents", documentHandler.FindAll)
	protected.GET("/documents/order/:order_id", documentHandler.FindByOrderID)
	protected.GET("/documents/spj/:spj_id", documentHandler.FindBySPJID)
	protected.GET("/documents/:id", documentHandler.FindByID)
	protected.POST("/documents", documentHandler.Create)
	protected.PUT("/documents/:id", documentHandler.Update)
	protected.DELETE("/documents/:id", documentHandler.Delete)

	protected.GET("/iot-inaproc", iotInaprocHandler.FindAll)
	protected.GET("/iot-inaproc/kode/:kode", iotInaprocHandler.FindByKode)
	protected.GET("/iot-inaproc/:id", iotInaprocHandler.FindByID)
	protected.POST("/iot-inaproc", iotInaprocHandler.Create)
	protected.PUT("/iot-inaproc/:id", iotInaprocHandler.Update)
	protected.DELETE("/iot-inaproc/:id", iotInaprocHandler.Delete)
}