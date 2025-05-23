// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iam2/AccessBoundaryPolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iam2

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceIAM2AccessBoundaryPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAM2AccessBoundaryPolicyCreate,
		Read:   resourceIAM2AccessBoundaryPolicyRead,
		Update: resourceIAM2AccessBoundaryPolicyUpdate,
		Delete: resourceIAM2AccessBoundaryPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIAM2AccessBoundaryPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name of the policy.`,
			},
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The attachment point is identified by its URL-encoded full resource name.`,
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Rules to be applied.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_boundary_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `An access boundary rule in an IAM policy.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"availability_condition": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `The availability condition further constrains the access allowed by the access boundary rule.`,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"expression": {
													Type:        schema.TypeString,
													Required:    true,
													Description: `Textual representation of an expression in Common Expression Language syntax.`,
												},
												"description": {
													Type:     schema.TypeString,
													Optional: true,
													Description: `Description of the expression. This is a longer text which describes the expression,
e.g. when hovered over it in a UI.`,
												},
												"location": {
													Type:     schema.TypeString,
													Optional: true,
													Description: `String indicating the location of the expression for error reporting,
e.g. a file name and a position in the file.`,
												},
												"title": {
													Type:     schema.TypeString,
													Optional: true,
													Description: `Title for the expression, i.e. a short string describing its purpose.
This can be used e.g. in UIs which allow to enter the expression.`,
												},
											},
										},
									},
									"available_permissions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `A list of permissions that may be allowed for use on the specified resource.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"available_resource": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The full resource name of a Google Cloud resource entity.`,
									},
								},
							},
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The description of the rule.`,
						},
					},
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The display name of the rule.`,
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The hash of the resource. Used internally during updates.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIAM2AccessBoundaryPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAM2AccessBoundaryPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	etagProp, err := expandIAM2AccessBoundaryPolicyEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	rulesProp, err := expandIAM2AccessBoundaryPolicyRules(d.Get("rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(rulesProp)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM2BasePath}}policies/{{parent}}/accessboundarypolicies?policyId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AccessBoundaryPolicy: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating AccessBoundaryPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = IAM2OperationWaitTime(
		config, res, "Creating AccessBoundaryPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create AccessBoundaryPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AccessBoundaryPolicy %q: %#v", d.Id(), res)

	return resourceIAM2AccessBoundaryPolicyRead(d, meta)
}

func resourceIAM2AccessBoundaryPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM2BasePath}}policies/{{parent}}/accessboundarypolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IAM2AccessBoundaryPolicy %q", d.Id()))
	}

	if err := d.Set("display_name", flattenIAM2AccessBoundaryPolicyDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("etag", flattenIAM2AccessBoundaryPolicyEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("rules", flattenIAM2AccessBoundaryPolicyRules(res["rules"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessBoundaryPolicy: %s", err)
	}

	return nil
}

func resourceIAM2AccessBoundaryPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAM2AccessBoundaryPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	etagProp, err := expandIAM2AccessBoundaryPolicyEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	rulesProp, err := expandIAM2AccessBoundaryPolicyRules(d.Get("rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM2BasePath}}policies/{{parent}}/accessboundarypolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating AccessBoundaryPolicy %q: %#v", d.Id(), obj)
	headers := make(http.Header)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PUT",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
		Headers:   headers,
	})

	if err != nil {
		return fmt.Errorf("Error updating AccessBoundaryPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating AccessBoundaryPolicy %q: %#v", d.Id(), res)
	}

	err = IAM2OperationWaitTime(
		config, res, "Updating AccessBoundaryPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceIAM2AccessBoundaryPolicyRead(d, meta)
}

func resourceIAM2AccessBoundaryPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM2BasePath}}policies/{{parent}}/accessboundarypolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting AccessBoundaryPolicy %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "AccessBoundaryPolicy")
	}

	err = IAM2OperationWaitTime(
		config, res, "Deleting AccessBoundaryPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting AccessBoundaryPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceIAM2AccessBoundaryPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^(?P<parent>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{parent}}/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIAM2AccessBoundaryPolicyDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM2AccessBoundaryPolicyEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM2AccessBoundaryPolicyRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"description":          flattenIAM2AccessBoundaryPolicyRulesDescription(original["description"], d, config),
			"access_boundary_rule": flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRule(original["accessBoundaryRule"], d, config),
		})
	}
	return transformed
}
func flattenIAM2AccessBoundaryPolicyRulesDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRule(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["available_resource"] =
		flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailableResource(original["availableResource"], d, config)
	transformed["available_permissions"] =
		flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailablePermissions(original["availablePermissions"], d, config)
	transformed["availability_condition"] =
		flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition(original["availabilityCondition"], d, config)
	return []interface{}{transformed}
}
func flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailableResource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailablePermissions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expression"] =
		flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionExpression(original["expression"], d, config)
	transformed["title"] =
		flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionTitle(original["title"], d, config)
	transformed["description"] =
		flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionDescription(original["description"], d, config)
	transformed["location"] =
		flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionLocation(original["location"], d, config)
	return []interface{}{transformed}
}
func flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionExpression(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionTitle(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIAM2AccessBoundaryPolicyDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandIAM2AccessBoundaryPolicyRulesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedAccessBoundaryRule, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRule(original["access_boundary_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAccessBoundaryRule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["accessBoundaryRule"] = transformedAccessBoundaryRule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIAM2AccessBoundaryPolicyRulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAvailableResource, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailableResource(original["available_resource"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailableResource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["availableResource"] = transformedAvailableResource
	}

	transformedAvailablePermissions, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailablePermissions(original["available_permissions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailablePermissions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["availablePermissions"] = transformedAvailablePermissions
	}

	transformedAvailabilityCondition, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition(original["availability_condition"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailabilityCondition); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["availabilityCondition"] = transformedAvailabilityCondition
	}

	return transformed, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailableResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailablePermissions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2AccessBoundaryPolicyRulesAccessBoundaryRuleAvailabilityConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
