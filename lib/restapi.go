package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type RestClient struct {
	URL     string
	Method  string
	Timeout int
	Headers map[string]string
	Options map[string]string
	Request interface{}
}

func (r *RestClient) SetURL(url string) *RestClient {
	r.URL = url
	return r
}

func (r *RestClient) SetMethod(method string) *RestClient {
	allowedMethods := []string{"get", "post", "delete", "patch", "put", "options", "head"}
	if _, hasSlice := FindSlice(allowedMethods, strings.ToLower(method)); !hasSlice {
		method = "get"
	}
	r.Method = method
	return r
}

func (r *RestClient) SetTimeout(timeout int) *RestClient {
	if timeout <= 1 {
		timeout = 5 //static default timeout.
	}
	r.Timeout = timeout
	return r
}

func (r *RestClient) SetHeaders(headers map[string]string) *RestClient {
	for hname, hval := range headers {
		r = r.AddHeader(hname, hval)
	}
	return r
}

func (r *RestClient) AddHeader(headerName, headerValue string) *RestClient {
	if len(r.Headers) == 0 {
		r.Headers = make(map[string]string)
	}
	r.Headers[headerName] = headerValue
	return r
}

func (r *RestClient) SetOptions(options map[string]string) *RestClient {
	for oname, oval := range options {
		r = r.AddOption(oname, oval)
	}
	return r
}

func (r *RestClient) AddOption(optionName, optionValue string) *RestClient {
	if len(r.Options) == 0 {
		r.Options = make(map[string]string)
	}
	r.Options[optionName] = optionValue
	return r
}

func (r *RestClient) SetRequest(param interface{}) *RestClient {
	r.Request = param
	return r
}

func (r *RestClient) Execute() (httpBody string, httpStatus int) {
	jsonValue, _ := json.Marshal(r.Request)

	if len(r.Method) == 0 {
		r.Method = "GET"
	}
	if r.Timeout == 0 {
		r.Timeout = 5
	}

	// create request structure
	req, err := http.NewRequest(r.Method, r.URL, strings.NewReader(string(jsonValue)))
	if err != nil {
		return
	}

	for hname, hval := range r.Headers {
		req.Header.Set(hname, hval)
	}
	req.Close = true // this is required to prevent too many files open

	// Create HTTP Connection
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(r.Timeout) * time.Second,
	}

	// Now hit to destionation endpoint
	res, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Call URL Failed : %s", err.Error())
		if res != nil {
			buff := new(bytes.Buffer)
			buff.ReadFrom(res.Body)
			httpBody = buff.String()
			httpStatus = res.StatusCode
			log.Errorf("Body : %s", httpBody)
		} else {
			httpBody = "Call URL Failed : " + err.Error()
		}
		return
	}
	defer res.Body.Close()

	buff := new(bytes.Buffer)
	buff.ReadFrom(res.Body)
	httpBody = buff.String()
	httpStatus = res.StatusCode
	return
}
