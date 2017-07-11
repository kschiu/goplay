package main

import (
  "fmt" // pkg used to print
  "net/http" // pkg to get pages
  "net/smtp" //pkg to send email
  "io/ioutil"
  "strings" // pkg help with string manipulation
)


const (
  SMTP_SERVER = "smtp.gmail.com"
)

type Sender struct {
  username  string
  password  string
}


func main() {
  sender := Sender{"EMAIL", "PASSWORD"}

  recipients := []string{"EMAIL"}

  body := getVolunteerStatus()

  sender.sendMail(recipients, body)
}

// Helper methods


// Case Insensitive Contains
func csContains(a, b string) bool {
  return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}

func getVolunteerStatus() string {
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

  // Cast body into string, from a bytearray
  // Then check if page says "stay tuned still"
  if csContains(string(body), "Stay Tuned For More Info") {
    return "Volunteering not open yet"
  } else {
    // Do something to notify me to sign up
    return "Go Sign Up For Volunteering"
  }
}


func (sender Sender) sendMail(Recipient []string, body string) {
  msg := getVolunteerStatus()

  // 587 is the port
  err := smtp.SendMail(
    SMTP_SERVER + ":587",
    smtp.PlainAuth("", sender.username, sender.password, SMTP_SERVER),
    sender.username,
    Recipient,
    []byte(msg))

  // Cast msg into a bytearray

  if err != nil {
    fmt.Println("Err sending:", err)
  } else {
    fmt.Println("Sent Mail")
  }
}