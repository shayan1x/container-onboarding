
### 2. Using a Different Logging Driver

**Q:** How can you configure Docker Compose to use a different logging driver, such as `syslog`?

**A:** Specify the logging driver in the Docker Compose file:

```yaml
version: '3'
services:
  app:
    image: my_app
    logging:
      driver: "syslog"
      options:
        syslog-address: "tcp://192.168.1.100:514"
