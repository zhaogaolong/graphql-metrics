#!/bin/sh

echo prefix: $prefix

if [ -z $prefix ]; then
  prefix=vendor
fi

graphql_path=$prefix/github.com/graph-gophers/graphql-go/graphql.go

if ! grep -q MySchema $graphql_path; then
  cat patch/graphql.txt >> $graphql_path
fi
