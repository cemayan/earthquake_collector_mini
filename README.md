# earthquake_collector-mini

-  It pulls earthquake data from http://udim.koeri.boun.edu.tr/zeqmap/xmlt/son24saat.xml and writes it to a topic in **Kafka** if there is a new earthquake.

## Kafka & Zookeeper &  Scraper 

**Start** :

```bash
cd deployment
docker-compose up --build
```

**Stop**:

```bash
cd deployment
docker-compose down -v --rmi all --remove-orphans
```

---

### **Kafka** 

You
You can see kafka posts events with commands to help.
**Bash**:

```bash
docker exec -it kafka_eq  /bin/sh
```

**EartQuake Comsumer**:
```bash
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic earthquake --from-beginning
```

**EartQuake Producer**:

```bash
kafka-console-producer.sh --broker-list localhost:9092 --topic earthquake
```


