# graphql-metrics

graphql Prometheus metrics

# start

    make install
    make dev

# test request

    $ curl -v -d '{"query":"query Users{\n  users{\n    name\n    age\n  }\n}\n\n","variables":null,"operationName":"Users"}' localhost:3000/graphql

    $ curl -v -d '{"query":"\nmutation addUser{\n  AddUser(input:{name:\"conny\", age:24})\n}","variables":null,"operationName":"addUser"}' localhost:3000/graphql

# metrics

    curl -v localhost:3000/metrics

```
...
# TYPE graphql_query_mutation_method_total counter
graphql_query_mutation_method_total{method="AddUser",service="medusa",type="MUTATION"} 1
graphql_query_mutation_method_total{method="__schema",service="medusa",type="QUERY"} 2
graphql_query_mutation_method_total{method="users",service="medusa",type="QUERY"} 1
...
```
