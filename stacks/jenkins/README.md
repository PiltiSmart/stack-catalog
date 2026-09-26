# Jenkins CI/CD Automation Stack

Enterprise Jenkins automation controller engine with Docker-out-of-Docker host socket mount.

## Ports
- **80 (Web UI)**: Forwarded to container port 8080.
- **50000 (Agent Listener)**: Inbound JNLP agent connection port.

## Volumes
- ./data: Jenkins persistent home directory (/var/jenkins_home).
- /var/run/docker.sock: Mounted from host for Docker build pipelines.
