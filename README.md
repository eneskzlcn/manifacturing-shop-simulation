### Manufacturing Shop Simulation Story

A machine tool in a manufacturing shop is turning out parts at the 
rate of every x minutes. As they are finished, the parts are sent to an 
inspector who takes at most y minutes, and at least z minutes (uniform distribution) 
to examine each one and rejects about d% of the parts as faulty.
System runs a simulation for m faulty parts to leave the system.

Parameters:

x: Stands for turning out rate of one part in the shop.

y: Maximum examine time of a part by inspector.

z: Minimum examine time of a part by inspector.

d: Possibility to examine a part as faulty in percentage.

m: Count of faulty parts to leave the system/ end the simulation.

### About Project

**manufacturing-shop-simulation** is a cli tool that provides you
to simulate the behaviour of the manufacturing shop described above.

### How To Build
You can simply build the application with the following command:
```shell
make build
```
or
```shell
go build -o bin/simulation ./cmd/simulation
```
To build docker image, you can use 
```shell
make dockerize
```
that gives you an image has the tagged name eneskzlcn/manufacturing-shop-simulation:latest. Or you can simply use
```shell
docker build -t manufacturing-shop-simulation:latest .
```

### How To Run

- If you have go 1.18 available in your system just run the command to compile
and run the program;
```shell
make run
```
or you can directly run one of the following commands:
```shell
go build -o bin/simulation ./cmd/simulation && ./bin/simulation
//or
go run ./cmd/simulation/main.go
```
- If you have docker available on your system you can simply pull the image from my github
repository and run.

- Or you can build the image stands on project directory yourself and use it. Suppose that your
terminal on the project's working directory;
```shell
docker build -t manufacturing-shop-simulation:latest .

docker run manufacturing-shop-simulation:latest -x=2 -y=5 -z=3 -m=100 -d=20
```
