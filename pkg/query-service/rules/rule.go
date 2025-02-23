package rules

import (
	"context"
	"go.signoz.io/query-service/utils/labels"
	"time"
)

// A Rule encapsulates a vector expression which is evaluated at a specified
// interval and acted upon (currently used for alerting).
type Rule interface {
	ID() string
	Name() string
	Type() RuleType

	Labels() labels.BaseLabels
	Annotations() labels.BaseLabels
	Condition() *RuleCondition
	State() AlertState
	ActiveAlerts() []*Alert

	Eval(context.Context, time.Time, *Queriers) (interface{}, error)
	String() string
	// Query() string
	SetLastError(error)
	LastError() error
	SetHealth(RuleHealth)
	Health() RuleHealth
	SetEvaluationDuration(time.Duration)
	GetEvaluationDuration() time.Duration
	SetEvaluationTimestamp(time.Time)
	GetEvaluationTimestamp() time.Time

	SendAlerts(ctx context.Context, ts time.Time, resendDelay time.Duration, interval time.Duration, notifyFunc NotifyFunc)
}
