# Copyright 2022 Clivern. All rights reserved.
# Use of this source code is governed by the MIT
# license that can be found in the LICENSE file.

resource "digitalocean_droplet" "mint" {

  image = var.image

  name = "${var.name}"

  region = var.region

  ipv6 = true

  monitoring = true

  size = var.droplet_size

  ssh_keys = [data.digitalocean_ssh_key.main.id]

  vpc_uuid = digitalocean_vpc.mint_infra.id
}

resource "digitalocean_floating_ip" "mint" {

  droplet_id = digitalocean_droplet.mint.id

  region     = digitalocean_droplet.mint.region
}
