# Docker Containers and Namespaces: Challenge Questions

## Scenario

You need to create Docker containers and understand how to use and manage different namespaces for various tasks. Answer the following questions to explore these concepts.

### 1. Running a Container with Host PID Namespace

**Q:** How can you run a Docker container with the host's PID namespace and what is a use case for this configuration?

**A:** To run a container with the host's PID namespace, use the `--pid` option:

```bash
docker run -d \
  --name my_htop \
  --pid host \
  alpine:latest \
  sh -c "apk add --no-cache htop && htop"
