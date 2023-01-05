BUNDLE_NAME=login-code-challenge

# Run back-end
be:
	@cd server && go run *.go

# Build back-end
build:
	@cd server && go build -o $(BUNDLE_NAME)

# Run front-end
fe:
	@cd client && npm run start
