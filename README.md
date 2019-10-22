# terraform-provider-tumblr 
[![Build Status](https://github.com/rfiestas/terraform-provider-tumblr/workflows/Test/badge.svg)](https://github.com/rfiestas/terraform-provider-tumblr/actions)
[![codecov](https://codecov.io/gh/rfiestas/terraform-provider-tumblr/branch/master/graph/badge.svg)](https://codecov.io/gh/rfiestas/terraform-provider-tumblr)
![Scrutinizer code quality (GitHub/Bitbucket)](https://img.shields.io/scrutinizer/quality/g/rfiestas/terraform-provider-tumblr)
[![Go Report Card](https://goreportcard.com/badge/github.com/rfiestas/terraform-provider-tumblr)](https://goreportcard.com/report/github.com/rfiestas/terraform-provider-tumblr)

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/rfiestas/terraform-provider-tumblr)
![GitHub](https://img.shields.io/github/license/rfiestas/terraform-provider-tumblr)

A [Terraform](https://www.terraform.io) Custom Provider for [tumblr](https://www.tumblr.com).

## Description

This is a custom terraform provider for managing common resources within the tumblr site platform, such as Transformations, Orchestrations, Writers etc.

## Supported Resources

Currently, the following tumblr resources are supported (or partially supported) for configuration via `terraform`:

* `tumblr_post_text`
* `tumblr_post_photo`
* `tumblr_post_quote`
* `tumblr_post_link`
* `tumblr_post_chat`
* `tumblr_post_audio`
* `tumblr_post_video`

## Requirements

* [hashicorp/terraform](https://github.com/hashicorp/terraform)

## Usage

### Install

```
go build -o ~/.terraform.d/plugins/terraform-provider-tumblr
```

### Provider Configuration

The provider only requires some configuration settings. Get your secrets from your tumblr account.
Use [tumblr/settings](https://www.tumblr.com/settings/apps) to get the *OAuth Consumer Key* and *OAuth Consumer Secret* 
and then validate on [oauth page](https://api.tumblr.com/console/calls/user/info)

#### `tumblr`

```
provider "tumblr" {
  consumer_key      = "XXXXXXXXXXXXXXXXXXXXXX"
  consumer_secret   = "XXXXXXXXXXXXXXXXXXXXXX"
  user_token        = "XXXXXXXXXXXXXXXXXXXXXX"
  user_token_secret = "XXXXXXXXXXXXXXXXXXXXXX"
}
```

Alternatively you can use environment variables

```
export CONSUMER_KEY="XXXXXXXXXXXXXXXXXXXXXX"
export CONSUMER_SECRET="XXXXXXXXXXXXXXXXXXXXXX"
export USER_TOKEN="XXXXXXXXXXXXXXXXXXXXXX"
export USER_TOKEN_SECRET="XXXXXXXXXXXXXXXXXXXXXX"
```

### Resource Configuration

For documentation on each supported resource, refer to the [blog](https://terraform-provider-for.tumblr.com/).

## Contributing

Bug reports, suggestions, code additions/changes etc. are very welcome! When making code changes, please branch off of `master` and then raise a pull request so it can be reviewed and merged.

### Running Acceptance Tests

In order to test the provider, you can run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note* Make sure `CONSUMER_KEY`, `CONSUMER_SECRET`, `USER_TOKEN` and  `USER_TOKEN_SECRET` variables are set.

Acceptance tests create real resources, set your testing blog setting the variable `TestBlog` on `tumblr/schema_tumblr_post_test.go` file.

```sh
$ make testacc
```
