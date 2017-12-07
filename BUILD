# Go boilerplate
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_push",
    "container_bundle",
)

go_binary(
    name = "hello",
    srcs = ["main.go"],
    importpath = ".",
)

go_image(
    name = "go_image",
    srcs = ["main.go"],
    importpath = ".",
)

container_push(
  name = "push",
  format = "Docker",
  image = ":go_image",
  registry = "index.docker.io",
  repository = "orangesys/bazel-go-hello-world",
  stamp = True,
)

container_bundle(
    name = "bundle_to_push",
    images = {
        "index.docker.io/orangesys/bazel-go-hello-world:{VERSION}": ":go_image",
         "index.docker.io/orangesys/bazel-go-hello-world:latest": ":go_image",
    },
    stamp = True,
)
load(
    "@io_bazel_rules_docker//contrib:push-all.bzl",
    docker_pushall = "docker_push",
)

docker_pushall(
    name = "push_bundle",
    bundle = ":bundle_to_push",
)