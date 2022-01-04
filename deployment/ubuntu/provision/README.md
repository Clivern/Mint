### Provision the Server

1. Create `terraform.tfvars` and define the required vars.

```bash
# Use any region from https://docs.digitalocean.com/products/platform/availability-matrix/
region = "nyc3"
# You can add and get the ssh name from digitalocean dashboard or doctl cli
ssh_key = "clivern"
name = "mint.sh"
droplet_size = "s-1vcpu-2gb"
image = "ubuntu-20-04-x64"
```

2. Create a digitalocean token from digitalocean dashboard and then run `terraform init`.

3. Run `terraform plan` to check the changes.

4. Run `terraform apply` to perform the action.
