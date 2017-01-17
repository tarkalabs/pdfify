GO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
docker build -t vagmi/pdfify:latest .
