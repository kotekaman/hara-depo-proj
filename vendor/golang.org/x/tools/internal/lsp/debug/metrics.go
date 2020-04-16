// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"golang.org/x/tools/internal/lsp/debug/tag"
	"golang.org/x/tools/internal/telemetry/event"
	"golang.org/x/tools/internal/telemetry/export/metric"
)

var (
	// the distributions we use for histograms
	bytesDistribution        = []int64{1 << 10, 1 << 11, 1 << 12, 1 << 14, 1 << 16, 1 << 20}
	millisecondsDistribution = []float64{0.1, 0.5, 1, 2, 5, 10, 50, 100, 500, 1000, 5000, 10000, 50000, 100000}

	receivedBytes = metric.HistogramInt64{
		Name:        "received_bytes",
		Description: "Distribution of received bytes, by method.",
		Keys:        []event.Key{tag.RPCDirection, tag.Method},
		Buckets:     bytesDistribution,
	}

	sentBytes = metric.HistogramInt64{
		Name:        "sent_bytes",
		Description: "Distribution of sent bytes, by method.",
		Keys:        []event.Key{tag.RPCDirection, tag.Method},
		Buckets:     bytesDistribution,
	}

	latency = metric.HistogramFloat64{
		Name:        "latency",
		Description: "Distribution of latency in milliseconds, by method.",
		Keys:        []event.Key{tag.RPCDirection, tag.Method},
		Buckets:     millisecondsDistribution,
	}

	started = metric.Scalar{
		Name:        "started",
		Description: "Count of RPCs started by method.",
		Keys:        []event.Key{tag.RPCDirection, tag.Method},
	}

	completed = metric.Scalar{
		Name:        "completed",
		Description: "Count of RPCs completed by method and status.",
		Keys:        []event.Key{tag.RPCDirection, tag.Method, tag.StatusCode},
	}
)

func registerMetrics(m *metric.Exporter) {
	receivedBytes.Record(m, tag.ReceivedBytes)
	sentBytes.Record(m, tag.SentBytes)
	latency.Record(m, tag.Latency)
	started.CountInt64(m, tag.Started)
	completed.CountFloat64(m, tag.Latency)
}