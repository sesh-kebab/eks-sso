# Spec

## Abstract
This application will allow users to manage certain aspects of their Kubernetes cluster using a web service. 
The web service will be protected by an SSO provider (Auth0) and have its own notion of users and permissions. 
Currently, the web service will only be able to manage namespaces for a Kubernetes cluster. 
The scope of this functionality may change in the future. 

## Authentication
This section details how we will secure the web service. 

* Users must have an Auth0 account setup to use our service.
* Users will need to setup an Application in their Auth0 account for the web service to use (we will have AMAZING documentation on how to do this). 
I believe users will need to configure the OAuth2 request to have the `openid` scope. 
The `openid` scope should contain a username or email we can use as the `userID` in our application (see: https://auth0.com/docs/protocols/oauth2).
* The service's frontend will use the Auth0 React Lock (https://www.npmjs.com/package/auth0-react-lock) to perform the handshake with Auth0. 
* Every route in the frontend will be protected by the Auth0 React Lock.
* Users will need to configure the service's Kubernetes yaml files to connect the web service with their Auth0 Application. They will also need to specify which connections to display in the Auth0 React Lock. 
This may look something like:
```
auth0:
  domain: "foo"
  client_id: "bar"
  connections:
    - github
    - facebook
    - ...
```
* When the Auth0 handshake is completed by the frontend, the frontend will pass the data returned by Auth0 (I believe this is the `hashHandler` function part of the Auth0 React Lock) to the backend using the `POST /user` endpoint. 
* During the `POST /user` endpoint, the backend will first verify that the data sent by the frontend is valid (I believe this can be done by sending a request to Auth0).
With the `openid` scope, the backend should be able to extract a `userID` from the data. 
Once the `userID` is extracted, the backend will do 1 of 2 things: If the `userID` does not exists in the database, the backend will create a new user for the `userID` and generate a token for that `userID`. If the `userID` already exists, the backend will lookup that user's token from our database. Finally, the backend will send a response back to the frontend containing the `userID` and `token`. From then on, each request sent by the frontend to the backend must be authenticated with the correct `userID` and `token`. 


## Permissions Model
The permissions model will use a simple RBAC model, which has the following concepts:
* **User**: A user is created after their first succesful Auth0 handshake.
A user can be assigned 0-many roles.
* **Role**: A role contains 0-many permissions. 
By default, the web service will create an `Administrator` and `Guest` role. 
Each time a new user is created, they are granted the `Guest` role. 
Each time a new namespace is created, a corresponding role with the same name and custom permissions is also created. 
* **Permission**: A permission allows a role to perform an action. 
If there are no permissions which allow said action, the action is denied. 


**Permissions Table**

| Permission/Role   | Guest     | <Namespace Role>     | Administrator | 
|-------------------|-----------|----------------------|---------------|
| View Dashboard    | Allow     | -                    | Allow         |
| View Cluster      | -         | -                    | Allow         |
| View Namespace    | -         | Self                 | Allow         |
| Create Namespace  | -         | -                    | Allow         |
| Delete Namespace  | -         | -                    | Allow         |
| View User         | Self      | -                    | Allow         |
| Grant User Role   | -         | -                    | Allow         |
| Revoke User Role  | Self      | -                    | Allow         |
| Delete User       | Self      | -                    | Allow         |


## First Administrator
There must be a user with the Administrator role when the service first deployed.
After a user performs a successful Auth0 handshake, the service will determine if 
this is the only user in the database, and if so, will grant said user the Administrator role.
This is a potential security concern if someone were to delete every user. 
We will also need to come up with safeguards against deleting users to the point where there are 0 users with the Administrator role. 

## API
High-level overview of the API methods provided by the backend. 

**Cluster**

| Method | Path     | Params | Description                                    |
|--------|----------|--------|------------------------------------------------|
| GET    | /cluster |        | Returns detailed information about the cluster |


**Users**

| Method | Path                        | Params      | Description                               |
|--------|-----------------------------|-------------|-------------------------------------------|
| GET    | /users                      | role:string | Returns a list of users                   |
| POST   | /users                      |             | Register a new user     |
| GET    | /users/:userID              |             | Returns detailed information about a user |
| DELETE    | /users/:userID           |             | Delete a user |
| POST   | /users/:userID/role         |             | Grants a role to a user                   |
| DELETE | /users/:userID/role/:roleID |             | Revokes a role from a user                |


**Namespace**

| Method | Path                                   | Params | Description                                            |
|--------|----------------------------------------|--------|--------------------------------------------------------|
| GET    | /namespace                             |        | Returns a list of namespaces                           |
| POST   | /namespace                             |        | Create a new namespace                                 |
| GET    | /namespace/:namespaceID                |        | Returns detailed information about a namespace |
| DELETE | /namespace/:namespaceID                |        | Delete a namespace                         |
| GET    | /namespace/:namespaceID/kubeconfig.yml |        | Returns a raw kubeconfig.yml for a namespace   |


