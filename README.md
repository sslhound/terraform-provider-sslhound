# terraform-provider-sslhound

A [terraform](https://www.terraform.io/) for [www.sslhound.com](https://www.sslhound.com/).

# Usage

Configure the provider:

```hcl-terraform
provider "sslhound" {
    token = "YOUR.TOKE_GOES_HERE"
}
```

Init the plugin:

    $ terraform init

The "**sslhound_endpoint**" resource is used to create and manage checks:

```hcl-terraform
resource "sslhound_endpoint" "www_sshould_com_443" {
  endpoint = "www.sslhound.com:443"
  protocol = "https"
}
```

## Import

Resources can also be imported:

    $ terraform import sslhound_endpoint.www_sshould_com_443 "www.sslhound.com:443"

# Release

* [ ] Did schema change and schema versions increment?
* [ ] Did Terraform required version change?

    $ GOOS=linux GOARCH=amd64 go build -o terraform-provider-sslhound_v1.0.0_linux_amd64 main.go
    $ GOOS=darwin GOARCH=amd64 go build -o terraform-provider-sslhound_v1.0.0_darwin_amd64 main.go
    $ GOOS=windows GOARCH=amd64 go build -o terraform-provider-sslhound_v1.0.0_windows_amd64.exe main.go
