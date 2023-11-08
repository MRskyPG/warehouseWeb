DROP FUNCTION CHECKLOGIN(_login varchar(20), _password varchar(20));
DROP function search(id_prod int, id_place int, prod_name text, cl_name text,
                     cl_surname text, cl_email text, depart_address text);
DROP function remove_product(prod_id int);
DROP PROCEDURE INSERT_PRODUCT(prod_name text, cl_name text, cl_surname text, dp_address text, _email text);
DROP function get_id_product(prod_name text);
DROP function change_placement(id_prod int);
DROP FUNCTION GET_POSITION();
drop function is_not_full();
DROP function select_products_expired(expire_date date);
DROP sequence client_id_seq;
DROP SEQUENCE prod_id_sequence;
DROP table workers cascade;
DROP table clients cascade;
DROP table warehouse cascade;
DROP table products cascade;
