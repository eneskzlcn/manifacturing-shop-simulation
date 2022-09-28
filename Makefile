run:
	go build -o bin/simulation ./cmd/simulation && ./bin/simulation

clean:
	rm -rf bin