//nolint:lll
package internal

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	v1 "github.com/hoomy-official/go-unifi-protect/api/v1"
	protect "github.com/hoomy-official/go-unifi-protect/pkg"
)

const (
	microsec = 1000
)

type Collector struct {
	apiClient protect.API
	timeout   time.Duration

	// If true, any error encountered during collection is reported as an
	// invalid metric (see NewInvalidMetric). Otherwise, errors are ignored
	// and the collected metrics will be incomplete. (Possibly, no metrics
	// will be collected at all.) While that's usually not desired, it is
	// appropriate for the common "mix-in" of process metrics, where process
	// metrics are nice to have, but failing to collect them should not
	// disrupt the collection of the remaining metrics.
	reportErrors bool

	minDetectionSpan time.Duration

	sensorInfoGauge              *prometheus.Desc
	temperatureGauge             *prometheus.Desc
	lightGauge                   *prometheus.Desc
	humidityGauge                *prometheus.Desc
	batteryStatusPercentageGauge *prometheus.Desc
	bluetoothSignalStrengthGauge *prometheus.Desc
	bluetoothSignalQualityGauge  *prometheus.Desc
	isUpdatingGauge              *prometheus.Desc
	isDownloadingFWGauge         *prometheus.Desc
	isAdoptingGauge              *prometheus.Desc
	isRestoringGauge             *prometheus.Desc
	isAdoptedGauge               *prometheus.Desc
	isAdoptedByOtherGauge        *prometheus.Desc
	isProvisionedGauge           *prometheus.Desc
	isRebootingGauge             *prometheus.Desc
	isSSHEnabledGauge            *prometheus.Desc
	canAdoptGauge                *prometheus.Desc
	isAttemptingToConnectGauge   *prometheus.Desc
	isMotionDetectedGauge        *prometheus.Desc
	isOpenedGauge                *prometheus.Desc
	isConnectedGauge             *prometheus.Desc
	upSinceGauge                 *prometheus.Desc
	lastSeenGauge                *prometheus.Desc
	connectedSinceGauge          *prometheus.Desc
}

func NewCollector(apiClient protect.API, minDetectionSpan time.Duration, timeout time.Duration, reportError bool) *Collector {
	return &Collector{
		apiClient:    apiClient,
		reportErrors: reportError,
		timeout:      timeout,

		minDetectionSpan: minDetectionSpan,

		sensorInfoGauge:              prometheus.NewDesc("sensor_info", "Sensor info.", []string{"id", "name", "firmwareVersion", "hardwareRevision", "nvr_mac", "brand", "type", "model", "market_name"}, nil),
		temperatureGauge:             prometheus.NewDesc("sensor_temperature_celsius", "Sensor monitor for temperature (input).", []string{"id", "name"}, nil),
		lightGauge:                   prometheus.NewDesc("sensor_light_lux", "Sensor monitor for light (input).", []string{"id", "name"}, nil),
		humidityGauge:                prometheus.NewDesc("sensor_humidity_percentage", "Sensor monitor for humidity (input).", []string{"id", "name"}, nil),
		batteryStatusPercentageGauge: prometheus.NewDesc("sensor_battery_status_percentage", "Sensor battery status.", []string{"id", "name", "is_low"}, nil),
		bluetoothSignalStrengthGauge: prometheus.NewDesc("sensor_bluetooth_signal_strength", "Sensor bluetooth signal strength (input).", []string{"id", "name"}, nil),
		bluetoothSignalQualityGauge:  prometheus.NewDesc("sensor_bluetooth_signal_quality", "Sensor bluetooth signal quality (input).", []string{"id", "name"}, nil),
		isUpdatingGauge:              prometheus.NewDesc("sensor_is_updating", "Sensor IsUpdatingGauge status (input).", []string{"id", "name"}, nil),
		isDownloadingFWGauge:         prometheus.NewDesc("sensor_is_downloading_f_w", "Sensor IsDownloadingFWGauge status (input).", []string{"id", "name"}, nil),
		isAdoptingGauge:              prometheus.NewDesc("sensor_is_adopting", "Sensor IsAdoptingGauge status (input).", []string{"id", "name"}, nil),
		isRestoringGauge:             prometheus.NewDesc("sensor_is_restoring", "Sensor IsRestoringGauge status (input).", []string{"id", "name"}, nil),
		isAdoptedGauge:               prometheus.NewDesc("sensor_is_adopted", "Sensor IsAdoptedGauge status (input).", []string{"id", "name"}, nil),
		isAdoptedByOtherGauge:        prometheus.NewDesc("sensor_is_adopted_by_other", "Sensor IsAdoptedByOtherGauge status (input).", []string{"id", "name"}, nil),
		isProvisionedGauge:           prometheus.NewDesc("sensor_is_provisioned", "Sensor IsProvisionedGauge status (input).", []string{"id", "name"}, nil),
		isRebootingGauge:             prometheus.NewDesc("sensor_is_rebooting", "Sensor IsRebootingGauge status (input).", []string{"id", "name"}, nil),
		isSSHEnabledGauge:            prometheus.NewDesc("sensor_is_ssh_enabled", "Sensor IsSshEnabledGauge status (input).", []string{"id", "name"}, nil),
		canAdoptGauge:                prometheus.NewDesc("sensor_can_adopt", "Sensor CanAdoptGauge status (input).", []string{"id", "name"}, nil),
		isAttemptingToConnectGauge:   prometheus.NewDesc("sensor_is_attempting_to_connect", "Sensor IsAttemptingToConnectGauge status (input).", []string{"id", "name"}, nil),
		isMotionDetectedGauge:        prometheus.NewDesc("sensor_is_motion_detected", "Sensor IsMotionDetectedGauge status (input).", []string{"id", "name", "detected_period"}, nil),
		isOpenedGauge:                prometheus.NewDesc("sensor_is_opened", "Sensor IsOpenedGauge status (input).", []string{"id", "name", "detected_period"}, nil),
		isConnectedGauge:             prometheus.NewDesc("sensor_is_connected", "Sensor IsConnectedGauge status (input).", []string{"id", "name"}, nil),
		upSinceGauge:                 prometheus.NewDesc("sensor_up_since_gauge", "Sensor UpSince status (input).", []string{"id", "name"}, nil),
		lastSeenGauge:                prometheus.NewDesc("sensor_last_seen_gauge", "Sensor LastSeen status (input).", []string{"id", "name"}, nil),
		connectedSinceGauge:          prometheus.NewDesc("sensor_connected_since_gauge", "Sensor ConnectedSince status (input).", []string{"id", "name"}, nil),
	}
}

func (c *Collector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.sensorInfoGauge
	descs <- c.temperatureGauge
	descs <- c.lightGauge
	descs <- c.humidityGauge
	descs <- c.batteryStatusPercentageGauge
	descs <- c.bluetoothSignalStrengthGauge
	descs <- c.bluetoothSignalQualityGauge
	descs <- c.isUpdatingGauge
	descs <- c.isDownloadingFWGauge
	descs <- c.isAdoptingGauge
	descs <- c.isRestoringGauge
	descs <- c.isAdoptedGauge
	descs <- c.isAdoptedByOtherGauge
	descs <- c.isProvisionedGauge
	descs <- c.isRebootingGauge
	descs <- c.isSSHEnabledGauge
	descs <- c.canAdoptGauge
	descs <- c.isAttemptingToConnectGauge
	descs <- c.isMotionDetectedGauge
	descs <- c.isOpenedGauge
	descs <- c.isConnectedGauge
	descs <- c.upSinceGauge
	descs <- c.lastSeenGauge
	descs <- c.connectedSinceGauge
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	var sensors []v1.Sensor
	if err := c.apiClient.V1().Sensors.List(ctx, &sensors); err != nil {
		c.reportError(ch, nil, err)
		return
	}

	for _, sensor := range sensors {
		ch <- prometheus.MustNewConstMetric(c.sensorInfoGauge, prometheus.GaugeValue, 1, sensor.ID, sensor.Name, sensor.FirmwareVersion, sensor.HardwareRevision, sensor.NvrMac, "unifi", sensor.Type, sensor.ModelKey, sensor.MarketName)
		ch <- prometheus.MustNewConstMetric(c.temperatureGauge, prometheus.GaugeValue, sensor.Stats.Temperature.Value, sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.lightGauge, prometheus.GaugeValue, float64(sensor.Stats.Light.Value), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.humidityGauge, prometheus.GaugeValue, float64(sensor.Stats.Humidity.Value), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.bluetoothSignalQualityGauge, prometheus.GaugeValue, float64(sensor.BluetoothConnectionState.SignalQuality), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.bluetoothSignalStrengthGauge, prometheus.GaugeValue, float64(sensor.BluetoothConnectionState.SignalStrength), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isUpdatingGauge, prometheus.GaugeValue, boolToInt(sensor.IsUpdating), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isDownloadingFWGauge, prometheus.GaugeValue, boolToInt(sensor.IsDownloadingFW), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isAdoptingGauge, prometheus.GaugeValue, boolToInt(sensor.IsAdopting), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isRestoringGauge, prometheus.GaugeValue, boolToInt(sensor.IsRestoring), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isAdoptedGauge, prometheus.GaugeValue, boolToInt(sensor.IsAdopted), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isAdoptedByOtherGauge, prometheus.GaugeValue, boolToInt(sensor.IsAdoptedByOther), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isProvisionedGauge, prometheus.GaugeValue, boolToInt(sensor.IsProvisioned), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isRebootingGauge, prometheus.GaugeValue, boolToInt(sensor.IsRebooting), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isSSHEnabledGauge, prometheus.GaugeValue, boolToInt(sensor.IsSSHEnabled), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.canAdoptGauge, prometheus.GaugeValue, boolToInt(sensor.CanAdopt), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isAttemptingToConnectGauge, prometheus.GaugeValue, boolToInt(sensor.IsAttemptingToConnect), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isConnectedGauge, prometheus.GaugeValue, boolToInt(sensor.IsConnected), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.upSinceGauge, prometheus.GaugeValue, float64(sensor.UpSince), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.lastSeenGauge, prometheus.GaugeValue, float64(sensor.LastSeen), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.connectedSinceGauge, prometheus.GaugeValue, float64(sensor.ConnectedSince), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.batteryStatusPercentageGauge, prometheus.GaugeValue, float64(sensor.BatteryStatus.Percentage), sensor.ID, sensor.Name, strconv.FormatBool(sensor.BatteryStatus.IsLow))
		ch <- prometheus.MustNewConstMetric(c.isMotionDetectedGauge, prometheus.GaugeValue, boolToInt(time.Now().Before(time.UnixMicro(sensor.MotionDetectedAt*microsec).Add(c.minDetectionSpan))), fmt.Sprintf("%.0f", c.minDetectionSpan.Seconds()), sensor.ID, sensor.Name)
		ch <- prometheus.MustNewConstMetric(c.isOpenedGauge, prometheus.GaugeValue, boolToInt(time.Now().Before(time.UnixMicro(sensor.OpenStatusChangedAt*microsec).Add(c.minDetectionSpan)) || sensor.IsOpened), fmt.Sprintf("%.0f", c.minDetectionSpan.Seconds()), sensor.ID, sensor.Name)
	}
}

func (c *Collector) reportError(ch chan<- prometheus.Metric, desc *prometheus.Desc, err error) {
	if !c.reportErrors {
		return
	}

	if desc == nil {
		desc = prometheus.NewInvalidDesc(err)
	}

	ch <- prometheus.NewInvalidMetric(desc, err)
}

func boolToInt(updating bool) float64 {
	if updating {
		return 1
	}

	return 0
}
