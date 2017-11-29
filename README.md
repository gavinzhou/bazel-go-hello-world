# bazel hello world 

## build go container image with bazel

```bash
bazel build :go_image
```

## run container with bazel

Need Docker Daemon

```bash
bazel run :go_image
```
