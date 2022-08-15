#!/bin/bash

for i in {1..500}; do
curl --header "Content-Type: application/json" --request GET  localhost:8080/api/v1/monkey/ban_1gt4ti4gnzjre341pqakzme8z94atcyuuawoso8gqwdx5m4a77wu1mxxighh?format=png
echo " -- $i"
done