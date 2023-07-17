package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func randomString(n int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    s := make([]rune, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}

func fuzzDomain(domain string) {
    for {
        path := randomString(rand.Intn(10) + 1)
        url := fmt.Sprintf("%s/%s", domain, path)

        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        fmt.Printf("Status code for %s: %d\n", url, resp.StatusCode)
        resp.Body.Close()
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    domain := "https://longflourishrc.com"
    fuzzDomain(domain)
}
