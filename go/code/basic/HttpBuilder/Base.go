package HttpBuilder

import (
    "fmt"
    "net/http"
)

type BasicHttp struct {}

var httpContext *HttpContext

func BuildHttpService() {
    var basic BasicHttp
    //http.NewServeMux()
    //http.HandleFunc()
    _ = http.ListenAndServe(":8080", &basic)
}

func (b *BasicHttp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    r := req.URL.Path
    fmt.Println(r)
    httpContext = &HttpContext{req, w}
    _, _ = httpContext.response.Write([]byte("qq"))
}
