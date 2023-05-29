# Variables for Docker building
DOCKER_BUILD_CMD = docker build -f Dockerfile --build-arg service=

# Default rule (builds the docker images for all services)
.PHONY: docker-build
docker-build: auth_service-docker project_service-docker gateway-docker

# Rules for building individual services
auth_service-docker:
	$(DOCKER_BUILD_CMD)auth_service -t auth_service .

project_service-docker:
	$(DOCKER_BUILD_CMD)project_service -t project_service .

gateway-docker:
	$(DOCKER_BUILD_CMD)gateway -t gateway .

