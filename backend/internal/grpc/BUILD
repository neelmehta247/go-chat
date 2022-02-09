load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "grpc",
    srcs = ["server.go"],
    importpath = "github.com/neelmehta247/go-chat/backend/internal/grpc",
    visibility = ["//visibility:public"],
    deps = [
        "//backend/internal/grpc/service",
        "//proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection",
    ],
)
