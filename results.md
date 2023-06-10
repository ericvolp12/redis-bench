# Redis Benchmark Results

I have been exploring both DragonflyDB and Redis-Stack as options for in-memory datastores.

In this exploration I decided to run a few benchmarks in Go using the `github.com/redis/go-redis/v9` client.

## Local Tests (Docker NAT)

All tests were performed on localhost behind Docker NAT, on a `AMD EPYC 7302P 16-Core Processor` inside a Proxmox VM with 24GB of allocated RAM available.

| Test         | Inserts | Value Size | Reads | Pipeline Size | redis-stack | dragonflydb |
|--------------|---------|------------|-------|---------------|-------------|-------------|
| No-Pipeline  | 100k    | 5b         | 5m    | N/A           | 64.117s     | 66.052s     |
| No-Pipeline  | 100k    | 1kb        | 5m    | N/A           | 69.527s     | 70.019s     |
| No-Pipeline  | 100k    | 10kb       | 5m    | N/A           | 101.712s    | 90.429s     |
| Pipeline     | 500k    | 5b         | 25m   | 10k           | 16.583s     | 35.217s     |
| Pipeline     | 500k    | 1kb        | 25m   | 10k           | 42.633s     | 66.865s     |
| Pipeline     | 500k    | 1kb        | 25m   | 1k            | 39.503s     | 68.294s     |
| Pipeline     | 100k    | 10kb       | 5m    | 1k            | 49.189s     | 26.303s     |

From these results we can see that `redis-stack` handles pipelined reads with higher throughput than `dragonflydb`. The tests used pipelines with 10,000 commands in each to prevent I/O errors.

The `redis-stack` utilized 100% of its single CPU core during execution while the `dragonflydb` utilized around 45-60% of the 16 CPU cores assigned to the VM for the duration of the test.

Redis recommends running `Redis Cluster` on a single host to shard out keys more effectively and increase throughput, I've yet to test that but given how little CPU utilization the single `redis-stack` consumed, I'm interested in conducting further testing.

### Versions
Docker Commands
```
docker run --rm -it --name=dragonfly_bench -p 6385:6379 docker.dragonflydb.io/dragonflydb/dragonfly
docker run --rm -it --name redis-stack -p 6380:6379 redis/redis-stack-server:latest
```

```
DragonflyDB: df-v1.3.0-f80afca9c23e2f30373437520a162c591eaa2005
Redis: 6.2.12 - oss
```


## GCP Tests (Host Network)

Tests below were run on GCP `t2d` instances (AMD EPYC Milan) with nothing else running on them, all containers were run in host network mode to avoid docker proxy bottlenecks.

Times shown are an average of 3 executions with DBs recreated in between each run.

### `t2d-standard-2`

| Test Name   | Inserts | Value Size | Reads      | Pipeline Size | Repetitions | redis-stack (write) | dragonfly (write) | redis-stack (read) | dragonfly (read) |
|-------------|---------|------------|------------|---------------|-------------|---------------------|-------------------|--------------------|------------------|
| No-Pipeline | 100,000 | 5          | 5,000,000  | -1            | 3           | 4.106s              | 6.427s            | 44.442s            | 78.345s          |
| No-Pipeline | 100,000 | 1,000      | 5,000,000  | -1            | 3           | 4.002s              | 6.516s            | 45.847s            | 84.976s          |
| No-Pipeline | 100,000 | 10,000     | 5,000,000  | -1            | 3           | 5.058s              | 8.190s            | 77.472s            | 133.798s         |
| Pipeline    | 500,000 | 5          | 25,000,000 | 10,000        | 3           | 1.488s              | 8.813s            | 15.443s            | 34.237s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 10,000        | 3           | 2.247s              | 11.023s           | 36.344s            | 56.709s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 1,000         | 3           | 2.393s              | 8.879s            | 35.333s            | 78.564s          |
| Pipeline    | 100,000 | 10,000     | 5,000,000  | 1,000         | 3           | 1.588s              | 3.848s            | 36.134s            | 36.020s          |


### `t2d-standard-4`

| Test Name   | Inserts | Value Size | Reads      | Pipeline Size | Repetitions | redis-stack (write) | dragonfly (write) | redis-stack (read) | dragonfly (read) |
|-------------|---------|------------|------------|---------------|-------------|---------------------|-------------------|--------------------|------------------|
| No-Pipeline | 100,000 | 5          | 5,000,000  | -1            | 3           | 4.307s              | 7.780s            | 46.0463s           | 85.603s          |
| No-Pipeline | 100,000 | 1,000      | 5,000,000  | -1            | 3           | 4.482s              | 7.989s            | 50.594s            | 90.717s          |
| No-Pipeline | 100,000 | 10,000     | 5,000,000  | -1            | 3           | 6.260s              | 9.310s            | 85.990s            | 142.115s         |
| Pipeline    | 500,000 | 5          | 25,000,000 | 10,000        | 3           | 1.538s              | 12.625s           | 15.214s            | 37.685s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 10,000        | 3           | 2.290s              | 14.342s           | 37.069s            | 64.967s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 1,000         | 3           | 2.470s              | 12.218s           | 35.964s            | 142.487s         |
| Pipeline    | 100,000 | 10,000     | 5,000,000  | 1,000         | 3           | 1.644s              | 4.160s            | 37.156s            | 41.889s          |

### `t2d-standard-8`

| Test Name   | Inserts | Value Size | Reads      | Pipeline Size | Repetitions | redis-stack (write) | dragonfly (write) | redis-stack (read) | dragonfly (read) |
|-------------|---------|------------|------------|---------------|-------------|---------------------|-------------------|--------------------|------------------|
| No-Pipeline | 100,000 | 5          | 5,000,000  | -1            | 3           | 3.602s              | 7.468s            | 27.297s            | 22.250s          |
| No-Pipeline | 100,000 | 1,000      | 5,000,000  | -1            | 3           | 3.678s              | 7.619s            | 29.751s            | 24.723s          |
| No-Pipeline | 100,000 | 10,000     | 5,000,000  | -1            | 3           | 4.765s              | 8.887s            | 43.024s            | 40.605s          |
| Pipeline    | 500,000 | 5          | 25,000,000 | 10,000        | 3           | 1.480s              | 15.327s           | 13.444s            | 12.918s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 10,000        | 3           | 2.266s              | 17.051s           | 32.299s            | 21.063s          |
| Pipeline    | 500,000 | 1,000      | 25,000,000 | 1,000         | 3           | 2.332s              | 14.921s           | 30.711s            | 39.722s          |
| Pipeline    | 100,000 | 10,000     | 5,000,000  | 1,000         | 3           | 1.472s              | 3.873s            | 38.814s            | 14.283s          |

### `t2d-standard-16`

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
