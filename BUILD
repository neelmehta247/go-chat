load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/neelmehta247/go-chat
# gazelle:build_file_name BUILD,BUILD.bazel
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
