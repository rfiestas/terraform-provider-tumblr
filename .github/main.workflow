workflow "Build and Publish" {
  resolves = ["Lint", "Coverage"]
  on = "push"
}

action "Lint" {
  uses = "actions/action-builder/shell@master"
  runs = "make"
  args = "lint"
}

action "Coverage" {
  uses = "actions/action-builder/shell@master"
  runs = "make"
  args = "coveralls"
  secrets = ["COVERALLS_TOKEN"]
}
