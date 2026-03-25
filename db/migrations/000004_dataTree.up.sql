CREATE TABLE dataTree(
                              id BIGSERIAL PRIMARY KEY,

                              species TEXT NOT NULL,          -- espécie da árvore
                              height DOUBLE PRECISION,        -- altura (metros)

                              latitude DOUBLE PRECISION NOT NULL,
                              longitude DOUBLE PRECISION NOT NULL,

                              diameter DOUBLE PRECISION,      -- opcional
                              age INTEGER,                    -- opcional

                              health TEXT,                    -- ex: "boa", "ruim"
                              status TEXT,                    -- ex: "ativa", "removida"

                              created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);