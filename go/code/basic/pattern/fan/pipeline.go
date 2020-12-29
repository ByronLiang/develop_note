package fan

import (
    "fmt"
    "time"
)

var text string

func CreatThread(chantings chan string, done chan struct{}, word string) chan string {
    next := make(chan string)
    go func() {
        var c = make([]string, 0)
        for {
            select {
            case data := <-chantings:
                if data == "" {
                    next <- word
                } else {
                    next <- fmt.Sprintf("%s%s", data, word)
                }
                c = append(c, word)
            case <-done:
                fmt.Println("thread: ", word, c)
                return
            }
        }
    }()
    return next
}

func New()  {
    done := make(chan struct{})
    start := make(chan string)
    a := CreatThread(start, done, "A")
    b := CreatThread(a, done, "B")
    c := CreatThread(b, done, "C")
    start <- ""
    for i := 0; i < 2; i++ {
        text = <- c
        start <- text
    }
    text = <- c
    close(done)
    time.Sleep(1 * time.Second)
    fmt.Println(text)
}

func beginner(done chan struct{}, word string) chan string {
    c := make(chan string)
    go func() {
        defer close(c)
        for {
            select {
            case c<-word:
            case <-done:
                return
           }
        }
    }()
    return c
}

func midPack(upChan chan string, done chan struct{}, word string) chan string {
    c := make(chan string)
    go func() {
        defer close(c)
        for {
            select {
            case data := <-upChan:
                c <- fmt.Sprintf("%s%s", data, word)
            case <-done:
                return
            }
        }
    }()
    return c
}

func NewPip()  {
    done := make(chan struct{})
    a := beginner(done, "A")
    b := midPack(a, done, "B")
    c := midPack(b, done, "C")
    r := 5
    text := ""
    for data := range c {
        text = fmt.Sprintf("%s%s", text, data)
        r--
        if r == 0 {
            close(done)
            break
        }
    }
    fmt.Println(text)
}
