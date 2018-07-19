# Developing EKS-SSO

EKS-SSO is a web application that consists of a backend, written in Go and a frontend, written in React.


## Setup

You will need the following installed and configured:

1. Go `v1.10.3` development environment 
2. Node `v10.6.0` development environment
3. Kubernetes cli `kubectl`
4. Aws cli authentication plugin `heptio-authenticator-aws`


## Running the service locally

1. Add aws IAM credentials in ~/.aws/credentials
1. Add configuration to connect to a existing eks cluster via `kubectl` in ~/.kube/config
2. Ensure kubectl's current context point to the eks cluster
3. Run the backend

```console
cd ./src
go run main.go  --kube-config-path="/Users/<username>/.kube/config" --cluster-name="seshi" --auth0-client-id="<client-id>" --auth0-connection="<ldap-connection-name>"
```

4. Run the frontend

```console
cd ./src/ui
yarn start
```

By default this will start your server portion on port `8080` and serve the UI on port `3000`. In development, any relative requests from the UI are routed to port `8080` via configuration set in `package.json`. In production, the backed will serve the frontend and expose rest api required from the frontend on the same server port.


## Production build

A production build consists of packaging the application as a docker image. To build the docker image, execute from project root

```console
make build
```

This will build a docker image. If you wanted to push the built image you can run

```console
make release
```

which will push the image to docker hub with the latest and commit tags.