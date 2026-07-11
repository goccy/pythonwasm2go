# pythonwasm2go

wasm2go sources for python.wasm — a build of
[CPython](https://github.com/python/cpython) 3.14.6 transpiled to Go by
[wasm2go](https://github.com/goccy/wasm2go), consumed by
[go-python](https://github.com/goccy/go-python).

The bundle is generated and released by
[python-wasm](https://github.com/goccy/python-wasm) and vendored here with
`make bundle`, which downloads the release tarball and verifies it against its
sha256 manifest and the upstream GitHub SLSA attestation before extracting it in
place.

## License

This bundle (`base/`, `p0/`–`p2/`, `data.bin`, and the generated Go glue) is a
derivative work of CPython, so it is distributed under CPython's license — the
Python Software Foundation License Agreement — reproduced in full in
[LICENSE](./LICENSE). CPython is Copyright (c) 2001-present Python Software
Foundation; All Rights Reserved. See <https://docs.python.org/3/license.html>.
