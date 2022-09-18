# Please Build file for the unit_testing Go module.
# Author: Andrew Jarombek
# Date: 9/17/2022

go_library(
    name = "license_plates",
    srcs = ["license_plates.go"],
    deps = ["//:go_download"],
    labels = ["unit_testing"],
)

go_test(
    name = "test",
    srcs = ["license_plates_test.go"],
    deps = [
        ":license_plates",
        "//:go_download",
    ],
    labels = [
        "unit_testing",
        "test",
    ],
)