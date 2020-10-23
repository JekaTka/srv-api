CREATE TABLE "users"(
    "id" UUID NOT NULL,
    "nickname" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "users_email_id_index" ON
    "users"("email", "id");
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
CREATE TABLE "currencies"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "currencies_id_name_index" ON
    "currencies"("id", "name");
ALTER TABLE
    "currencies" ADD PRIMARY KEY("id");
CREATE TABLE "balances"(
    "id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "currency_id" UUID NOT NULL,
    "balance" DOUBLE PRECISION NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "balances_user_id_currency_id_id_index" ON
    "balances"("user_id", "currency_id", "id");
ALTER TABLE
    "balances" ADD PRIMARY KEY("id");
CREATE TABLE "balance_history"(
    "id" UUID NOT NULL,
    "balance_id" UUID NOT NULL,
    "amount" DOUBLE PRECISION NOT NULL,
    "balance_before" DOUBLE PRECISION NOT NULL,
    "balance_after" DOUBLE PRECISION NOT NULL,
    "transaction_id" UUID NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "balance_history_balance_id_transaction_id_index" ON
    "balance_history"("balance_id", "transaction_id");
ALTER TABLE
    "balance_history" ADD PRIMARY KEY("id");
CREATE TABLE "transaction_types"(
    "id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "transaction_types_id_name_index" ON
    "transaction_types"("id", "name");
ALTER TABLE
    "transaction_types" ADD PRIMARY KEY("id");
CREATE TABLE "transactions"(
    "id" UUID NOT NULL,
    "transaction_type_id" UUID NOT NULL,
    "amount" DOUBLE PRECISION NOT NULL,
    "currency_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "status" BOOLEAN NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "transactions_id_transaction_type_id_currency_id_user_id_index" ON
    "transactions"(
        "id",
        "transaction_type_id",
        "currency_id",
        "user_id"
    );
ALTER TABLE
    "transactions" ADD PRIMARY KEY("id");
CREATE TABLE "games"(
    "id" UUID NOT NULL,
    "start_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "end_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "balance" DOUBLE PRECISION NOT NULL,
    "total_won" DOUBLE PRECISION NOT NULL,
    "currency_id" UUID NOT NULL,
    "hex_count" INTEGER NOT NULL,
    "logs" jsonb NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX "games_id_currency_id_index" ON
    "games"("id", "currency_id");
ALTER TABLE
    "games" ADD PRIMARY KEY("id");
CREATE TABLE "game_hexes"(
    "id" UUID NOT NULL,
    "game_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "buy_price" DOUBLE PRECISION NOT NULL,
    "won" DOUBLE PRECISION NOT NULL
);
CREATE INDEX "game_hexes_id_game_id_user_id_index" ON
    "game_hexes"("id", "game_id", "user_id");
ALTER TABLE
    "game_hexes" ADD PRIMARY KEY("id");
ALTER TABLE
    "balances" ADD CONSTRAINT "balances_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "game_hexes" ADD CONSTRAINT "game_hexes_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "balance_history" ADD CONSTRAINT "balance_history_balance_id_foreign" FOREIGN KEY("balance_id") REFERENCES "balances"("id");
ALTER TABLE
    "balances" ADD CONSTRAINT "balances_currency_id_foreign" FOREIGN KEY("currency_id") REFERENCES "currencies"("id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_currency_id_foreign" FOREIGN KEY("currency_id") REFERENCES "currencies"("id");
ALTER TABLE
    "games" ADD CONSTRAINT "games_currency_id_foreign" FOREIGN KEY("currency_id") REFERENCES "currencies"("id");
ALTER TABLE
    "transactions" ADD CONSTRAINT "transactions_transaction_type_id_foreign" FOREIGN KEY("transaction_type_id") REFERENCES "transaction_types"("id");
ALTER TABLE
    "balance_history" ADD CONSTRAINT "balance_history_transaction_id_foreign" FOREIGN KEY("transaction_id") REFERENCES "transactions"("id");
ALTER TABLE
    "game_hexes" ADD CONSTRAINT "game_hexes_game_id_foreign" FOREIGN KEY("game_id") REFERENCES "games"("id");