# Please Build file for the go_tool Go module.
# Author: Andrew Jarombek
# Date: 9/17/2022

filegroup(
    name = "test_files",
    srcs = glob(["*.go"]),
)

go_test(
    name = "test",
    srcs = [":test_files"],
    deps = [
        "//:go_download",
        "//:testify",
    ],
    labels = [
        "go_tool",
        "test"
    ],
)

sh_binary(
    name = "go_run",
    srcs = ["run.sh"],
    labels = [
       "go_tool",
       "shell"
    ],
)
