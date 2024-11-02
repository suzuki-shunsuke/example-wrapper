# example-wrapper

This is an example to wrap Terraform command using [aqua](https://aquaproj.github.io/).

First, [please install aqua](https://aquaproj.github.io/docs/install).

And run `aqua i -l`.

```sh
aqua i -l
```

1. Run `go run ./cmd/terraform version`:

> [!NOTE]
> This approach doesn't need aqua.

```console
$ go run ./cmd/terraform
2024/11/02 10:15:21 [INFO] Terraform wrapper
Terraform v1.9.8
on darwin_arm64
```

2. Install the wrapper by `go install` and run Terraform:

> [!WARNING]
> `$HOME/go/bin` must take precedence over `$(aqua root-dir)/bin`.

```console
$ go install ./cmd/terraform
$ terraform version
2024/11/02 10:15:21 [INFO] Terraform wrapper
Terraform v1.9.8
on darwin_arm64
```

## LICENSE

[MIT](LICENSE)
