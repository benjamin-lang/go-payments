package db

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Open() error {
    var err error

    DB, err = gorm.Open("postgres",
        "host=springfield "+
            "port=5432 "+
            "user=a6IwWTd1sXWL6Wv2eXl4nIn "+
            "password=2wgQu7CGXMvVlGvcgPO1p72vnZPSkMNaA9nnCNNTLvFhLo "+
            "dbname=payments "+
            "sslmode=disable")

    return err
}

func Close() error {
    return DB.Close()
}
