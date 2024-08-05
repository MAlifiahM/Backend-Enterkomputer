CREATE DATABASE IF NOT EXISTS enterkomputer;

USE enterkomputer;

CREATE TABLE IF NOT EXISTS products (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    category VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    variant VARCHAR(50),
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO products (category, name, variant, price) VALUES
    -- minuman
    ('minuman', 'Jeruk', 'DINGIN', 12000),
    ('minuman', 'Jeruk', 'PANAS', 10000),
    ('minuman', 'Teh', 'MANIS', 8000),
    ('minuman', 'Teh', 'TAWAR', 5000),
    ('minuman', 'Kopi', 'DINGIN', 8000),
    ('minuman', 'Kopi', 'PANAS', 6000),
    ('minuman', 'EXTRA ES BATU', NULL, 2000),
    -- makanan
    ('makanan', 'Mie', 'GORENG', 15000),
    ('makanan', 'Mie', 'KUAH', 15000),
    ('makanan', 'Nasi Goreng', NULL, 15000),
    -- promo
    ('promo', 'Nasi Goreng + Jeruk Dingin', NULL, 23000);

CREATE TABLE IF NOT EXISTS printers (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO printers (name, description) VALUES
     ('A', 'Kasir'), -- printer kasir
     ('B', 'Dapur'), -- printer dapur ( makanan )
     ('C', 'Bar' ); -- printer bar ( minuman )

CREATE TABLE IF NOT EXISTS orders (
      id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
      table_number VARCHAR(50),
      total_price INT NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_items (
       id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
       order_id INT NOT NULL,
       product_id INT NOT NULL,
       quantity INT NOT NULL,
       price INT NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
       FOREIGN KEY (order_id) REFERENCES orders(id),
       FOREIGN KEY (product_id) REFERENCES products(id)
);