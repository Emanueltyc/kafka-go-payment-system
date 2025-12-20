# payment system

Payment system using Go and Kafka (in progress)<br>
Database: Postgres 16

- [ ]  Order Service
  - [ ]  create consumer for payment events
- [ ]  Payment Service
  - [ ]  create gateway mock call

### Clone the repository

```bash
  git clone https://github.com/Emanueltyc/kafka-go-payment-system
  cd kafka-go-payment-system
```

### Configure environment variables

#### Linux/Unix
```bash
  cp .env.example .env
```
#### Windows
```bash
  copy .env.example .env
```

### Run Docker

```bash
  docker compose up --build -d
```
