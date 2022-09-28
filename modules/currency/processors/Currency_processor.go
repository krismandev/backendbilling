package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/currency/datastruct"
	"backendbillingdashboard/modules/currency/models"
)

func GetListCurrency(conn *connections.Connections, req datastruct.CurrencyRequest) ([]datastruct.CurrencyDataStruct, error) {
	var output []datastruct.CurrencyDataStruct
	var err error

	// grab mapping data from model
	currencyList, err := models.GetCurrencyFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, currency := range currencyList {
		single := CreateSingleCurrencyStruct(currency)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleCurrencyStruct(currency map[string]string) datastruct.CurrencyDataStruct {
	var single datastruct.CurrencyDataStruct
	single.CurrencyCode, _ = currency["currency_code"]
	single.CurrencyName, _ = currency["currency_name"]
	single.Default, _ = currency["default"]
	single.LastUpdateUsername, _ = currency["last_update_username"]
	single.LastUpdateDate, _ = currency["last_update_date"]

	return single
}

func InsertCurrency(conn *connections.Connections, req datastruct.CurrencyRequest) (datastruct.CurrencyDataStruct, error) {
	var output datastruct.CurrencyDataStruct
	var err error

	err = models.InsertCurrency(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single currency
	single, err := models.GetSingleCurrency(req.CurrencyCode, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleCurrencyStruct(single)
	return output, err
}

func UpdateCurrency(conn *connections.Connections, req datastruct.CurrencyRequest) (datastruct.CurrencyDataStruct, error) {
	var output datastruct.CurrencyDataStruct
	var err error

	err = models.UpdateCurrency(conn, req)
	if err != nil {
		return output, err
	}

	// jika tidak ada error, return single instance of single currency
	single, err := models.GetSingleCurrency(req.CurrencyCode, conn, req)
	if err != nil {
		return output, err
	}

	output = CreateSingleCurrencyStruct(single)
	return output, err
}

func DeleteCurrency(conn *connections.Connections, req datastruct.CurrencyRequest) error {
	err := models.DeleteCurrency(conn, req)
	return err
}

func GetListBalance(conn *connections.Connections, req datastruct.BalanceRequest) ([]datastruct.BalanceDataStruct, error) {
	var output []datastruct.BalanceDataStruct
	var err error

	// grab mapping data from model
	balanceList, err := models.GetBalanceFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, balance := range balanceList {
		single := CreateSingleBalanceStruct(balance)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleBalanceStruct(balance map[string]string) datastruct.BalanceDataStruct {
	var single datastruct.BalanceDataStruct
	single.BalanceType, _ = balance["balance_type"]
	single.BalanceName, _ = balance["balance_name"]
	single.Exponent, _ = balance["exponent"]
	single.BalanceCategory, _ = balance["balance_category"]
	single.LastUpdateUsername, _ = balance["last_update_username"]
	single.LastUpdateDate, _ = balance["last_update_date"]

	return single
}
