package internal

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"net/url"
)

type config struct {
	client houndClient
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		TerraformVersion: "0.12.8",
		Schema: map[string]*schema.Schema{
			"host": {
				Description: "SSL Hound host to make API calls to",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"token": {
				Description: "SSL Hound token to use",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sslhound_endpoint": resourceEndpoint(),
		},
		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	host := d.Get("host").(string)
	token := d.Get("token").(string)
	if token == "" {
		return nil, fmt.Errorf("required configuration parameter missing: token")
	}
	if host == "" {
		host = "www.sslhound.com"
	}
	u, err := url.Parse(fmt.Sprintf("https://%s/api/endpoints", host))
	if err != nil {
		return nil, fmt.Errorf("unable to configure sslhound provider: %s", err)
	}
	return config{
		client: houndClient{
			token:   token,
			baseURL: u.String(),
		},
	}, nil
}
