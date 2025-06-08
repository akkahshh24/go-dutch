CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar(50) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "groups" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "description" varchar(255),
  "created_by" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "members" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int NOT NULL,
  "group_id" int NOT NULL,
  "joined_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "expenses" (
  "id" SERIAL PRIMARY KEY,
  "group_id" int NOT NULL,
  "paid_by" int NOT NULL,
  "amount" decimal(10,2) NOT NULL,
  "description" varchar(255),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "balances" (
  "id" SERIAL PRIMARY KEY,
  "lender" int NOT NULL,
  "borrower" int NOT NULL,
  "group_id" int NOT NULL,
  "amount" decimal(10,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "members" ("user_id");

CREATE INDEX ON "members" ("group_id");

CREATE INDEX ON "expenses" ("group_id");

CREATE INDEX ON "expenses" ("paid_by");

CREATE INDEX ON "balances" ("lender", "group_id");

CREATE INDEX ON "balances" ("borrower", "group_id");

ALTER TABLE "groups" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("paid_by") REFERENCES "users" ("id");

ALTER TABLE "balances" ADD FOREIGN KEY ("lender") REFERENCES "users" ("id");

ALTER TABLE "balances" ADD FOREIGN KEY ("borrower") REFERENCES "users" ("id");

ALTER TABLE "balances" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");