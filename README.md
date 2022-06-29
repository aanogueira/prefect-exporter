# Prefect Prometheus exporter

Prometheus exporter exposing Prefect metrics.

## Configuration

The exporter can be configured using env variables or command flags.

| **KEY**            | **Description**                                                                    |
| ------------------ | ---------------------------------------------------------------------------------- |
| `LISTEN_PORT`      | listen on addr:port (default `:8080`), omit addr to listen on all interfaces       |
| `METRICS_PATH`     | path for metrics, default `/metrics`                                               |
| `SCRAPE_DELAY`     | scrape delay in seconds, default `300`                                             |
| `GRAPHQL_ENDPOINT` | Prefect Apollo url (GraphQL endpoint), default `https://prefect-apollo.example.io` |
