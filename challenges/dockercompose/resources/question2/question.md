
### 2. Scaling Services with Replicas

**Q:** How can you scale a service to run multiple replicas in Docker Compose?

**A:** Use the `replicas` option under the `deploy` key:

```yaml
version: '3.8'
services:
  web:
    image: nginx
    deploy:
      replicas: 3
