# Serverless Golang Template

This image is a HTML template that takes in arguments to build an HTML document from a template that is provided in the code.

Edit the template in `template/main.go` (`HtmlTemplate`).

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

To make deployment a one time command that doesn't pollute your environment or anything, it's contained inside a Dockerfile, so edit the Dockerfile to add `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.

Then run `make deploy` to deploy the lambda.

The makefile is just sugar coating the commands a little bit to make it nicer. If that doesn't work for you, just run:

```
docker build . -t <image-name>
docker run --rm <image-name>
```

The --rm flag is just there to cleanup the container after completion, if you run into trouble, you can remove it.

The output of the docker run will have your URL, so after that make a POST request to the URL similar to the example:

```
curl -X POST <URL> -d '{"name":"Robert", "age": 18}'
```