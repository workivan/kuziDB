package store

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// TableServiceContentParser отражает структуру с указанием дефолтного пути до файлов со служебной информацией о таблицах
type TableServiceContentParser struct {
	serviceContentPath string
}

// Parse метод парсит файлы по дефолтному пути в мапу
func (parser TableServiceContentParser) Parse() (map[string]*TableServiceContent, error) {
	m := make(map[string]*TableServiceContent)

	bytesFileContent, err := os.ReadFile(parser.serviceContentPath)

	if err != nil {
		err = fmt.Errorf("can't open file: get error %w", err)
		return m, err
	}

	fileContent := string(bytesFileContent)

	if len(fileContent) == 0 {
		err = fmt.Errorf("file cant be empty")
		return m, err
	}

	_, contents, _ := strings.Cut(fileContent, ":")

	for _, tablesContent := range strings.Split(contents, ";") {
		if len(tablesContent) > 0 {
			tableContent := strings.Split(tablesContent, ",")

			if len(tableContent) != 4 {
				err = fmt.Errorf("bad service file")
				return m, err
			}

			//парсим юзер колонки в список структур
			colums := tableContent[2]
			columnSlice := []ColumnInfo{}
			for _, column := range strings.Split(colums, "/") {
				f := strings.Split(column, " ") //это сплит по инфе о конкретной колонке (я уже не знаю как нормально именовать переменные)

				if len(f) != 3 {
					err = fmt.Errorf("bad userColumn info")
					return m, err
				}

				columnSlice = append(columnSlice, ColumnInfo{f[0], f[1], f[2]})
			}

			tableName := tableContent[0]
			m[tableName] = &TableServiceContent{tableName, time.Now(), columnSlice, tableContent[3]}
		}
	}

	return m, err
}

// Что привело к безопасности?
// Почему не взламываешь?
// где учился, как учился, что круто, что нет.
// Какие задачи решаются
// Что самое не интересное в работе.
// Ваня сказал что ты какой то доклад где то рассказывал.
