package processors

import (
	"billingdashboard/connections"
	"billingdashboard/modules/currency/datastruct"
	"billingdashboard/modules/currency/models"
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
