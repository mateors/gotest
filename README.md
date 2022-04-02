# golang test

```
1. repository layer:
should call https://gitlab.com/api/graphql
with query (where $n is param)
query last_projects($n: Int = 5) {
  projects(last:$n) {
    nodes {
      name
      description
      forksCount
    }
  }
}

Note: manually you can test query via 
https://gitlab.com/-/graphql-explorer

2a. service layer
Design a return schema and endpoint. Could be GraphQL, or plain REST API with JSON.
2b. service layer return
return join names with a ", " delimiter and the sum of all forks

3. write tests for service layer

4. print results via call that service from main.go
For the service repository layer, if you are not familiar with GraphQL, mock it with JSON based on the response you get from the explorer.

(optional) 5. handling of the environment variables (e.g. Token for GraphQL endpoint, etc)
(optional) 6. good logging 

```
## Step-1
run the server first
> cd /server

> go build

> ./service

> go test -v

## Step-2
run the client
> cd /client

> go build

> ./client
