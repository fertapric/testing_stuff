workflow "Cirrus" {
  on = "check_suite"
  resolves = ["docker://cirrusactions/email:latest"]
}

action "docker://cirrusactions/email:latest" {
  uses = "docker://cirrusactions/email:latest"
  secrets = ["GITHUB_TOKEN", "MAIL_FROM", "MAIL_HOST", "MAIL_USERNAME", "MAIL_PASSWORD"]
  env = {
    APP_NAME = "Cirrus CI"
  }
}
