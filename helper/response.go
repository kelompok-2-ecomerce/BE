package helper

import (
	"net/http"
	"strings"
)

func PrintSuccessReponse(message string, data ...interface{}) interface{} {
	resp := map[string]interface{}{}
	if message != "" {
		resp["message"] = message
	}

	if len(data) == 0 {
		return resp
	} else if len(data) < 2 {
		if data[0] != "" {
			resp["data"] = data[0]
		}
	} else {
		if data[0] != "" {
			resp["data"] = data[0]
		}
		resp["token"] = data[1].(string)
	}

	return resp
}

func PrintErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := -1
	if msg != "" {
		resp["message"] = msg
	}

	if strings.Contains(msg, "server") || strings.Contains(msg, "Error") {
		resp["message"] = "data tidak bisa diolah"
		code = http.StatusInternalServerError
	} else if strings.Contains(msg, "format") {
		code = http.StatusBadRequest
	} else if strings.Contains(msg, "tidak ditemukan") {
		code = http.StatusNotFound
	} else if strings.Contains(msg, "password") {
		code = http.StatusUnauthorized
	} else if strings.Contains(msg, "sudah terdaftar") {
		code = http.StatusConflict
	} else if strings.Contains(msg, "belum terdaftar") {
		code = http.StatusNotFound
	} else if strings.Contains(msg, "required") {
		code = http.StatusBadRequest
	} else {
		resp["message"] = "data tidak bisa diolah"
		code = http.StatusInternalServerError
	}

	return code, resp
}
