ENV_LOCAL_TEST=\
				APP_NAME="Go API Boilerplate" \
				APP_PORT=3000 \
				APP_PREFORK=false \
				APP_TIMEOUT=10 \
				DB_USER=root \
				DB_PASSWORD= \
				DB_HOST=localhost \
				DB_PORT=3306 \
				DB_NAME=go_api_boilerplate_test \
				POOL_IDLE=5 \
				POOL_MAX=100 \
				POOL_LIFETIME=3000 \
				LOG_LEVEL=6 \
				JWT_SECRET_KEY=secretkey

test.unit:
				APP_ENV=test go test ./test/unit -v -coverprofile=./coverage/unit-coverage.out -coverpkg=github.com/givxl33t/go-fiber-boilerplate/...
test.integration:
				APP_ENV=test $(ENV_LOCAL_TEST) go test ./test/integration -v -coverprofile=./coverage/integration-coverage.out -coverpkg=github.com/givxl33t/go-fiber-boilerplate/...

ifneq (,$(wildcard .env))
    include .env
endif

DATABASE_URL="mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

migrate.create:
				migrate create -ext sql -dir database/migrations $(name)

migrate.up:
				migrate -database $(DATABASE_URL) -path database/migrations up

migrate.down:
				migrate -database $(DATABASE_URL) -path database/migrations down

migrate.force:
				migrate -database $(DATABASE_URL) -path database/migrations force $(version)
