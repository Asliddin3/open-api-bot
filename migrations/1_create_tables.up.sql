CREATE TABLE IF NOT EXISTS "users"(
    "id" serial PRIMARY KEY,
    "tg_id" INT PRIMARY KEY,
    "tg_name" VARCHAR(50) NOT NULL,
    "step" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE Table if NOT EXISTS "user_search"(
    "user_id" BIGINT,
    
);
