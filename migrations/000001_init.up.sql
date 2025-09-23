create table foo (
    id serial primary key,
    value text not null
);

create table bar (
    id serial primary key,
    value text not null,
    foo_id serial not null,

    constraint fk_foo foreign key (foo_id) references foo(id)
);

create table baz (
    id serial primary key,
    value text not null
);

create table foo_baz (
    foo_id serial not null,
    baz_id serial not null,

    constraint fk_foo foreign key (foo_id) references foo(id),
    constraint fk_baz foreign key (baz_id) references baz(id),
    primary key (foo_id, baz_id)
);
