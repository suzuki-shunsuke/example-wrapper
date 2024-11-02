# example-wrapper

This is an example to wrap Terraform command using [aqua](https://aquaproj.github.io/).

First, [please install aqua](https://aquaproj.github.io/docs/install).
This example builds the wrapper in Go, so Go is necessary.
In this approach, the wrapper is installed by aqua, but the registry isn't standard reigstry, so you need to configure and allow the policy.

```sh
aqua policy allow
```

And run `aqua i -l`.

```sh
aqua i -l
```

Then run `terraform`.

```console
$ terraform version
2024/11/02 10:30:14 [INFO] Terraform wrapper
Terraform v1.9.8
on darwin_arm64
```

## LICENSE

[MIT](LICENSE)
