
### 3. Creating a Network with a Specific IP Range

# Creating a Network with a Specific IP Range

**Q:** How can you create a custom network in Docker Compose with a specific IP range, and how does it affect service assignment?

**A:** Define a custom network with a specific IP range in the `docker-compose.yml` file:

```yaml
version: '3'
services:
  web:
    image: nginx
    networks:
      custom_network:
        ipv4_address: 172.16.1.2
  app:
    image: my_app
    networks:
      custom_network:
        ipv4_address: 172.16.1.3
networks:
  custom_network:
    ipam:
      config:
        - subnet: 172.16.1.0/24
