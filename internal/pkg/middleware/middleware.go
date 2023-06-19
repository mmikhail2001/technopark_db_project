package middleware

import (
	"fmt"
	"net/http"
)

type HTTPMiddleware struct {
}

func NewHTTPMiddleware() *HTTPMiddleware {
	return &HTTPMiddleware{}
}

func (m *HTTPMiddleware) AccessLogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// log.Printf("url: [%s], method: [%s], query: [%s]\n", r.URL.Path, r.Method, r.URL.Query())
		fmt.Printf("%s %s\n", r.URL.Path, r.URL.Query())

		// var buf bytes.Buffer
		// _, err := buf.ReadFrom(r.Body)
		// if err != nil {
		// 	http.Error(w, "Error reading request body", http.StatusInternalServerError)
		// 	return
		// }

		// log.Printf("request body: %s\n", buf.String())

		// r.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))

		// recorder := httptest.NewRecorder()

		// h.ServeHTTP(recorder, r)
		h.ServeHTTP(w, r)

		// responseBody, err := ioutil.ReadAll(recorder.Body)
		// if err != nil {
		// 	http.Error(w, "Error reading response body", http.StatusInternalServerError)
		// 	return
		// }

		// log.Printf("response body: %s\n", string(responseBody))

		// // записать ответ в исходный ResponseWriter
		// for k, v := range recorder.Header() {
		// 	w.Header()[k] = v
		// }
		// w.WriteHeader(recorder.Code)
		// w.Write(responseBody)
	})
}
