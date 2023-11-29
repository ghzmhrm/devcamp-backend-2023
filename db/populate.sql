CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);

INSERT INTO account (id, name, age) VALUES ('1', 'Andi', 20);
INSERT INTO account (id, name, age) VALUES ('2', 'Budi', 21);
INSERT INTO account (id, name, age) VALUES ('3', 'Caca', 22);