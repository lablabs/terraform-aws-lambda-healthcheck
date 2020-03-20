provider "aws" {
  alias   = "alarm"
  version = "~> 2.0"
  region  = var.region
}
