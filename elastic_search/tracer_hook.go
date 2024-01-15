package elastic_search

import (
	"io"
	"net/http"

	"github.com/joaosoft/clean-infrastructure/tracer"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type tracerHook struct {
	spanName  string
	transport estransport.Interface
}

func newTracerHook(spanName string, client *elasticsearch.Client) estransport.Interface {
	return &tracerHook{
		spanName:  spanName,
		transport: client.Transport,
	}

}

func (th *tracerHook) Perform(request *http.Request) (*http.Response, error) {
	if request.URL.Path == "/" {
		// productCheck, it should be ignored!
		return th.transport.Perform(request)
	}

	span, ctx := opentracing.StartSpanFromContext(request.Context(), th.spanName)
	defer span.Finish()
	request = request.WithContext(ctx)

	carrier := opentracing.HTTPHeadersCarrier(request.Header)
	_ = opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)

	ext.HTTPMethod.Set(span, request.Method)
	ext.HTTPUrl.Set(span, request.URL.Path)

	if request.Body != nil {
		body, _ := io.ReadAll(request.Body)
		span.SetTag(tracer.TracerTagBody, string(body))
	}

	response, err := th.transport.Perform(request)
	if err != nil {
		ext.Error.Set(span, true)
		span.SetTag(tracer.TracerTagError, err.Error())
	}

	if response != nil {
		ext.HTTPStatusCode.Set(span, uint16(response.StatusCode))
		if response.StatusCode >= http.StatusBadRequest {
			ext.Error.Set(span, true)
			if response.Body != nil {
				body, _ := io.ReadAll(response.Body)
				span.SetTag(tracer.TracerTagBody, string(body))
			}
		}
	}

	return response, nil
}
