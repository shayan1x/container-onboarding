### 1. Finding the Cgroup Path for a Docker Container

**Q:** How can you locate the cgroup path for a running Docker container?


```bash
docker run -d \
  --name my_nginx \
  --memory 512m \
  --cpus 1 \
  nginx
