terraform {
  required_providers {
    jackdemo = {
      source  = "idcos/jackdemo"
    }
  }
}

provider "jackdemo" {
}

resource "jackdemo_jack" "test" {
  instance_name  = "helloworld2"
  disk_size = 102
  tags = "jack2"
}

data "jackdemo_ecs" "test" {
  name = "ecs"
}
