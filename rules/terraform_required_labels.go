package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/project"
)

// TerraformRequiredLabelsRule checks whether a terraform version has required_version attribute
type TerraformRequiredLabelsRule struct {
	tflint.DefaultRule
}

// NewTerraformRequiredLabelsRule returns new rule with default attributes
func NewTerraformRequiredLabelsRule() *TerraformRequiredLabelsRule {
	return &TerraformRequiredLabelsRule{}
}

// Name returns the rule name
func (r *TerraformRequiredLabelsRule) Name() string {
	return "terraform_required_labels"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformRequiredLabelsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TerraformRequiredLabelsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *TerraformRequiredLabelsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check Checks whether required_version is set
func (r *TerraformRequiredLabelsRule) Check(runner tflint.Runner) error {
	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}
	if len(files) == 0 {
		// This rule does not run on non-Terraform directory.
		return nil
	}

	body, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "locals",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{{Name: "labels"}},
				},
			},
		},
	}, &tflint.GetModuleContentOption{ExpandMode: tflint.ExpandModeNone})
	if err != nil {
		return err
	}

	var exists bool

	for _, block := range body.Blocks {
		_, ok := block.Body.Attributes["labels"]
		exists = exists || ok
	}

	if !exists {
		return runner.EmitIssue(
			r,
			`terraform "labels" attribute is required`,
			hcl.Range{},
		)
	}

	return nil
}