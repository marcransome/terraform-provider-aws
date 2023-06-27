package flex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ExpandFrameworkStringMap(ctx context.Context, set types.Map) map[string]*string {
	if set.IsNull() || set.IsUnknown() {
		return nil
	}

	var m map[string]*string

	if set.ElementsAs(ctx, &m, false).HasError() {
		return nil
	}

	return m
}

func ExpandFrameworkStringValueMap(ctx context.Context, set types.Map) map[string]string {
	if set.IsNull() || set.IsUnknown() {
		return nil
	}

	var m map[string]string

	if set.ElementsAs(ctx, &m, false).HasError() {
		return nil
	}

	return m
}

// FlattenFrameworkStringValueMapLegacy has no Plugin SDK equivalent as schema.ResourceData.Set can be passed string value maps directly.
// A nil map is converted to an empty (non-null) Map.
func FlattenFrameworkStringValueMapLegacy(_ context.Context, m map[string]string) types.Map {
	elems := make(map[string]attr.Value, len(m))

	for k, v := range m {
		elems[k] = types.StringValue(v)
	}

	return types.MapValueMust(types.StringType, elems)
}
