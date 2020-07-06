package HttpBuilder

import (
    "fmt"
    "net/http"
)

type BasicHttp struct {}

var context *HttpContext

func BuildHttpService() {
    var basic BasicHttp
    //http.NewServeMux()
    //http.HandleFunc()
    _ = http.ListenAndServe(":8080", &basic)
}

func (b *BasicHttp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    r := req.URL.Path
    fmt.Println(r)
    context = &HttpContext{req, w}
    _, _ = context.response.Write([]byte("qq"))
}
