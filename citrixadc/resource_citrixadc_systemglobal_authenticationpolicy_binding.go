package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSystemglobal_authenticationpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemglobal_authenticationpolicy_bindingFunc,
		Read:          readSystemglobal_authenticationpolicy_bindingFunc,
		Delete:        deleteSystemglobal_authenticationpolicy_bindingFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
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
			"builtin": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"feature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"nextfactor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createSystemglobal_authenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemglobal_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	systemglobal_authenticationpolicy_binding := system.Systemglobalauthenticationpolicybinding{
		//Builtin:                d.Get("builtin").([]interface{}),
		//Feature:                d.Get("feature").(string),
		Globalbindtype:         d.Get("globalbindtype").(string),
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Nextfactor:             d.Get("nextfactor").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource(service.Systemglobal_authenticationpolicy_binding.Type(), &systemglobal_authenticationpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policyname)

	err = readSystemglobal_authenticationpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemglobal_authenticationpolicy_binding but we can't read it ?? %s", policyname)
		return nil
	}
	return nil
}

func readSystemglobal_authenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemglobal_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()


	log.Printf("[DEBUG] citrixadc-provider: Reading systemglobal_authenticationpolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType:             "systemglobal_authenticationpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing systemglobal_authenticationpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if v["policyname"].(string) == policyname {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams policyname not found in array")
		log.Printf("[WARN] citrixadc-provider: Clearing systemglobal_authenticationpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("builtin", data["builtin"])
	d.Set("feature", data["feature"])
	d.Set("globalbindtype", data["globalbindtype"])
	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("nextfactor", data["nextfactor"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteSystemglobal_authenticationpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemglobal_authenticationpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()

	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))

	err := client.DeleteResourceWithArgs(service.Systemglobal_authenticationpolicy_binding.Type(), "" , args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
