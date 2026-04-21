CREATE TABLE articles
(
    id            SERIAL PRIMARY KEY,
    category_id   INT          REFERENCES categories (id) ON DELETE SET NULL,
    title         VARCHAR(255) NOT NULL,
    slug          VARCHAR(255) NOT NULL UNIQUE,
    summary       TEXT,
    content       TEXT         NOT NULL,
    cover_image   VARCHAR(500),
    status        VARCHAR(20)  NOT NULL DEFAULT 'draft',
    published_at  TIMESTAMPTZ,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    search_vector TSVECTOR GENERATED ALWAYS AS (
        setweight(to_tsvector('simple', coalesce(title, '')), 'A') ||
        setweight(to_tsvector('simple', coalesce(summary, '')), 'B') ||
        setweight(to_tsvector('simple', coalesce(content, '')), 'C')
        ) STORED
);

CREATE INDEX idx_articles_search_vector ON articles USING GIN (search_vector);
CREATE INDEX idx_articles_status ON articles (status);
CREATE INDEX idx_articles_published_at ON articles (published_at DESC);
CREATE INDEX idx_articles_category_id ON articles (category_id);