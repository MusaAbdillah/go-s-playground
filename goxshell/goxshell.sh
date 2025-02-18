#!/bin/sh
string="John"

if [[ -z "$string" ]]; then
  echo "String is empty"
elif [[ -n "$string" ]]; then
  echo "String is not empty, $string"
  #-p, --parents no error if existing, make parent directories as needed
  mkdir -p "$string"
fi

