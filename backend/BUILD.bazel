load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "backend",
    embed = [":backend_lib"],
    importpath = "github.com/neelmehta247/go-chat/backend",
    visibility = ["//visibility:public"],
)

go_library(
    name = "backend_lib",
    srcs = ["main.go"],
    importpath = "github.com/neelmehta247/go-chat/backend",
    visibility = ["//visibility:private"],
    deps = ["//proto"],
)
