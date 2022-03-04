-- auto-generated definition
create table casbin_rule
(
    id     int          not null,
    p_type varchar(255) not null,
    v0     varchar(255) not null,
    v1     varchar(255) not null,
    v2     varchar(255) not null,
    v3     varchar(255) not null,
    v4     varchar(255) not null,
    v5     varchar(255) not null
);

insert into community_dev.casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5)
values  (2, 'p', 'community_r_1', '/api/:id/visitor', 'GET', null, null, null),
        (3, 'p', 'community_c_2', '/api/:id/visitor', 'POST', null, null, null),
        (4, 'p', 'community_u_3', '/api/:id/visitor', 'PUT', null, null, null),
        (5, 'p', 'community_d_4', '/api/:id/visitor', 'DELETE', null, null, null),
        (52, 'g', 'admin', 'community_r_1', null, null, null, null),
        (53, 'g', 'admin', 'community_c_2', null, null, null, null),
        (54, 'g', 'admin', 'community_u_3', null, null, null, null),
        (55, 'g', 'admin', 'community_d_4', null, null, null, null),
        (56, 'g', 'ckx', 'admin', null, null, null, null);

