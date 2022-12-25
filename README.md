# docker-netinspect

## Summary

Get containers' IP address

## Usage

```bash

$ docker ps

CONTAINER ID   IMAGE     COMMAND                  CREATED             STATUS             PORTS     NAMES
6c6d4d5a8f1c   nginx     "/docker-entrypoint.…"   About an hour ago   Up About an hour   80/tcp    nginx-test_2
5fa4bee5cac5   nginx     "/docker-entrypoint.…"   About an hour ago   Up About an hour   80/tcp    nginx-test

$ ./bin/docker-netinspect 
nginx-test_2    172.17.0.3
nginx-test      172.17.0.2

$ ./bin/docker-netinspect 2
nginx-test_2    172.17.0.3
```
