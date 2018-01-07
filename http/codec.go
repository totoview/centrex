// Code generated DO NOT EDIT.

package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/totoview/centrex/pb"
)

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("EncodeHttpError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func DecodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req pb.Login
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
