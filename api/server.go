package api

import (
    "fmt"
    "net/http"
)

func StartServer(port int) error {

    return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
