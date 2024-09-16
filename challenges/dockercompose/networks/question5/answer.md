
### 6. Configuring Network Aliases

**Q:** How can you configure network aliases for a service in Docker Compose, and what is the benefit?

**A:** Set network aliases in the `docker-compose.yml` file:

```yaml
version: '3'
services:
  web:
    image: nginx
    networks:
      custom_network:
        aliases:
          - webalias
networks:
  custom_network:
