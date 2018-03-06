package main

import (
    "github.com/hashicorp/terraform/helper/schema"
)

func resourceRollbar() *schema.Resource {
    return &schema.Resource{
        Create: resourceRollbarCreate,
        Read:   resourceRollbarRead,
        Update: resourceRollbarUpdate,
        Delete: resourceRollbarDelete,

        Schema: map[string]*schema.Schema{
            "access_token": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceRollbarCreate(d *schema.ResourceData, m interface{}) error {
    access_token := d.Get("access_token").(string)
    d.SetId(access_token)
    return nil
}

func resourceRollbarRead(d *schema.ResourceData, m interface{}) error {
/*
  client := m.(*MyClient)

  // Attempt to read from an upstream API
  obj, ok := client.Get(d.Id())

  // If the resource does not exist, inform Terraform. We want to immediately
  // return here to prevent further processing.
  if !ok {
  d.SetId("")
  return nil
  }

  d.Set("access_token", obj.access_token)
*/
  return nil
}

func resourceRollbarUpdate(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceRollbarDelete(d *schema.ResourceData, m interface{}) error {
  // d.SetId("") is automatically called assuming delete returns no errors, but
  // it is added here for explicitness.
  d.SetId("")
  return nil
}
