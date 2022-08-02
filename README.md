# DevOps - Getting started!
## Web application written in Go using the Fiber/v2 routing module
![title-page](https://imgur.com/GvXp6po.jpg)

## Quickstart
- Clone this repository
- Navigate into the project root directory
- Run the command `go run main.go`

---

## Building the project
- EP1 - [Build Test Deploy](https://youtu.be/mw3AaxtEVSY)
- EP2 - [GitHub Action Automation](https://youtu.be/SfWvEwQUnD8)
- EP3 - [Infrastructure as Code (IaC)](https://youtu.be/j1x4TDOfvjA)
- EP4 - [Custom domains in Azure](https://youtu.be/cr32MQx3d80)

---

## Dockerfile build
- Ensure podman or Docker are running on your local machine
- Using either of the tools, run the commands shown below:
```bash
$ docker build --tag docker.io/doa-ep1:v1 .
$ docker run -dt -p 3000:3000/tcp docker.io/library/doa-ep1:v1
```
In a browser, navigate to `localhost:3000`