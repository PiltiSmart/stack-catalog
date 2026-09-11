# PiltiSmart Stack Catalog

Official centralized microservice and IoT stack catalog repository for the PiltiSmart Enterprise CLI (`ps`).

The `ps` CLI pulls dynamic compose manifests and configuration blueprints directly from this repository during deployment (`ps install <stack-name>`).

---

## 📦 Available Stacks

| Stack ID | Name | Version | Components |
|---|---|---|---|
| [`tb-stack`](stacks/tb-stack/) | **ThingsBoard Microservices Stack** | `3.8.1` | TimescaleDB (`tb-db`), ThingsBoard Core (`tb`), ThingsBoard Edge (`edge-tb`) |

---

## ⚡ Install the 'ps' CLI Globally (Linux & macOS)

Install the standalone enterprise `ps` CLI with a single command:

```bash
curl -fsSL https://raw.githubusercontent.com/PiltiSmart/stack-catalog/main/install.sh | bash
```

> See [INSTALL.md](INSTALL.md) for manual download links for Linux, Apple Silicon / Intel Mac, and Windows.

---

## 🚀 Deploying Stacks via 'ps' CLI

Once `ps` is installed, you can orchestrate any stack from this catalog on your nodes:

```bash
# 1. Pre-flight health and dependency checks:
ps doctor

# 2. Discover stacks dynamically from this repository:
ps catalog

# 3. Install and deploy the ThingsBoard 3-component stack:
ps install tb-stack

# 4. Inspect live running container status:
ps status

# 5. Teardown or restart stack:
ps stack down
ps stack restart
```

---

## 🔨 Building the CLI Locally from Source

```bash
cd cli
./build.sh
# Or install globally:
make install
```

