# earthquake_collector-mini

- Deprem verilerini verilen aralıkta http://udim.koeri.boun.edu.tr/zeqmap/xmlt/son24saat.xml adresinden çekerek yeni bir deprem varsa bunu **Kafka**'da bir topic'e(**earthquake**) yazar.


## Kafka & Zookeeper &  Scraper 

**Start** :

```bash
docker-compose up --build
```

**Stop**:

```bash
docker-compose down -v --rmi all --remove-orphans
```

---

### **Kafka** 

Yardımcı olması için aşağıdaki komutlarla kafkaya yazılan eventleri görebilirsiniz.

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


