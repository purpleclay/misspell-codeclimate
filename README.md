# misspell-codeclimate

Turn that [misspell](https://github.com/client9/misspell) report into a GitLab compatible Code Climate [report](https://docs.gitlab.com/ee/ci/testing/code_quality.html#implementing-a-custom-tool) for annotated merge requests.

[![Build status](https://img.shields.io/github/workflow/status/purpleclay/misspell-codeclimate/ci?style=flat-square&logo=go)](https://github.com/purpleclay/misspell-codeclimate/actions?workflow=ci)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/purpleclay/misspell-codeclimate?style=flat-square)](https://goreportcard.com/report/github.com/purpleclay/misspell-codeclimate)
[![Go Version](https://img.shields.io/github/go-mod/go-version/purpleclay/misspell-codeclimate.svg?style=flat-square)](go.mod)
[![codecov](https://codecov.io/gh/purpleclay/misspell-codeclimate/branch/main/graph/badge.svg)](https://codecov.io/gh/purpleclay/misspell-codeclimate)

## Quick Start

Generate a misspell report by scanning your repository:

```sh
misspell -locale UK docs > misspell-report.txt
```

Then generate a Code Climate report:

```sh
misspell-codeclimate --file misspell-report.txt > codeclimate-report.json
```

## Install

### Homebrew

To use [Homebrew](https://brew.sh/):

```sh
brew install purpleclay/tap/misspell-codeclimate
```

### Scoop

To use [Scoop](https://scoop.sh/):

```sh
scoop bucket add purpleclay https://github.com/purpleclay/scoop-bucket.git
scoop install misspell-codeclimate
```

### Apt

To install using the [apt](https://ubuntu.com/server/docs/package-management) package manager:

```sh
echo 'deb [trusted=yes] https://fury.purpleclay.dev/apt/ /' | sudo tee /etc/apt/sources.list.d/purpleclay.list
sudo apt update
sudo apt install -y misspell-codeclimate
```

You may need to install the `ca-certificates` package if you encounter [trust issues](https://gemfury.com/help/could-not-verify-ssl-certificate/) with regards to the gemfury certificate:

```sh
sudo apt update && sudo apt install -y ca-certificates
```

### Yum

To install using the yum package manager:

```sh
echo '[purpleclay]
name=purpleclay
baseurl=https://fury.purpleclay.dev/yum/
enabled=1
gpgcheck=0' | sudo tee /etc/yum.repos.d/purpleclay.repo
sudo yum install -y misspell-codeclimate
```

### Aur

To install from the [aur](https://archlinux.org/) using [yay](https://github.com/Jguer/yay):

```sh
yay -S misspell-codeclimate-bin
```

### Bash Script

To install the latest version using a bash script:

```sh
curl https://raw.githubusercontent.com/purpleclay/misspell-codeclimate/main/scripts/install | bash
```

A specific version can be downloaded by using the `-v` flag. By default the script uses `sudo`, which can be turned off by using the `--no-sudo` flag.

```sh
curl https://raw.githubusercontent.com/purpleclay/misspell-codeclimate/main/scripts/install | bash -s -- -v v0.1.0 --no-sudo
```
