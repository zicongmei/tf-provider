terraform {
  required_providers {
    zicong= {
      source  = "registry.terraform.io/zicong/zicong"
    }
  }
}

provider "zicong" {
  foo = "zicong_value"
}