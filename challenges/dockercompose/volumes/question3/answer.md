
### 4. Volume Mount Options

**Q:** How can you specify mount options such as read-only for volumes in Docker Compose?

**A:** Use mount options in the volume definition:

```yaml
version: '3'
services:
  web:
    image: nginx
    volumes:
      - my_volume:/usr/share/nginx/html:ro
volumes:
  my_volume:
