CREATE TABLE IF NOT EXISTS order_schema.orders (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL,
    currency CHAR(3) NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    payment_method VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION order_schema.set_updated_at() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_set_updated_at BEFORE UPDATE ON order_schema.orders FOR EACH ROW EXECUTE FUNCTION order_schema.set_updated_at();

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON order_schema.orders (user_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON order_schema.orders (status);