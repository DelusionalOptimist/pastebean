include .env

# run it as process in your host system 🤮
# make pastebean-dirty
dev-db:
	docker run --name=postgres -p=5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_DB=pastebeandb -d postgres:latest || true;

build-all:
	cd create; CGO_ENABLED=0 GOOS=linux go build -o create main.go; cd ..;
	cd read; CGO_ENABLED=0 GOOS=linux go build -o read main.go; cd ..;
	cd update; CGO_ENABLED=0 GOOS=linux go build -o update main.go; cd ..;
	cd delete; CGO_ENABLED=0 GOOS=linux go build -o delete main.go; cd ..;

run-create:
	DB_DRIVER=$(DB_DRIVER) DB_SOURCE=$(DEV_DB_SOURCE) ./create/create

run-read:
	DB_DRIVER=$(DB_DRIVER) DB_SOURCE=$(DEV_DB_SOURCE) ./read/read

run-update:
	DB_DRIVER=$(DB_DRIVER) DB_SOURCE=$(DEV_DB_SOURCE) ./update/update

run-delete:
	DB_DRIVER=$(DB_DRIVER) DB_SOURCE=$(DEV_DB_SOURCE) ./delete/delete

pastebean-dirty: dev-db build-all run-create run-read run-update run-delete

# --------------------------------------------------------------------------------------------------------------------------------------------

# Uses docker compose to setup containerized dev environments
# make pastebean-dc-dev

pastebean-dc-dev: build-all
	docker-compose -f ./dev-docker-compose.yaml up --build

# --------------------------------------------------------------------------------------------------------------------------------------------

# prodding the bean
# the only difference is that the binary is built inside the container
docker-build-prod:
	DOCKER_BUILDKIT=1 docker build --no-cache --file create/Dockerfile -t pastebean-create:latest . ;
	DOCKER_BUILDKIT=1 docker build --no-cache --file read/Dockerfile -t pastebean-read:latest . ;
	DOCKER_BUILDKIT=1 docker build --no-cache --file update/Dockerfile -t pastebean-update:latest . ;
	DOCKER_BUILDKIT=1 docker build --no-cache --file update/Dockerfile -t pastebean-delete:latest . ;

pastebean:
	docker-compose -f ./docker-compose.yaml up
