# Redis Benchmark Results

I have been exploring both DragonflyDB and Redis-Stack as options for in-memory datastores.

In this exploration I decided to run a few benchmarks in Go using the `github.com/redis/go-redis/v9` client.

## GCP Tests (Host Network)

Tests below were run on GCP instances with nothing else running on them, all containers were run in host network mode to avoid docker proxy bottlenecks.

Times shown are an average of 3 executions with DBs recreated in between each run.

### `t2d-standard-2` (AMD Epyc Milan)

| Test Name   | Inserts | Value Size | Reads      | Pipeline Size | Repetitions | redis-stack (write) | dragonfly (write) | redis-stack (read) | dragonfly (read) |
|-------------|---------|------------|------------|---------------|-------------|---------------------|-------------------|--------------------|------------------|
| No-Pipeline | 100,000 | 5          | 5,000,000  | -1            | 3           | 4.106s              | 6.427s            | 44.442s            | 78.345s          |
| No-Pipeline | 100,000 | 1,000      | 5,000,000  | -1            | 3           | 4.002s              | 6.516s            | 45.847s            | 84.976s          |
| No-Pipeline | 100,000 | 10,000     | 5,000,000  | -1            | 3           | 5.058s              | 8.190s            | 77.472s            | 133.798s         |
| Pipeline    | 500,000 | 5          | 25,000,000 | 10,000        | 3           | 1.488s              | 8.813s            | 15.443s            | 34.237s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 10,000        | 3           | 2.247s              | 11.023s           | 36.344s            | 56.709s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 1,000         | 3           | 2.393s              | 8.879s            | 35.333s            | 78.564s          |
| Pipeline    | 100,000 | 10,000     | 5,000,000  | 1,000         | 3           | 1.588s              | 3.848s            | 36.134s            | 36.020s          |

### `t2d-standard-4` (AMD Epyc Milan)

| Test Name   | Inserts | Value Size | Reads      | Pipeline Size | Repetitions | redis-stack (write) | dragonfly (write) | redis-stack (read) | dragonfly (read) |
|-------------|---------|------------|------------|---------------|-------------|---------------------|-------------------|--------------------|------------------|
| No-Pipeline | 100,000 | 5          | 5,000,000  | -1            | 3           | 3.453s              | 5.540s            | 23.100s            | 27.747s          |
| No-Pipeline | 100,000 | 1,000      | 5,000,000  | -1            | 3           | 3.603s              | 5.586s            | 25.138s            | 30.318s          |
| No-Pipeline | 100,000 | 10,000     | 5,000,000  | -1            | 3           | 4.536s              | 6.565s            | 38.234s            | 56.348s          |
| Pipeline    | 500,000 | 5          | 25,000,000 | 10,000        | 3           | 1.438s              | 7.524s            | 13.398s            | 17.811s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 10,000        | 3           | 2.141s              | 8.338s            | 34.349s            | 31.938s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 1,000         | 3           | 2.156s              | 7.699s            | 27.274s            | 52.570s          |
| Pipeline    | 100,000 | 10,000     | 5,000,000  | 1,000         | 3           | 1.315s              | 2.403s            | 36.558s            | 19.660s          |

### `t2a-standard-4` (Ampere Altra ARM)

| Test Name   | Inserts | Value Size | Reads      | Pipeline Size | Repetitions | redis-stack (write) | dragonfly (write) | redis-stack (read) | dragonfly (read) |
|-------------|---------|------------|------------|---------------|-------------|---------------------|-------------------|--------------------|------------------|
| No-Pipeline | 100,000 | 5          | 5,000,000  | -1            | 3           | 3.215s              | 4.361s            | 31.554s            | 45.531s          |
| No-Pipeline | 100,000 | 1,000      | 5,000,000  | -1            | 3           | 3.463s              | 4.664s            | 35.010s            | 50.393s          |
| No-Pipeline | 100,000 | 10,000     | 5,000,000  | -1            | 3           | 5.090s              | 6.630s            | 60.131s            | 85.508s          |
| Pipeline    | 500,000 | 5          | 25,000,000 | 10,000        | 3           | 1.667s              | 6.310s            | 16.955s            | 24.440s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 10,000        | 3           | 2.560s              | 7.050s            | 39.049s            | 55.376s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 1,000         | 3           | 2.629s              | 7.001s            | 37.631s            | 65.552s          |
| Pipeline    | 100,000 | 10,000     | 5,000,000  | 1,000         | 3           | 1.828s              | 2.253s            | 37.783s            | 36.004s          |

### `t2d-standard-8` (AMD Epyc Milan)

| Test Name   | Inserts | Value Size | Reads      | Pipeline Size | Repetitions | redis-stack (write) | dragonfly (write) | redis-stack (read) | dragonfly (read) |
|-------------|---------|------------|------------|---------------|-------------|---------------------|-------------------|--------------------|------------------|
| No-Pipeline | 100,000 | 5          | 5,000,000  | -1            | 3           | 3.602s              | 7.468s            | 27.297s            | 22.250s          |
| No-Pipeline | 100,000 | 1,000      | 5,000,000  | -1            | 3           | 3.678s              | 7.619s            | 29.751s            | 24.723s          |
| No-Pipeline | 100,000 | 10,000     | 5,000,000  | -1            | 3           | 4.765s              | 8.887s            | 43.024s            | 40.605s          |
| Pipeline    | 500,000 | 5          | 25,000,000 | 10,000        | 3           | 1.480s              | 15.327s           | 13.444s            | 12.918s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 10,000        | 3           | 2.266s              | 17.051s           | 32.299s            | 21.063s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 1,000         | 3           | 2.332s              | 14.921s           | 30.711s            | 39.722s          |
| Pipeline    | 100,000 | 10,000     | 5,000,000  | 1,000         | 3           | 1.472s              | 3.873s            | 38.814s            | 14.283s          |

### `t2d-standard-16` (AMD Epyc Milan)

| Test Name   | Inserts | Value Size | Reads      | Pipeline Size | Repetitions | redis-stack (write) | dragonfly (write) | redis-stack (read) | dragonfly (read) |
|-------------|---------|------------|------------|---------------|-------------|---------------------|-------------------|--------------------|------------------|
| No-Pipeline | 100,000 | 5          | 5,000,000  | -1            | 3           | 3.815s              | 8.345s            | 29.162s            | 21.100s          |
| No-Pipeline | 100,000 | 1,000      | 5,000,000  | -1            | 3           | 4.054s              | 8.197s            | 31.351s            | 23.126s          |
| No-Pipeline | 100,000 | 10,000     | 5,000,000  | -1            | 3           | 5.419s              | 9.741s            | 44.718s            | 33.335s          |
| Pipeline    | 500,000 | 5          | 25,000,000 | 10,000        | 3           | 1.304s              | 20.265s           | 12.301s            | 12.927s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 10,000        | 3           | 2.072s              | 21.752s           | 30.933s            | 20.841s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 1,000         | 3           | 2.091s              | 19.244s           | 27.827s            | 39.162s          |
| Pipeline    | 100,000 | 10,000     | 5,000,000  | 1,000         | 3           | 1.339s              | 4.534s            | 35.128s            | 13.063s          |

### Versions
Docker Commands
```
docker run --rm -it --network host --name=dragonfly_bench   docker.dragonflydb.io/dragonflydb/dragonfly
docker run --rm -it --network host --name redis-stack       redis/redis-stack-server:latest
```

```
DragonflyDB: df-v1.3.0-f80afca9c23e2f30373437520a162c591eaa2005
Redis: 6.2.12 - oss
```

## Running

To run the tests:
- Install docker on your system
- Pull the required docker images (the binary doesn't currently do that for you right now)
  - `docker pull redis/redis-stack-server:latest`
  - `docker pull docker.dragonflydb.io/dragonflydb/dragonfly:latest` 
- Build the binary with `make bench`
- Execute the bench binary with `./redis-bench`

The bench binary uses the Docker API to create containers for each test.

If you'd like to test other backends, update the `backendImageMap` and `backends` properties in `cmd/bench/main.go` and rebuild the binary.

Results will be dumped into a `results.txt` file after all tests are finished in the format of a Markdown Table like the ones above.
