# PiltiSmart Stack Catalog

Official centralized microservice and IoT stack catalog repository for the PiltiSmart Enterprise CLI (`ps`).

The `ps` CLI pulls dynamic compose manifests and configuration blueprints directly from this repository during deployment (`ps install <stack-name>`).

---

## 📦 Available Stacks

| Stack ID | Name | Version | Components |
|---|---|---|---|
| [`tb-stack`](stacks/tb-stack/) | **ThingsBoard Microservices Stack** | `3.8.1` | TimescaleDB (`tb-db`), ThingsBoard Core (`tb`), ThingsBoard Edge (`edge-tb`) |

---

## 🚀 Quick Installation via 'ps' CLI

Install any stack from this catalog using the compiled `ps` CLI:

```bash
# Pre-flight health and dependency checks:
ps doctor

# Install the ThingsBoard 3-component stack dynamically from GitHub:
ps install tb-stack

# Inspect running stack status:
ps status
```
