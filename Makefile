include .env
default: run

build:
	go build -C ./cmd/main-service -o main-service
	du -sh ./cmd/main-service

start-prod:
	./cmd/main-service/main-service

run:
	go run -C ./cmd/main-service main.go

dev:
	make swag
	make run

install:
	go install

update:
	go get -u ./...

clean:
	go clean -C ./cmd/main-service

swag:
	cd cmd/main-service && swag init -g ../../internal/modules/controllers/server.go -o ../../docs -d ./ --pd --parseInternal

swag-fmt:
	swag fmt

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

lint-clean:
	golangci-lint cache clean

tidy:
	go mod tidy

install-devtools:
	sh install-dev-dependencies.sh

diff-migrations:
	cd ./internal/database && go run gen.go ${name}

hash:
	cd ./internal/database && atlas migrate hash ./internal/database/migrations

up-migrations:
	cd ./internal/database/migrations && goose postgres "user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable host=${DB_HOST} port=${DB_PORT}"  up

down-migration:
	cd ./internal/database/migrations && goose postgres "user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable host=${DB_HOST} port=${DB_PORT}"  down

reset-migrations:
	cd ./internal/database/migrations && goose postgres "user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable host=${DB_HOST} port=${DB_PORT}" reset

create-migration:
	cd ./internal/database/migrations && goose postgres "user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable host=${DB_HOST} port=${DB_PORT}" create ${name} sql
	hash

fieldalignment:
	fieldalignment -fix ./...

#docker-build:
#	docker build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} \
#                  --build-arg GITHUB_USERNAME=${GITHUB_USERNAME} \
#                  --build-arg GOPRIVATE=github.com/go-boilerplate/* \
#                  -t backend-workspace-main-service:latest .

docker-build:
	docker build  --driver=docker-container \
				  --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} \
                  --build-arg GITHUB_USERNAME=${GITHUB_USERNAME} \
                  --build-arg GOPRIVATE=github.com/go-boilerplate/* \
                  --cache-to type=local,dest=/tmp/buildx-cache-new,mode=max \
                  --cache-from type=local,src=/tmp/buildx-cache \
                  -t backend-workspace-main-service:latest .

gen-ent:
	go generate ./internal/modules/ent/generate.go