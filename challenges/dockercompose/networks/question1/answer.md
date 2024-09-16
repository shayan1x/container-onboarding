# Creating a Default Network

**Q:** How does Docker Compose create and use a default network for services, and what is the behavior of this network?

**A:** By default, Docker Compose creates a network for the services defined in the `docker-compose.yml` file. This network is a bridge network with a name derived from the project directory and is used for internal communication between containers.

Example `docker-compose.yml`:

```yaml
version: '3'
services:
  web:
    image: nginx
  app:
    image: my_app
