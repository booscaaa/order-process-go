CREATE TABLE product (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  name varchar(100) NOT NULL,
  description  varchar(500),
  price numeric(15,2)
);