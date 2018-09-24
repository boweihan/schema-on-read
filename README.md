# SCHEMA-ON-READ

This project is an on-going experiment to set up an infrastructure for schema-on-read analytics.

### Pipeline

Ideal Properties for a Pipeline include:
* Low Event Latency
* Scalability
* Interative Querying
* Versioning
* Monitoring
* Testing
* Fault Tolerance

Types of Data
* Raw Data (JSON blob)
* Processed Data (Schema-applied)
* Cooked Data (Processed data aggregated/summarized)

Eras
* Flat File Era
* Database Era
* Data Lake Era
* Serverless Era

Event Sourcing
* Stream Processing System (Kafka, Amazon Kinesis, Google PubSub)
* Message Encoding (JSON/Protocol Buffers)
* Handle Delivery Failure, Queueing, Batching, Prioritization
* Handle Auditing (lightweight sequential counting is probably enough)
* Strip PII for GDPR

### Implementation

1. Accept events via API
2. Send to PubSub
3. Ingest PubSub messages using DataFlow
4. Group DataFlow events into fixed batches, convert to String, output to AVRO files on Storage
5. In parallel, pull DataFlow events / perform ETL / send to BigQuery
6. Autoscaling

### Notes

Data gathering
* Tracking specifications as a best practice (Conditions/Properties/Definitions)
* Server vs Client tracking (Trusted Source/Ad Blocking/Versioning/Testing/Data Availability)
 
