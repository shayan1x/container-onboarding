# Scenario: Managing Nginx Container with Resource Limits

## Task

You need to run an Nginx container with specific memory and CPU limits. After running the container, you should be able to retrieve information about its cgroup and namespace settings.

### Step 1: Run the Nginx Container with Resource Limits

Use the following command to run an Nginx container with memory and CPU limits:

```bash
docker run -d \
  --name my_nginx \
  --memory 512m \
  --cpus 1 \
  nginx
