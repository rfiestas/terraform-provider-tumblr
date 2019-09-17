# terraform-provider-tumblr 
[![Build Status](https://travis-ci.com/rfiestas/terraform-provider-tumblr.svg?branch=master)](https://travis-ci.com/rfiestas/terraform-provider-tumblr)
[![Coverage Status](https://coveralls.io/repos/github/rfiestas/terraform-provider-tumblr/badge.svg?branch=master)](https://coveralls.io/github/rfiestas/terraform-provider-tumblr?branch=master)

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

PENDING

## License
`terraform-provider-tumblr` is provided *"as-is"* under the [Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0).
