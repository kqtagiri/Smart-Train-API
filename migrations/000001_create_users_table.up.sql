CREATE TABLE users(
    id SERIAL PRIMARY KEY, 
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    login VARCHAR(50) NOT NULL,
    password VARCHAR(24) NOT NULL,
    balance DECIMAL(9,2),
    UNIQUE(login)
);