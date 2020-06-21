package HttpBuilder

import "net/http"

//上下文对象
type HttpContext struct {
    request     *http.Request
    response    http.ResponseWriter
}
