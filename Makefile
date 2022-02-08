.PHONY: run gazelle gazelle-update-repos generate

run:
	bazel run //backend

gazelle:
	bazel run //:gazelle

gazelle-update-repos:
	bazel run //:gazelle-update-repos

generate:
	bazel run //proto:copyall
