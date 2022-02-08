load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "hello_proto",
    srcs = ["hello.proto"],
    visibility = ["//visibility:private"],
)

go_proto_library(
    name = "hello_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/neelmehta247/go-chat/proto",
    proto = ":hello_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "proto",
    embed = [":hello_go_proto"],
    importpath = "github.com/neelmehta247/go-chat/proto",
    visibility = ["//visibility:public"],
)
