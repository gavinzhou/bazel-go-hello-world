load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library", "go_prefix")

go_prefix("github.com/gavinzhou/bazel-go-hello-world")

gazelle(
    name = "gazelle",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/gavinzhou/bazel-go-hello-world",
    visibility = ["//visibility:private"],
    deps = ["@com_github_golang_glog//:go_default_library"],
)

go_binary(
    name = "bazel-go-hello-world",
    importpath = "github.com/gavinzhou/bazel-go-hello-world",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
