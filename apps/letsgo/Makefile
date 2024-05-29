# podman-build:
# 	@echo "starting docker build"
# 	podman build -t docker.io/manishm1907/letsgom .
# podman-push:
# 	@echo "pushing to docker hub"
docker-build:
	@echo "starting docker build"
	docker build -t docker.io/manishm1907/letsgom .
docker-push:
	@echo "pushing to docker hub"
	docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}
	docker push docker.io/manishm1907/letsgom
# 	@podman container prune -f