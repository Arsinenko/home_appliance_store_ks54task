create table Stores (
                        id serial primary key,
                        address text not null,
                        created_at timestamp not null,
                        updated_at timestamp,
                        is_alive bool not null
);

create table Roles (
                       id serial primary key,
                       name varchar(50) not null,
                       created_at timestamp not null
);

create table Accounts(
                         id serial primary key,
                         login varchar(50) not null,
                         password varchar(50) not null,
                         created_at timestamp not null,
                         is_alive bool not null
);

create table Employees
(
    id         serial primary key,
    account_id integer not null references Accounts (id),
    role_id    integer not null references Roles (id),
    created_at timestamp not null,
    is_alive bool not null
);

create table Customers(
                          id serial primary key,
                          account_id integer not null references Accounts(id),
                          balance decimal not null,
                          created_at timestamp not null,
                          is_alive bool not null
);

create table Suppliers(
                          id integer primary key,
                          account_id integer not null references Accounts(id),
                          created_at timestamp not null,
                          is_alive bool not null
);

create table Goods(
                      id serial primary key,
                      article text not null,
                      price decimal not null,
                      name text not null,
                      quantity integer not null,
                      is_alive bool not null
);

create table Goods_Suppliers(
                                id integer primary key,
                                supplier_id integer not null references Suppliers(id),
                                good_id integer not null references Goods(id),
                                created_at timestamp not null,
                                is_alive bool not null
)