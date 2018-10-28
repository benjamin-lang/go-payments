package main

import (
    "grid/go-payments/db"
    "grid/go-payments/models"
    "grid/go-payments/routes"
    "log"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/render"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

//https://itnext.io/structuring-a-production-grade-rest-api-in-golang-c0229b3feedc
//https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
//https://flaviocopes.com/golang-tutorial-rest-api/
//https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f
//https://outcrawl.com/go-microservices-cqrs-docker/
//http://gorm.io/docs/sql_builder.html#content-inner
//https://dev.to/aspittel/how-i-built-an-api-with-mux-go-postgresql-and-gorm-5ah8
//https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/
func Routes() *chi.Mux {
    router := chi.NewRouter()
    router.Use(
        render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
        middleware.Logger,                             // Log API request calls
        middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
        middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
        middleware.Recoverer,                          // Recover from panics without crashing server
    )

    router.Route("/v1", func(r chi.Router) {
        r.Mount("/api/payments", routes.PaymentsRoutes())
        r.Mount("/api/imports", routes.ImportsRoutes())
    })

    return router
}

func main() {
    router := Routes()

    if err1 := db.Open(); err1 != nil {
        panic("failed to connect database")
    }

    defer db.Close()

    db.DB.AutoMigrate(&models.PaymentDto{})

    walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
        log.Printf("%s %s\n", method, route) // Walk and print out all routes
        return nil
    }

    if err := chi.Walk(router, walkFunc); err != nil {
        log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
    }

    log.Fatal(http.ListenAndServe(":8080", router)) // Note, the port is usually gotten from the environment.
}
