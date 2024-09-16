# Docker Containers: Challenge Question

## Scenario

You need to configure a Docker container so that it uses the same hostname as the host system by leveraging the UTS namespace. Answer the following question to explore this concept.

### Ensuring a Container Uses the Same Hostname as the Host System

**Q:** How can you run a Docker container so that it shares the same hostname as the host system using the UTS namespace, and why might you want to do this?

**A:** To ensure that a Docker container uses the same hostname as the host system, you need to run the container with the `--hostname` option set to the host's hostname and ensure the container uses the host’s UTS namespace:

1. First, get the hostname of the host system:

    ```bash
    hostname
    ```

2. Run the Docker container with the same hostname:

    ```bash
    docker run -d \
      --name my_shared_hostname_container \
      --uts host \
      nginx
    ```

**Why You Might Want to Do This:** Sharing the hostname between the host and container can be useful in scenarios where applications need to resolve or rely on a consistent hostname for network operations, logging, or identity purposes. This setup can simplify interactions between the container and host, particularly when dealing with software that requires hostname-based configurations or when integrating with legacy systems that expect a specific hostname.

**Verification:** To verify that the container is using the same hostname as the host system, execute the following command inside the container:

```bash
docker exec -it my_shared_hostname_container hostname
