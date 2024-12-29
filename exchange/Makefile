# Используем для генерации доки openapi, чтобы локально посмотреть что там есть
# Не забываем, что эти файлы надо удалить, ибо документация генерируется при деплое
swagger:
	swag init -o docs -d cmd,./ --parseInternal --parseDependency

# Используем для проверки ошибок при генерации доки openapi. Команда не оставляет после себя артефактов
check-swagger:
	swag init -o tmp -g cmd/main.go --parseInternal --parseDependency
	rm -rf tmp

# Используем, чтобы прогнать код линтером (следим, чтобы локальный golangci был всегда последней версии)
lint:
	golangci-lint run -v

# Используем, чтобы обновить моки для тестов
mockery:
	find . -type f -name 'mock_*' -exec rm {} +
	find . -type f -name "mockWrappers.go" -execdir mv -n -- {} mockWrappers.txt \; # Костылина, чтобы не падала генерация моков
	mockery
	find . -type f -name "mockWrappers.txt" -execdir mv -n -- {} mockWrappers.go \;

# Используем для получения тест-кавераджа для каждого файла и проекта в целом
test-coverage-number: mockery
	go test -tags=prod -v -coverprofile=profile.cov ./cmd...
	go tool cover -func profile.cov
	rm profile.cov

# Используем для получения графического тест-кавераджа. Просматривая код в браузере мы можем видеть сколько тесткейсов
# Проходятся по одному и тому же участку кода
test-coverage-html: mockery
	go test -tags=prod -v -coverprofile=profile.cov ./cmd...
	go tool cover -html profile.cov
	rm profile.cov

# Используем для прогона тестов текущего сервиса
test: mockery
	go test -tags=prod ./...

# Перед деплоем дергаем эту команду, чтобы проверить, что код готов к сливанию с другой веткой
# Если этого не делать, завалится пайплайн и все равно придется править :)
deploy-check: test lint check-swagger

# Миграции для postgresql

migration-pg-create-file: # Создать файл миграции
	goose -dir migrations/pgsql create rename_me sql

migration-pg-status: # Показать статус миграций
	goose -dir migrations/pgsql postgres "postgres://user:password@127.0.0.1:5433/db" status

migration-pg-reset: # Сбросить все миграции
	goose -dir migrations/pgsql postgres "postgres://user:password@127.0.0.1:5433/db" reset

# Миграции для clickhouse

migration-cl-create-file: # Создать файл миграции
	goose -dir migrations/clickhouse create rename_me sql

migration-cl-status: # Показать статус миграций
	goose -dir migrations/clickhouse clickhouse "tcp://user:password@127.0.0.1:9003/default" status

migration-cl-reset: # Сбросить все миграции
	goose -dir migrations/clickhouse clickhouse "tcp://user:password@127.0.0.1:9003/default" reset

download-dependencies-if-needed: # Скачиваем все необходимые зависимости для локального запуска
	if [ ! -f "GeoLite2-City.mmdb" ]; then curl -o GeoLite2-City.mmdb "https://s3.sspnet.tech/api/v1/buckets/public/objects/download?prefix=GeoLite2-City.mmdb"; fi
