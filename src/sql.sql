CREATE TABLE ranking(id SERIAL PRIMARY KEY, nome VARCHAR(100), pontos INTEGER, email VARCHAR(100));

INSERT INTO ranking('Lucas', 10, 'lucas@teste.com');

SELECT * FROM ranking;