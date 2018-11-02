package models

import (
    "net/url"
    "strings"
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

type ContinuationToken struct {
    Timestamp time.Time
    Id        uuid.UUID
}

type PageDTO struct {
    TotalCount        int
    HasNext           bool
    ContinuationToken *ContinuationToken
    NextPageUrl       *string
    Payments          []PaymentDto
}

func TokenFromString(tokenStr string) *ContinuationToken {

    parts := strings.Split(tokenStr, "_")

    if len(parts) != 2 {
        return nil
    }

    timestamp, parseTimeError := time.Parse(time.RFC3339, parts[0])
    id, parseIdError := uuid.Parse(parts[1])

    if parseTimeError != nil || parseIdError != nil {
        return nil
    }

    token := ContinuationToken{Timestamp: timestamp, Id: id}

    return &token
}

func TokenFromPaymentDto(dto PaymentDto) *ContinuationToken {

    token := ContinuationToken{Timestamp: dto.DateOccurred, Id: dto.Id}

    return &token
}

func UrlParamFromToken(token *ContinuationToken, requestUrl *url.URL) *string {

    if token == nil {
        return nil
    }

    var nextPageUrl = *requestUrl
    query := nextPageUrl.Query()
    query.Del("continuationToken")
    query.Add("continuationToken", token.Timestamp.Format(time.RFC3339)+"_"+token.Id.String())
    nextPageUrl.RawQuery = query.Encode()

    urlStr := nextPageUrl.String()

    return &urlStr
}
