package main

import (
  "fmt" // pkg used to print
  "net/http" // pkg to get pages
  "io/ioutil"
  "strings" // pkg help with string manipulation
)

func main() {
  resp, err := http.Get("https://www.sfoutsidelands.com/get-involved/volunteers/")
  // Expect that functions return two things:
  // first - result from fn
  // second - an error, if there is one

  fmt.Println("Error fetching webpage:", err)

  body, err := ioutil.ReadAll(resp.Body)
  // resp.Body is a datastream that we need to convert
  // into something we can parse, using ioutil

  fmt.Println("Error reading body:", err)

  fmt.Println(csContains(string(body), "Stay Tuned For More Info"))
  // Cast body into string, from a bytearray
  // Then check if page says "stay tuned still"
}

// Case Insensitive Contains
func csContains(a, b string) bool {
  return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}