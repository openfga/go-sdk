module example1

go 1.21.9

require github.com/openfga/go-sdk v0.3.5

require golang.org/x/sync v0.7.0 // indirect

// To reference local build, uncomment below and run `go mod tidy`
//replace github.com/openfga/go-sdk v0.3.5 => ../../
