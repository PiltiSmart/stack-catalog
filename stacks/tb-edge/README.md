# ThingsBoard Edge Gateway (`tb-edge`)

Autonomous local edge computing gateway for remote deployments.

## Installation Order: 3rd
**Prerequisite:** Requires `tb-app` (and `tb-db`) to be deployed first to establish cloud routing keys and RPC host connectivity.

## Ports
- **8082**: ThingsBoard Edge Web UI (mapped from container 8080).
- **1884**: ThingsBoard Edge MQTT Broker (mapped from container 1883).
- **5683-5688/udp**: CoAP protocol ports.
