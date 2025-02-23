// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tracetranslator

import (
	"go.opentelemetry.io/collector/model/pdata"
	conventions "go.opentelemetry.io/collector/translator/conventions/v1.5.0"
)

// Some of the keys used to represent OTLP constructs as tags or annotations in other formats.
const (
	TagMessage = "message"

	TagSpanKind = "span.kind"

	TagError         = "error"
	TagHTTPStatusMsg = "http.status_message"

	TagW3CTraceState = "w3c.tracestate"
)

const (
	// Deprecated: Use conventions.OtelStatusCode
	TagStatusCode = conventions.OtelStatusCode
	// Deprecated: Use conventions.OtelStatusDescription
	TagStatusMsg = conventions.OtelStatusDescription
)

// Constants used for signifying batch-level attribute values where not supplied by OTLP data but required
// by other protocols.
const (
	ResourceNoServiceName = "OTLPResourceNoServiceName"
)

// OpenTracingSpanKind are possible values for TagSpanKind and match the OpenTracing
// conventions: https://github.com/opentracing/specification/blob/main/semantic_conventions.md
// These values are used for representing span kinds that have no
// equivalents in OpenCensus format. They are stored as values of TagSpanKind
type OpenTracingSpanKind string

const (
	OpenTracingSpanKindUnspecified OpenTracingSpanKind = ""
	OpenTracingSpanKindClient      OpenTracingSpanKind = "client"
	OpenTracingSpanKindServer      OpenTracingSpanKind = "server"
	OpenTracingSpanKindConsumer    OpenTracingSpanKind = "consumer"
	OpenTracingSpanKindProducer    OpenTracingSpanKind = "producer"
	OpenTracingSpanKindInternal    OpenTracingSpanKind = "internal"
)

// StatusCodeFromHTTP takes an HTTP status code and return the appropriate OpenTelemetry status code
// See: https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/http.md#status
func StatusCodeFromHTTP(httpStatusCode int) pdata.StatusCode {
	if httpStatusCode >= 100 && httpStatusCode < 399 {
		return pdata.StatusCodeUnset
	}
	return pdata.StatusCodeError
}
