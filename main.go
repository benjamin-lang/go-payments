package main

import (
    "net/http"
    "os"

    "github.com/go-chi/cors"
    log "github.com/sirupsen/logrus"
    "grid/go-payments/db"
    "grid/go-payments/models"
    "grid/go-payments/routes"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/render"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func Routes() *chi.Mux {
    router := chi.NewRouter()

    corsPolicy := cors.New(cors.Options{
        // AllowedOrigins: []string{"http://localhost:4200"}, // Use this to allow specific origin hosts
        AllowedOrigins: []string{"*"},
        // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    })

    router.Use(
        render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
        middleware.Logger,                             // Log API request calls
        middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
        middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
        middleware.Recoverer,                          // Recover from panics without crashing server
        corsPolicy.Handler,
    )

    router.Route("/v1", func(r chi.Router) {
        r.Mount("/api/payments", routes.PaymentsRoutes())
        r.Mount("/api/imports", routes.ImportsRoutes())
    })

    return router
}

func init() {
    log.SetFormatter(&log.TextFormatter{})
    log.SetOutput(os.Stdout)
    log.SetLevel(log.InfoLevel)
}

func main() {
    router := Routes()

    if err1 := db.Open(); err1 != nil {
        log.Panic("failed to connect database")
    }

    defer db.Close()

    db.DB.AutoMigrate(&models.PaymentDto{})

    walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
        log.Infof("%s %s\n", method, route) // Walk and print out all routes
        return nil
    }

    if err := chi.Walk(router, walkFunc); err != nil {
        log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
    }

    log.Fatal(http.ListenAndServe(":8080", router)) // Note, the port is usually gotten from the environment.
}
