package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceFileParser(t *testing.T) {
	t.Run("normFile", func(t *testing.T) {
		assertions := assert.New(t)
		parserNorm := TableServiceContentParser{"/home/crowdyara/go/src/my-project/kuziDB/internal/.serviceData/norm_file.txt"}

		contentMap, _ := parserNorm.Parse()
		testContent := contentMap["golang"]

		// 1) норм файл, 2)файла нет, 3)файл с говной, 4)пустой файл
		assertions.NotNil(testContent, "TestContent must not be null")
	})

	t.Run("noFile", func(t *testing.T) {
		parserNoFile := TableServiceContentParser{" "}
		_, err := parserNoFile.Parse()

		if err == nil {
			t.Errorf("func must return err")
		}

	})

	t.Run("badfile", func(t *testing.T) {
		parserBadFile := TableServiceContentParser{"/home/crowdyara/go/src/my-project/kuziDB/internal/.serviceData/govno_file.txt"}
		_, err := parserBadFile.Parse()

		if err == nil {
			t.Errorf("func must return err")
		}

	})

	t.Run("badfile", func(t *testing.T) {
		parserBadFile := TableServiceContentParser{"/home/crowdyara/go/src/my-project/kuziDB/internal/.serviceData/empty_file.txt"}
		_, err := parserBadFile.Parse()

		if err == nil {
			t.Errorf("func must return err")
		}

	})

}
