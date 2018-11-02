package routes

import (
    "net/http"
    "os"
    "time"

    "github.com/go-chi/chi"
    "github.com/gocarina/gocsv"
    "github.com/google/uuid"
    "github.com/marcsantiago/StringToFloat"
    log "github.com/sirupsen/logrus"
    "grid/go-payments/db"
    "grid/go-payments/models"
)

type PaymentCSV struct {
    EventId      uuid.UUID `csv:"EventId"`
    DateOccurred DateTime  `csv:"DateOccured"`
    BookingType  string    `csv:"BookingType"`
    Category     string    `csv:"Category"`
    Subcategory  string    `csv:"Subcategory"`
    Content      string    `csv:"Content"`
    Description  string    `csv:"Description"`
}

func ImportsRoutes() *chi.Mux {
    router := chi.NewRouter()
    router.Post("/", ImportCSV)
    return router
}

func ImportCSV(w http.ResponseWriter, r *http.Request) {
    paymentsFile, err := os.OpenFile("F:/workspace/go/src/grid/go-payments/events.csv", os.O_RDWR, os.ModePerm)
    if err != nil {
        panic(err)
    }
    defer paymentsFile.Close()

    var payments []*PaymentCSV

    if err := gocsv.UnmarshalFile(paymentsFile, &payments); err != nil {
        panic(err)
    }

    for _, payment := range payments {
        value, err := stringtofloat.Convert(payment.Content)
        if err != nil {
            panic(err)
        }

        paymentDto := models.PaymentDto{Id: payment.EventId,
            DateOccurred: payment.DateOccurred.Time,
            Type:         models.PaymentType(payment.BookingType),
            Category:     models.PaymentCategory(payment.Category),
            SubCategory:  models.PaymentSubCategory(payment.Subcategory),
            Value:        value,
            Note:         payment.Description}

        go db.DB.Save(&paymentDto)
    }

    log.Infof("%v payments imported from csv", len(payments))
}

type DateTime struct {
    time.Time
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
    date.Time, err = time.Parse("2006-01-02", csv)
    return err
}
