load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "github.com/neelmehta247/go-chat/backend/_internal/grpc",
    visibility = ["//visibility:public"],
    deps = ["//proto"],
)
