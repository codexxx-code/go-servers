-- +goose Up
-- +goose StatementBegin
-- +goose ENVSUB ON
CREATE TABLE IF NOT EXISTS ssp_billings_consumer
(
    msg String
)
    ENGINE = Kafka
SETTINGS
    kafka_broker_list = '${CLICKHOUSE_KAFKA_BOOTSTRAP_SERVERS}',
    kafka_topic_list = '${ANALYTIC_SSP_BILLINGS_TOPIC_NAME}',
    kafka_group_name = '${CLICKHOUSE_KAFKA_CONSUMER_GROUP}',
    kafka_format = 'JSONAsString',
    kafka_num_consumers = ${CLICKHOUSE_KAFKA_NUM_CONSUMERS};
-- +goose ENVSUB OFF
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ssp_billings_consumer;
-- +goose StatementEnd
