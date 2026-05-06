#!/usr/bin/env bash
set -euo pipefail

go test ./internal/cases -coverprofile=coverage-cases.out
coverage=$(go tool cover -func=coverage-cases.out | awk '/total:/ {print substr($3, 1, length($3)-1)}')

required=90

awk -v coverage="$coverage" -v required="$required" 'BEGIN {
  if (coverage < required) {
    printf("coverage %.2f is below required %.2f\n", coverage, required)
    exit 1
  }
  printf("coverage %.2f meets required %.2f\n", coverage, required)
}'
