### 5. Creating a Network with Specific DNS Settings and Attaching Containers

**Q:** How can you create a Docker network with custom DNS settings and attach a container to it?

**A:** To create a network with custom DNS settings:

1. Create the network with DNS settings:

    ```bash
    docker network create \
      --driver bridge \
      --dns 8.8.8.8 \
      --dns-search example.com \
      my_dns_network
    ```

2. Run a container attached to the network:

    ```bash
    docker run -d \
      --name my_dns_container \
      --network my_dns_network \
      nginx
    ```

**Use Case:** Custom DNS settings can be applied to networks to control DNS resolution within containers, which is useful for managing domain names and resolving services within a specific namespace.