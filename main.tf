terraform {
  required_providers {
    heroku = {
      source = "heroku/heroku"
    }
  }

  cloud {
    organization = "cohky"

    workspaces {
      name = "invite"
    }
  }
}


provider "heroku" {}

variable "APP_ENV" {}
variable "DISCORD_TOKEN" {}

resource "heroku_app" "invite-ydkk" {
  acm = false
  buildpacks = [
    "heroku/go"
  ]
  config_vars = {
    APP_ENV : var.APP_ENV
    DISCORD_TOKEN : var.DISCORD_TOKEN
  }
  internal_routing      = false
  name                  = "invite-ydkk"
  region                = "us"
  sensitive_config_vars = {}
  space                 = null
  stack                 = "heroku-20"
}

resource "heroku_build" "invite-ydkk" {
  app_id = heroku_app.invite-ydkk.id
  source {
    path = "./cmd"
  }
}

resource "heroku_formation" "worker" {
  app_id     = heroku_app.invite-ydkk.id
  quantity   = 1
  size       = "Free"
  type       = "worker"
  depends_on = [heroku_build.invite-ydkk]
}