# Please Build file for the go-programming repository.
# Author: Andrew Jarombek
# Date: 9/17/2022

package(default_visibility = ["PUBLIC"])

go_toolchain(
    name = "go_download",
    version = "1.18",
)

go_module(
    name = "testify",
    module = "github.com/stretchr/testify",
    version = "v1.7.0",
    deps = [":go_download"],
)