build:
	go build -o main.out main.go

run:
	go run main.go

clean:
	go clean
	rm -f main.out
