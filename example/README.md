## Examples of using the OpenFGA Go SDK

A set of examples on how to call the OpenFGA Go SDK

### Examples
Example 1:
A bare bones example. It creates a store, and runs a set of calls against it including creating a model, writing tuples and checking for access.


### Running the Examples

Prerequisites:
- `docker`
- `make`
- `go` 1.21+

#### Run using a published SDK

Steps
1. Clone/Copy the example folder
2. If you have an OpenFGA server running, you can use it, otherwise run `make run-openfga` to spin up an instance (you'll need to switch to a different terminal after - don't forget to close it when done)
3. Run `make run` to run the example

#### Run using a local unpublished SDK build

Steps
1. Build the SDK
2. In the Example `go.mod`, uncomment out the part that replaces the remote SDK with the local one, e.g.
```
// To refrence local build, uncomment below and run `go mod tidy`
replace github.com/openfga/go-sdk v0.3.2 => ../../
```
3. If you have an OpenFGA server running, you can use it, otherwise run `make run-openfga` to spin up an instance (you'll need to switch to a different terminal after - don't forget to close it when done)
4. Run `make run` to run the example