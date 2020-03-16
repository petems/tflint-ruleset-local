package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_NewLocalFileNoChmodNeeded(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "local_file" "key_file" {
  filename = "ssh-key.pem"
  content  = <<EOT
-----BEGIN CERTIFICATE-----
MIIDpjCCAo6gAwIBAgIUItD3L/bBJsosfPrXY6wrqX06iTAwDQYJKoZIhvcNAQEL
BQAwFjEUMBIGA1UEAxMLZXhhbXBsZS5jb20wHhcNMjAwMzEzMTQzNzMxWhcNMjUw
MzEyMTQzODAxWjAtMSswKQYDVQQDEyJleGFtcGxlLmNvbSBJbnRlcm1lZGlhdGUg
QXV0aG9yaXR5MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1veq6qgz
X8X7efKNQLF7BzTKd5iFm7MypSZTpfd6kunUSKCrLoIPH+oNTUbxXLsGXPxsKvSt
b4DNoZ/XJkCPjTjNY3m11AWDD3Yg/Ons/KBPlfIwPW/c0tQs3N1t+b83lSWbU98B
Ft/pmfQelsG2lP+N7YqGTYGkShhdgO1BApJizjlO0xOyrlnKqUZrm3ccIII+iHHo
5CIHnwZoFXTrixuWDumE6nsCI7nQw4LJuuNCkOQfdVvVrcnWXK8fwRpHsZjcX4fL
v6JpSkVkIfj3zpp47b2zhdwPi8MTthvlHcDwU7+iseHsClGDhJ0FfSOpvnwQ4Wis
mHlPbCYMCzVXVQIDAQABo4HUMIHRMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBQTW6RW6565S3W0gqr8G+KFQADmVjAfBgNVHSMEGDAW
gBSPUACzYtuTIA5VNhoGitB84NPOVjA7BggrBgEFBQcBAQQvMC0wKwYIKwYBBQUH
MAKGH2h0dHA6Ly8xMjcuMC4wLjE6ODIwMC92MS9wa2kvY2EwMQYDVR0fBCowKDAm
oCSgIoYgaHR0cDovLzEyNy4wLjAuMTo4MjAwL3YxL3BraS9jcmwwDQYJKoZIhvcN
AQELBQADggEBAEwrVmDoIkamedgRvLdiyUla+DP6L1FCLlg/G+MhyGqdaDdI9zZm
oEfF7b1BtgKG+G2GrCIyZdmafCkZbRnfn+qQLsPd8rHFrhqCmr8PKJckRMXFWniJ
p5Bd1N9pziVvnctsu9JatGTMzxYvvj14UJri9aMSfCcpDscxKz9sqh+l8QCxC9qJ
bIjLj4hXgw7ggHGYVjhcqM8ifloGOsTZ1DAvNWEhoVRzw4t2083Ro0g9dS9i08VB
nrgae+OMIdV+B6Xw14GXXqpIEe4al+vN+6l9hhGPal3W0qKNvAzxue8GRDil2D4b
eQj3+9rzqbUdkaIhZosSX9/iF32FEpCztt0=
-----END CERTIFICATE-----
EOT

  provisioner "local-exec" {
    command = "chmod 600 ssh-key.pem"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewLocalFileNoChmodNeeded(),
					Message: "instance type is t2.micro",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 21},
						End:      hcl.Pos{Line: 3, Column: 31},
					},
				},
			},
		},
	}

	rule := NewLocalFileNoChmodNeeded()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
