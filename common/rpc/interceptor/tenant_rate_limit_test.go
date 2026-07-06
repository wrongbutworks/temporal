package interceptor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/server/common/headers"
)

func TestErrTenantRateLimitExceeded_CauseAndScope(t *testing.T) {
	require.Equal(t, enumspb.RESOURCE_EXHAUSTED_CAUSE_TENANT_RPS_LIMIT, ErrTenantRateLimitExceeded.Cause)
	require.Equal(t, enumspb.RESOURCE_EXHAUSTED_SCOPE_NAMESPACE, ErrTenantRateLimitExceeded.Scope)
	require.Equal(t, "tenant rate limit exceeded", ErrTenantRateLimitExceeded.Message)
}

func TestNoopTenantRateLimitInterceptor_AllowReturnsNil(t *testing.T) {
	i := NewNoopTenantRateLimitInterceptor()
	require.NoError(t, i.Allow(nil, "SomeAPI", headers.NewGRPCHeaderGetter(context.Background())))
}
