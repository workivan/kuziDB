package store

import (
	"time"
)

type TableServiceContent struct {
	nameOfTable  string
	timeOfCreate time.Time
	colums       []ColumnInfo // Как хранить инфу по типу колонны(инт не инт) я не придумал да и пофек пока
	pathToData   string
}

type ColumnInfo struct {
	name       string
	columnType string
	nullable   string
}
