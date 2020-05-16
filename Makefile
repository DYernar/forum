GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=forum
    BINARY_UNIX=$(BINARY_NAME)_unix
    
    all: test build
    build: 
			$(GOBUILD) -o $(BINARY_NAME) -v
    test: 
			$(GOTEST) -v ./...
    clean: 
			$(GOCLEAN)
			rm -f $(BINARY_NAME)
			rm -f $(BINARY_UNIX)
    run:
			$(GOBUILD) -o $(BINARY_NAME) -v ./...
			./$(BINARY_NAME)
    deps:
			$(GOGET) github.com/mattn/go-sqlite3 v2.0.3+incompatible
			$(GOGET) github.com/satori/go.uuid v1.2.0
			$(GOGET) golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
    
    
    # Cross compilation
    build-linux:
			CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
    docker-build:
			docker build -t forum .
			docker run -it -p 9090:9090 forum 
