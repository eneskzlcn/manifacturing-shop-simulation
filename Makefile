run:
	go build -o bin/simulation ./cmd/simulation && ./bin/simulation

build:
	go build -o bin/simulation ./cmd/simulation

dockerize:
	docker build -t eneskzlcn/manufacturing-shop-simulation:latest .

clean:
	rm -rf bin