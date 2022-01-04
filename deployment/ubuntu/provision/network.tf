# Copyright 2022 Clivern. All rights reserved.
# Use of this source code is governed by the MIT
# license that can be found in the LICENSE file.

resource "digitalocean_vpc" "mint_infra" {

    name = "mint-infra-vpc"

    region = var.region
}
