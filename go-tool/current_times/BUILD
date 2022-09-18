# Please Build file for the current_times Go module.
# Author: Andrew Jarombek
# Date: 9/17/2022

go_library(
    name = "current_times",
    srcs = ["current_times.go"],
    deps = ["//:go_download"],
    labels = ["go_tool"],
)

go_test(
    name = "test",
    srcs = ["current_times_test.go"],
    deps = [
        ":current_times",
        "//:go_download",
        "//:testify",
    ],
    labels = [
        "go_tool",
        "test",
    ],
)