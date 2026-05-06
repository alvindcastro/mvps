#!/usr/bin/env bash
set -euo pipefail

go test ./internal/cases ./internal/triage ./internal/reviewerqueue ./internal/platform/httpserver -coverprofile=coverage.out
coverage=$(go tool cover -func=coverage.out | awk '/total:/ {print substr($3, 1, length($3)-1)}')

required=85

awk -v coverage="$coverage" -v required="$required" 'BEGIN {
  if (coverage < required) {
    printf("coverage %.2f is below required %.2f\n", coverage, required)
    exit 1
  }
  printf("coverage %.2f meets required %.2f\n", coverage, required)
}'
