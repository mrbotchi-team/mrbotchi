package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteResponse はいい感じにレスポンスを返す関数。
func WriteResponse(w http.ResponseWriter, status int, contentType string, body []byte) {
	w.Header().Set("Content-Type", contentType+"; charset=utf-8")
	w.WriteHeader(status)

	w.Write(body)
}

// WriteJSONResponse はいい感じにJSONのレスポンスを返す関数。
func WriteJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	body, err := json.Marshal(payload)
	if nil != err {
		// もうめんどうみきれよう
		// なんで自分が定義した構造体をエンコードできないんですか(呆れ)
		log.Println("error: ", err)

		WriteResponse(w, http.StatusInternalServerError, "text/plain", []byte("The encoding of the error message failed! WTF?????\n Anyway, here's the error ID:"))
		return
	}
	WriteResponse(w, status, "application/json", body)
}
