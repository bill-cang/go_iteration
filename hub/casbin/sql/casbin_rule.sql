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

-- rbac 模拟数据
insert into community_dev.casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5)
values  (2, 'p', 'community_r_1', '/api/*/visitor', 'GET', null, null, null),
        (3, 'p', 'community_c_2', '/api/[\w]+/visitor', 'POST', null, null, null),
        (4, 'p', 'community_u_3', '/api/:id/visitor', 'PUT', null, null, null),
        (5, 'p', 'community_d_4', '/api/:id/visitor', 'DELETE', null, null, null),
        (52, 'g', 'admin', 'community_r_1', null, null, null, null),
        (53, 'g', 'admin', 'community_c_2', null, null, null, null),
        (54, 'g', 'admin', 'community_u_3', null, null, null, null),
        (55, 'g', 'admin', 'community_d_4', null, null, null, null),
        (56, 'g', 'ckx', 'admin', null, null, null, null);

-- rbac 带域 模拟数据
insert into community_dev.casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5, ptype)
values  (2, 'p', 'community', 'visitor', '/api/:officeID/visitor', '(GET|LOCK)', null, null, null),
        (3, 'p', 'community', 'visitor', '/api/[\\w]+/visitor', 'POST', null, null, null),
        (4, 'p', 'community', 'visitor', '/api/*/visitor', 'PUT', null, null, null),
        (5, 'p', 'community', 'visitor', '/api/:id/visitor', 'DELETE', null, null, null),
        (52, 'g', 'admin', 'community', 'visitor', null, null, null, null),
        (53, 'g', 'admin', 'community', 'station', null, null, null, null),
        (56, 'g', 'ckx', 'admin', 'visitor', null, null, null, null),
        (57, 'g', 'commander', 'ckx', 'visitor', null, null, '', null),
        (59, 'p', 'community', 'station', '/api/:officeID/station', 'GET', null, null, null),
        (60, 'p', 'community', 'station', '/api/[\\w]+/station', 'POST', null, null, null),
        (61, 'p', 'community', 'station', '/api/:id/station', 'PUT', null, null, null),
        (62, 'p', 'community', 'station', '/api/:id/station', 'DELETE', null, null, null);

