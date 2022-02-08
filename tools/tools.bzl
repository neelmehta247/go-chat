def gen_cmds(label, lsfiles):
    return ["cp -f {f} $BUILD_WORKSPACE_DIRECTORY/{dir}".format(f = f.short_path, dir = label.label.package) for f in lsfiles]

def _copy_genfiles_impl(ctx):
    files_to_copy = []
    cmds = []
    for label in ctx.attr.proto:
        files = label.output_groups.go_generated_srcs.to_list()
        files_to_copy = files_to_copy + files
        cmds = cmds + gen_cmds(label, files)

    cps = "\n".join(cmds)

    script = """
set -ex
{cps}""".format(cps = cps)
    ctx.actions.write(
        output = ctx.outputs.executable,
        content = script,
    )
    runfiles = ctx.runfiles(files = files_to_copy)
    return [DefaultInfo(runfiles = runfiles)]

copy_genfiles = rule(
    implementation = _copy_genfiles_impl,
    attrs = {
        "proto": attr.label_list(),
    },
    executable = True,
)
