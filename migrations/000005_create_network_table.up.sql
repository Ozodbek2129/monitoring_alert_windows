CREATE TABLE IF NOT EXISTS network(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    current_situation INT
);