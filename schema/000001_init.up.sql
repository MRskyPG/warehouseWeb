create table if not exists clients (
                                       id serial PRIMARY KEY,
                                       name text,
                                       surname text,
                                       email text
);
create table if not exists warehouse (
                                         id_placement serial PRIMARY KEY,
                                         product_id int default null
);
-- drop table warehouse cascade ;

INSERT INTO warehouse (product_id)
SELECT null
FROM generate_series(1, 15) AS g (id);

create table if not exists products (
                                        id serial PRIMARY KEY,
                                        id_placement int unique not null ,
                                        id_product int not null,
                                        product_name text,
                                        registration_date date not null,
                                        client_id int not null,
                                        depart_address text);

create table if not exists workers(
                                      user_id serial PRIMARY KEY,
                                      name varchar(20) not null,
    surname varchar(20) not null,
    login varchar(20) not null unique,
    password varchar(20) not null
    );


create sequence if not exists prod_id_sequence
increment 1
minvalue 1
start 1
cycle;

create sequence if not exists client_id_seq
increment 1
minvalue 1
start 1
cycle;


CREATE OR REPLACE FUNCTION GET_POSITION()
    RETURNS integer AS $$
BEGIN
RETURN (
    SELECT res FROM (
                        select id_placement as res, ROW_NUMBER() OVER (ORDER BY id_placement) AS id from warehouse
                        where product_id is null
                    ) as data WHERE data.id = 1);
END;
$$ LANGUAGE plpgsql;

create or replace function change_placement(id_prod int) returns int as $$
    declare
pos int = get_position();
begin
if (pos is not null)
then
update warehouse
set product_id = null where product_id = id_prod;
update products
set id_placement = pos where id = id_prod;
return pos;
else
return -1;
end if;
end;
    $$ language plpgsql;
----------------------------------------------------------------------------



create or replace function get_id_product(prod_name text) returns integer as $$
begin
        if exists (select 1 from products where product_name = prod_name)
        then
            return (select id_product from products where product_name = prod_name);
else
            return nextval('prod_id_sequence');
end if;
end;
$$ language plpgsql;


create or replace function is_not_full() returns boolean
    as $$
begin
        if (get_position() is not null )
        then return True;
        else return False;
end if;
end;
    $$ language plpgsql;


CREATE OR REPLACE procedure INSERT_PRODUCT(prod_name text, cl_name text, cl_surname text, dp_address text, _email text)
    AS $$
        declare
pos integer = get_position();
            id_prod integer = get_id_product(prod_name);
            cl_id int;
            cur_id int = nextval('products_id_seq');
BEGIN
            if (pos is null)
            then
                return;
end if;
            if (not exists(select 1 from clients where name = cl_name and surname = cl_surname and clients.email = _email))
            then
                cl_id = nextval('client_id_seq');
insert into clients values (cl_id, cl_name, cl_surname, _email);
else cl_id = (select cl.id from clients as cl where cl.name = cl_name and cl.surname = cl_surname and cl.email = _email);
end if;
insert into products (id, id_placement, id_product, product_name, registration_date, client_id, depart_address)
values (cur_id, pos, id_prod, prod_name, current_date, cl_id, dp_address);
update warehouse
set product_id = cur_id where id_placement = pos;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE function remove_product(prod_id int) returns bool
    AS $$
        declare
BEGIN
delete from products as pd where pd.id = prod_id;
update warehouse
set product_id = null where warehouse.product_id = prod_id;
return True;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE PROCEDURE delete_products_expired(expire_date date)
    AS $$
        declare
BEGIN
select remove_product(id) from products where ((expire_date - interval '1 month') >= registration_date);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE function select_products_expired(expire_date date) RETURNS table (id_uniq int, id_placement int, name text) 
    AS $$
select pd.id, pd.id_placement, pd.product_name from products as pd where ((expire_date - interval '1 month') >= registration_date);
$$ LANGUAGE sql;





CREATE OR REPLACE FUNCTION search(id_prod int default -1, id_place int default -1,
                                  prod_name text default 'not_exist', cl_name text default 'not_exist',
                                  cl_surname text default 'not_exist', cl_email text default 'not_exist',
                                  dp_address text default 'not_exist')
    RETURNS table (id_uniq int, id_placement int, name text) 
AS $$
        select pd.id, pd.id_placement, pd.product_name from (products as pd join clients as cl on pd.client_id = cl.id)
                where (pd.id_product = id_prod or id_prod = -1 )and
                      (pd.id_placement = id_place or id_place = -1) and
                      (pd.product_name = prod_name or prod_name = 'not_exist') and
                      (cl.name = cl_name or cl_name = 'not_exist') and
                      (cl.surname = cl_surname or cl_surname = 'not_exist') and
                      (cl.email = cl_email or cl_email = 'not_exist') and 
                      (pd.depart_address = dp_address or dp_address = 'not_exist');
$$ LANGUAGE SQL;

CREATE OR REPLACE  FUNCTION CHECKLOGIN(_login varchar(20), _password varchar(20))
  RETURNS BOOLEAN
AS $func$
SELECT EXISTS (
    SELECT 1
    FROM workers w
    WHERE w.login = _login
      AND w.password = _password
);
$func$ LANGUAGE sql;

insert into workers (name, surname, login, password)
VALUES ('andrei', 'degtyarev', 'aboba', 'aboba'),
       ('mikhail', 'rogalsky', 'vizzcon', 'vizzcon'),
       ('admin', 'admin', 'admin', 'admin');

call insert_product('Яблоки', 'Михаил', 'Рогальский', 'ул.Карла Маркса 53, Новосибирск', 'example@gmail.com');
call insert_product('Набор инструментов', 'Михаил', 'Рогальский', 'ул.Карла Маркса 53, Новосибирск', 'example@gmail.com');
call insert_product('Перфоратор', 'Михаил', 'Рогальский', 'ул.Карла Маркса 53, Новосибирск', 'example@gmail.com');
call insert_product('Велосипед', 'Андрей', 'Дегтярев', 'ул.Карла Маркса 40, Новосибирск', 'other_example@gmail.com');