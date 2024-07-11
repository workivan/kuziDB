package store

// TableServiceContentParser отражает структуру с указанием дефолтного пути до файлов со служебной информацией о таблицах
type TableServiceContentParser struct {
	DefaultFolder string
}

// Parse метод парсит файлы cо служебной информацией по дефолтному пути в мапу
func (parser TableServiceContentParser) Parse() map[string]*TableServiceContent {
	m := make(map[string]*TableServiceContent)
	m["test"] = &TableServiceContent{"name"}
	return m
}
