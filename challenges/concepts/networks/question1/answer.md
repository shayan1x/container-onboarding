# Docker Networking: Challenge Questions

## Scenario

You need to create and manage Docker networks and attach them to containers. Answer the following questions to explore different aspects of Docker networking.

### 1. Creating and Using a Bridge Network

**Q:** How can you create a bridge network in Docker and attach a container to it?

**A:** To create a bridge network and attach a container:

1. Create the bridge network:

    ```bash
    docker network create my_bridge_network
    ```

2. Run a container and attach it to the bridge network:

    ```bash
    docker run -d \
      --name my_bridge_container \
      --network my_bridge_network \
      nginx
    ```

**Use Case:** Bridge networks are used for communication between containers on the same host, isolating them from external networks.