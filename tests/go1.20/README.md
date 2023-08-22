# Test for explicitly setting/inferring go version from toolchain

This uses a syntax feature added in go 1.20 (conversion of slice to an
array with length) to ensure that the build fails when using go 1.18
and succeeds with go 1.20.
