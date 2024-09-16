
### 2. Mounting Host Directories

**Q:** How can you mount a host directory as a volume in Docker Compose, and what is an example configuration?

**A:** To mount a host directory as a volume, use the following syntax:

```yaml
version: '3'
services:
  app:
    image: my_app
    volumes:
      - ./host_directory:/app/data
