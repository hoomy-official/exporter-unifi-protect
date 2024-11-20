# Exporter for UniFi Protect

> ⚠️ Sensors are the only component exposed as metrics. Feel free to contribute to add more.

Exporter for UniFi Protect is a Prometheus exporter designed to help monitor your UniFi Protect setup by exporting
relevant metrics to Prometheus. This enables better integration with Prometheus monitoring solutions, ensuring you have
insights into the performance and status of your UniFi Protect cameras and other devices.

<!-- TOC -->

* [Exporter for UniFi Protect](#exporter-for-unifi-protect)
  * [Overview](#overview)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
    * [Using Docker](#using-docker)
    * [Using Go](#using-go)
  * [Configuration](#configuration)
  * [Usage](#usage)
  * [Integrating with Prometheus](#integrating-with-prometheus)
  * [CLI Reference](#cli-reference)
  * [Contributing](#contributing)
  * [License](#license)

<!-- TOC -->

## Overview

This project facilitates the monitoring of UniFi Protect devices by exposing metrics in a consumable format for
Prometheus. It works by interfacing directly with UniFi Protect APIs to fetch metrics and exposes them via an HTTP
endpoint.

## Prerequisites

* Go (to build and run the project outside Docker)
* Docker (optional, for containerized setup)
* Access to a UniFi Protect setup

## Installation

Clone the repository to get started with the exporter:

```shell
git clone https://github.com/hoomy-official/exporter-unifi-protect.git
cd exporter-unifi-protect
```

### Using Docker

To use the Docker setup:

1. Build the Docker image: `docker build -t exporter-unifi-protect .`
2. Run using Docker Compose (adjust `compose/docker-compose.yml` as needed): `docker-compose up -d`

### Using Go

Alternatively, if you prefer a non-containerized setup:

1. Ensure Go is installed.
2. Build the project: `go build ./cmd`
3. Run the built binary.

## Configuration

Configuration is managed via `.env` and `.envrc` files. Refer to these files for environment variables that can be
adjusted to match your UniFi Protect setup and preferences.

## Usage

Once the exporter is running, it will regularly fetch metrics from your UniFi Protect devices and expose them on an HTTP
endpoint. You can directly access this endpoint to see the raw metrics.

To use the exporter, simply run the following command (after building it, if you're not using Docker):

```shell
./exporter
```

You can use the `--help` flag to get a list of configuration options available:

```shell
./exporter --help
```

## Integrating with Prometheus

To integrate the exporter with Prometheus, add the exporter's HTTP endpoint to your Prometheus configuration file, under
the `scrape_configs` section:

```yaml
scrape_configs:
  - job_name: 'unifi-protect'
    static_configs:
      - targets: [ '<exporter-host>:<port>' ]
```

Replace `<exporter-host>` and `<port>` with the appropriate values for your setup.

## Integrate with Grafana

To integrate the metrics from Prometheus into Grafana, ensure the Prometheus is added as a Datasource. Next you can import the example dashboard from `./compose/grafana/unifi-protect/unifi-protect-dashboard.json`

## CLI Reference

The exporter provides several command-line options for configuring its behavior:

* `--help`: Displays help information.
* `--unifi-address <address>`: Specifies the address of the UniFi Protect server.
* `--unifi-port <port>`: Specifies the port of the UniFi Protect server (default: 443).
* `--unifi-username <username>`: Username for authenticating with UniFi Protect.
* `--unifi-password <password>`: Password for authenticating with UniFi Protect.
* `--metrics-path <path>`: Path under which to expose metrics (default: `/metrics`).
* `--scrape-interval <interval>`: Interval (in seconds) between fetches from UniFi Protect.

## Contributing

We welcome contributions to improve the exporter. Feel free to fork the repository, make changes, and submit a pull
request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - [see the LICENSE file for details](./LICENSE.md).
