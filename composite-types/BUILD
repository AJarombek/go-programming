# Please Build file for the composite_types Go module.
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
        "composite_types",
        "test"
    ],
)