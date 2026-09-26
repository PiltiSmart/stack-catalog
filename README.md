# PiltiSmart Stack Catalog

Official centralized microservice and IoT stack catalog repository for the PiltiSmart Enterprise CLI (pilti).

The pilti CLI pulls dynamic compose manifests and configuration blueprints directly from this repository during deployment (pilti install <stack-name>).

---

## 🚀 Available Stacks

| Stack ID | Name | Version | Components |
|---|---|---|---|
| [	b-stack](stacks/tb-stack/) | **ThingsBoard Microservices Stack** | 3.8.1 | TimescaleDB (	b-db), ThingsBoard Core (	b), ThingsBoard Edge (edge-tb) |
| [jenkins](stacks/jenkins/) | **Jenkins CI/CD Automation** | lts | Jenkins LTS Controller (jenkins) |
| [piltiservices](stacks/piltiservices/) | **PiltiSmart Microservices** | 7.10.7 | PiltiSmart Microservices Backend (piltiservices-test) |
| [kafka](stacks/kafka/) | **Apache Kafka Broker** | 4.1.1 | Apache Kafka KRaft Broker & Controller (kafka) |
| [pilticloud](stacks/pilticloud/) | **PiltiSmart Cloud Gateway** | 8.4.41 | PiltiCloud PMX Service (piltiCloud) |

---

## ⚡ Quick Installation via 'pilti' CLI

Install any stack from this catalog using the pilti CLI:

`ash
# Pre-flight health and dependency checks:
pilti doctor

# List available stacks in remote catalog:
pilti catalog

# Install individual tools / stacks:
pilti install jenkins
pilti install piltiservices
pilti install kafka
pilti install pilticloud
pilti install tb-stack

# Inspect running container status:
pilti status
`
