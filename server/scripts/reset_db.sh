#!/bin/bash

if [[ -z $DATABASE_URL ]]; then
  echo "DATABASE_URL not set"
  exit 1
fi


psql $DATABASE_URL << EOF
  drop schema public cascade;
  create schema public;
EOF