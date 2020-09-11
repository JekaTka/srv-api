DROP INDEX CONCURRENTLY IF EXISTS users_hash_idx, users_nickname_idx, users_email_idx, currencies_hash_idx;
DROP TABLE IF EXISTS balances, currencies, users;
