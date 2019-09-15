package internal

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceEndpoint() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        resourceEndpointCreate,
		Read:          resourceEndpointRead,
		Update:        resourceEndpointUpdate,
		Delete:        resourceEndpointDelete,
		Exists:        resourceEndpointExists,
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Description: "Host and optional port",
				Type:        schema.TypeString,
				Required:    true,
				Optional:    false,
			},
			"protocol": {
				Description: "The protocol",
				Type:        schema.TypeString,
				Required:    true,
				Optional:    false,
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceEndpointState,
		},
	}
}

func resourceEndpointCreate(d *schema.ResourceData, m interface{}) error {
	endpoint := d.Get("endpoint").(string)
	if endpoint == "" {
		return fmt.Errorf("could not create endpoint: invalid endpoint")
	}
	protocol := d.Get("protocol").(string)
	if protocol == "" {
		protocol = "https"
	}
	cli := m.(config).client
	resp, err := cli.createEndpoint(endpoint, protocol)
	if err != nil {
		return err
	}
	d.SetId(resp.ID)
	if err = d.Set("endpoint", resp.Endpoint); err != nil {
		return err
	}
	return nil
}

func resourceEndpointRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	cli := m.(config).client
	endpoints, err := cli.listEndpoints()
	if err != nil {
		return err
	}
	for _, endpoint := range endpoints.Endpoints {
		if id == endpoint.ID || id == endpoint.Endpoint {
			d.SetId(endpoint.ID)
			if err = d.Set("endpoint", endpoint.Endpoint); err != nil {
				return err
			}
			if err = d.Set("protocol", endpoint.Protocol); err != nil {
				return err
			}
			return nil
		}
	}
	d.SetId("")
	return nil
}

func resourceEndpointUpdate(d *schema.ResourceData, m interface{}) error {
	newEndpoint := ""
	if d.HasChange("endpoint") {
		newEndpoint = d.Get("endpoint").(string)
	}

	id := d.Id()
	cli := m.(config).client
	endpoints, err := cli.listEndpoints()
	if err != nil {
		return err
	}

	for _, endpoint := range endpoints.Endpoints {
		if id == endpoint.ID {
			oldEndpoint := endpoint.Endpoint
			if newEndpoint != "" && newEndpoint != oldEndpoint {
				createEndpointResp, err := cli.createEndpoint(newEndpoint, endpoint.Protocol)
				if err != nil {
					return err
				}
				d.SetId(createEndpointResp.ID)
				if err = d.Set("endpoint", createEndpointResp.Endpoint); err != nil {
					return err
				}
				return nil
			}
			return nil
		}
	}
	d.SetId("")
	return nil
}

func resourceEndpointDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()
	cli := m.(config).client

	return cli.deleteEndpoint(id)
}

func resourceEndpointExists(d *schema.ResourceData, m interface{}) (bool, error) {
	id := d.Id()
	cli := m.(config).client
	endpoints, err := cli.listEndpoints()
	if err != nil {
		return false, err
	}
	for _, endpoint := range endpoints.Endpoints {
		if id == endpoint.ID {
			return true, nil
		}
	}
	return false, nil
}

func resourceEndpointState(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	nameOrID := d.Id()
	if err := resourceEndpointRead(d, m); err != nil {
		return nil, err
	}
	if d.Id() == "" {
		return nil, fmt.Errorf("team with ID or name %s not found", nameOrID)
	}
	return []*schema.ResourceData{d}, nil

}
