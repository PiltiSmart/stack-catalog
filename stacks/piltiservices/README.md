# PiltiSmart Microservices Backend Stack

PiltiSmart core backend API services integrated with Infisical secret management.

## Ports
- **80 (HTTP / Web API)**: Main service API endpoint.

## Configuration
- ./.piltiservices.env: Infisical project secrets configuration.
- ./logs: Persistent application logs.

## Healthcheck
Configured to poll http://127.0.0.1:80/pilti/piltiUrls every 10s.
