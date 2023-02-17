package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/tm"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcTmglobal_auditsyslogpolicy_binding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createTmglobal_auditsyslogpolicy_bindingFunc,
		Read:          readTmglobal_auditsyslogpolicy_bindingFunc,
		Delete:        deleteTmglobal_auditsyslogpolicy_bindingFunc,
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
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createTmglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Get("policyname").(string)
	tmglobal_auditsyslogpolicy_binding := tm.Tmglobalauditsyslogpolicybinding{
		Gotopriorityexpression: d.Get("gotopriorityexpression").(string),
		Policyname:             d.Get("policyname").(string),
		Priority:               d.Get("priority").(int),
	}

	err := client.UpdateUnnamedResource(service.Tmglobal_auditsyslogpolicy_binding.Type(), &tmglobal_auditsyslogpolicy_binding)
	if err != nil {
		return err
	}

	d.SetId(policyname)

	err = readTmglobal_auditsyslogpolicy_bindingFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this tmglobal_auditsyslogpolicy_binding but we can't read it ?? %s", policyname)
		return nil
	}
	return nil
}

func readTmglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client
	policyname := d.Id()

	log.Printf("[DEBUG] citrixadc-provider: Reading tmglobal_auditsyslogpolicy_binding state %s", policyname)

	findParams := service.FindParams{
		ResourceType:             "tmglobal_auditsyslogpolicy_binding",
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
		log.Printf("[WARN] citrixadc-provider: Clearing tmglobal_auditsyslogpolicy_binding state %s", policyname)
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
		log.Printf("[WARN] citrixadc-provider: Clearing tmglobal_auditsyslogpolicy_binding state %s", policyname)
		d.SetId("")
		return nil
	}
	// Fallthrough

	data := dataArr[foundIndex]

	d.Set("gotopriorityexpression", data["gotopriorityexpression"])
	d.Set("policyname", data["policyname"])
	setToInt("priority", d, data["priority"])

	return nil

}

func deleteTmglobal_auditsyslogpolicy_bindingFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmglobal_auditsyslogpolicy_bindingFunc")
	client := meta.(*NetScalerNitroClient).client

	policyname := d.Id()
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", policyname))

	err := client.DeleteResourceWithArgs(service.Tmglobal_auditsyslogpolicy_binding.Type(), "", args)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
