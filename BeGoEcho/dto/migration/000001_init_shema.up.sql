CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "hased_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_verified" bool DEFAULT false,
  "full_name" varchar,
  "phone_float" varchar UNIQUE,
  "role" int NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "short_desc" varchar NOT NULL,
  "desc" varchar,
  "article" varchar,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "discount_applied" bigint DEFAULT null
);

CREATE TABLE "product_variant" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "img_main" varchar NOT NULL,
  "imgs_detail" varchar[] NOT NULL,
  "variant_on_product" bigint NOT NULL,
  "qty" int NOT NULL DEFAULT 0,
  "is_deleted" bool NOT NULL DEFAULT false,
  "price" int NOT NULL DEFAULT 0
);

CREATE TABLE "collections" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "short_desc" varchar NOT NULL,
  "desc" varchar,
  "article" varchar,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "parent" bigint DEFAULT null,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "discounts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "begin_at" timestamptz NOT NULL DEFAULT ('0001-01-01 00:00:00Z'),
  "expire_at" timestamptz NOT NULL DEFAULT ('0001-01-01 00:00:00Z'),
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "type" int DEFAULT null,
  "value" float DEFAULT null
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "summary" float NOT NULL DEFAULT 0,
  "order_status" int NOT NULL DEFAULT 0,
  "customer_name" varchar NOT NULL,
  "customer_phone" varchar NOT NULL,
  "note_from_customer" varchar,
  "response_from_staff" varchar,
  "shipping_address" varchar NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_variants" (
  "id" bigserial PRIMARY KEY,
  "variant" bigint NOT NULL,
  "belong_to_order" bigserial NOT NULL,
  "qty" int NOT NULL DEFAULT 1,
  "summary" float NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "conversations" (
  "id" bigserial PRIMARY KEY,
  "client_id" bigint,
  "client_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "is_deleted" bool NOT NULL DEFAULT false
);

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "content" varchar NOT NULL,
  "belong_to_conversation" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "accounts"."role" IS 'constant: 0-client 1-staff 2-admin';

COMMENT ON COLUMN "products"."short_desc" IS 'simple desc at list and detail';

COMMENT ON COLUMN "products"."desc" IS 'desc describe more detail at detail page - should be rich text';

COMMENT ON COLUMN "products"."article" IS 'an article introduce about product - should be rich text';

COMMENT ON COLUMN "product_variant"."img_main" IS 'img url main img';

COMMENT ON COLUMN "product_variant"."imgs_detail" IS 'list img for gallery';

COMMENT ON COLUMN "collections"."short_desc" IS 'simple desc at list and detail';

COMMENT ON COLUMN "collections"."desc" IS 'desc describe more detail at detail page - should be rich text';

COMMENT ON COLUMN "collections"."article" IS 'an article introduce about collection - should be rich text';

COMMENT ON COLUMN "orders"."order_status" IS 'constant: -1:canceled 0:confirmed 1:pendingPayment 2:shipping 3:shipped';

ALTER TABLE "products" ADD FOREIGN KEY ("discount_applied") REFERENCES "discounts" ("id");

ALTER TABLE "product_variant" ADD FOREIGN KEY ("variant_on_product") REFERENCES "products" ("id");

CREATE TABLE "products_collections" (
  "products_id" bigserial,
  "collections_id" bigserial,
  PRIMARY KEY ("products_id", "collections_id")
);

ALTER TABLE "products_collections" ADD FOREIGN KEY ("products_id") REFERENCES "products" ("id");

ALTER TABLE "products_collections" ADD FOREIGN KEY ("collections_id") REFERENCES "collections" ("id");


CREATE TABLE "products_categories" (
  "products_id" bigserial,
  "categories_id" bigserial,
  PRIMARY KEY ("products_id", "categories_id")
);

ALTER TABLE "products_categories" ADD FOREIGN KEY ("products_id") REFERENCES "products" ("id");

ALTER TABLE "products_categories" ADD FOREIGN KEY ("categories_id") REFERENCES "categories" ("id");


ALTER TABLE "categories" ADD FOREIGN KEY ("parent") REFERENCES "categories" ("id");

ALTER TABLE "order_variants" ADD FOREIGN KEY ("variant") REFERENCES "product_variant" ("id");

ALTER TABLE "order_variants" ADD FOREIGN KEY ("belong_to_order") REFERENCES "orders" ("id");

ALTER TABLE "conversations" ADD FOREIGN KEY ("client_id") REFERENCES "accounts" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("belong_to_conversation") REFERENCES "conversations" ("id");
