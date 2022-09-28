package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	accountRoute "backendbillingdashboard/modules/account/routes"
	categoryRoute "backendbillingdashboard/modules/category/routes"
	companyRoute "backendbillingdashboard/modules/company/routes"
	configRoute "backendbillingdashboard/modules/config/routes"
	currencyRoute "backendbillingdashboard/modules/currency/routes"
	exchangeRateRoute "backendbillingdashboard/modules/exchange-rate/routes"
	invoiceTypeRoute "backendbillingdashboard/modules/invoice-type/routes"
	invoiceRoute "backendbillingdashboard/modules/invoice/routes"
	itemPriceRoute "backendbillingdashboard/modules/item-price/routes"
	itemRoute "backendbillingdashboard/modules/item/routes"
	paymentMethodRoute "backendbillingdashboard/modules/payment-method/routes"
	paymentRoute "backendbillingdashboard/modules/payment/routes"
	serverAccountRoute "backendbillingdashboard/modules/server-account/routes"
	serverDataRoute "backendbillingdashboard/modules/server-data/routes"
	serverRoute "backendbillingdashboard/modules/server/routes"
	stubRoute "backendbillingdashboard/modules/stub/routes"
)

// InitRoutes handle all route requests
func InitRoutes(conn *connections.Connections, version, builddate string) {
	core.InitRoutes(conn, version, builddate)

	// another new module route will be registered here
	stubRoute.InitRoutes(conn)
	companyRoute.InitRoutes(conn)
	accountRoute.InitRoutes(conn)
	itemRoute.InitRoutes(conn)
	categoryRoute.InitRoutes(conn)
	itemPriceRoute.InitRoutes(conn)
	serverRoute.InitRoutes(conn)
	invoiceRoute.InitRoutes(conn)
	serverDataRoute.InitRoutes(conn)
	invoiceTypeRoute.InitRoutes(conn)
	configRoute.InitRoutes(conn)
	paymentRoute.InitRoutes(conn)
	paymentMethodRoute.InitRoutes(conn)
	serverAccountRoute.InitRoutes(conn)
	currencyRoute.InitRoutes(conn)
	exchangeRateRoute.InitRoutes(conn)
	// ...
}
