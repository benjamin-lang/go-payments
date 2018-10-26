package payments

import (
    "encoding/json"
    "net/http"
    "time"

    "test/db"

    "github.com/go-chi/chi"
    "github.com/google/uuid"
)

type PaymentType string
type PaymentCategory string
type PaymentSubCategory string

type PaymentDto struct {
    Id           uuid.UUID          `json:"id"`
    DateOccurred time.Time          `json:"dateOccurred"`
    Type         PaymentType        `json:"type"`
    Category     PaymentCategory    `json:"category"`
    SubCategory  PaymentSubCategory `json:"subcategory"`
    Value        float64            `json:"value,string"`
    Note         string             `json:"note"`
}

func Routes() *chi.Mux {
    router := chi.NewRouter()
    router.Post("/", CreatePayment)
    router.Get("/", GetPayments)
    router.Delete("/", DeleteAll)
    router.Get("/{paymentId}", GetAPayment)
    router.Put("/{paymentId}", CreatePaymentWithId)
    return router
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
    var paymentDto PaymentDto
    json.NewDecoder(r.Body).Decode(&paymentDto)
    paymentDto.Id = uuid.New()
    db.DB.Save(&paymentDto)

    respondCreatedWithLocationHeader(w, r.URL.String(), paymentDto.Id.String())
}

func CreatePaymentWithId(w http.ResponseWriter, r *http.Request) {
    var payment PaymentDto
    json.NewDecoder(r.Body).Decode(&payment)
    payment.Id = uuid.MustParse(chi.URLParam(r, "paymentId"))
    db.DB.Save(&payment)

    respondCreatedWithLocationHeader(w, r.URL.String(), payment.Id.String())
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
    var payments []PaymentDto
    db.DB.Find(&payments)
    respondWithJSON(w, http.StatusOK, &payments)
}

func GetAPayment(w http.ResponseWriter, r *http.Request) {
    var payment PaymentDto
    payment.Id = uuid.MustParse(chi.URLParam(r, "paymentId"))
    db.DB.Find(&payment)
    respondWithJSON(w, http.StatusOK, &payment)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
    var payments []PaymentDto
    db.DB.Delete(payments)
}
