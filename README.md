# PiltiSmart Stack Catalog

Official centralized microservice, IoT, and DevOps stack catalog repository for the PiltiSmart Enterprise CLI (`pilti`).

The `pilti` CLI pulls dynamic compose manifests and configuration blueprints directly from this repository during deployment (`pilti install <stack-id>`).

---

## 🚀 Available Stacks & Tools

| Stack ID | Name | Version | Dependencies | Description |
|---|---|---|---|---|
| [`tb-db`](stacks/tb-db/) | **TimescaleDB / PostgreSQL** | `pg17` | *None* (Install 1st) | Telemetry & relational time-series database |
| [`tb-app`](stacks/tb-app/) | **ThingsBoard Core Application** | `3.8.1` | **`tb-db`** | ThingsBoard IoT orchestrator & server |
| [`tb-edge`](stacks/tb-edge/) | **ThingsBoard Edge Gateway** | `3.9.1EDGE` | **`tb-app`** | Remote autonomous IoT Edge instance |
| [`jenkins`](stacks/jenkins/) | **Jenkins CI/CD Automation** | `lts` | *None* | CI/CD build automation controller |
| [`piltiservices`](stacks/piltiservices/) | **PiltiSmart Microservices** | `v7.10.7` | *None* | Specialized API microservices backend |
| [`kafka`](stacks/kafka/) | **Apache Kafka Broker** | `4.1.1` | *None* | KRaft distributed event streaming broker |
| [`pilticloud`](stacks/pilticloud/) | **PiltiSmart Cloud Gateway** | `v8.4.41` | *None* | Hybrid cloud sync tunnel (PMX) |

---

## ⚡ Installation Workflow & Dependency Hierarchy

ThingsBoard components must be deployed in prerequisite order:

1. **Step 1: Deploy Database (`tb-db`)**
   ```bash
   pilti install tb-db
   ```
2. **Step 2: Deploy ThingsBoard Core (`tb-app`)**
   ```bash
   pilti install tb-app
   # Note: If tb-db is not running, pilti will abort with a dependency error!
   ```
3. **Step 3: Deploy ThingsBoard Edge (`tb-edge`)**
   ```bash
   pilti install tb-edge
   # Note: Requires tb-app to be deployed first!
   ```

### Other Standalone Tools:
```bash
pilti install jenkins
pilti install kafka
pilti install piltiservices
pilti install pilticloud
```

---

## 🔍 Pre-Flight Diagnostics & Health
```bash
# Verify Docker, compose, RAM, disk, and network ports:
pilti doctor

# Check live container statuses:
pilti list
pilti status
```
