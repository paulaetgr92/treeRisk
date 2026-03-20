CREATE TABLE data_geojson (
                              id BIGSERIAL PRIMARY KEY,
                              latitude DOUBLE PRECISION NOT NULL,
                              longitude DOUBLE PRECISION NOT NULL,
                              species TEXT,
                              height DOUBLE PRECISION,
                              diameter DOUBLE PRECISION,
                              age INTEGER,
                              health TEXT,
                              status TEXT,
                              created_at TIMESTAMPTZ DEFAULT NOW()
);