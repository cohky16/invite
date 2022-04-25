terraform {
  required_providers {
    heroku = {
      source = "heroku/heroku"
    }
  }
}

provider "heroku" {}

resource "heroku_app" "invite-ydkk" {
  acm = false
  buildpacks = [
    "heroku/go"
  ]
  config_vars           = {}
  internal_routing      = false
  name                  = "invite-ydkk"
  region                = "us"
  sensitive_config_vars = {}
  space                 = null
  stack                 = "heroku-20"
}

resource "heroku_formation" "worker" {
  app_id   = heroku_app.invite-ydkk.id
  quantity = 1
  size     = "Free"
  type     = "worker"
}