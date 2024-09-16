# Setting CPU and Memory Limits

**Q:** How can you set CPU and memory limits for a service in Docker Compose?

**A:** Use the `deploy` key to set resource constraints in the `docker-compose.yml` file:

```yaml
version: '3.8'
services:
  web:
    image: nginx
    deploy:
      resources:
        limits:
          cpus: "0.5"
          memory: "512M"
        reservations:
          cpus: "0.2"
          memory: "256M"
