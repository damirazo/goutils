package main
// ========================================================================================
// Простейший веб фреймворк
//
// @author damirazo<me@damirazo.ru>
//
//
// Пример использования:
//
// func main() {
// 		routes := []Route{
// 			Route{Url: "/", Handler: requestHello},
// 			Route{Url: "/test/", Handler: requestIp},
// 		}
// 
// 		s := Server(routes)
// 		s.Listen()
// }
// 
// 
// func requestHello(writer http.ResponseWriter, _ *http.Request) {
// 		fmt.Fprintf(writer, "Hello, World!")
// }
// 
// func requestIp(writer http.ResponseWriter, request *http.Request) {
// 		fmt.Fprintf(writer, "Я тебя вычислю!11 Твой адрес: " + request.RemoteAddr)
// }
//
// ========================================================================================

import (
	"net/http"
)


type server struct {
	Routes []Route
}

func Server(routes []Route) *server {
	s := &server{}
	s.Routes = routes

	for _, route := range s.Routes {
		http.HandleFunc(route.Url, route.Handler)
	}

	return s
}

func (self *server) Listen() {
	http.ListenAndServe(":8080", nil)
}


type Route struct {
	Url string
	Handler func(http.ResponseWriter, *http.Request)
}
