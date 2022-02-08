.PHONY: build gazelle gazelle-update-repos generate

build:
	bazel build //backend

gazelle:
	bazel run //:gazelle

gazelle-update-repos:
	bazel run //:gazelle-update-repos

generate:
	bazel run //proto:copyall
