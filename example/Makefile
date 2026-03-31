all: run

project_name=example1
openfga_version=latest

restore:
	go mod tidy

run: restore
	go run ${project_name}/${project_name}.go

run-streamed-list-objects:
	cd streamed_list_objects && go run .

run-api-executor-streaming:
	cd api_executor_streaming && go run .

run-openfga:
	docker pull docker.io/openfga/openfga:${openfga_version} && \
		docker run -p 8080:8080 docker.io/openfga/openfga:${openfga_version} run
