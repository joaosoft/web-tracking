
-- migrate up
CREATE SCHEMA IF NOT EXISTS "tracking";

CREATE TABLE "tracking"."category" (
	id_category 		TEXT PRIMARY KEY,
	"name"		        TEXT NOT NULL UNIQUE,
	created_at			TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE "tracking"."action" (
	id_action 		    TEXT PRIMARY KEY,
	"name"		        TEXT NOT NULL UNIQUE,
	created_at			TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE "tracking"."category_action" (
	fk_category 	    TEXT NOT NULL,
	fk_action	        TEXT NOT NULL,
	created_at			TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY(fk_category, fk_action),
	FOREIGN KEY(fk_category) REFERENCES "tracking"."category"(id_category),
	FOREIGN KEY(fk_action) REFERENCES "tracking"."action"(id_action)
);

CREATE TABLE "tracking"."event" (
	id_event 		    TEXT PRIMARY KEY,
	fk_category		    TEXT NOT NULL,
	fk_action  			TEXT NOT NULL,
	label 				TEXT,
	value              	INTEGER,
	latitude            NUMERIC(14, 11),
	longitude           NUMERIC(14, 11),
	country             TEXT,
	city                TEXT,
	street              TEXT,
	meta_data           JSONB,
	created_at			TIMESTAMP NOT NULL DEFAULT NOW(),
	FOREIGN KEY(fk_category) REFERENCES "tracking"."category"(id_category),
	FOREIGN KEY(fk_action) REFERENCES "tracking"."action"(id_action)
);

-- migrate down
DROP TABLE "tracking"."event";
DROP TABLE "tracking"."category_action";
DROP TABLE "tracking"."action";
DROP TABLE "tracking"."category";

DROP SCHEMA "tracking";
