# Copyright 2022 Clivern. All rights reserved.
# Use of this source code is governed by the MIT
# license that can be found in the LICENSE file.

variable do_token {
    type = string
    description = "Digitalocean API Token"
}

variable "ssh_key" {
    type = string
    default = "clivern"
    description = "Digitalocean Account SSH Key Name"
}

variable "region" {
    type    = string
    default = "nyc3"
    description = "Digitalocean Region"
}

variable "name" {
    type = string
    default = "mint.sh"
    description = "Server Name"
}

variable "droplet_size" {
    type = string
    default = "s-1vcpu-2gb"
    description = "Server Size"
}

variable "image" {
    type = string
    default = "ubuntu-20-04-x64"
    description = "Server Image"
}
