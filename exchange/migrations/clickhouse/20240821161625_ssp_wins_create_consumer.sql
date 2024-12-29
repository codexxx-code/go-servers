-- +goose Up
-- +goose StatementBegin
-- +goose ENVSUB ON
CREATE TABLE IF NOT EXISTS ssp_wins_consumer
(
    msg String
)
    ENGINE = Kafka(
                      '${CLICKHOUSE_KAFKA_BOOTSTRAP_SERVERS}', -- Брокеры
                      '${SSP_WINS_TOPIC_NAME}', -- Топик
                      '${CLICKHOUSE_KAFKA_CONSUMER_GROUP}', -- ConsumerGroupId
                      'JSONAsString'
                  )
    SETTINGS kafka_num_consumers = ${CLICKHOUSE_KAFKA_NUM_CONSUMERS};
-- +goose ENVSUB OFF
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ssp_wins_consumer;
-- +goose StatementEnd
