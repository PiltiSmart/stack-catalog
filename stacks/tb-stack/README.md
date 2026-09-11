# ThingsBoard Microservices Stack (`tb-stack`)

This stack encapsulates the complete 3-component ThingsBoard IoT ecosystem:

1. **`tb-db`**: TimescaleDB HA (PostgreSQL 17) database storage.
2. **`tb`**: ThingsBoard Core application backend (`piltismartsolutions/thingsboard-3.8.1:v-4.1.2`).
3. **`edge-tb`**: ThingsBoard Edge gateway (`thingsboard/tb-edge:3.9.1EDGE`) with dedicated PostgreSQL.

---

## 🔧 Environment Configuration

All credentials and ports are dynamically parameterized via environment variables:

| Variable | Description | Default |
|---|---|---|
| `TB_HTTP_PORT` | ThingsBoard Core Web UI port | `80` |
| `TB_MQTT_PORT` | ThingsBoard Core MQTT broker port | `1883` |
| `TB_RPC_PORT` | ThingsBoard Core internal RPC port | `7070` |
| `EDGE_WEB_PORT` | ThingsBoard Edge Web UI port | `8082` |
| `EDGE_MQTT_PORT` | ThingsBoard Edge MQTT port | `1884` |
| `TB_DB_PORT` | TimescaleDB host port | `5432` |
| `POSTGRES_PASSWORD` | TimescaleDB password | `qwer1234` |
