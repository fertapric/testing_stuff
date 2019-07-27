workflow "Cirrus" {
  on = "check_suite"
  resolves = ["Email"]
}

action "Email" {
  uses = "docker://fertapric/elixir-ci-email:latest"
  secrets = ["GITHUB_TOKEN", "MAIL_FROM", "MAIL_HOST", "MAIL_USERNAME", "MAIL_PASSWORD"]
  env = {
    APP_NAME = "Cirrus CI"
    ELIXIR_CI_MAIL = "fernando.tapia@jobandtalent.com"
  }
}
