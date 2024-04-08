// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package confmap // import "go.opentelemetry.io/collector/confmap"

import (
	"context"

	"go.uber.org/zap"
)

// ConverterSettings are the settings to initialize a Converter.
// Any Converter should take this as a parameter in its constructor.
type ConverterSettings struct {
	Logger *zap.Logger
}

// Converter is a converter interface for the confmap.Conf that allows distributions
// (in the future components as well) to build backwards compatible config converters.
type Converter interface {
	// Convert applies the conversion logic to the given "conf".
	Convert(ctx context.Context, conf *Conf) error
}
