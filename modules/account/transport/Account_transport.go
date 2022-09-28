package transport

import (
	"backendbillingdashboard/core"
	er "backendbillingdashboard/error"
	dt "backendbillingdashboard/modules/account/datastruct"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// AccountDecodeRequest is use for ...
func AccountDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request dt.AccountRequest

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

// AccountListEncodeResponse is use for ...
func AccountListEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

// AccountSingleEncodeResponse is use for ...
func AccountSingleEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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

// AccountDecodeRequest is use for ...
func RootParentAccountDecodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request dt.RootParentAccountRequest

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

// AccountListEncodeResponse is use for ...
func RootParentAccountListEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
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
