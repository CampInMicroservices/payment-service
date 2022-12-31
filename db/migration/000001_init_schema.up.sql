CREATE TABLE "payments" (
    "id"            BIGSERIAL PRIMARY KEY,
    "booking_id"    BIGINT NOT NULL,
    "price"         DOUBLE PRECISION NOT NULL,
    "paid"          BOOLEAN NOT NULL,
    "created_at"    TIMESTAMP NOT NULL DEFAULT(now())
);