package server

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {

	r := Routes()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	addr := flag.String("addr", ":9901", "Server address")

	srv := &http.Server{
		Handler:      logging(infoLog)(r),
		Addr:         *addr,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
		ErrorLog:     errorLog,
	}

	infoLog.Printf("Server started on %s\n", srv.Addr)
	errorLog.Fatal(srv.ListenAndServe())

}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			next.ServeHTTP(w, r)
		})
	}
}
