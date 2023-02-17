package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/rewrite"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func resourceCitrixAdcRewriteglobal_rewritepolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createRewriteglobal_rewritepolicy_bindingFunc,
		Read:          readRewriteglobal_rewritepolicy_bindingFunc,
		Delete:        deleteRewriteglobal_rewritepolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"globalbindtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createRewriteglobal_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createRewriteglobal_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	priority := strconv.Itoa(d.Get("priority").(int))
	type_bindpoint := d.Get("type").(string)
	bindingId := fmt.Sprintf("%s,%s,%s", policyname, priority, type_bindpoint)
	rewriteglobal_rewritepolicy_binding := rewrite.Rewriteglobalrewritepolicybinding{
		Globalbindtype:         d.Get("globalbindtype").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
		Type:                   d.Get("type").(string),
	}

	_, err := client.AddResource(service.Rewriteglobal_rewritepolicy_binding.Type(), bindingId, &rewriteglobal_rewritepolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readRewriteglobal_rewritepolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this rewriteglobal_rewritepolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readRewriteglobal_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readRewriteglobal_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	policyname := idSlice[0]
	priority := idSlice[1]
	type_bindpoint := idSlice[2]

	log.Printf("[DEBUG] citrixadc-provider: Reading rewriteglobal_rewritepolicy_binding state %s", bindingId)

	argsMap := make(map[string]string)
	argsMap["type"] = type_bindpoint

	findParams := service.FindParams{
		ResourceType: "rewriteglobal_rewritepolicy_binding",
		ArgsMap:      argsMap,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)

	// Unexpected error
	if err != nil {
		log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		return err
	}

	// Resource is missing
	if len(dataArr) == 0 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		log.Printf("[WARN] citrixadc-provider: Clearing rewriteglobal_rewritepolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	foundIndex := -1
	for i, rewriteglobal_rewritepolicy_binding := range dataArr {
		if rewriteglobal_rewritepolicy_binding["policyname"] != policyname {
			continue
		} else if rewriteglobal_rewritepolicy_binding["priority"] != priority {
			continue
		} else if rewriteglobal_rewritepolicy_binding["type"] != type_bindpoint {
			continue
		}
		foundIndex = i
		break
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams rewriteglobal_rewritepolicy_binding not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing rewriteglobal_rewritepolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("globalbindtype", data["globalbindtype"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])
	d.Set("type", data["type"])

	return nil

}

func deleteRewriteglobal_rewritepolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteRewriteglobal_rewritepolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 3)

	policyname := idSlice[0]
	priority := idSlice[1]
	type_bindpoint := idSlice[2]

	argsMap := make(map[string]string)
	argsMap["policyname"] = url.QueryEscape(policyname)
	argsMap["priority"] = url.QueryEscape(priority)
	argsMap["type"] = url.QueryEscape(type_bindpoint)

	err := client.DeleteResourceWithArgsMap(service.Rewriteglobal_rewritepolicy_binding.Type(), "", argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
