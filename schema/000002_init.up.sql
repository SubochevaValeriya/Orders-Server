CREATE TABLE ORDERS
(
    order_id serial not null unique,
    data json
);