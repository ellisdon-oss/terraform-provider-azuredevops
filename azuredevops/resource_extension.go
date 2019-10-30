package azuredevops

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/extensionmanagement"
	"strings"
	"time"
)

func resourceExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionCreate,
		Update: resourceExtensionUpdate,
		Delete: resourceExtensionDelete,
		Read:   resourceExtensionRead,
		Importer: &schema.ResourceImporter{
			State: resourceExtensionImport,
		},

		Schema: map[string]*schema.Schema{
			"publisher": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "none",
			},
		},
	}
}

func resourceExtensionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	extensionClient, err := extensionmanagement.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	extensionName := d.Get("name").(string)
	publisherName := d.Get("publisher").(string)

	extension, err := extensionClient.GetInstalledExtensionByName(config.Context, extensionmanagement.GetInstalledExtensionByNameArgs{
		ExtensionName: &extensionName,
		PublisherName: &publisherName,
	})

	if err != nil {
		if strings.Contains(err.Error(), "TF1590003") {
			d.SetId("")
			return nil
		}
		return err
	}

	if *extension.ExtensionId == "" {
		d.SetId("")
		return nil
	}

	d.Set("publisher", *extension.PublisherId)
	d.Set("name", *extension.ExtensionId)

	d.Set("state", *extension.InstallState.Flags)
	d.SetId(fmt.Sprintf("%s-%s-%s", config.Organization, *extension.PublisherId, *extension.ExtensionId))

	return nil
}

func resourceExtensionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	extensionName := d.Get("name").(string)
	publisherName := d.Get("publisher").(string)
	state := extensionmanagement.ExtensionStateFlags(d.Get("state").(string))
	version := d.Get("version").(string)

	newExtension := extensionmanagement.InstallExtensionByNameArgs{
		PublisherName: &publisherName,
		ExtensionName: &extensionName,
		Version:       &version,
	}

	extensionClient, err := extensionmanagement.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	extension, err := extensionClient.InstallExtensionByName(config.Context, newExtension)

	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)

	readExtension, err := extensionClient.GetInstalledExtensionByName(config.Context, extensionmanagement.GetInstalledExtensionByNameArgs{
		ExtensionName: &extensionName,
		PublisherName: &publisherName,
	})

	if err != nil {
		return err
	}

	readExtension.InstallState.Flags = &state

	_, err = extensionClient.UpdateInstalledExtension(config.Context, extensionmanagement.UpdateInstalledExtensionArgs{
		Extension: readExtension,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s-%s-%s", config.Organization, *extension.PublisherId, *extension.ExtensionId))
	return resourceExtensionRead(d, meta)
}

func resourceExtensionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	extensionName := d.Get("name").(string)
	publisherName := d.Get("publisher").(string)
	state := extensionmanagement.ExtensionStateFlags(d.Get("state").(string))

	extensionClient, err := extensionmanagement.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	extension, err := extensionClient.GetInstalledExtensionByName(config.Context, extensionmanagement.GetInstalledExtensionByNameArgs{
		ExtensionName: &extensionName,
		PublisherName: &publisherName,
	})

	if err != nil {
		return err
	}

	extension.InstallState.Flags = &state

	_, err = extensionClient.UpdateInstalledExtension(config.Context, extensionmanagement.UpdateInstalledExtensionArgs{
		Extension: extension,
	})

	if err != nil {
		return err
	}

	return resourceExtensionRead(d, meta)
}

func resourceExtensionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	extensionName := d.Get("name").(string)
	publisherName := d.Get("publisher").(string)

	extensionClient, err := extensionmanagement.NewClient(config.Context, config.Connection)

	if err != nil {
		return err
	}

	err = extensionClient.UninstallExtensionByName(config.Context, extensionmanagement.UninstallExtensionByNameArgs{
		PublisherName: &publisherName,
		ExtensionName: &extensionName,
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceExtensionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := d.Id()
	if name == "" {
		return nil, fmt.Errorf("extension cannot be empty")
	}

	res := strings.Split(name, "/")
	if len(res) != 2 {
		return nil, fmt.Errorf("the format has to be in <publisher-name>/<extension-name>")
	}

	d.Set("publisher", res[0])
	d.Set("name", res[1])

	return []*schema.ResourceData{d}, nil
}
