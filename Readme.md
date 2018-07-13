# Serverless Golang Template

First edit the Dockerfile to add AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY.
Edit the template in `template/main.go` (HtmlTemplate).

Add the variables you want to insert in the template like this

```
<h1> Hi {{ .Name }}! You are {{ .Age }} years old.
```

Then Edit the struct (in this example, the `Person` Struct), to match (with type).

```
type Person struct {
    Name string
    Age int
}
```

Then run `make deploy` to deploy the lambda.