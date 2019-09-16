workflow "Build and Publish" {
  on = "push"
  resolves = ["Coverage"]
}

action "Coverage" {
  uses = "actions/action-builder/shell@master"
  runs = "make"
  args = "cover"
}
