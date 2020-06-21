package HttpBuilder

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
)

type BasicHttp struct {}

var context *HttpContext

func ByteBuff() {
    var readInput *bufio.Reader
    readInput = bufio.NewReader(os.Stdin)
    fmt.Println("Input your name ?")
    name, _ := readInput.ReadString('\n')
    fmt.Printf("your name is %s", name)
}

func BuildHttpService()  {
    var basic BasicHttp
    // http.NewServeMux()
    // http.HandleFunc()
    _ = http.ListenAndServe(":8080", &basic)
}

func (b *BasicHttp) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    r := req.URL.Path
    fmt.Println(r)
    _, _ = context.response.Write([]byte("qq"))
}
