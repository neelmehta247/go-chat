load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "service",
    srcs = ["hello_service.go"],
    importpath = "github.com/neelmehta247/go-chat/backend/internal/grpc/service",
    visibility = ["//backend:__subpackages__"],
    deps = ["//proto"],
)
