package puppetdb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestFactNames performs a test on the FactNames endpoint and verifies the expected response is returned.
func TestFactNames(t *testing.T) {
	// Test FactNames
	setupGetResponder(t, factnames, "", "factnames-response.json")
	actual, err := pdbClient.FactNames()
	require.Nil(t, err)
	require.Equal(t, expectedFactNames, actual)
}

// TestFacts performs a test on the Facts method, and verified the expected response is returned,
func TestFacts(t *testing.T) {
	// Test with query
	query := `["=", "certname", "foobar.puppetlabs.net"]`
	setupGetResponder(t, facts, "query="+query, "facts-response.json")
	actual, err := pdbClient.Facts(query)
	require.Nil(t, err)
	require.Equal(t, expectedFacts, actual)
}

var expectedFactNames = []string{"agent_canary", "agent_specified_environment", "aio_agent_build", "aio_agent_version"}
var expectedFacts = []Fact{
	{Name: "id", Value: "root", Certname: "foobar.puppetlabs.net", Environment: "production"},
	{Name: "os", Value: map[string]interface{}{"architecture": "x86_64", "distro": map[string]interface{}{"codename": "Core", "description": "CentOS Linux release 7.4.1708 (Core)", "id": "CentOS", "release": map[string]interface{}{"full": "7.4.1708", "major": "7", "minor": "4"}, "specification": ":core-4.1-amd64:core-4.1-noarch"}, "family": "RedHat", "hardware": "x86_64", "name": "CentOS", "release": map[string]interface{}{"full": "7.4.1708", "major": "7", "minor": "4"}, "selinux": map[string]interface{}{"config_mode": "permissive", "config_policy": "targeted", "current_mode": "permissive", "enabled": true, "enforced": false, "policy_version": "28"}}, Certname: "foobar.puppetlabs.net", Environment: "production"},
	{Name: "gid", Value: "root", Certname: "foobar.puppetlabs.net", Environment: "production"},
}