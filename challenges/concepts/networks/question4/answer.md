### 4. Connecting a Container to Multiple Networks

**Q:** How can you connect a Docker container to multiple networks and verify its network connections?

**A:** To connect a container to multiple networks:

1. Create multiple networks:

    ```bash
    docker network create my_network1
    docker network create my_network2
    ```

2. Run a container and connect it to both networks:

    ```bash
    docker run -d \
      --name my_multi_network_container \
      --network my_network1 \
      nginx
    ```

    ```bash
    docker network connect my_network2 my_multi_network_container
    ```

3. Verify connections:

    ```bash
    docker inspect my_multi_network_container
    ```

**Use Case:** Connecting containers to multiple networks allows them to communicate with different segments of your infrastructure, useful for complex applications requiring access to various network resources.