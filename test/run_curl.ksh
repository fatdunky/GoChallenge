#!/bin/ksh

#curl -H "Content-Type: application/json" -X POST -d @../config/request.json localhost:8080

curl -H "Content-Type: application/json" -X POST -d @../config/request.json https://5-dot-stan-challenge.appspot.com/:8080
