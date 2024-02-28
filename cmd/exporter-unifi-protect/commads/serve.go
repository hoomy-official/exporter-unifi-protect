package commads

import (
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/hoomy-official/go-unifi-protect/pkg"

	"github.com/hoomy-official/exporter-unifi-protect/internal"
	"github.com/hoomy-official/go-shared/pkg/buildinfo"
	"github.com/hoomy-official/go-shared/pkg/cmd"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Serve struct {
	Username string   `env:"USERNAME" help:"Username used to authenticate API calls. It should be a local regular user" required:""`
	Password string   `env:"PASSWORD" help:"Password used to authenticate API calls" required:""`
	Host     *url.URL `env:"HOST" help:"Host of the Dream Machine" required:""`

	Timeout          time.Duration     `default:"5s" help:"Max duration for collecting data"`
	MinDetectionSpan time.Duration     `default:"1m" help:"Minimum detection span"`
	Labels           map[string]string `help:"extra labels with value by device ID, for example '123456=room=kitchen'"`

	MetricsAddr string `env:"METRICS_ADDR" default:":8080" help:"TCP address for the server to listen on, in the form \"host:port\"."`
	MetricsPath string `default:"/metrics" help:"Path to receive requests related to metrics"`

	MetricsReadTimeout       time.Duration `default:"1s"`
	MetricsWriteTimeout      time.Duration `default:"1s"`
	MetricsIdleTimeout       time.Duration `default:"30s"`
	MetricsReadHeaderTimeout time.Duration `default:"2s"`
}

func (s Serve) Run(common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	cl := pkg.NewClient(s.Host, pkg.NewAuth(s.Username, s.Password), logger)

	reg := prometheus.NewRegistry()
	reg.MustRegister(internal.NewCollector(cl, s.MinDetectionSpan, s.Timeout, true))
	reg.MustRegister(buildinfo.NewCollector(common.Version.BuildInfo))
	reg.MustRegister(collectors.NewBuildInfoCollector())
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))

	http.Handle(s.MetricsPath, promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	server := http.Server{
		Addr:              s.MetricsAddr,
		ReadTimeout:       s.MetricsReadTimeout,
		WriteTimeout:      s.MetricsWriteTimeout,
		IdleTimeout:       s.MetricsIdleTimeout,
		ReadHeaderTimeout: s.MetricsReadHeaderTimeout,
	}

	return server.ListenAndServe()
}
