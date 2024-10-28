CREATE TABLE "users" (
                         "id" INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
                         "username" varchar(30) UNIQUE NOT NULL,
                         "password_hash" varchar(255) UNIQUE NOT NULL,
                         "role_id" int UNIQUE NOT NULL DEFAULT 0,
                         "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "food" (
                        "id" INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
                        "name" varchar(50) UNIQUE NOT NULL,
                        "calories" int NOT NULL DEFAULT 0,
                        "fat" int NOT NULL DEFAULT 0,
                        "carbs" int NOT NULL DEFAULT 0,
                        "protein" int NOT NULL DEFAULT 0,
                        "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "meals" (
                         "id" INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
                         "name" varchar(50) UNIQUE NOT NULL,
                         "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "roles" (
                         "id" INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
                         "name" varchar(20) UNIQUE NOT NULL,
                         "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "logs" (
                        "id" INT GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY NOT NULL,
                        "user_id" int NOT NULL,
                        "food_id" int NOT NULL,
                        "meal_id" int NOT NULL DEFAULT 0,
                        "quantity" int NOT NULL DEFAULT 0,
                        "created_at" timestamp DEFAULT (now())
);

ALTER TABLE "logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("food_id") REFERENCES "food" ("id");

ALTER TABLE "logs" ADD FOREIGN KEY ("meal_id") REFERENCES "meals" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");