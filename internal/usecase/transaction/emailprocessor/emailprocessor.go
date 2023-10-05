package emailprocessor

import (
	"context"
	"time"
)

type EmailTemplateProcessor interface {
	ProcessEmail(ctx context.Context, fromDate time.Time)
}

type TransactionSummarizer interface {
	WithEmailTemplateProcessor(etp EmailTemplateProcessor)
}
