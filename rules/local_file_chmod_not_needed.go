package rules

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/hcl/v2"
	tflint "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

const pattern = `chmod`

var containsChmod = regexp.MustCompile(pattern)

// LocalFileNoChmodNeeded checks whether a local file runs chmod
type LocalFileNoChmodNeeded struct {
	resourceType  string
	attributeName string
}

// NewLocalFileNoChmodNeeded returns new rule with default attributes
func NewLocalFileNoChmodNeeded() *LocalFileNoChmodNeeded {
	return &LocalFileNoChmodNeeded{
		resourceType:  "local_file",
		attributeName: "provisioner",
	}
}

// Name returns the rule name
func (r *LocalFileNoChmodNeeded) Name() string {
	return "local_file_no_chmod_needed"
}

// Enabled returns whether the rule is enabled by default
func (r *LocalFileNoChmodNeeded) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *LocalFileNoChmodNeeded) Severity() string {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *LocalFileNoChmodNeeded) Link() string {
	return ""
}

// Check checks whether ...
func (r *LocalFileNoChmodNeeded) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("local_file", "provisioner", func(attribute *hcl.Attribute) error {
		var provisoner string
		err := runner.EvaluateExpr(attribute.Expr, &provisoner)

		return runner.EnsureNoError(err, func() error {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("provisoner is present %s", provisoner),
				attribute.Expr.Range(),
				tflint.Metadata{Expr: attribute.Expr},
			)
		})
	})
}
