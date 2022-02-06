# Merge Service
This implements an interval merging functionality as a backend service in Golang. The requests to merge can be send via REST API. The merged intervals can be found in the response from the server.


# Architecture
The application is set up as a microservice with a REST API to merge intervals. You can build it with your toolchain or use the docker variant.

## API
- [api description](api/README.md)

## Algorithm
Here you can find the implemented [algorithm](doc/Algorithm.md) for this service.

## Build
You can build with your own toolchain or use the Docker version (preferred).
### local
If you didn't change anything you don't need to build the application and can rather jump to "Start the service".
Download the dependencies first. This will download all the needed go modules needed to build the service.
```shell
make download-deps 
```
Now build the service.
```shell
make build-local
```

### Docker
```shell
make build
```

## Start the service
### local
If you build the service with your toolchain:
```shell
make start-local
```
If you want to start with metrics:
```shell
make start-local-with-metrics
```
### Docker
If you used the docker variant:
```shell
make start
```
If you want to start with metrics:
```shell
make start-with-metrics
```

## Test
Run the unit test with code coverage result.
```shell
make test-local
```

# Open points
- OpenApi spec missing in ./api
- Integration test missing