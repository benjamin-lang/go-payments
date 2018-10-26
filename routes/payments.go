package routes

import (
    "encoding/json"
    "github.com/go-chi/chi"
    "github.com/google/uuid"
    "grid/go-payments/db"
    "grid/go-payments/middleware"
    "grid/go-payments/models"
    "net/http"
)

func PaymentsRoutes() *chi.Mux {
    router := chi.NewRouter()
    router.Post("/", CreatePayment)
    router.Get("/", GetPayments)
    router.Delete("/", DeleteAll)
    router.Get("/{paymentId}", GetAPayment)
    router.Put("/{paymentId}", CreatePaymentWithId)
    return router
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
    var paymentDto models.PaymentDto
    json.NewDecoder(r.Body).Decode(&paymentDto)
    paymentDto.Id = uuid.New()
    db.DB.Save(&paymentDto)

    middleware.RespondCreatedWithLocationHeader(w, r.URL.String(), paymentDto.Id.String())
}

func CreatePaymentWithId(w http.ResponseWriter, r *http.Request) {
    var payment models.PaymentDto
    json.NewDecoder(r.Body).Decode(&payment)
    payment.Id = uuid.MustParse(chi.URLParam(r, "paymentId"))
    db.DB.Save(&payment)

    middleware.RespondCreatedWithLocationHeader(w, r.URL.String(), payment.Id.String())
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
    var payments []models.PaymentDto
    db.DB.Find(&payments)
    middleware.RespondWithJSON(w, http.StatusOK, &payments)
}

func GetAPayment(w http.ResponseWriter, r *http.Request) {
    var payment models.PaymentDto
    payment.Id = uuid.MustParse(chi.URLParam(r, "paymentId"))
    db.DB.Find(&payment)
    middleware.RespondWithJSON(w, http.StatusOK, &payment)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
    var payments []models.PaymentDto
    db.DB.Delete(payments)
}
