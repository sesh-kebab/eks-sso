# build backend
FROM golang:alpine AS build-env-server
ADD ./pkg /go/src/github.com/quintilesims/eks-sso/pkg/
ADD ./src/main.go /go/src/github.com/quintilesims/eks-sso/src/main.go
ADD ./vendor /go/src/github.com/quintilesims/eks-sso/vendor/
RUN cd /go/src/github.com/quintilesims/eks-sso/src && go build -o eks-sso

# build frontend
FROM node:8.11.3-alpine AS build-env-ui
ADD ./src/ui/ /usr/src/app/
WORKDIR /usr/src/app
RUN npm install
RUN npm run build

# package executable and static assets 
FROM alpine
RUN apk add --update ca-certificates
WORKDIR /app
COPY --from=build-env-ui /usr/src/app/build/ /app/ui/build/
COPY --from=build-env-server /go/src/github.com/quintilesims/eks-sso/src/eks-sso /app/
ENTRYPOINT ./eks-sso