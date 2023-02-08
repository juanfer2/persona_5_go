# Configuration
copy-config:
	cp .env.example .env

  cp src/shared/infrastructure/persistence/postgres/config/database_env.yml src/shared/infrastructure/persistence/postgres/config/database_development.yml
  cp src/shared/infrastructure/persistence/postgres/config/database_env.yml src/shared/infrastructure/persistence/postgres/config/database_test.yml
  cp src/shared/infrastructure/persistence/postgres/config/database_env.yml src/shared/infrastructure/persistence/postgres/config/database_staging.yml
