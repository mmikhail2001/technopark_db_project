package httptools

import (
	"context"
	"log"
	"net/http"

	"github.com/mailru/easyjson"
	"github.com/mmikhail2001/technopark_db_project/internal/pkg"
)

func getEasyJSON(v any) ([]byte, error) {
	easyJsonStruct, ok := v.(easyjson.Marshaler)
	if !ok {
		return []byte{}, pkg.ErrGetEasyJSON
	}
	bytes, err := easyjson.Marshal(easyJsonStruct)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func Response(ctx context.Context, w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := getEasyJSON(someStruct)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ResponseError(ctx context.Context, w http.ResponseWriter, statusCode int, err error) {
	errorStruct := ErrResponseDTO{
		Message: err.Error(),
	}

	Response(ctx, w, statusCode, errorStruct)
}
