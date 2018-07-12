# build executable
FROM golang:alpine AS build-env
ADD ./pkg /go/src/github.com/quintilesims/eks-sso/pkg/
ADD ./src/main.go /go/src/github.com/quintilesims/eks-sso/src/main.go
ADD ./vendor /go/src/github.com/quintilesims/eks-sso/vendor/
RUN cd /go/src/github.com/quintilesims/eks-sso/src && go build -o eks-sso

# package executable and static assets 
FROM alpine
RUN apk add --update ca-certificates
WORKDIR /app
COPY ./src/ui/build/ /app/ui/build/
COPY --from=build-env /go/src/github.com/quintilesims/eks-sso/src/eks-sso /app/
ENTRYPOINT ./eks-sso