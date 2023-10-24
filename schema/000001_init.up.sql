create table if not exists workers(
                                      user_id serial PRIMARY KEY,
                                      name varchar(20) not null,
    surname varchar(20) not null,
    login varchar(20) not null unique,
    password varchar(20) not null
    );
insert into workers (name, surname, login, password)
VALUES ('andrei', 'degtyarev', 'aboba', 'aboba'),
       ('mikhail', 'rogalsky', 'vizzcon', 'vizzcon');

CREATE FUNCTION CHECKLOGIN(_login varchar(20), _password varchar(20))
    RETURNS BOOLEAN
AS $func$
SELECT EXISTS (
    SELECT 1
    FROM workers w
    WHERE w.login = _login
      AND w.password = _password
);
$func$ LANGUAGE sql;