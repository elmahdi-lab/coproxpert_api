-- Create "organizations" table
CREATE TABLE "public"."organizations" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "name" text NULL,
  "is_enabled" boolean NULL DEFAULT true,
  "user_id" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create index "uni_organizations_name" to table: "organizations"
CREATE UNIQUE INDEX "uni_organizations_name" ON "public"."organizations" ("name");
-- Create "complaints" table
CREATE TABLE "public"."complaints" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "title" text NULL,
  "description" text NULL,
  "user_id" uuid NULL,
  "is_resolved" boolean NULL,
  PRIMARY KEY ("id")
);
-- Create "contacts" table
CREATE TABLE "public"."contacts" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "user_id" uuid NULL,
  "phone_number" text NULL,
  "address" text NULL,
  "city" text NULL,
  "province" text NULL,
  "zip_code" text NULL,
  "country" text NULL,
  PRIMARY KEY ("id")
);
-- Create "base_models" table
CREATE TABLE "public"."base_models" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "units" table
CREATE TABLE "public"."units" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "name" text NULL,
  "type" text NULL,
  "building_id" uuid NULL,
  "user_id" uuid NULL,
  "is_enabled" boolean NULL,
  PRIMARY KEY ("id")
);
-- Create index "uni_units_name" to table: "units"
CREATE UNIQUE INDEX "uni_units_name" ON "public"."units" ("name");
-- Create "messages" table
CREATE TABLE "public"."messages" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "content" text NULL,
  "user_id" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create "notifications" table
CREATE TABLE "public"."notifications" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "user_id" uuid NULL,
  "is_read" boolean NULL DEFAULT false,
  "message" text NULL,
  PRIMARY KEY ("id")
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "username" text NULL,
  "password" text NULL,
  "tries" bigint NULL DEFAULT 0,
  "lock_expires_at" timestamptz NULL,
  "is_verified" boolean NULL DEFAULT false,
  "password_reset_token" text NULL,
  "reset_token_expires_at" timestamptz NULL,
  "token" text NULL,
  "token_expires_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "uni_users_username" to table: "users"
CREATE UNIQUE INDEX "uni_users_username" ON "public"."users" ("username");
-- Create "files" table
CREATE TABLE "public"."files" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "user_id" uuid NULL,
  "bucket_name" text NULL,
  "provider" text NOT NULL,
  "file_type" text NOT NULL,
  "public_url" text NULL,
  "private_url" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_files_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "unit_groups" table
CREATE TABLE "public"."unit_groups" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "name" text NOT NULL,
  "organization_id" uuid NULL,
  "user_id" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create "maintenances" table
CREATE TABLE "public"."maintenances" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "unit_group_id" uuid NOT NULL,
  "unit_id" uuid NOT NULL,
  "type" text NOT NULL,
  "comment" character varying(255) NULL,
  "start_date" timestamptz NOT NULL,
  "end_date" timestamptz NOT NULL,
  "is_done" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_unit_groups_maintenances" FOREIGN KEY ("unit_group_id") REFERENCES "public"."unit_groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "chk_maintenances_end_date" CHECK (end_date > start_date)
);
-- Create "permissions" table
CREATE TABLE "public"."permissions" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "user_id" uuid NULL,
  "role" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_permissions" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "features" table
CREATE TABLE "public"."features" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "name" text NULL,
  PRIMARY KEY ("id")
);
-- Create "subscriptions" table
CREATE TABLE "public"."subscriptions" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "organization_id" uuid NULL,
  "subscription_type" text NULL,
  "expires_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "subscription_features" table
CREATE TABLE "public"."subscription_features" (
  "subscription_id" uuid NOT NULL,
  "feature_id" uuid NOT NULL,
  PRIMARY KEY ("subscription_id", "feature_id"),
  CONSTRAINT "fk_subscription_features_feature" FOREIGN KEY ("feature_id") REFERENCES "public"."features" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_subscription_features_subscription" FOREIGN KEY ("subscription_id") REFERENCES "public"."subscriptions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "resolutions" table
CREATE TABLE "public"."resolutions" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "user_id" uuid NULL,
  "organization_id" uuid NULL,
  "unit_group_id" uuid NULL,
  "status" smallint NULL DEFAULT 0,
  "percentage_required" bigint NULL DEFAULT 51,
  PRIMARY KEY ("id")
);
-- Create "votes" table
CREATE TABLE "public"."votes" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "user_id" uuid NULL,
  "resolution_id" uuid NULL,
  "is_approved" boolean NULL DEFAULT false,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_resolutions_votes" FOREIGN KEY ("resolution_id") REFERENCES "public"."resolutions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
