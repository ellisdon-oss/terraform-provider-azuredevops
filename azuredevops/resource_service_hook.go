package azuredevops

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/servicehooks"
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
			"custom_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

	hookClient := servicehooks.NewClient(config.Context, config.Connection)

	subscriptionId, _ := uuid.Parse(d.Id())

	serviceHook, err := hookClient.GetSubscription(config.Context, servicehooks.GetSubscriptionArgs{
		SubscriptionId: &subscriptionId,
	})

	if err != nil {
		return err
	}

	if serviceHook.Id.String() == "" {
		d.SetId("")
		return nil
	}

	consumer := []map[string]interface{}{
		{
			"action_id": *serviceHook.ConsumerActionId,
			"id":        *serviceHook.ConsumerId,
			"inputs":    *serviceHook.ConsumerInputs,
		},
	}

	d.Set("consumer", consumer)

	inputs := *serviceHook.PublisherInputs
	delete(inputs, "tfsSubscriptionId")
	publisher := []map[string]interface{}{
		{
			"id":     *serviceHook.PublisherId,
			"inputs": inputs,
		},
	}

	d.Set("publisher", publisher)

	d.Set("event_type", *serviceHook.EventType)
	d.SetId(serviceHook.Id.String())

	return nil
}

func resourceServiceHookCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	actionID := d.Get("consumer.0.action_id").(string)
	consumerID := d.Get("consumer.0.id").(string)
	consumerInputs := convertInterfaceToStringMap(d.Get("consumer.0.inputs").(map[string]interface{}))
	publisherID := d.Get("publisher.0.id").(string)
	publisherInputs := convertInterfaceToStringMap(d.Get("publisher.0.inputs").(map[string]interface{}))
	eventType := d.Get("event_type").(string)

	newServiceHook := servicehooks.Subscription{
		ConsumerActionId: &actionID,
		ConsumerId:       &consumerID,
		ConsumerInputs:   &consumerInputs,
		PublisherId:      &publisherID,
		PublisherInputs:  &publisherInputs,
		EventType:        &eventType,
	}

	connection := config.Connection
	if v := d.Get("custom_path"); v.(string) != "" {
		connection.BaseUrl = v.(string)
	}

	hookClient := servicehooks.NewClient(config.Context, connection)

	serviceHook, err := hookClient.CreateSubscription(config.Context, servicehooks.CreateSubscriptionArgs{
		Subscription: &newServiceHook,
	})

	if err != nil {
		return err
	}

	d.SetId(serviceHook.Id.String())
	return resourceServiceHookRead(d, meta)
}

func resourceServiceHookUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	actionID := d.Get("consumer.0.action_id").(string)
	consumerID := d.Get("consumer.0.id").(string)
	consumerInputs := convertInterfaceToStringMap(d.Get("consumer.0.inputs").(map[string]interface{}))
	publisherID := d.Get("publisher.0.id").(string)
	publisherInputs := convertInterfaceToStringMap(d.Get("publisher.0.inputs").(map[string]interface{}))
	eventType := d.Get("event_type").(string)

	newServiceHook := servicehooks.Subscription{
		ConsumerActionId: &actionID,
		ConsumerId:       &consumerID,
		ConsumerInputs:   &consumerInputs,
		PublisherId:      &publisherID,
		PublisherInputs:  &publisherInputs,
		EventType:        &eventType,
	}

	connection := config.Connection
	if v := d.Get("custom_path"); v != nil {
		connection.BaseUrl = v.(string)
	}

	hookClient := servicehooks.NewClient(config.Context, connection)

	serviceHookID, _ := uuid.Parse(d.Id())
	serviceHook, err := hookClient.ReplaceSubscription(config.Context, servicehooks.ReplaceSubscriptionArgs{
		Subscription:   &newServiceHook,
		SubscriptionId: &serviceHookID,
	})

	if err != nil {
		return err
	}

	d.SetId(serviceHook.Id.String())

	return resourceServiceHookRead(d, meta)
}

func resourceServiceHookDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	hookClient := servicehooks.NewClient(config.Context, config.Connection)

	serviceHookID, _ := uuid.Parse(d.Id())

	err := hookClient.DeleteSubscription(config.Context, servicehooks.DeleteSubscriptionArgs{
		SubscriptionId: &serviceHookID,
	})

	if err != nil {
		return err
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
