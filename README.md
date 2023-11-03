# dep-doctor

`dep-doctor` is a tool to diagnose whether your software dependency libraries are maintained.

Today, most software relies heavily on external libraries. Vulnerabilities in those libraries can be detected by vulnerability scanners ([dependabot](https://docs.github.com/en/code-security/dependabot), [trivy](https://aquasecurity.github.io/trivy), [Grype](https://github.com/anchore/grype), etc) if they are publicly available.

However, some libraries have archived their source code repositories or have had their development stopped, although not explicitly. `dep-doctor` will notify you of those libraries in the dependencies file.

![overview](doc/images/dep-doctor_overview.png "dep-doctor overview")

## Support dependencies files

| language | package manager | dependencies file (e.g.) | status |
| -------- | ------------- | -- | :----: |
| Go | golang | go.mod | :heavy_check_mark: |
| JavaScript | npm | package-lock.json | :heavy_check_mark: |
| JavaScript | yarn | yarn.lock | :heavy_check_mark: |
| PHP | composer | composer.lock | :heavy_check_mark: |
| Python | pip | requirements.txt | :heavy_check_mark: |
| Python | pipenv | Pipfile.lock | :heavy_check_mark: |
| Python | poetry | poetry.lock | (later) |
| Ruby | bundler | Gemfile.lock | :heavy_check_mark: |
| Rust | cargo | Cargo.lock | :heavy_check_mark: |
| Swift | cocoapods | Podfile.lock | :heavy_check_mark: |

## Support repository hosting services

Only GitHub.com

## Install

### Homebrew (macOS and Linux)

```console
$ brew tap kyoshidajp/dep-doctor
$ brew install kyoshidajp/dep-doctor/dep-doctor
```

### Binary packages

[Releases](https://github.com/kyoshidajp/dep-doctor/releases)

## How to use

`GITHUB_TOKEN` must be set as an environment variable before execution.

```console
Usage:
  dep-doctor diagnose [flags]

Flags:
  -f, --file string      dependencies file path
  -h, --help             help for diagnose
  -i, --ignores string   ignore dependencies (separated by a space)
  -p, --package string   package manager
      --strict           exit with non-zero if warnings exist
  -y, --year int         max years of inactivity (default 5)
```

For example:

```console
$ dep-doctor diagnose -p bundler -f /path/to/Gemfile.lock
concurrent-ruby
dotenv
faker
i18n
method_source
paperclip
......
[error] paperclip (archived): https://github.com/thoughtbot/paperclip
Diagnosis completed! 6 dependencies.
1 error, 0 warn (0 unknown), 0 info (0 ignored)
```

## Report level

| level | e.g. |
| :---: | :---------- |
| *error* | Source code repository is already archived. |
| *warn* | Source code repository is not active or unknown. |
| *info* | Other reasons. (specified to be ignored) | |

## How it works

![how_works](doc/images/how_works.png "dep-doctor how works")

## Author
Katsuhiko YOSHIDA
