package tags

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Terraform Plugin Framework variants of tags schemas.

func TagsAttribute() tfsdk.Attribute {
	return tfsdk.Attribute{
		Type:     types.MapType{ElemType: types.StringType},
		Optional: true,
	}
}

func TagsAttributeComputedOnly() tfsdk.Attribute {
	return tfsdk.Attribute{
		Type:     types.MapType{ElemType: types.StringType},
		Computed: true,
	}
}

var (
	Null    = types.Map{ElemType: types.StringType, Null: true}
	Unknown = types.Map{ElemType: types.StringType, Unknown: true}
)
