package jaeger_trace

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
)

func NewJaegerTracer(serviceName string, jaegerHostPort string) (opentracing.Tracer, io.Closer) {

	cfg := &jaegerConfig.Configuration {
		Sampler: &jaegerConfig.SamplerConfig{
			Type  : "const", //固定采样
			Param : 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans           : true,
			LocalAgentHostPort : jaegerHostPort,
		},

		ServiceName: serviceName,
	}

	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
