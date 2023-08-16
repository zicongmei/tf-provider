terraform {
  required_providers {
    zicong = {
      version = "0.0.1"
      source = "zicong/zicong"
    }
  }
}

provider "zicong" {
}

resource "resource1" "object1" {
  provider = zicong

  resource_key="key1"
  resource_value="value2"
}

resource "resource2" "object2" {
  provider = zicong

  resource_key="key2"
  resource_value="value22"
}