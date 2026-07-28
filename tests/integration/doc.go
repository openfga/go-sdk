// Package integration provides integration tests for the OpenFGA Go SDK.
//
// These tests use testcontainers-go to spin up a real OpenFGA server with
// preshared-key authentication enabled. They verify that authentication
// headers are correctly sent across all request paths:
//
//   - Standard API calls (Check, ListObjects, etc.)
//   - StreamedListObjects (streaming client method)
//   - APIExecutor.Execute (low-level executor)
//   - APIExecutor.ExecuteStreaming (low-level streaming executor)
//
// # Running
//
// This directory is a separate Go module, so run the tests from within it
// (or use the Makefile target from the repo root):
//
//	make test-integration
//
//	cd tests/integration && go test -v -tags integration -timeout 120s ./...
//
// # Prerequisites
//
//   - Docker must be running (testcontainers-go manages containers automatically)

//go:build integration

package integration
