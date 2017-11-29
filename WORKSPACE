git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.6.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

go_rules_dependencies()

go_register_toolchains()

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    tag = "v0.3.0",
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

container_pull(
  name = "distroless_base",
  registry = "gcr.io",
  repository = "distroless/base",
)

container_pull(
  name = "java_base",
  registry = "gcr.io",
  repository = "distroless/java",
)

container_pull(
  name = "alpine_base",
  registry = "index.docker.io",
  repository = "alpine",
  tag = "3.5"
)

container_pull(
  name = "ubuntu_base",
  registry = "index.docker.io",
  repository = "library/ubuntu",
  tag = "16.04"
)

container_pull(
    name = "bitnami_minideb",
    registry ="index.docker.io",
    repository = "bitnami/minideb-extras",
    tag = "jessie",
)
