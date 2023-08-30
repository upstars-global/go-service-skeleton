all: gen-swagger gen-sqlc

gen-sqlc:
	./sqlc.sh

gen-mocks:
	mockery --all --keeptree --dir .
