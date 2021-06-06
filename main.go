package main

import (
    "fmt"
    "net/http"
    "github.com/pkg/errors"
)

type wrapError struct {
    errMsg string
    err error
}


func ErrorInfo(sql string) error {
    return errors.Wrapf(sql,fmt.Sprintf("sql:%s %+v",sql))
}

type errorString string

func (e errorString) Error() string {
    return string(e)
}

func New(text string) error {
    return errorString(text)
}

var ErrorType = New("code.NotFound")


func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}

func main() {
   if errors.Is(ErrorType, New("code.NotFound")) {
        ErrorInfo("code.NotFound")
    } 

//    http.HandleFunc("/", HelloHandler)
 //   http.ListenAndServe(":9000", nil)
}


