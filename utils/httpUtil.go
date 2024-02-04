package httpUtil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type RequestBodyType string

const (
	FORM_DATA RequestBodyType = "FORM_DATA"
	RAW_JSON  RequestBodyType = "RAW_JSON"
)

type Request struct {
	Url         string
	Headers     *map[string]string
	QueryParams interface{}
	ReqBody     interface{}
	RespBody    interface{}
	ReqBodyType RequestBodyType
	Ctx         context.Context
}

func (r *Request) Get() (int, error) {
	if r.QueryParams != nil {
		r.queryStringUrl()
	}

	jsonValue, err := json.Marshal(r.ReqBody)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	jsonBytes := bytes.NewBuffer(jsonValue)

	var req *http.Request
	var reqErr error
	if r.Ctx != nil {
		req, reqErr = http.NewRequestWithContext(r.Ctx, "GET", r.Url, jsonBytes)
	} else {
		req, reqErr = http.NewRequest("GET", r.Url, jsonBytes)
	}
	if reqErr != nil {
		return http.StatusInternalServerError, reqErr
	}
	r.addHeaders(r.Headers, req)
	resp, err := r.sendRequest(req, r.RespBody)
	if err != nil {
		return http.StatusConflict, err
	}

	return resp.StatusCode, nil
}

func (r *Request) Post() (int, error) {
	if r.QueryParams != nil {
		r.queryStringUrl()
	}
	var req *http.Request
	var reqErr error
	if r.Ctx != nil {
		req, reqErr = http.NewRequestWithContext(r.Ctx, "POST", r.Url, r.getPayload())
	} else {
		req, reqErr = http.NewRequest("POST", r.Url, r.getPayload())
	}
	if reqErr != nil {
		return 0, reqErr
	}
	r.addHeaders(r.Headers, req)
	resp, err := r.sendRequest(req, r.RespBody)
	if err != nil {
		return http.StatusConflict, err
	}
	return resp.StatusCode, nil
}

func (r *Request) Put() (int, error) {
	if r.QueryParams != nil {
		r.queryStringUrl()
	}

	var req *http.Request
	var reqErr error
	if r.Ctx != nil {
		req, reqErr = http.NewRequestWithContext(r.Ctx, "PUT", r.Url, r.getPayload())
	} else {
		req, reqErr = http.NewRequest("PUT", r.Url, r.getPayload())
	}
	if reqErr != nil {
		return 0, reqErr
	}
	r.addHeaders(r.Headers, req)
	resp, err := r.sendRequest(req, r.RespBody)
	if err != nil {
		return http.StatusConflict, err
	}
	return resp.StatusCode, nil
}

func (r *Request) Delete() (int, error) {
	if r.QueryParams != nil {
		r.queryStringUrl()
	}

	var req *http.Request
	var reqErr error
	if r.Ctx != nil {
		req, reqErr = http.NewRequestWithContext(r.Ctx, "DELETE", r.Url, nil)
	} else {
		req, reqErr = http.NewRequest("DELETE", r.Url, nil)
	}
	if reqErr != nil {
		return 0, reqErr
	}
	r.addHeaders(r.Headers, req)
	resp, err := r.sendRequest(req, r.RespBody)
	if err != nil {
		return http.StatusConflict, err
	}
	return resp.StatusCode, nil
}

func (r *Request) Patch() (int, error) {
	if r.QueryParams != nil {
		r.queryStringUrl()
	}

	var req *http.Request
	var reqErr error
	if r.Ctx != nil {
		req, reqErr = http.NewRequestWithContext(r.Ctx, "PATCH", r.Url, r.getPayload())
	} else {
		req, reqErr = http.NewRequest("PATCH", r.Url, r.getPayload())
	}
	if reqErr != nil {
		return 0, reqErr
	}
	r.addHeaders(r.Headers, req)
	resp, err := r.sendRequest(req, r.RespBody)
	if err != nil {
		return http.StatusConflict, err
	}
	return resp.StatusCode, nil
}

func (r *Request) addHeaders(headers *map[string]string, request *http.Request) {
	if headers != nil {
		for key, value := range *headers {
			request.Header.Set(key, value)
		}
	}
}

func (r *Request) mapBytesToStruct(body io.ReadCloser, structBody interface{}) error {
	if structBody == nil {
		return nil
	}

	byteBody, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	// Check if result is string
	if val, ok := structBody.(*string); ok {
		*val = string(byteBody)
		return nil
	}

	if err := json.Unmarshal(byteBody, &structBody); err != nil {
		return err
	}

	return nil
}

func (r *Request) sendRequest(request *http.Request, response interface{}) (*http.Response, error) {
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	resp, respErr := client.Do(request)
	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode == http.StatusNoContent ||
		resp.StatusCode == http.StatusResetContent {
		return resp, nil
	}

	defer resp.Body.Close()
	if response != nil {
		if err := r.mapBytesToStruct(resp.Body, response); err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func structToUrlValues(obj interface{}) url.Values {
	v := reflect.ValueOf(obj)
	urlValues := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("form")
		if key == "" {
			continue
		}
		switch v.Field(i).Kind() {
		case reflect.Ptr:
			if !v.Field(i).IsNil() {
				value := fmt.Sprintf("%v", v.Field(i).Elem().Interface())
				urlValues.Add(key, value)
			}
		case reflect.Slice:
			value := fmt.Sprintf("%v", v.Field(i))
			if value == "[]" {
				continue
			}
			value = value[1 : len(value)-1]
			values := strings.Split(value, " ")
			for index, stringValue := range values {
				urlValues.Add(fmt.Sprintf("%s[%d]", key, index), stringValue)
			}
		default:
			value := fmt.Sprintf("%v", v.Field(i))
			urlValues.Add(key, value)
		}
	}

	return urlValues
}

func (r *Request) queryStringUrl() {
	queryParams := structToUrlValues(r.QueryParams)
	u, _ := url.ParseRequestURI(r.Url)
	u.RawQuery = queryParams.Encode()
	r.Url = u.String()
}

func (r *Request) getPayload() *bytes.Buffer {
	if r.ReqBodyType == FORM_DATA {
		formData := structToUrlValues(r.ReqBody)
		return bytes.NewBufferString(formData.Encode())
	} else {
		jsonValue, _ := json.Marshal(r.ReqBody)
		return bytes.NewBuffer(jsonValue)
	}
}
