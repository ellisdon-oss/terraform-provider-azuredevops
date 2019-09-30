package azuredevops

import (
	"fmt"
	"github.com/ellisdon/azuredevops-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceServiceHook() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceHookCreate,
		Update: resourceServiceHookUpdate,
		Delete: resourceServiceHookDelete,
		Read:   resourceServiceHookRead,
		Importer: &schema.ResourceImporter{
			State: resourceServiceHookImport,
		},

		Schema: map[string]*schema.Schema{
			"publisher": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inputs": &schema.Schema{
							Type:     schema.TypeMap,
							Required: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"consumer": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inputs": &schema.Schema{
							Type:     schema.TypeMap,
							Required: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"action_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceHookRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serviceHook, _, err := config.SubscriptionClient.SubscriptionsApi.GetSubscription(config.Context, config.Organization, d.Id(), config.ApiVersion)

	if err != nil {
		return err
	}

	if serviceHook.Id == "" {
		d.SetId("")
		return nil
	}

	consumer := []map[string]interface{}{
		{
			"action_id": serviceHook.ConsumerActionId,
			"id":        serviceHook.ConsumerId,
			"inputs":    serviceHook.ConsumerInputs,
		},
	}

	d.Set("consumer", consumer)

	inputs := serviceHook.PublisherInputs
	delete(inputs, "tfsSubscriptionId")
	publisher := []map[string]interface{}{
		{
			"id":     serviceHook.PublisherId,
			"inputs": inputs,
		},
	}

	d.Set("publisher", publisher)

	d.Set("event_type", serviceHook.EventType)
	d.SetId(serviceHook.Id)

	return nil
}

func resourceServiceHookCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newServiceHook := azuredevops.Subscription{
		ConsumerActionId: d.Get("consumer.0.action_id").(string),
		ConsumerId:       d.Get("consumer.0.id").(string),
		ConsumerInputs:   d.Get("consumer.0.inputs").(map[string]interface{}),
		PublisherId:      d.Get("publisher.0.id").(string),
		PublisherInputs:  d.Get("publisher.0.inputs").(map[string]interface{}),
		EventType:        d.Get("event_type").(string),
	}

	switch newServiceHook.EventType {
	case "ms.vss-release.deployment-completed-event":
		config.SubscriptionClient.ChangeBasePath("https://vsrm.dev.azure.com")
	default:
		config.SubscriptionClient.ChangeBasePath("https://dev.azure.com")
	}

	serviceHook, _, err := config.SubscriptionClient.SubscriptionsApi.CreateSubscription(config.Context, config.Organization, config.ApiVersion, newServiceHook)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(serviceHook.Id)
	config.SubscriptionClient.ChangeBasePath("https://dev.azure.com")
	return resourceServiceHookRead(d, meta)
}

func resourceServiceHookUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	newServiceHook := azuredevops.Subscription{
		ConsumerActionId: d.Get("consumer.0.action_id").(string),
		ConsumerId:       d.Get("consumer.0.id").(string),
		ConsumerInputs:   d.Get("consumer.0.inputs").(map[string]interface{}),
		PublisherId:      d.Get("publisher.0.id").(string),
		PublisherInputs:  d.Get("publisher.0.inputs").(map[string]interface{}),
		EventType:        d.Get("event_type").(string),
	}

	switch newServiceHook.EventType {
	case "ms.vss-release.deployment-completed-event":
		config.SubscriptionClient.ChangeBasePath("https://vsrm.dev.azure.com")
	default:
		config.SubscriptionClient.ChangeBasePath("https://dev.azure.com")
	}

	serviceHook, _, err := config.SubscriptionClient.SubscriptionsApi.ReplaceSubscription(config.Context, config.Organization, d.Id(), config.ApiVersion, newServiceHook)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	d.SetId(serviceHook.Id)

	config.SubscriptionClient.ChangeBasePath("https://dev.azure.com")
	return resourceServiceHookRead(d, meta)
}

func resourceServiceHookDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := config.SubscriptionClient.SubscriptionsApi.DeleteSubscription(config.Context, config.Organization, d.Id(), config.ApiVersion)

	if err != nil {
		return errors.New(string(err.(azuredevops.GenericOpenAPIError).Body()))
	}

	return nil
}

func resourceServiceHookImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	name := d.Id()
	if name == "" {
		return nil, fmt.Errorf("serice hook id cannot be empty")
	}

	d.SetId(name)

	return []*schema.ResourceData{d}, nil
}
