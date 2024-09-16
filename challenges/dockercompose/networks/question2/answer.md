
### 2. Creating and Using a Custom Network

**Q:** How can you create and use a custom network in Docker Compose, and how does it affect service communication?

**A:** Define a custom network in the `docker-compose.yml` file:

```yaml
version: '3'
services:
  web:
    image: nginx
    networks:
      - custom_network
  app:
    image: my_app
    networks:
      - custom_network
networks:
  custom_network:
