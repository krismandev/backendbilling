package core

import (
	"backendbillingdashboard/config"
	conf "backendbillingdashboard/config"
	lib "backendbillingdashboard/lib"
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	log "github.com/sirupsen/logrus"
)

// GlobalEncodeResponse is used for
func GlobalEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	log.WithFields(GetLogFieldValues(ctx, "GlobalEncodeResponse"))
	var body []byte

	body, err := json.Marshal(&response)
	log.Infof("Send Response %s", body[:])
	if err != nil {
		return err
	}

	var e = response.(GlobalJSONResponse).ResponseCode
	w = WriteHTTPResponse(w, e)
	_, err = w.Write(body)

	return err
}

// HandleJWT will return error when UseJWT is active and bearer token is invalid
func HandleJWT(ctx context.Context) (int, error) {
	if conf.Param.UseJWT {
		tokenAuth, ok := ctx.Value(httptransport.ContextKeyRequestAuthorization).(string)
		if !strings.Contains(tokenAuth, "Bearer ") {
			return ErrNotAuthorized, errors.New(DescNotAuthorized)
		}
		tokenAuthKey := strings.Replace(tokenAuth, "Bearer ", "", 1)
		tokenAuthKey = strings.Trim(tokenAuthKey, " ")
		remoteIP, ok := ctx.Value(httptransport.ContextKeyRequestRemoteAddr).(string)
		remoteIP, _, _ = net.SplitHostPort(remoteIP)
		forwardedFor := ctx.Value(httptransport.ContextKeyRequestXForwardedFor).(string)
		log.Infof("["+remoteIP+"] ReceivedRequest : %s ; ForwardedFor : "+forwardedFor, tokenAuthKey)
		if !ok {
			log.Errorf("["+remoteIP+"] Error : %s", "Invalid/Not Found/Expired Auth")
			return ErrNotAuthorized, errors.New(DescNotAuthorized)
		}
		claim, tokenValid := lib.ValidateToken(tokenAuthKey, remoteIP, config.Param.AppName, conf.Param.JWTSecretKey)

		if !tokenValid {
			log.Errorf("["+remoteIP+"] Error : %s", "Invalid Token")
			return ErrNotAuthorized, errors.New(DescNotAuthorized)
		}
		if claim.ExpiresAt < time.Now().Unix() {
			log.Errorf("["+remoteIP+"] Error : %s", "Token Expired")
			return ErrNotAuthorized, errors.New(DescNotAuthorized)
		}
	}

	return 0, nil
}

// WriteHTTPResponse is used for
func WriteHTTPResponse(w http.ResponseWriter, e int) http.ResponseWriter {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Connection", "close")

	if e < 8000 {
		w.WriteHeader(http.StatusOK)
	} else if e < 9000 {
		w.WriteHeader(http.StatusUnauthorized)
	} else if e <= 9999 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	return w
}

// GetRequestInformation is use for ...
func GetRequestInformation(ctx context.Context, r *http.Request) context.Context {
	ctx = context.WithValue(ctx, ContextTransactionID, lib.GetTransactionid(true))
	return ctx
}

// LogBodyRequest will log the body request/response, and chunk the log if the request is too large
func LogBodyRequest(body []byte, msg string) {
	if len(msg) == 0 {
		msg = "Received Request"
	}
	msg = strings.ToUpper(msg)

	if len(body) > config.Param.MaxBodyLogLength {
		// chunk body log request
		log.Infof("["+msg+"] (Chunked %d first chars) %s", config.Param.MaxBodyLogLength, body[:config.Param.MaxBodyLogLength])
	} else {
		log.Infof("["+msg+"] %s", body[:])
	}
}
