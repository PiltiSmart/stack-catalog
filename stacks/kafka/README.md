# Apache Kafka Broker Stack

Apache Kafka 4.1.1 running in KRaft mode (no ZooKeeper required).

## Ports
- **9092**: PLAINTEXT Client Broker listener.

## Listeners
- **Internal**: PLAINTEXT://:9092,CONTROLLER://:9093
- **Advertised**: PLAINTEXT://10.70.70.4:9092
