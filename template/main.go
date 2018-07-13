package main

import (
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "html/template"
    "bytes"
    "encoding/json"
    "strings"
)

// Edit this file with your html template.
// Add the variables you want to insert like this: {{ .Variable }}
var HtmlTemplate = `
<head>
<style>
body {
    background-color: linen;
}

h1 {
    color: maroon;
    margin-left: 40px;
} 
</style>
</head>
<html>
  <body>
    <h1>Hello {{ .Name }}! You are {{ .Age }} years old!</h1>
    <h2>BANANA</h2>
  </body>
</html>
`

// Example of struct you can create to store the variables
// you are going to add to the template.
type Person struct {
    Name string
    Age int
}

func BuildPage(htmlTemplate string, person Person) *bytes.Buffer {
    var bodyBuffer bytes.Buffer
    t := template.New("template")
    var templates = template.Must(t.Parse(htmlTemplate))
    templates.Execute(&bodyBuffer, person)
    return &bodyBuffer
}

// Useful Link regarding request payloads
// https://github.com/aws/aws-lambda-go/blob/master/events/apigw.go
func ParseRequest(request events.APIGatewayProxyRequest) Person {
    var p Person
    dec := json.NewDecoder(strings.NewReader(request.Body))
    dec.Decode(&p)
    return p
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    person := ParseRequest(request)
    return events.APIGatewayProxyResponse{
    Headers:    map[string]string{"content-type": "text/html"},
    Body:       BuildPage(HtmlTemplate, person).String(),
    StatusCode: 200,
    }, nil
}

func main() {
    lambda.Start(Handler)
}