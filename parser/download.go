package parser
// ========================================================================================
// Функции для скачивания файла
//
// @author damirazo<me@damirazo.ru>
// ========================================================================================

import (
	"os"
	"io"
	"path"
	"strings"
	"net/http"
)

// Скачивание файла
// @param string url          Ссылка на файл
// @param string destination  Путь для сохранения скачанного файла
// @return nil or error
func DownloadFile(url string, destination string) error {
	segments := strings.Split(url, "/")
	fileName := segments[len(segments)-1]

	filePath := path.Join(destination, fileName)

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	output, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	return nil
}
