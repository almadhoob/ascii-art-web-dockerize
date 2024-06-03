docker image build --file Dockerfile --tag docker-ascii . \
docker container run --publish 8080:8080 -d docker-ascii