# Sample Golang with GitHub Actions

[![Go](https://github.com/djaque/sampleGolangGHA/actions/workflows/ci.yml/badge.svg)](https://github.com/djaque/sampleGolangGHA/actions/workflows/ci.yml)

This is an example on golang to test the github actions

The app is very easy, contains a single handler listening for a single route.

Also add a single test to check the process.

Then add the workflow steps:

- Build
- Tests
- Build Docker image
- Push image to dockerhub

You can check this workflow on
https://github.com/djaque/sampleGolangGHA/blob/main/.github/workflows/ci.yml

For comments use: djaque@gmail.com