terraform {
  required_providers {
    zicong = {
      version = "0.0.1"
      source = "zicong/zicong"
    }
  }
}

provider "zicong" {
  base_url = "zicong_value"
}

resource "item1" "default" {
  provider = zicong
  foo {
    bar {
      number = 1
    }
  }

}