package idefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/marty-macfly/goidefix"
	"github.com/marty-macfly/goidefix/services/project"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Description:   "Manages project.",
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The id of the project.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of project (must be unique).",
			},
			"company_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The company ID associated to the CI.",
			},
			"parent_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ID of the parent project.",
			},
			"wbs_france": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"wbs_vietnam": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"wbs_singapour": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"wbs_maurice": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"wbs_luxembourg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"wbs_hong_kong": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"wbs_chine": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"wbs_canada": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"wbs_belgique": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The WBS of this project",
			},
			"contract_number": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Contract number",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*goidefix.Idefix)

	project, err := client.Project.Create(ctx, &project.CreateRequest{
		Name:           d.Get("name").(string),
		CompanyID:      d.Get("company_id").(int),
		ParentID:       d.Get("parent_id").(int),
		WbsFrance:      d.Get("wbs_france").(string),
		WbsVietnam:     d.Get("wbs_vietnam").(string),
		WbsSingapour:   d.Get("wbs_singapour").(string),
		WbsMaurice:     d.Get("wbs_maurice").(string),
		WbsLuxembourg:  d.Get("wbs_luxembourg").(string),
		WbsHongKong:    d.Get("wbs_hong_kong").(string),
		WbsChine:       d.Get("wbs_chine").(string),
		WbsCanada:      d.Get("wbs_canada").(string),
		WbsBelgique:    d.Get("wbs_belgique").(string),
		ContractNumber: d.Get("contract_number").(string),
		TypeName:       "Suivi",
		InvoiceType:    "FDT",
		InitialBudget:  "0",
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(project.ID)

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*goidefix.Idefix)

	project, err := client.Project.Read(ctx, &project.ReadRequest{
		ID: d.Id(),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	if project == nil {
		d.SetId("")

		return diags
	}

	d.SetId(d.Id())
	d.Set("name", project.Name)
	d.Set("company_id", project.CompanyID)
	d.Set("parent_id", project.ParentID)

	return diags
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*goidefix.Idefix)

	_, err := client.Project.Update(ctx, &project.UpdateRequest{
		ID:             d.Id(),
		Name:           d.Get("name").(string),
		CompanyID:      d.Get("company_id").(int),
		ParentID:       d.Get("parent_id").(int),
		WbsFrance:      d.Get("wbs_france").(string),
		WbsVietnam:     d.Get("wbs_vietnam").(string),
		WbsSingapour:   d.Get("wbs_singapour").(string),
		WbsMaurice:     d.Get("wbs_maurice").(string),
		WbsLuxembourg:  d.Get("wbs_luxembourg").(string),
		WbsHongKong:    d.Get("wbs_hong_kong").(string),
		WbsChine:       d.Get("wbs_chine").(string),
		WbsCanada:      d.Get("wbs_canada").(string),
		WbsBelgique:    d.Get("wbs_belgique").(string),
		ContractNumber: d.Get("contract_number").(string),
		TypeName:       "Suivi",
		InvoiceType:    "FDT",
		InitialBudget:  "0",
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*goidefix.Idefix)

	_, err := client.Project.Delete(ctx, &project.DeleteRequest{
		ID: d.Id(),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
