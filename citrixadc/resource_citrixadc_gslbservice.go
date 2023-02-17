package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"bytes"
	"fmt"
	"log"
	"strconv"
)

func resourceCitrixAdcGslbservice() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGslbserviceFunc,
		Read:          readGslbserviceFunc,
		Update:        updateGslbserviceFunc,
		Delete:        deleteGslbserviceFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cnameentry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookietimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hashid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"healthmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxaaausers": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxclient": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"monitornamesvc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"naptrdomainttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"naptrorder": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"naptrpreference": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"naptrreplacement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"naptrservices": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"publicip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publicport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"servername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sitename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sitepersistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"siteprefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"viewip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"viewname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lbmonitorbinding": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: false,
				Set:      lbmonitorMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"monitor_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"monstate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In createGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	var gslbserviceName string
	if v, ok := d.GetOk("servicename"); ok {
		gslbserviceName = v.(string)
	} else {
		gslbserviceName = resource.PrefixedUniqueId("tf-gslbservice-")
		d.Set("servicename", gslbserviceName)
	}
	gslbservice := gslb.Gslbservice{
		Appflowlog:       d.Get("appflowlog").(string),
		Cip:              d.Get("cip").(string),
		Cipheader:        d.Get("cipheader").(string),
		Clttimeout:       d.Get("clttimeout").(int),
		Cnameentry:       d.Get("cnameentry").(string),
		Comment:          d.Get("comment").(string),
		Cookietimeout:    d.Get("cookietimeout").(int),
		Downstateflush:   d.Get("downstateflush").(string),
		Hashid:           d.Get("hashid").(int),
		Healthmonitor:    d.Get("healthmonitor").(string),
		Ip:               d.Get("ip").(string),
		Ipaddress:        d.Get("ipaddress").(string),
		Maxaaausers:      d.Get("maxaaausers").(int),
		Maxbandwidth:     d.Get("maxbandwidth").(int),
		Maxclient:        d.Get("maxclient").(int),
		Monitornamesvc:   d.Get("monitornamesvc").(string),
		Monthreshold:     d.Get("monthreshold").(int),
		Naptrdomainttl:   d.Get("naptrdomainttl").(int),
		Naptrorder:       d.Get("naptrorder").(int),
		Naptrpreference:  d.Get("naptrpreference").(int),
		Naptrreplacement: d.Get("naptrreplacement").(string),
		Naptrservices:    d.Get("naptrservices").(string),
		Port:             d.Get("port").(int),
		Publicip:         d.Get("publicip").(string),
		Publicport:       d.Get("publicport").(int),
		Servername:       d.Get("servername").(string),
		Servicename:      d.Get("servicename").(string),
		Servicetype:      d.Get("servicetype").(string),
		Sitename:         d.Get("sitename").(string),
		Sitepersistence:  d.Get("sitepersistence").(string),
		Siteprefix:       d.Get("siteprefix").(string),
		State:            d.Get("state").(string),
		Svrtimeout:       d.Get("svrtimeout").(int),
		Viewip:           d.Get("viewip").(string),
		Viewname:         d.Get("viewname").(string),
		Weight:           d.Get("weight").(int),
	}

	_, err := client.AddResource(service.Gslbservice.Type(), gslbserviceName, &gslbservice)
	if err != nil {
		return err
	}

	if err := updateLbmonitorBindinds(d, meta); err != nil {
		return err
	}

	d.SetId(gslbserviceName)

	err = readGslbserviceFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this gslbservice but we can't read it ?? %s", gslbserviceName)
		return nil
	}
	return nil
}

func readGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] netscaler-provider:  In readGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Id()
	log.Printf("[DEBUG] netscaler-provider: Reading gslbservice state %s", gslbserviceName)
	data, err := client.FindResource(service.Gslbservice.Type(), gslbserviceName)
	if err != nil {
		log.Printf("[WARN] netscaler-provider: Clearing gslbservice state %s", gslbserviceName)
		d.SetId("")
		return nil
	}
	d.Set("servicename", data["servicename"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("cip", data["cip"])
	d.Set("cipheader", data["cipheader"])
	setToInt("clttimeout", d, data["clttimeout"])
	d.Set("cnameentry", data["cnameentry"])
	d.Set("comment", data["comment"])
	d.Set("cookietimeout", data["cookietimeout"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("hashid", data["hashid"])
	d.Set("healthmonitor", data["healthmonitor"])
	d.Set("ip", data["ipaddress"]) //ip is not returned, but it ipaddress is returned by NITRO
	d.Set("ipaddress", data["ipaddress"])
	d.Set("maxaaausers", data["maxaaausers"])
	d.Set("maxbandwidth", data["maxbandwidth"])
	setToInt("maxclient", d, data["maxclient"])
	d.Set("monitornamesvc", data["monitornamesvc"])
	d.Set("monthreshold", data["monthreshold"])
	d.Set("naptrdomainttl", data["naptrdomainttl"])
	d.Set("naptrorder", data["naptrorder"])
	d.Set("naptrpreference", data["naptrpreference"])
	d.Set("naptrreplacement", data["naptrreplacement"])
	d.Set("naptrservices", data["naptrservices"])
	d.Set("port", data["port"])
	d.Set("publicip", data["publicip"])
	d.Set("publicport", data["publicport"])
	d.Set("servername", data["servername"])
	d.Set("servicename", data["servicename"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sitename", data["sitename"])
	d.Set("sitepersistence", data["sitepersistence"])
	d.Set("siteprefix", data["siteprefix"])
	d.Set("state", data["state"])
	d.Set("svrtimeout", data["svrtimeout"])
	d.Set("viewip", data["viewip"])
	d.Set("viewname", data["viewname"])
	d.Set("weight", data["weight"])

	if err := readLbmonitorBindinds(d, meta); err != nil {
		return err
	}

	return nil

}

func updateGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In updateGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Get("servicename").(string)

	gslbservice := gslb.Gslbservice{
		Servicename: d.Get("servicename").(string),
	}
	stateChange := false
	hasChange := false
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG]  netscaler-provider: Appflowlog has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Appflowlog = d.Get("appflowlog").(string)
		hasChange = true
	}
	if d.HasChange("cip") {
		log.Printf("[DEBUG]  netscaler-provider: Cip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cip = d.Get("cip").(string)
		hasChange = true
	}
	if d.HasChange("cipheader") {
		log.Printf("[DEBUG]  netscaler-provider: Cipheader has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cipheader = d.Get("cipheader").(string)
		hasChange = true
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Clttimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Clttimeout = d.Get("clttimeout").(int)
		hasChange = true
	}
	if d.HasChange("cnameentry") {
		log.Printf("[DEBUG]  netscaler-provider: Cnameentry has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cnameentry = d.Get("cnameentry").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  netscaler-provider: Comment has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("cookietimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Cookietimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Cookietimeout = d.Get("cookietimeout").(int)
		hasChange = true
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG]  netscaler-provider: Downstateflush has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Downstateflush = d.Get("downstateflush").(string)
		hasChange = true
	}
	if d.HasChange("hashid") {
		log.Printf("[DEBUG]  netscaler-provider: Hashid has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Hashid = d.Get("hashid").(int)
		hasChange = true
	}
	if d.HasChange("healthmonitor") {
		log.Printf("[DEBUG]  netscaler-provider: Healthmonitor has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Healthmonitor = d.Get("healthmonitor").(string)
		hasChange = true
	}
	if d.HasChange("ip") {
		log.Printf("[DEBUG]  netscaler-provider: Ip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Ipaddress = d.Get("ip").(string) //use ipaddress during Update
		hasChange = true
	}
	if d.HasChange("ipaddress") {
		log.Printf("[DEBUG]  netscaler-provider: Ipaddress has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Ipaddress = d.Get("ipaddress").(string)
		hasChange = true
	}
	if d.HasChange("maxaaausers") {
		log.Printf("[DEBUG]  netscaler-provider: Maxaaausers has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxaaausers = d.Get("maxaaausers").(int)
		hasChange = true
	}
	if d.HasChange("maxbandwidth") {
		log.Printf("[DEBUG]  netscaler-provider: Maxbandwidth has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxbandwidth = d.Get("maxbandwidth").(int)
		hasChange = true
	}
	if d.HasChange("maxclient") {
		log.Printf("[DEBUG]  netscaler-provider: Maxclient has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Maxclient = d.Get("maxclient").(int)
		hasChange = true
	}
	if d.HasChange("monitornamesvc") {
		log.Printf("[DEBUG]  netscaler-provider: Monitornamesvc has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Monitornamesvc = d.Get("monitornamesvc").(string)
		hasChange = true
	}
	if d.HasChange("monthreshold") {
		log.Printf("[DEBUG]  netscaler-provider: Monthreshold has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Monthreshold = d.Get("monthreshold").(int)
		hasChange = true
	}
	if d.HasChange("naptrdomainttl") {
		log.Printf("[DEBUG]  netscaler-provider: Naptrdomainttl has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrdomainttl = d.Get("naptrdomainttl").(int)
		hasChange = true
	}
	if d.HasChange("naptrorder") {
		log.Printf("[DEBUG]  netscaler-provider: Naptrorder has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrorder = d.Get("naptrorder").(int)
		hasChange = true
	}
	if d.HasChange("naptrpreference") {
		log.Printf("[DEBUG]  netscaler-provider: Naptrpreference has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrpreference = d.Get("naptrpreference").(int)
		hasChange = true
	}
	if d.HasChange("naptrreplacement") {
		log.Printf("[DEBUG]  netscaler-provider: Naptrreplacement has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrreplacement = d.Get("naptrreplacement").(string)
		hasChange = true
	}
	if d.HasChange("naptrservices") {
		log.Printf("[DEBUG]  netscaler-provider: Naptrservices has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Naptrservices = d.Get("naptrservices").(string)
		hasChange = true
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG]  netscaler-provider: Port has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Port = d.Get("port").(int)
		hasChange = true
	}
	if d.HasChange("publicip") {
		log.Printf("[DEBUG]  netscaler-provider: Publicip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Publicip = d.Get("publicip").(string)
		hasChange = true
	}
	if d.HasChange("publicport") {
		log.Printf("[DEBUG]  netscaler-provider: Publicport has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Publicport = d.Get("publicport").(int)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  netscaler-provider: Servername has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("servicename") {
		log.Printf("[DEBUG]  netscaler-provider: Servicename has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servicename = d.Get("servicename").(string)
		hasChange = true
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG]  netscaler-provider: Servicetype has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Servicetype = d.Get("servicetype").(string)
		hasChange = true
	}
	if d.HasChange("sitename") {
		log.Printf("[DEBUG]  netscaler-provider: Sitename has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Sitename = d.Get("sitename").(string)
		hasChange = true
	}
	if d.HasChange("sitepersistence") {
		log.Printf("[DEBUG]  netscaler-provider: Sitepersistence has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Sitepersistence = d.Get("sitepersistence").(string)
		hasChange = true
	}
	if d.HasChange("siteprefix") {
		log.Printf("[DEBUG]  netscaler-provider: Siteprefix has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Siteprefix = d.Get("siteprefix").(string)
		hasChange = true
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG]  netscaler-provider: State has changed for gslbservice %s, starting update", gslbserviceName)
		stateChange = true
	}
	if d.HasChange("svrtimeout") {
		log.Printf("[DEBUG]  netscaler-provider: Svrtimeout has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Svrtimeout = d.Get("svrtimeout").(int)
		hasChange = true
	}
	if d.HasChange("viewip") {
		log.Printf("[DEBUG]  netscaler-provider: Viewip has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Viewip = d.Get("viewip").(string)
		hasChange = true
	}
	if d.HasChange("viewname") {
		log.Printf("[DEBUG]  netscaler-provider: Viewname has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Viewname = d.Get("viewname").(string)
		hasChange = true
	}
	if d.HasChange("weight") {
		log.Printf("[DEBUG]  netscaler-provider: Weight has changed for gslbservice %s, starting update", gslbserviceName)
		gslbservice.Weight = d.Get("weight").(int)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Gslbservice.Type(), gslbserviceName, &gslbservice)
		if err != nil {
			return fmt.Errorf("Error updating gslbservice %s", gslbserviceName)
		}
	}

	if stateChange {
		err := doGslbServiceStateChange(d, client)
		if err != nil {
			return fmt.Errorf("Error enabling/disabling glsb service %s", gslbserviceName)
		}
	}

	if d.HasChange("lbmonitorbinding") {
		if err := updateLbmonitorBindinds(d, meta); err != nil {
			return err
		}

	}

	return readGslbserviceFunc(d, meta)
}

func deleteGslbserviceFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In deleteGslbserviceFunc")
	client := meta.(*NetScalerNitroClient).client
	gslbserviceName := d.Id()
	err := client.DeleteResource(service.Gslbservice.Type(), gslbserviceName)
	if err != nil {
		return err
	}

	if err := deleteLbmonitorBindinds(d, meta); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func doGslbServiceStateChange(d *schema.ResourceData, client *service.NitroClient) error {
	log.Printf("[DEBUG]  netscaler-provider: In doGslbServiceStateChange")

	// We need a new instance of the struct since
	// ActOnResource will fail if we put in superfluous attributes
	gslbService := basic.Service{
		Name: d.Get("servicename").(string),
	}

	newstate := d.Get("state")

	// Enable action
	if newstate == "ENABLED" {
		err := client.ActOnResource(service.Service.Type(), gslbService, "enable")
		if err != nil {
			return err
		}
	} else if newstate == "DISABLED" {
		// Add attributes relevant to the disable operation
		gslbService.Delay = d.Get("delay").(int)
		err := client.ActOnResource(service.Service.Type(), gslbService, "disable")
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\").", newstate)
	}

	return nil
}

func readLbmonitorBindinds(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readLbmonitorBindinds")

	client := meta.(*NetScalerNitroClient).client
	servicename := d.Get("servicename").(string)
	lbmonitorBindings, _ := client.FindResourceArray("gslbservice_lbmonitor_binding", servicename)
	log.Printf("lbmonitorBindings %v\n", lbmonitorBindings)

	processedBindings := make([]interface{}, len(lbmonitorBindings))

	for i, m := range lbmonitorBindings {
		processedBindings[i] = make(map[string]interface{})
		if d, ok := m["weight"]; ok {
			if intval, err := strconv.Atoi(d.(string)); err != nil {
				processedBindings[i].(map[string]interface{})["weight"] = intval
			} else {
				return err
			}
		}
		if d, ok := m["monitor_name"]; ok {
			processedBindings[i].(map[string]interface{})["monitor_name"] = d.(string)
		}
		if d, ok := m["monstate"]; ok {
			processedBindings[i].(map[string]interface{})["monstate"] = d.(string)
		}
	}

	updatedSet := schema.NewSet(lbmonitorMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("lbmonitorbinding", updatedSet); err != nil {
		return err
	}

	return nil
}

func updateLbmonitorBindinds(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readLbmonitorBindinds")

	oldSet, newSet := d.GetChange("lbmonitorbinding")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleLbmonitorBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}

	for _, binding := range add.List() {
		if err := addSingleLbmonitorBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func deleteLbmonitorBindinds(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateLbmonitorBindinds")
	if bindings, ok := d.GetOk("lbmonitorbinding"); ok {
		for _, binding := range bindings.(*schema.Set).List() {
			if err := deleteSingleLbmonitorBinding(d, meta, binding.(map[string]interface{})); err != nil {
				return err
			}
		}
	}
	return nil
}

func addSingleLbmonitorBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleLbmonitorBinding")

	client := meta.(*NetScalerNitroClient).client

	bindingStruct := gslb.Gslbservicemonitorbinding{}
	servicename := d.Get("servicename").(string)
	bindingStruct.Servicename = servicename

	if d, ok := binding["weight"]; ok {
		bindingStruct.Weight = uint32(d.(int))
	}

	if d, ok := binding["monitor_name"]; ok {
		bindingStruct.Monitorname = d.(string)
	}

	if d, ok := binding["monstate"]; ok {
		bindingStruct.Monstate = d.(string)
	}

	if _, err := client.UpdateResource("gslbservice_lbmonitor_binding", servicename, bindingStruct); err != nil {
		return err
	}
	return nil
}

func deleteSingleLbmonitorBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleLbmonitorBinding")
	client := meta.(*NetScalerNitroClient).client

	// Construct args from binding data
	args := make([]string, 0, 1)

	if d, ok := binding["monitor_name"]; ok {
		s := fmt.Sprintf("monitor_name:%s", d.(string))
		args = append(args, s)
	}

	servicename := d.Get("servicename").(string)

	if err := client.DeleteResourceWithArgs("gslbservice_lbmonitor_binding", servicename, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting lb monitor binding %v\n", binding)
		return err
	}

	return nil
}

func lbmonitorMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In lbmonitorMappingHash")
	var buf bytes.Buffer

	// All keys added in alphabetical order.
	m := v.(map[string]interface{})
	if d, ok := m["weight"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}

	if d, ok := m["monitor_name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}

	if d, ok := m["monstate"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	return hashcode.String(buf.String())
}
