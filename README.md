# example-wrapper

This is examples to wrap Terraform command using [aqua](https://aquaproj.github.io/).

There are several ways to implement the wrapper:

1. [Shell script wrapper](https://github.com/suzuki-shunsuke/example-wrapper/tree/main)
1. [Go wrapper](https://github.com/suzuki-shunsuke/example-wrapper/tree/p2)
1. [Install the wrapper using aqua](https://github.com/suzuki-shunsuke/example-wrapper/tree/p3)

## Shell script wrapper

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
