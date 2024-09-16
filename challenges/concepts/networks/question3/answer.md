### 3. Creating a Macvlan Network and Attaching a Container

**Q:** How can you create a Macvlan network and attach a container to it, assigning the container a specific IP address?

**A:** To create a Macvlan network and attach a container:

1. Create the Macvlan network:

    ```bash
    docker network create \
      --driver macvlan \
      --subnet=192.168.1.0/24 \
      --gateway=192.168.1.1 \
      -o parent=eth0 \
      my_macvlan_network
    ```

2. Run a container with a specific IP address on the Macvlan network:

    ```bash
    docker run -d \
      --name my_macvlan_container \
      --network my_macvlan_network \
      --ip 192.168.1.100 \
      nginx
    ```

**Use Case:** Macvlan networks allow containers to appear as separate physical devices on the network, which can be useful for applications needing direct network access or specific IP assignments.