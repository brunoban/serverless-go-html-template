FROM golang:1.11beta1-stretch
RUN mkdir -p /go/src/template \
    && apt-get update \
    && curl -sL https://deb.nodesource.com/setup_10.x | bash - \
    && apt-get install -y nodejs \
    && npm install -g serverless \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ENV AWS_ACCESS_KEY_ID=SOMETHINGSOMETHING
ENV AWS_SECRET_ACCESS_KEY=SOMETHINGSOMETHING
COPY ./ /go/src/template/
WORKDIR /go/src/template
CMD make build && serverless deploy