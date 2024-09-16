### 3. Listing Cgroups Hierarchy for Docker Containers

**Q:** How can you list the cgroup hierarchy to see the structure and relationships between different cgroups?

**A:** Use the `systemd-cgls` command to view the hierarchical tree of cgroups:

```bash
systemd-cgls