create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('Alan Turing', 'Alan Mathison Turing foi um matemático, cientista da computação, lógico, criptoanalista, filósofo e biólogo teórico britânico.'),
('Grace Hopper', 'Grace Murray Hopper foi almirante e analista de sistemas da Marinha dos Estados Unidos nas décadas de 1940 e 1950, criadora da linguagem de programação de alto nível Flow-Matic — base para a criação do COBOL — e uma das primeiras programadoras do computador Harvard Mark I em 1944.');