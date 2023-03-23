FROM Go:Alpine
WORKDIR /app
COPY . WORKDIR
CMD ["go run", "main.go"]