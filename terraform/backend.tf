terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "lemaire"

    workspaces {
      name = "aloesia"
    }
  }
}
