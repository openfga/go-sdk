all: run

project_name=example1
openfga_version=latest

restore:
	go mod tidy

run: restore
	go run ${project_name}/${project_name}.go

run-openfga:
	docker pull docker.io/openfga/openfga:${openfga_version} && \
		docker run -p 8080:8080 docker.io/openfga/openfga:${openfga_version}
