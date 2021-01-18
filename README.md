Workflow as a framwork demo
=============================

Requirements
------------
+ docker
+ go (tested in 1.16Beta1, but should work fine for 1.14.x and 1.15.x)

Run the demo
------------
1. Start temporal server locally
```bash
$ docker-compose up -d
```
2. Start worker 
```bash
$ go run worker/main.go
```
3. Start workflow (In a separate terminal)
```bash
$ go run trigger-linear/main.go
```
or 
```bash
$ go run trigger-parallel/main.go
```
4. Watch logs in the worker tab (step 2.) and overall progress on the [temporal ui](http://127.0.0.1:8088/namespaces/default)
5. Enjoy your cookies 🍪
