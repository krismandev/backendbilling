package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/company/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	CompanyRoute(conn)
}

// CompanyRoute is used for
func CompanyRoute(conn *connections.Connections) {
	companyRoute := mux.NewRouter()
	companyRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListCompanyEndpoint(conn),
		transport.CompanyDecodeRequest,
		transport.CompanyListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	companyRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateCompanyEndpoint(conn),
		transport.CompanyDecodeRequest,
		transport.CompanySingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	companyRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateCompanyEndpoint(conn),
		transport.CompanyDecodeRequest,
		transport.CompanySingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	companyRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteCompanyEndpoint(conn),
		transport.CompanyDecodeRequest,
		transport.CompanySingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/company", companyRoute)
}
