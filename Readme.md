# Memory Leaker
Service that will help you understand how your platform reacts to a memory leak.

# Usage
The service will allocate memory at a rate of 10MB per second. You can change the rate by passing the `--mb-per-second` flag. For example:

```bash
docker run wincus/memory-leaker --mb-per-second 1024
```

Will result on a process allocating 1GB per second.