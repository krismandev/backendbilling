package transport

import (
	"billingdashboard/core"
	er "billingdashboard/error"
	dt "billingdashboard/modules/currency/datastruct"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CurrencyDecodeRequest is use for ...
func CurrencyDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request dt.CurrencyRequest

	//decode request body
	body, err := ioutil.ReadAll(r.Body)
	core.LogBodyRequest(body, "Received Request")

	// only decode json if the body length > 0
	if len(body) > 0 {
		if err = json.Unmarshal(body, &request); err != nil {
			return er.Error(err, core.ErrInvalidFormat).Rem("Failed decoding json message"), nil
		}
	}
	return request, nil
}

// CurrencyListEncodeResponse is use for ...
func CurrencyListEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	body, err := json.Marshal(&response)
	core.LogBodyRequest(body, "Send Response")

	if err != nil {
		return err
	}

	var e = response.(core.GlobalListResponse).ResponseCode
	w = core.WriteHTTPResponse(w, e)
	_, err = w.Write(body)

	return err
}

// CurrencySingleEncodeResponse is use for ...
func CurrencySingleEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	body, err := json.Marshal(&response)
	core.LogBodyRequest(body, "Send Response")

	if err != nil {
		return err
	}

	var e = response.(core.GlobalSingleResponse).ResponseCode
	w = core.WriteHTTPResponse(w, e)
	_, err = w.Write(body)

	return err
}
