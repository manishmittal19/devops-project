# podman-build:
# 	@echo "starting docker build"
# 	podman build -t docker.io/vishnu930417/devopsfactory .
# podman-push:
# 	@echo "pushing to docker hub"
docker-build:
	@echo "starting docker build"
	docker build -t docker.io/vishnu930417/devopsfactory .
docker-push:
	@echo "pushing to docker hub"
	docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}
	docker push docker.io/vishnu930417/devopsfactory
# podman-prune:
# 	@podman container prune -f