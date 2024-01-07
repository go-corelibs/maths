[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/maths)
[![codecov](https://codecov.io/gh/go-corelibs/maths/graph/badge.svg?token=0LMMOkCFIX)](https://codecov.io/gh/go-corelibs/maths)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/maths)](https://goreportcard.com/report/github.com/go-corelibs/maths)

# maths - mathematics and number-type utilities

maths is a package for general purpose mathematics and other number-type things.

# Installation

``` shell
> go get github.com/go-corelibs/maths@latest
```

# Examples

## Clamp, ToInt

``` go
func main() {
    clamped := maths.Clamp(2.5, 0.0, 1.0)
    // clamped == float64(1.0)
    integer := maths.ToInt(10.0)
    // integer == int(10)
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2023 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
