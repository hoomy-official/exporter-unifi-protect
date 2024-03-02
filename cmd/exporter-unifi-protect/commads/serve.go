package commads

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/exporter-toolkit/web"
	"go.uber.org/zap"

	"github.com/hoomy-official/exporter-unifi-protect/internal"
	"github.com/hoomy-official/go-shared/pkg/buildinfo"
	"github.com/hoomy-official/go-shared/pkg/cmd"
	u "github.com/hoomy-official/go-shared/pkg/net/url"
	"github.com/hoomy-official/go-shared/pkg/zapadapter"
	"github.com/hoomy-official/go-unifi-protect/pkg"
)

type Serve struct {
	Username string   `env:"UNIFI_USERNAME" help:"Username used to authenticate API calls. It should be a local regular user" required:""`
	Password string   `env:"UNIFI_PASSWORD" help:"Password used to authenticate API calls" required:""`
	Host     *url.URL `env:"UNIFI_HOST" help:"Host of the Dream Machine" required:""`

	Timeout          time.Duration     `default:"5s" help:"Max duration for collecting data"`
	MinDetectionSpan time.Duration     `default:"1m" help:"Minimum detection span"`
	Labels           map[string]string `help:"extra labels with value by device ID, for example '123456=room=kitchen'"`

	Web struct {
		ExternalURL     string   `name:"external-url" help:"The URL under which the exporter is externally reachable (for example, if the exporter is served via a reverse proxy). Used for generating relative and absolute links back to the exporter itself. If the URL has a path portion, it will be used to prefix all HTTP endpoints served by the exporter. If omitted, relevant URL components will be derived automatically."`
		RoutePrefix     *string  `name:"route-prefix" help:"Prefix for the internal routes of web endpoints. Defaults to path of --web.external-url."`
		SystemdSocket   bool     `name:"systemd-socket" help:"Use systemd socket activation listeners instead of port listeners (Linux only)."`
		ListenAddresses []string `name:"listen-addresses" default:":9090" help:"Addresses on which to expose metrics and web interface. Repeatable for multiple addresses."`
		ConfigFile      string   `name:"config.file" help:"Path to configuration file that can enable TLS or authentication. See: https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md"`
	} `embed:"" prefix:"web."`

	MetricsReadTimeout       time.Duration `default:"1s"`
	MetricsWriteTimeout      time.Duration `default:"1s"`
	MetricsIdleTimeout       time.Duration `default:"30s"`
	MetricsReadHeaderTimeout time.Duration `default:"2s"`

	externalURL *url.URL
}

func (s *Serve) GetExternalURL() *url.URL {
	if s.externalURL == nil {
		eurl, err := u.ComputeExternalURL(s.Web.ExternalURL, (s.Web.ListenAddresses)[0])
		if err != nil {
			panic(fmt.Errorf("cannot compute external url based on the listening address: %w", err))
		}

		s.externalURL = eurl
	}

	return s.externalURL
}

func (s *Serve) PrefixRoute(routes ...string) string {
	var prefixedRoute string
	var err error

	if s.Web.RoutePrefix != nil {
		prefixedRoute, err = url.JoinPath(*s.Web.RoutePrefix, routes...)
	} else {
		prefixedRoute, err = url.JoinPath(s.GetExternalURL().Path, routes...)
	}

	if err != nil {
		prefixedRoute = ""
	}

	// routePrefix must always be at least '/'.
	prefixedRoute = "/" + strings.Trim(prefixedRoute, "/")

	// routePrefix requires path to have trailing "/" in order
	// for browsers to interpret the path-relative path correctly, instead of stripping it.
	if prefixedRoute != "/" {
		prefixedRoute += "/"
	}

	return prefixedRoute
}

func (s Serve) Run(common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	logger.Info(
		"Starting server...",
		zap.String("name", common.Version.Name()),
		zap.String("version", common.Version.Version()),
		zap.String("commit", common.Version.Commit()),
		zap.String("date", common.Version.Date()),
		zap.String("build-source", common.Version.BuildSource()),
	)

	cl := pkg.NewClient(s.Host, pkg.NewAuth(s.Username, s.Password), logger)
	reg := prometheus.NewRegistry()

	logger.Debug(
		"Registering unifi-protect collector",
		zap.Duration("min-detection-span", s.MinDetectionSpan),
		zap.Duration("timeout", s.Timeout),
	)
	reg.MustRegister(internal.NewCollector(cl, s.MinDetectionSpan, s.Timeout, true))

	logger.Debug("Registering common build info collector")
	reg.MustRegister(buildinfo.NewCollector(common.Version.BuildInfo))

	logger.Debug("Registering golang build info collector")
	reg.MustRegister(collectors.NewBuildInfoCollector())

	logger.Debug("Registering golang collector")
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
	))

	if s.Web.ExternalURL == "" && s.Web.SystemdSocket {
		return fmt.Errorf(
			"cannot automatically infer external URL with systemd socket listener. Please provide --web.external-url",
		)
	}

	// Match Prometheus behavior and redirect over externalURL for root path only
	// if routePrefix is different than "/"
	if s.PrefixRoute() != "/" {
		http.HandleFunc("/", redirectOverExternalURL(s))
	}
	http.Handle(s.PrefixRoute("/metrics"), promhttp.HandlerFor(reg, promhttp.HandlerOpts{EnableOpenMetrics: true}))
	http.HandleFunc(s.PrefixRoute("-", "healthy"), AlwaysHealthy(logger))
	http.HandleFunc(s.PrefixRoute(), StatusPage(logger))

	srv := &http.Server{
		ReadTimeout:       s.MetricsReadTimeout,
		WriteTimeout:      s.MetricsWriteTimeout,
		IdleTimeout:       s.MetricsIdleTimeout,
		ReadHeaderTimeout: s.MetricsReadHeaderTimeout,
	}

	srvc := make(chan error)
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	go func() {
		flagConfig := &web.FlagConfig{
			WebListenAddresses: &s.Web.ListenAddresses,
			WebSystemdSocket:   &s.Web.SystemdSocket,
			WebConfigFile:      &s.Web.ConfigFile,
		}

		logger.Info(fmt.Sprintf("HTTP Server started on %s", s.GetExternalURL()))
		if er := web.ListenAndServe(srv, flagConfig, zapadapter.ZapAdapter("HTTP server", logger)); er != nil {
			defer close(srvc)
			srvc <- er
		}
	}()

	for {
		select {
		case <-term:
			logger.Info("Received SIGTERM, exiting gracefully...")
			return nil
		case er := <-srvc:
			logger.Error("The application did not end gracefully, exiting...", zap.Error(er))
			return fmt.Errorf("unexpected end: %w", er)
		}
	}
}

func redirectOverExternalURL(s Serve) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, s.GetExternalURL().String(), http.StatusFound)
	}
}

func StatusPage(logger *zap.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte(`<html>
<head>
	<title>Unifi Protect Exporter</title></head>
<body>
    <h1>Unifi Protect Exporter</h1>
    <p><a href="-/healthy">Healthy</a></p>
    <p><a href="metrics">Metrics</a></p>
</body>
</html>`))

		if err != nil {
			logger.Error("cannot write", zap.Error(err))
		}
	}
}

func AlwaysHealthy(logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Healthy"))

		if err != nil {
			logger.Error("cannot write", zap.Error(err))
		}
	}
}
