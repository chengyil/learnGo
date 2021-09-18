package handler

import (
	"encoding/json"
	"inventory-service/logic"
	"net/http"
	"net/textproto"
	"strings"
)

func HandleProducts() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case http.MethodGet:
			payload, err := json.Marshal(logic.GetProductsList())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set(textproto.CanonicalMIMEHeaderKey("content-type"), "application/json")
			w.Write(payload)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}
