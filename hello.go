package main

import (
  "fmt" // pkg used to print
  "net/http" // pkg to get pages
  "io/ioutil"
)

func main() {
  resp, err := http.Get("https://google.com")
  // Expect that functions return two things:
  // first - result from fn
  // second - an error, if there is one

  fmt.Println("Error fetching webpage:", err)

  body, err := ioutil.ReadAll(resp.Body)
  // resp.Body is a datastream that we need to convert
  // into something we can parse, using ioutil

  fmt.Println("Error reading body:", err)

  fmt.Println(string(body))
  // Cast body into string, from a bytearray
}
