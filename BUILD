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
  tag = "{VERSION}",
)