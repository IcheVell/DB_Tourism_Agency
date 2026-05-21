CREATE TABLE tourists (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name non_empty_varchar_255 NOT NULL,
    last_name non_empty_varchar_255 NOT NULL,
    middle_name VARCHAR(50),
    sex VARCHAR(10) NOT NULL,
    birth_date DATE NOT NULL,
    user_id BIGINT UNIQUE,

    CONSTRAINT chk_tourists_sex
        CHECK (sex IN ('male', 'female')),

    CONSTRAINT fk_tourists_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
);

CREATE TABLE identity_documents (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    document_type VARCHAR(20) NOT NULL,
    document_series non_empty_varchar_255 NOT NULL,
    document_number non_empty_varchar_255 NOT NULL,
    expiration_date DATE,
    issued_by non_empty_varchar_255 NOT NULL,
    issue_date DATE NOT NULL,
    citizenship non_empty_varchar_255 NOT NULL,
    tourist_id BIGINT NOT NULL,

    CONSTRAINT chk_identity_documents_document_type
        CHECK (document_type IN ('PASSPORT', 'BIRTH_CERTIFICATE', 'INTERNATIONAL_PASSPORT')),

    CONSTRAINT chk_identity_documents_dates
        CHECK (expiration_date IS NULL OR expiration_date > issue_date),

    CONSTRAINT uq_identity_documents_document
        UNIQUE (document_type, document_series, document_number, citizenship),

    CONSTRAINT fk_identity_documents_tourist_id
        FOREIGN KEY (tourist_id)
        REFERENCES tourists(id)
);

CREATE TABLE tourist_groups (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    arrival_date TIMESTAMPTZ NOT NULL,
    departure_date TIMESTAMPTZ NOT NULL,
    name non_empty_varchar_255 NOT NULL UNIQUE,

    CONSTRAINT chk_tourist_groups_arrival_less_departure
        CHECK (arrival_date < departure_date)
);

CREATE TABLE group_members (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tourist_group_id BIGINT NOT NULL,
    tourist_category_id BIGINT NOT NULL,
    tourist_id BIGINT NOT NULL,

    CONSTRAINT uq_group_members_group_tourist
        UNIQUE (tourist_group_id, tourist_id),

    CONSTRAINT fk_group_members_tourist_group_id
        FOREIGN KEY (tourist_group_id)
        REFERENCES tourist_groups(id),

    CONSTRAINT fk_group_members_category_id
        FOREIGN KEY (tourist_category_id)
        REFERENCES tourist_categories(id),

    CONSTRAINT fk_group_members_tourist_id
        FOREIGN KEY (tourist_id)
        REFERENCES tourists(id)
);

CREATE TABLE visas (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    number non_empty_varchar_255,
    destination_country non_empty_varchar_255 NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    submitted_at TIMESTAMPTZ,
    decision_at TIMESTAMPTZ,
    issued_at TIMESTAMPTZ,
    valid_from DATE,
    valid_until DATE,
    tourist_id BIGINT NOT NULL,

    CONSTRAINT chk_visas_status
        CHECK (status IN ('draft', 'submitted', 'rejected', 'issued', 'cancelled', 'expired')),

    CONSTRAINT chk_visas_timing
        CHECK ((submitted_at IS NULL OR created_at <= submitted_at) AND (decision_at IS NULL OR submitted_at IS NULL OR submitted_at <= decision_at) AND (issued_at IS NULL OR decision_at IS NULL OR decision_at <= issued_at)),

    CONSTRAINT chk_visas_valid
        CHECK (valid_from IS NULL OR valid_until IS NULL OR valid_from < valid_until),

    CONSTRAINT fk_visas_tourist_id
        FOREIGN KEY (tourist_id)
        REFERENCES tourists(id),

    CONSTRAINT chk_visas_issued_fields
        CHECK (status <> 'issued' OR (number IS NOT NULL AND issued_at IS NOT NULL AND valid_from IS NOT NULL AND valid_until IS NOT NULL))
);

CREATE TABLE child_companions (
    child_group_member_id BIGINT NOT NULL,
    adult_group_member_id BIGINT NOT NULL,

    CONSTRAINT pk_child_companions
        PRIMARY KEY (child_group_member_id, adult_group_member_id),

    CONSTRAINT fk_child_companions_child_id
        FOREIGN KEY (child_group_member_id)
        REFERENCES group_members(id),

    CONSTRAINT fk_child_companions_adult_id
        FOREIGN KEY (adult_group_member_id)
        REFERENCES group_members(id),

    CONSTRAINT chk_child_companions_not_same_member
        CHECK (child_group_member_id <> adult_group_member_id)
);

