# Please Build file for the goroutines Go module.
# Author: Andrew Jarombek
# Date: 9/18/2022

filegroup(
    name = "test_files",
    srcs = glob(["*.go"]),
)

filegroup(
    name = "data_files",
    srcs = glob(["*.json"]),
)

go_test(
    name = "test",
    srcs = [":test_files"],
    data = [":data_files"],
    deps = [
        "//:go_download",
        "//:testify",
    ],
    labels = [
        "goroutines",
        "test"
    ],
)