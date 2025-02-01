-- Create "invoices" table
CREATE TABLE "invoices" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "amount" integer NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "base_models" table
CREATE TABLE "base_models" (
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL
);
-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "username" text NULL,
  "first_name" text NULL,
  "last_name" text NULL,
  "phone_number" text NULL,
  "address" text NULL,
  "city" text NULL,
  "province" text NULL,
  "zip_code" text NULL,
  "country" text NULL,
  "is_claimed" boolean NULL DEFAULT false,
  "is_email_verified" boolean NULL DEFAULT false,
  "is_phone_verified" boolean NULL DEFAULT false,
  "tries" bigint NULL DEFAULT 0,
  "lock_expires_at" timestamptz NULL,
  "password_reset_token" text NULL,
  "reset_token_expires_at" timestamptz NULL,
  "password" text NULL,
  "refresh_token" uuid NULL DEFAULT uuid_generate_v4(),
  "refresh_token_expires_at" timestamptz NULL,
  "sign_in_provider" text NULL DEFAULT 'email',
  "provider_id" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_username" UNIQUE ("username")
);
-- Create "collections" table
CREATE TABLE "collections" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "description" text NOT NULL,
  "owner_id" uuid NOT NULL,
  "unit_id" uuid NOT NULL,
  "amount" numeric NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "complaints" table
CREATE TABLE "complaints" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "title" text NULL,
  "description" text NULL,
  "reporter_id" uuid NULL,
  "unit_group_id" uuid NULL,
  "is_resolved" boolean NULL,
  "resolved_at" text NULL,
  "response" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "contracts" table
CREATE TABLE "contracts" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "reports" table
CREATE TABLE "reports" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "description" text NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "files" table
CREATE TABLE "files" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "uploader_id" uuid NULL,
  "bucket_name" text NULL,
  "provider" text NOT NULL,
  "file_type" text NOT NULL,
  "public_url" text NULL,
  "private_url" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "fundraisings" table
CREATE TABLE "fundraisings" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "amount" numeric NOT NULL,
  "description" text NOT NULL,
  "owner_id" uuid NOT NULL,
  "unit_group_id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "notifications" table
CREATE TABLE "notifications" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "user_id" uuid NULL,
  "is_read" boolean NULL DEFAULT false,
  "message" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "inspections" table
CREATE TABLE "inspections" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "unit_group_id" uuid NOT NULL,
  "owner_id" uuid NOT NULL,
  "assigned_to" uuid NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "is_complete" boolean NULL DEFAULT false,
  "is_complete_at" timestamp NULL,
  "details" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "maintenances" table
CREATE TABLE "maintenances" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "unit_group_id" uuid NULL,
  "unit_id" uuid NULL,
  "type" text NOT NULL,
  "comment" character varying(255) NULL,
  "start_date" timestamptz NOT NULL,
  "end_date" timestamptz NOT NULL,
  "is_done" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "chk_maintenances_end_date" CHECK (end_date > start_date)
);
-- Create "unit_groups" table
CREATE TABLE "unit_groups" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" text NOT NULL,
  "owner_id" uuid NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_unit_groups_owner" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "assemblies" table
CREATE TABLE "assemblies" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" character varying(255) NOT NULL,
  "owner_id" uuid NOT NULL,
  "unit_group_id" uuid NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "active" boolean NULL DEFAULT true,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_assemblies_owner" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_assemblies_unit_group" FOREIGN KEY ("unit_group_id") REFERENCES "unit_groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "budgets" table
CREATE TABLE "budgets" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "owner_id" uuid NOT NULL,
  "unit_group_id" uuid NOT NULL,
  "amount" numeric NOT NULL,
  "year" bigint NOT NULL,
  "is_provisional" boolean NULL DEFAULT true,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_budgets_owner" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_budgets_unit_group" FOREIGN KEY ("unit_group_id") REFERENCES "unit_groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "messages" table
CREATE TABLE "messages" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "content" text NULL,
  "from_id" uuid NULL,
  "to_id" uuid NULL,
  "is_read" boolean NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_messages_from_user" FOREIGN KEY ("from_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_messages_to_user" FOREIGN KEY ("to_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "permissions" table
CREATE TABLE "permissions" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "user_id" uuid NULL,
  "entity_id" uuid NULL,
  "entity_name" text NOT NULL,
  "role" smallint NOT NULL DEFAULT 1,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_permissions" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "features" table
CREATE TABLE "features" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "subscriptions" table
CREATE TABLE "subscriptions" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "user_id" uuid NULL,
  "subscription_type" text NULL,
  "expires_at" timestamptz NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_subscription" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "subscription_features" table
CREATE TABLE "subscription_features" (
  "subscription_id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "feature_id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  PRIMARY KEY ("subscription_id", "feature_id"),
  CONSTRAINT "fk_subscription_features_feature" FOREIGN KEY ("feature_id") REFERENCES "features" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_subscription_features_subscription" FOREIGN KEY ("subscription_id") REFERENCES "subscriptions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "units" table
CREATE TABLE "units" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "name" text NOT NULL,
  "type" text NULL DEFAULT 'a',
  "unit_group_id" uuid NULL,
  "is_enabled" boolean NULL DEFAULT true,
  "owner_id" uuid NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_units_owner" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_units_unit_group" FOREIGN KEY ("unit_group_id") REFERENCES "unit_groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "resolutions" table
CREATE TABLE "resolutions" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "unit_group_id" uuid NULL,
  "status" smallint NULL DEFAULT 0,
  "percentage_required" bigint NULL DEFAULT 51,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "votes" table
CREATE TABLE "votes" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "user_id" uuid NULL,
  "resolution_id" uuid NULL,
  "is_approved" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_resolutions_votes" FOREIGN KEY ("resolution_id") REFERENCES "resolutions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
