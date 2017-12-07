[![CircleCI](https://circleci.com/gh/gavinzhou/bazel-go-hello-world.svg?style=svg)](https://circleci.com/gh/gavinzhou/bazel-go-hello-world)
# bazel hello world

## build binary

```bash
bazel build :hello
```

## build go container image with bazel

```bash
bazel build :go_image
```

## run container with bazel

Need Docker Daemon

```bash
bazel run :go_image
```
