workflow "Build and Publish" {
  resolves = [
    "Coverage",
    "Lint",
  ]
  on = "push"
}

action "Coverage" {
  uses = "actions/action-builder/shell@master"
  runs = "make"
  args = "coveralls"
  secrets = ["COVERALLS_TOKEN"]
}

action "Lint" {
  uses = "./.github/actions/terraform-provider-tumblr-action"
  runs = "lint"
}
