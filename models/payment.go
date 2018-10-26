package models

import (
    "time"

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
