swagger:
	swag init -o internal/docs -d internal --parseDependency --parseInternal

check-swagger:
	swag init -o tmp -d internal --parseDependency --parseInternal
	rm -rf tmp

lint:
	golangci-lint run -v

mockery:
	find . -type f -name 'mock_*' -exec rm {} +
	find . -type f -name "mockWrappers.go" -execdir mv -n -- {} mockWrappers.txt \; # Костылина, чтобы не падала генерация моков
	mockery
	find . -type f -name "mockWrappers.txt" -execdir mv -n -- {} mockWrappers.go \;

test-coverage-number: mockery
	go test -v -coverprofile=profile.cov ./app...
	go tool cover -func profile.cov
	rm profile.cov

test-coverage-html: mockery
	go test -v -coverprofile=profile.cov ./app...
	go tool cover -html profile.cov
	rm profile.cov

test: mockery
	go test ./...

deploy-check: test lint check-swagger

migration-pg-create-file: # Создать файл миграции
	goose -dir migrations/pgsql create rename_me sql

migration-pg-status: # Показать статус миграций
	goose -dir migrations/pgsql postgres "postgres://postgres:password@127.0.0.1:5432/generator" status

migration-pg-reset: # Сбросить все миграции
	goose -dir migrations/pgsql postgres "postgres://postgres:password@127.0.0.1:5432/generator" reset
