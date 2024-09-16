
### 4. Isolating Services with Multiple Networks

**Q:** How can you isolate services using multiple networks in Docker Compose, and what is an example configuration?

**A:** Use multiple networks to control communication between services:

```yaml
version: '3'
services:
  web:
    image: nginx
    networks:
      - front_end
  app:
    image: my_app
    networks:
      - back_end
      - front_end
  db:
    image: postgres
    networks:
      - back_end
networks:
  front_end:
  back_end:
