### 1. Finding the Cgroup Path for a Docker Container

**Q:** How can you locate the cgroup path for a running Docker container?

**A:** You can locate the cgroup path by first finding the container’s PID and then checking the cgroup filesystem:

1. Obtain the PID of the container:

    ```bash
    docker inspect --format '{{ .State.Pid }}' <container_name_or_id>
    ```

2. Use the PID to find the cgroup path:

    ```bash
    grep $(docker inspect --format '{{ .State.Pid }}' <container_name_or_id>) /sys/fs/cgroup/cgroup.procs
    ```

