create table hoge (
    id uuid primary key,
    value text not null
);

create table cancelled_hoge (
    id uuid primary key,
    value text not null,
    cancelled_at timestamp not null,

    constraint fk_hoge foreign key (id) references hoge(id)
);

create table piyo (
    id uuid primary key,
    value text not null
);

create table processed_hoge (
    id uuid primary key,
    value text not null,
    piyo_id uuid not null,
    processed_at timestamp not null,

    constraint fk_hoge foreign key (id) references hoge(id),
    constraint fk_piyo foreign key (piyo_id) references piyo(id)
);
