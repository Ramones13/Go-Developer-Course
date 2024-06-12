package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	f, err := os.Create("./lesson17/Task17-3/log.txt")
	if err != nil {

		fmt.Println(err)
	}
	defer f.Close()
	defer f.WriteString("Программа завершена")
	log.SetOutput(f)

	f.WriteString(fmt.Sprintf("Программа запущена\n"))

	//l := log.New(os.Stdout, "", 0)
	l := log.New(f, "", 0)
	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: authHandler(logHandler(mux)),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			msg123 := "пользователь не авторизован"
			log.Println(msg123, http.StatusUnauthorized)
			http.Error(w, msg123, http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//log.Println("url:", r.URL)
			l.Println("url:", r.URL)

			h.ServeHTTP(w, r)
		})
	}
}
