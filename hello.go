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

  // NO TERNARY OP IN GO??? :(
  if err != nil {
    fmt.Println("Error fetching webpage:", err)
  }

  body, err := ioutil.ReadAll(resp.Body)
  // resp.Body is a datastream that we need to convert
  // into something we can parse, using ioutil

  if err != nil {
    fmt.Println("Error reading body:", err)
  }

  // TODO: Make this send me an email or text

  // Cast body into string, from a bytearray
  // Then check if page says "stay tuned still"
  if csContains(string(body), "Stay Tuned For More Info") {
    fmt.Println("Volunteering not open yet")
  } else {
    // Do something to notify me to sign up
    fmt.Println("Go Sign Up For Volunteering")
  }
}

// Helper method


// Case Insensitive Contains
func csContains(a, b string) bool {
  return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}