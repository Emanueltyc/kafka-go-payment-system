CREATE TABLE IF NOT EXISTS payment.payments (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    order_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL,
    currency CHAR(3) NOT NULL,
    amount NUMERIC(12, 2) NOT NULL,
    payment_method VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION payment.set_updated_at() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_set_updated_at BEFORE UPDATE ON payment.payments FOR EACH ROW EXECUTE FUNCTION payment.set_updated_at();

CREATE INDEX IF NOT EXISTS idx_payments_order_id ON payment.payments (order_id);
CREATE INDEX IF NOT EXISTS idx_payments_user_id ON payment.payments (user_id);
CREATE INDEX IF NOT EXISTS idx_payments_status ON payment.payments (status);