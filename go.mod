module github.com/antonybholmes/go-edb-server-mailer

go 1.23

replace github.com/antonybholmes/go-mailer => ../go-mailer

replace github.com/antonybholmes/go-sys => ../go-sys

require (
	github.com/antonybholmes/go-sys v0.0.0-20241219150701-64d5bd9623f7
	github.com/rs/zerolog v1.33.0
)

require (
	github.com/aws/aws-sdk-go v1.55.5 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

require (
	github.com/antonybholmes/go-mailer v0.0.0-20241204182257-78d1094d458e
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/redis/go-redis/v9 v9.7.0
	golang.org/x/sys v0.28.0 // indirect
)
