package hello

import "net/http"

// Handler заглушка для тестов
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello!"))
}
