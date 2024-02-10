#!/bin/bash
echo "creating author"
curl -X POST -d '{"author": { "firstName": "Aidan", "lastName": "King", "country": "IE"} }'  http://localhost:8080/api/authors
echo ""

echo "reading author"
curl http://localhost:8080/api/authors/0
echo ""

echo "updating author"
curl -X PUT -d '{"author": {"firstName": "Aidan", "lastName": "King", "country": "IE2"} }'  http://localhost:8080/api/authors/0
echo ""

echo "delete author"
curl -X DELETE   http://localhost:8080/api/authors/0
echo ""

echo "listling all authors"
curl http://localhost:8080/api/authors
echo ""
