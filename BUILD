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
    install = ["..."],
    version = "v1.7.0",
    deps = [
        ":go_difflib",
        ":go_yaml",
        ":go_spew",
        ":objx",
    ],
)

go_module(
    name = "go_difflib",
    module = "github.com/pmezard/go-difflib",
    install = ["..."],
    version = "v1.0.0",
    deps = [":go_download"],
)

go_module(
    name = "go_yaml",
    module = "gopkg.in/yaml.v3",
    install = ["..."],
    version = "v3.0.0-20200313102051-9f266ea9e77c",
    deps = [":go_download"],
)

go_module(
    name = "go_spew",
    module = "github.com/davecgh/go-spew",
    install = ["..."],
    version = "v1.1.0",
    deps = [":go_download"],
)

go_module(
    name = "objx",
    module = "github.com/stretchr/objx",
    install = ["..."],
    version = "v0.1.0",
    deps = [":go_download"],
)