# Docker Containers and Host Network Namespace: Challenge Questions

## Scenario

You need to create Docker containers that utilize the host's network namespace. Answer the following questions to explore how to configure and use the host network namespace effectively.

### 1. Running a Container with Host Network Namespace

**Q:** How can you run a Docker container with the host's network namespace, and what is a use case for this configuration?

**A:** To run a container with the host's network namespace, use the `--network host` option:

```bash
docker run -d \
  --name my_host_network \
  nginx
