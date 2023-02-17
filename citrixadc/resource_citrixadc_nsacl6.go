package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNsacl6() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNsacl6Func,
		Read:          readNsacl6Func,
		Update:        updateNsacl6Func,
		Delete:        deleteNsacl6Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"acl6name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"acl6action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"aclaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destipv6": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destipv6val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destport": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"destportop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destportval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dfdhash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dfdprefix": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"established": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"icmpcode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmptype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocolnumber": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ratelimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"srcipop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcipv6": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcipv6val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcmacmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcport": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"srcportop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcportval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stateful": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vxlan": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNsacl6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNsacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nsacl6Name := d.Get("acl6name").(string)
	nsacl6 := ns.Nsacl6{
		Acl6action:     d.Get("acl6action").(string),
		Acl6name:       d.Get("acl6name").(string),
		Aclaction:      d.Get("aclaction").(string),
		Destipop:       d.Get("destipop").(string),
		Destipv6:       d.Get("destipv6").(bool),
		Destipv6val:    d.Get("destipv6val").(string),
		Destport:       d.Get("destport").(bool),
		Destportop:     d.Get("destportop").(string),
		Destportval:    d.Get("destportval").(string),
		Dfdhash:        d.Get("dfdhash").(string),
		Dfdprefix:      d.Get("dfdprefix").(int),
		Established:    d.Get("established").(bool),
		Icmpcode:       d.Get("icmpcode").(int),
		Icmptype:       d.Get("icmptype").(int),
		Interface:      d.Get("interface").(string),
		Logstate:       d.Get("logstate").(string),
		Newname:        d.Get("newname").(string),
		Priority:       d.Get("priority").(int),
		Protocol:       d.Get("protocol").(string),
		Protocolnumber: d.Get("protocolnumber").(int),
		Ratelimit:      d.Get("ratelimit").(int),
		Srcipop:        d.Get("srcipop").(string),
		Srcipv6:        d.Get("srcipv6").(bool),
		Srcipv6val:     d.Get("srcipv6val").(string),
		Srcmac:         d.Get("srcmac").(string),
		Srcmacmask:     d.Get("srcmacmask").(string),
		Srcport:        d.Get("srcport").(bool),
		Srcportop:      d.Get("srcportop").(string),
		Srcportval:     d.Get("srcportval").(string),
		State:          d.Get("state").(string),
		Stateful:       d.Get("stateful").(string),
		Td:             d.Get("td").(int),
		Ttl:            d.Get("ttl").(int),
		Type:           d.Get("type").(string),
		Vlan:           d.Get("vlan").(int),
		Vxlan:          d.Get("vxlan").(int),
	}

	_, err := client.AddResource(service.Nsacl6.Type(), nsacl6Name, &nsacl6)
	if err != nil {
		return err
	}

	d.SetId(nsacl6Name)

	err = readNsacl6Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nsacl6 but we can't read it ?? %s", nsacl6Name)
		return nil
	}
	return nil
}

func readNsacl6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNsacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nsacl6Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading nsacl6 state %s", nsacl6Name)
	data, err := client.FindResource(service.Nsacl6.Type(), nsacl6Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nsacl6 state %s", nsacl6Name)
		d.SetId("")
		return nil
	}
	d.Set("acl6action", data["acl6action"])
	d.Set("acl6name", data["acl6name"])
	d.Set("aclaction", data["aclaction"])
	d.Set("destipop", data["destipop"])
	d.Set("destipv6", data["destipv6"])
	d.Set("destipv6val", data["destipv6val"])
	d.Set("destport", data["destport"])
	d.Set("destportop", data["destportop"])
	d.Set("destportval", data["destportval"])
	d.Set("dfdhash", data["dfdhash"])
	d.Set("dfdprefix", data["dfdprefix"])
	d.Set("established", data["established"])
	d.Set("icmpcode", data["icmpcode"])
	d.Set("icmptype", data["icmptype"])
	d.Set("interface", data["interface"])
	d.Set("logstate", data["logstate"])
	d.Set("newname", data["newname"])
	setToInt("priority", d, data["priority"])
	d.Set("protocol", data["protocol"])
	d.Set("protocolnumber", data["protocolnumber"])
	d.Set("ratelimit", data["ratelimit"])
	d.Set("srcipop", data["srcipop"])
	d.Set("srcipv6", data["srcipv6"])
	d.Set("srcipv6val", data["srcipv6val"])
	d.Set("srcmac", data["srcmac"])
	d.Set("srcmacmask", data["srcmacmask"])
	d.Set("srcport", data["srcport"])
	d.Set("srcportop", data["srcportop"])
	d.Set("srcportval", data["srcportval"])
	d.Set("state", data["state"])
	d.Set("stateful", data["stateful"])
	d.Set("td", data["td"])
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])
	d.Set("vlan", data["vlan"])
	d.Set("vxlan", data["vxlan"])

	return nil

}

func updateNsacl6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateNsacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nsacl6Name := d.Get("acl6name").(string)

	nsacl6 := ns.Nsacl6{
		Acl6name: d.Get("acl6name").(string),
	}
	hasChange := false
	stateChange := false
	if d.HasChange("acl6action") {
		log.Printf("[DEBUG]  citrixadc-provider: Acl6action has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Acl6action = d.Get("acl6action").(string)
		hasChange = true
	}
	if d.HasChange("aclaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Aclaction has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Aclaction = d.Get("aclaction").(string)
		hasChange = true
	}
	if d.HasChange("destipop") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipop has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destipop = d.Get("destipop").(string)
		hasChange = true
	}
	if d.HasChange("destipv6") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipv6 has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destipv6 = d.Get("destipv6").(bool)
		hasChange = true
	}
	if d.HasChange("destipv6val") {
		log.Printf("[DEBUG]  citrixadc-provider: Destipv6val has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destipv6val = d.Get("destipv6val").(string)
		hasChange = true
	}
	if d.HasChange("destport") {
		log.Printf("[DEBUG]  citrixadc-provider: Destport has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destport = d.Get("destport").(bool)
		hasChange = true
	}
	if d.HasChange("destportop") {
		log.Printf("[DEBUG]  citrixadc-provider: Destportop has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destportop = d.Get("destportop").(string)
		hasChange = true
	}
	if d.HasChange("destportval") {
		log.Printf("[DEBUG]  citrixadc-provider: Destportval has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Destportval = d.Get("destportval").(string)
		hasChange = true
	}
	if d.HasChange("dfdhash") {
		log.Printf("[DEBUG]  citrixadc-provider: Dfdhash has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Dfdhash = d.Get("dfdhash").(string)
		hasChange = true
	}
	if d.HasChange("dfdprefix") {
		log.Printf("[DEBUG]  citrixadc-provider: Dfdprefix has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Dfdprefix = d.Get("dfdprefix").(int)
		hasChange = true
	}
	if d.HasChange("established") {
		log.Printf("[DEBUG]  citrixadc-provider: Established has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Established = d.Get("established").(bool)
		hasChange = true
	}
	if d.HasChange("icmpcode") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmpcode has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Icmpcode = d.Get("icmpcode").(int)
		hasChange = true
	}
	if d.HasChange("icmptype") {
		log.Printf("[DEBUG]  citrixadc-provider: Icmptype has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Icmptype = d.Get("icmptype").(int)
		hasChange = true
	}
	if d.HasChange("interface") {
		log.Printf("[DEBUG]  citrixadc-provider: Interface has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Interface = d.Get("interface").(string)
		hasChange = true
	}
	if d.HasChange("logstate") {
		log.Printf("[DEBUG]  citrixadc-provider: Logstate has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Logstate = d.Get("logstate").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("priority") {
		log.Printf("[DEBUG]  citrixadc-provider: Priority has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Priority = d.Get("priority").(int)
		hasChange = true
	}
	if d.HasChange("protocol") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocol has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Protocol = d.Get("protocol").(string)
		hasChange = true
	}
	if d.HasChange("protocolnumber") {
		log.Printf("[DEBUG]  citrixadc-provider: Protocolnumber has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Protocolnumber = d.Get("protocolnumber").(int)
		hasChange = true
	}
	if d.HasChange("ratelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Ratelimit has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Ratelimit = d.Get("ratelimit").(int)
		hasChange = true
	}
	if d.HasChange("srcipop") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipop has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcipop = d.Get("srcipop").(string)
		hasChange = true
	}
	if d.HasChange("srcipv6") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipv6 has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcipv6 = d.Get("srcipv6").(bool)
		hasChange = true
	}
	if d.HasChange("srcipv6val") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcipv6val has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcipv6val = d.Get("srcipv6val").(string)
		hasChange = true
	}
	if d.HasChange("srcmac") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcmac has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcmac = d.Get("srcmac").(string)
		hasChange = true
	}
	if d.HasChange("srcmacmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcmacmask has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcmacmask = d.Get("srcmacmask").(string)
		hasChange = true
	}
	if d.HasChange("srcport") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcport has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcport = d.Get("srcport").(bool)
		hasChange = true
	}
	if d.HasChange("srcportop") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcportop has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcportop = d.Get("srcportop").(string)
		hasChange = true
	}
	if d.HasChange("srcportval") {
		log.Printf("[DEBUG]  citrixadc-provider: Srcportval has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Srcportval = d.Get("srcportval").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  citrixadc-provider: State has changed for nsacl6 %s, starting update", nsacl6Name)
		stateChange = true
	}
	if d.HasChange("stateful") {
		log.Printf("[DEBUG]  citrixadc-provider: Stateful has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Stateful = d.Get("stateful").(string)
		hasChange = true
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG]  citrixadc-provider: Td has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Td = d.Get("td").(int)
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Ttl = d.Get("ttl").(int)
		hasChange = true
	}
	if d.HasChange("vlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vlan has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Vlan = d.Get("vlan").(int)
		hasChange = true
	}
	if d.HasChange("vxlan") {
		log.Printf("[DEBUG]  citrixadc-provider: Vxlan has changed for nsacl6 %s, starting update", nsacl6Name)
		nsacl6.Vxlan = d.Get("vxlan").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Nsacl6.Type(), nsacl6Name, &nsacl6)
		if err != nil {
			return fmt.Errorf("Error updating nsacl6 %s", nsacl6Name)
		}
	}

	if stateChange {
		err := doNsacl6StateSchange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling Nsacl6 %s", nsacl6Name)
		}
	}
	return readNsacl6Func(d, meta)
}

func deleteNsacl6Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNsacl6Func")
	client := meta.(*NetScalerNitroClient).client
	nsacl6Name := d.Id()
	err := client.DeleteResource(service.Nsacl6.Type(), nsacl6Name)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func doNsacl6StateSchange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doNsacl6StateSchange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	nsacl6 := ns.Nsacl6{
		Acl6name: d.Get("acl6name").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Nsacl6.Type(), nsacl6, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		err := client.ActOnResource(service.Nsacl6.Type(), nsacl6, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}