package routes

import (
    "net/http"
    "os"
    "time"
    "strconv"
    "strings"

    "github.com/go-chi/chi"
    "github.com/gocarina/gocsv"
    "github.com/google/uuid"
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

type PaymentING struct {
    BookingDate DateTime `csv:"Buchung"`
    ValueDate DateTime `csv:"Valuta"`
    Issuer string `csv:"Auftraggeber/Empfänger"`
    PostingText string `csv:"Buchungstext"`
    Ppurpose string `csv:"Verwendungszweck"`
    Amount string `csv:"Betrag"`
    Currency string `csv:"Währung"`
}

func ImportsRoutes() *chi.Mux {
    router := chi.NewRouter()
    router.Post("/", ImportCSV)
    router.Post("/ing", ImportIngCSV)
    return router
}

func ImportCSV(w http.ResponseWriter, r *http.Request) {
    paymentsFile, err := os.OpenFile("D:/workspace/go/src/grid/go-payments/events.csv", os.O_RDWR, os.ModePerm)
    if err != nil {
        panic(err)
    }
    defer paymentsFile.Close()

    var payments []*PaymentCSV

    if err := gocsv.UnmarshalFile(paymentsFile, &payments); err != nil {
        panic(err)
    }

    for _, payment := range payments {
        value, err := strconv.ParseFloat(normalizeEurope(payment.Content), 64)
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

func ImportIngCSV(w http.ResponseWriter, r *http.Request) {
    csvFile, err := os.OpenFile("C:/Umsatzanzeige_DE80500105175418832945_20190323.csv", os.O_RDWR, os.ModePerm)
    if err != nil {
        panic(err)
    }
    defer csvFile.Close()

    var payments []*PaymentING

    if err := gocsv.UnmarshalFile(csvFile, &payments); err != nil {
        panic(err)
    }

    for _, payment := range payments {

        //paymentDto := models.PaymentDto{Id: payment.EventId,
        //    DateOccurred: payment.ValueDate.Time,
        //    Type:         models.PaymentType(payment.BookingType),
        //    Category:     models.PaymentCategory(payment.Category),
        //    SubCategory:  models.PaymentSubCategory(payment.Subcategory),
        //    Value:        value,
        //    Note:         payment.Description}
        //
        //go db.DB.Save(&paymentDto)

        log.Infof("%v" , payment)
    }
}

type DateTime struct {
    time.Time
}

// Convert the CSV string as internal date
func (date *DateTime) UnmarshalCSV(csv string) (err error) {
    date.Time, err = time.Parse("2006-01-02", csv)

    if err != nil {
        date.Time, err = time.Parse("02.01.2006", csv)
        return err
    }

    return err
}

func normalizeEurope(old string) string {
	count := strings.Count(old, ".")
	s := strings.Replace(old, ",", ".", -1)
	return strings.Replace(s, ".", "", count)

}
