ARTIFACT='hinter-tui'

.PHONY: build 
build:
	go build -o ${ARTIFACT}

.PHONY: develop
develop: 
	go run main.go

.PHONY: clean
clean:
	rm ${ARTIFACT}