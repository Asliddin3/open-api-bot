CREATE TABLE IF NOT EXISTS "users"(
    "tg_id" INT PRIMARY KEY,
    "tg_name" VARCHAR(50) NOT NULL,
    "step" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE if NOT EXISTS "search_photo"(
    "id" serial PRIMARY KEY,
    "user_id" BIGINT,
    "search_result" TEXT,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    Foreign Key (user_id) REFERENCES users(tg_id)
);
CREATE TABLE if NOT EXISTS "search_video"(
    "id" serial PRIMARY KEY,
    "user_id" BIGINT,
    "search_result" TEXT,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    Foreign Key (user_id) REFERENCES users(tg_id)
);



