CREATE TABLE IF NOT EXISTS disk(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    current_situation FLOAT
);