all: casa

casa:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o casa casa.go

start-app:
	reflex -s -r '\.go$$' -- go run casa.go

start-view:
	reflex -r '\.qtpl$$' -- qtc -dir=view

.PHONT: casa start-app start-view
