package processors

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/exchange-rate/datastruct"
	"backendbillingdashboard/modules/exchange-rate/models"
)

func GetListExchangeRate(conn *connections.Connections, req datastruct.ExchangeRateRequest) ([]datastruct.ExchangeRateDataStruct, error) {
	var output []datastruct.ExchangeRateDataStruct
	var err error

	// grab mapping data from model
	dataList, err := models.GetExchangeRateFromRequest(conn, req)
	if err != nil {
		return output, err
	}

	for _, data := range dataList {
		single := CreateSingleExchangeRateStruct(data)
		output = append(output, single)
	}

	return output, err
}

func CreateSingleExchangeRateStruct(data map[string]string) datastruct.ExchangeRateDataStruct {
	var single datastruct.ExchangeRateDataStruct
	single.Date, _ = data["date"]
	single.FromCurrency, _ = data["from_currency"]
	single.ToCurrency, _ = data["to_currency"]
	single.ConvertValue, _ = data["convert_value"]
	single.LastUpdateUsername, _ = data["last_update_username"]
	single.LastUpdateDate, _ = data["last_update_date"]

	return single
}

func InsertExchangeRate(conn *connections.Connections, req datastruct.ExchangeRateRequest) error {
	var err error

	err = models.InsertExchangeRate(conn, req)
	if err != nil {
		return err
	}

	return err
}

func UpdateExchangeRate(conn *connections.Connections, req datastruct.ExchangeRateRequest) error {
	var err error

	err = models.UpdateExchangeRate(conn, req)
	if err != nil {
		return err
	}

	return err
}

func DeleteExchangeRate(conn *connections.Connections, req datastruct.ExchangeRateRequest) error {
	err := models.DeleteExchangeRate(conn, req)
	return err
}
