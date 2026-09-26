# ThingsBoard Core Application (`tb-app`)

Enterprise IoT server and device orchestration platform.

## Installation Order: 2nd
**Prerequisite:** Requires `tb-db` (TimescaleDB) to be installed and **RUNNING** first!
If `tb-db` is not running, installation will abort with a dependency error.

## Ports
- **80 / 8080**: ThingsBoard Web UI and REST API.
- **1883**: MQTT device telemetry broker.
- **7070**: Internal RPC communications.
