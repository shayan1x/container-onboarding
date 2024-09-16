### 2. Viewing Cgroup Resource Limits

**Q:** How can you view the CPU and memory limits set for a Docker container using cgroups?

**A:** For cgroup v2, you can check resource limits by reading the appropriate files in the cgroup hierarchy:

1. **CPU Limits:**

    ```bash
    cat /sys/fs/cgroup/cpu.max
    ```

2. **Memory Limits:**

    ```bash
    cat /sys/fs/cgroup/memory.max
    ```

   Make sure to locate the specific cgroup directory for your container, usually under `/sys/fs/cgroup/<container_cgroup>`.