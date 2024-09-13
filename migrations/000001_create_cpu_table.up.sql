CREATE TABLE IF NOT EXISTS cpu(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    current_situation FLOAT
);