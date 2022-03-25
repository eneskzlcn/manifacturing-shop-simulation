**The problem definition:**
A machine tool in a manufacturing shop is turning out parts at the 
rate of every 5 minutes. As they are finished, the parts are sent to an 
inspector who takes 6±4 minutes (uniform distribution) to examine each 
one and rejects about 10% of the parts as faulty. You are asked to run 
the simulation for 100 faulty parts to leave the system. Gather the 
statistics about average and maximum queue lengths.

**Problem Reference From:** Discrete Modelling And Simulation Lecture On Fatih Sultan
Mehmet Vakif University.

### **To Run The Program**

- If you have make and go 1.17 available in your system just run the command to compile
and run the program;
``` bash
    make run
```
- If you have docker available on your system you can simply run the command to compile
and run.
```bash
    docker run eneskzlcn/manifacturing-shop-simulation:latest
```
- Or you can build the image stands on project directory yourself and use it. Suppose that your
terminal on the project's working directory;
```bash
    docker build -t manifacturing-shop-simulation:latest
    docker run manifacturing-shop-simulation:latest
```

Docker image available in the link below.

https://hub.docker.com/repository/docker/eneskzlcn/manifacturing-shop-simulation