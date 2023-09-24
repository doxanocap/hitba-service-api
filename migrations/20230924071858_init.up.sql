CREATE TABLE "services" (
    "id" SERIAL PRIMARY KEY,
    "alias" varchar,
    "name_key" varchar unique,
    "price_per_unit" int,
    "description_key" varchar,
    "created_at" timestamp default now(),
    "updated_at" timestamp
);

CREATE TABLE "service_tariffs" (
    "id" SERIAL PRIMARY KEY,
    "service_id" int,
    "limit" int,
    "limitation_type" varchar,
    "price" int DEFAULT 0,
    "auto_pay" boolean,
    "is_active" boolean,
    "created_at" timestamp default now(),
    "updated_at" timestamp
);

CREATE TABLE "purchased_services" (
    "id" bigint PRIMARY KEY,
    "user_id" bigint,
    "tariff_id" int,
    "remaining_limit" int,
    "created_at" timestamp default now(),
    "expire_at" timestamp
);

ALTER TABLE "service_tariffs" ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id");
ALTER TABLE "purchased_services" ADD FOREIGN KEY ("tariff_id") REFERENCES "service_tariffs" ("id");
