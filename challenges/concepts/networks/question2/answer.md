### 2. Creating an Overlay Network and Attaching Containers

**Q:** How can you create an overlay network and attach multiple containers running on different Docker hosts to it?

**A:** To create an overlay network, you need a Docker Swarm setup. Once Swarm is initialized:

1. Create the overlay network:

    ```bash
    docker network create --driver overlay my_overlay_network
    ```

2. Deploy containers to the overlay network:

    ```bash
    docker service create \
      --name my_overlay_service \
      --network my_overlay_network \
      nginx
    ```

**Use Case:** Overlay networks are used for multi-host networking in Docker Swarm, enabling containers on different hosts to communicate securely.