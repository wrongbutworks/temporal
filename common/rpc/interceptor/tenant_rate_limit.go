package interceptor

import (
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"
	"go.temporal.io/server/common/headers"
	"go.temporal.io/server/common/namespace"
)

var (
	ErrTenantRateLimitExceeded = &serviceerror.ResourceExhausted{
		Cause:   enumspb.RESOURCE_EXHAUSTED_CAUSE_TENANT_RPS_LIMIT,
		Scope:   enumspb.RESOURCE_EXHAUSTED_SCOPE_NAMESPACE,
		Message: "tenant rate limit exceeded",
	}
)

type (
	// TenantRateLimitInterceptor rate-limits incoming Nexus requests by tenant.
	// OSS ships only the no-op; Cloud injects the enforcing implementation via fx.
	TenantRateLimitInterceptor interface {
		Allow(handlerNamespace *namespace.Namespace, apiName string, headerGetter headers.HeaderGetter) error
	}

	noopTenantRateLimitInterceptor struct{}
)

var _ TenantRateLimitInterceptor = (*noopTenantRateLimitInterceptor)(nil)

func NewNoopTenantRateLimitInterceptor() TenantRateLimitInterceptor {
	return &noopTenantRateLimitInterceptor{}
}

func (*noopTenantRateLimitInterceptor) Allow(*namespace.Namespace, string, headers.HeaderGetter) error {
	return nil
}
