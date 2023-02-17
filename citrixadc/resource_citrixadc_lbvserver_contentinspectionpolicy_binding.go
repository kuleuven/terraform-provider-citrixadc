package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
	"net/url"
	"strings"
)

func resourceCitrixAdcLbvserver_contentinspectionpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbvserver_contentinspectionpolicy_bindingFunc,
		Read:          readLbvserver_contentinspectionpolicy_bindingFunc,
		Delete:        deleteLbvserver_contentinspectionpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bindpoint": &schema.Schema{
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createLbvserver_contentinspectionpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createLbvserver_contentinspectionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	lbvserverName := d.Get("name").(string)
	policyName := d.Get("policyname").(string)
	bindingId := fmt.Sprintf("%s,%s", lbvserverName, policyName)
	lbvserver_contentinspectionpolicy_binding := lb.Lbvservercontentinspectionpolicybinding{
		Bindpoint:              d.Get("bindpoint").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Invoke:                 d.Get("invoke").(bool),
		Labelname:              d.Get("labelname").(string),
		Labeltype:              d.Get("labeltype").(string),
		Name:                   lbvserverName,
		Policyname:             policyName,
		Priority:               d.Get("priority").(int),
	}

	_, err := client.AddResource("lbvserver_contentinspectionpolicy_binding", lbvserverName, &lbvserver_contentinspectionpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(bindingId)

	err = readLbvserver_contentinspectionpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this lbvserver_contentinspectionpolicy_binding but we can't read it ?? %s", bindingId)
		return nil
	}
	return nil
}

func readLbvserver_contentinspectionpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readLbvserver_contentinspectionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	lbvserverName := idSlice[0]
	policyName := idSlice[1]

	log.Printf("[DEBUG] citrixadc-provider: Reading lbvserver_contentinspectionpolicy_binding state %s", bindingId)

	findParams := service.FindParams{
		ResourceType:             "lbvserver_contentinspectionpolicy_binding",
		ResourceName:             lbvserverName,
		ResourceMissingErrorCode: 258,
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
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_contentinspectionpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondIdComponent not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing lbvserver_contentinspectionpolicy_binding state %s", bindingId)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("bindpoint", data["bindpoint"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("invoke", data["invoke"])
	d.Set("labelname", data["labelname"])
	d.Set("labeltype", data["labeltype"])
	d.Set("name", data["name"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteLbvserver_contentinspectionpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteLbvserver_contentinspectionpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	bindingId := d.Id()
	idSlice := strings.SplitN(bindingId, ",", 2)

	lbvserverName := idSlice[0]
	policyName := idSlice[1]

	argsMap := make(map[string]string)
	argsMap["policyname"] = url.QueryEscape(policyName)

	if v, ok := d.GetOk("bindpoint"); ok {
		argsMap["bindpoint"] = url.QueryEscape(v.(string))
	}

	if v, ok := d.GetOk("priority"); ok {
		argsMap["priority"] = url.QueryEscape(fmt.Sprintf("%v", v))
	}
	err := client.DeleteResourceWithArgsMap("lbvserver_contentinspectionpolicy_binding", lbvserverName, argsMap)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
