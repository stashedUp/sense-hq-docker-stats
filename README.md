
# Sense Docker Container Stats

This is my experimental repo. This is not my official repo. My professional repo is https://github.com/warrensbox. 

## Step 1
Build application: `go build -o sense-stats main.go`
or `make`

## Step 2
Run application: `chmod 755 sense-stats`, `./sense-stats`

## Step 3
Get json request: `curl http://localhost:8080/stats` or test on browser

### Expected output if there are containers running:
```sh
➜  ~ curl http://localhost:8080/stats
[{"container":"17f508642f42","memory":{"raw":"484KiB / 1.939GiB","percent":"0.02%"},"cpu":"0.00%","io":{"network":"3.08kB / 0B","block":"164kB / 0B"},"pids":1},{"container":"b2424d11df69","memory":{"raw":"324KiB / 1.939GiB","percent":"0.02%"},"cpu":"0.00%","io":{"network":"3.29kB / 0B","block":"0B / 0B"},"pids":1}]
```

### Debug:
Make sure this application can find your docker installation. Here we assume that it's on `/usr/local/bin/docker`. Run `which docker` to find installation location.