package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	me "github.com/mrbotchi-team/mrbotchi/error"
)

type (
	// HTTPHandlerIF はHTTPのハンドラを定義するヤツ。
	HTTPHandlerIF interface {
		Get(w http.ResponseWriter, r *http.Request) (StatusCode int, Response interface{}, Error error)
		Post(w http.ResponseWriter, r *http.Request) (StatusCode int, Response interface{}, Error error)
		Put(w http.ResponseWriter, r *http.Request) (StatusCode int, Response interface{}, Error error)
		Delete(w http.ResponseWriter, r *http.Request) (StatusCode int, Response interface{}, Error error)
	}

	// HTTPHandler はHTTPHanderIFを実装したヤツ。
	HTTPHandler struct {
	}

	// HTTPHandlerFunc は関数。(そりゃそうだ!)
	HTTPHandlerFunc func(http.ResponseWriter, *http.Request) (StatusCode int, Response interface{}, Error error)
)

var methodNotAllowedError = &me.APIError{ID: "METHOD_NOT_ALLOWED", Message: "This method isnt allowed."}

// Get はGETリクエストを受ける関数。
func (h HTTPHandler) Get(w http.ResponseWriter, r *http.Request) (StatusCode int, Response interface{}, Error error) {
	return http.StatusMethodNotAllowed, nil, methodNotAllowedError
}

// Post はPOSTリクエストを受ける関数。
func (h HTTPHandler) Post(w http.ResponseWriter, r *http.Request) (StatusCode int, Response interface{}, Error error) {
	return http.StatusMethodNotAllowed, nil, methodNotAllowedError
}

// Put はPUTリクエストを受ける関数。
func (h HTTPHandler) Put(w http.ResponseWriter, r *http.Request) (StatusCode int, Response interface{}, Error error) {
	return http.StatusMethodNotAllowed, nil, methodNotAllowedError
}

// Delete はDELETEリクエストを受ける関数。
func (h HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) (StatusCode int, Response interface{}, Error error) {
	return http.StatusMethodNotAllowed, nil, methodNotAllowedError
}

func (hf HTTPHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 適当な関数を発火させる。
	status, res, err := hf(w, r)

	// エラー処理。エラーメッセージを返せそうなヤツはJSON形式で返して、
	// そうで無いものは適当な文字列を返す。いずれにせよステータスコードを返す。
	if nil != err {
		if err, ok := err.(*me.APIError); ok {
			log.Printf("error: %s: %s", err.ID, err.Message)

			errID := err.ID
			body, err := json.Marshal(err)
			if nil != err {
				// もうめんどうみきれよう
				// なんで自分が定義した構造体をエンコードできないんですか(呆れ)
				log.Println("error: ", err)

				writeResponse(w, http.StatusInternalServerError, "text/plain", []byte(fmt.Sprintf("The encoding of the error message failed! WTF?????\n Anyway, here's the error ID: %s", errID)))
				return
			}

			writeResponse(w, status, "application/json", body)
		} else {
			log.Println("error: ", err)

			writeResponse(w, status, "text/plain", []byte(http.StatusText(status)))
		}
		return
	}

	// ステータスコードを返す。
	// 204ならそのまま抜ける。
	w.WriteHeader(status)
	if nil == res && http.StatusNoContent == status {
		return
	}

	// くぅ～疲れましたw これにて終了です!
	// 実は、実装したらエラー処理の話を持ちかけられたのが始まりでした
	// 本当は面倒臭かったのですが←
	// 本番環境でクラッシュにするわけには行かないので流行りの処理(?)で挑んでみた所存ですw
	// 以下、シュビムワーゲン達のみんなへのメッセジをどぞ
	//
	// シュビムワーゲン「えー、Y.Tなら知っていますが、U.Tというのはわたくし初めて聞きました。」
	//
	// シュビムワーゲン「はい、えー超地球的存在です。」
	//
	// シュビムワーゲン「ええ。スゴいです。」
	//
	// では、
	//
	// ワシ「う~、じゃBGMのこの曲、お聴きになりますか?」
	// シュビムワーゲン達「まっさかぁ!」
	if body, ok := res.([]byte); ok {
		w.Write(body)
	} else {
		writeResponse(w, http.StatusInternalServerError, "text/plain", []byte(http.StatusText(http.StatusInternalServerError)))
	}
}

func writeResponse(w http.ResponseWriter, status int, contentType string, body []byte) {
	w.Header().Set("Content-Type", contentType+"; charset=utf-8")
	w.WriteHeader(status)

	w.Write(body)
}
