package tenanter

import (
	"os"
	"testing"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
)

var (
	aliTenant []Tenanter
	tcTenant  []Tenanter
	hwTenant  []Tenanter
	awsTenant []Tenanter
)

func TestMain(m *testing.M) {
	err := LoadCloudConfigs("../../config.yaml")
	if err != nil {
		panic(err)
	}

	if aliTenant, err = GetTenanters(pbtenant.CloudProvider_ali_cloud); err != nil {
		panic("get aliTenant failed")
	}
	if tcTenant, err = GetTenanters(pbtenant.CloudProvider_tencent_cloud); err != nil {
		panic("get tcTenantr failed")
	}
	if hwTenant, err = GetTenanters(pbtenant.CloudProvider_huawei_cloud); err != nil {
		panic("get hwTenant failed")
	}
	if awsTenant, err = GetTenanters(pbtenant.CloudProvider_aws_cloud); err != nil {
		panic("get awsTenant failed")
	}
	os.Exit(m.Run())
}
