package http
// ========================================================================================
// Структура для загрузки интернет страниц
//
// @author damirazo<me@damirazo.ru>
// ========================================================================================

import (
	"errors"
	"io/ioutil"
	netHttp "net/http"
)

var LoadPageError = errors.New("Произошла ошибка при загрузке страницы!")
var ReadPageError = errors.New("Произошла ошибка при чтении страницы!")

// Структура для загрузки и хранения информации об интернет странице
type page struct {
	Url string
	Response *netHttp.Response
	Body []byte
}

// Псевдо-конструктор
func Page(url string) *page {
	p := &page{Url: url}
	p.init()
	return p
}

// Первичная инициализация структуры
func (self *page) init() error {
	response, err := netHttp.Get(self.Url)
	if err != nil {
		return LoadPageError
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ReadPageError
	}

	self.Response = response
	self.Body = body

	return nil
}

// Получение тела страницы в строковом формате
func (self *page) Content() string {
	return string(self.Body)
}

// Код ответа от сервера
func (self *page) StatusCode() int {
	return self.Response.StatusCode
}

// Закрытие открытых дескрипторов
func (self *page) Close() {
	self.Response.Body.Close()
}
