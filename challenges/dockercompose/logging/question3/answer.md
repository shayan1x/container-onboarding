
### 7. Centralized Logging Configuration

**Q:** How can you configure centralized logging for multiple services in Docker Compose?

**A:** Define a common logging configuration for all services:

```yaml
version: '3'
services:
  web:
    image: nginx
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
  app:
    image: my_app
    logging:
      driver: "json-file"
      options:
        max-size: "10m"
        max-file: "3"
