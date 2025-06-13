-- Create stocks table for stock purchase tracker
CREATE TABLE IF NOT EXISTS stocks (
    id SERIAL PRIMARY KEY,
    exchange VARCHAR(10) NOT NULL,
    code VARCHAR(20) NOT NULL,
    state VARCHAR(10) NOT NULL,
    quantity INTEGER NOT NULL,
    currency VARCHAR(3) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_stocks_exchange ON stocks(exchange);
CREATE INDEX IF NOT EXISTS idx_stocks_code ON stocks(code);
CREATE INDEX IF NOT EXISTS idx_stocks_state ON stocks(state);
CREATE INDEX IF NOT EXISTS idx_stocks_created_at ON stocks(created_at);

-- Add some sample data for testing
INSERT INTO stocks (exchange, code, state, quantity, currency, price) VALUES
    ('ASX', 'BHP', 'BUY', 100, 'AUD', 45.50),
    ('NYSE', 'AAPL', 'BUY', 50, 'USD', 150.25),
    ('ASX', 'CBA', 'SELL', 200, 'AUD', 95.75),
    ('NYSE', 'GOOGL', 'BUY', 25, 'USD', 2750.00);

-- Create a function to update the updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create a trigger to automatically update the updated_at column
CREATE TRIGGER update_stocks_updated_at 
    BEFORE UPDATE ON stocks 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column(); 