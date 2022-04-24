package routes

import (
	"billingdashboard/connections"
	"billingdashboard/core"
	accountRoute "billingdashboard/modules/account/routes"
	categoryRoute "billingdashboard/modules/category/routes"
	companyRoute "billingdashboard/modules/company/routes"
	configRoute "billingdashboard/modules/config/routes"
	invoiceTypeRoute "billingdashboard/modules/invoice-type/routes"
	invoiceRoute "billingdashboard/modules/invoice/routes"
	itemPriceRoute "billingdashboard/modules/item-price/routes"
	itemRoute "billingdashboard/modules/item/routes"
	serverDataRoute "billingdashboard/modules/server-data/routes"
	serverRoute "billingdashboard/modules/server/routes"
	stubRoute "billingdashboard/modules/stub/routes"
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
	// ...
}
