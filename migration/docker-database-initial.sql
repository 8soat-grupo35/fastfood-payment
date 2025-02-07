CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO payments (order_id, status)  VALUES (1, 'AGUARDANDO');
INSERT INTO payments (order_id, status)  VALUES (2, 'RECUSADO');
INSERT INTO payments (order_id, status)  VALUES (3, 'APROVADO');