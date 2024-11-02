# example-wrapper

This is an example to wrap Terraform command using [aqua](https://aquaproj.github.io/).

First, [please install aqua](https://aquaproj.github.io/docs/install).

And run `aqua i -l`.

```sh
aqua i -l
```

1. Add `bin` to `$PATH:

```sh
export PATH=$PWD/bin:$PATH
```

If you use direnv,

```sh
direnv allow
```

2. Run `terraform`:

```console
$ terraform version
[INFO] Terraform wrapper
Terraform v1.9.8
on darwin_arm64
```

## LICENSE

[MIT](LICENSE)
