# Basic Logging Configuration

**Q:** How can you configure basic logging for a service in Docker Compose?

**A:** By default, Docker Compose uses the `json-file` logging driver. To specify a logging configuration in Docker Compose:

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
