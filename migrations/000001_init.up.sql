create table foo (
    id uuid primary key,
    value text not null
);

create table bar (
    id uuid primary key,
    value text not null,
    foo_id uuid not null,

    constraint fk_foo foreign key (foo_id) references foo(id)
);

create index idx_bar_foo_id on bar(foo_id);

create table baz (
    id uuid primary key,
    value text not null
);

create table foo_baz (
    foo_id uuid not null,
    baz_id uuid not null,

    constraint fk_foo foreign key (foo_id) references foo(id),
    constraint fk_baz foreign key (baz_id) references baz(id),
    primary key (foo_id, baz_id)
);

create index idx_foo_baz_foo_id on foo_baz(foo_id);

create index idx_foo_baz_baz_id on foo_baz(baz_id);
