INSERT INTO roles (id, name) VALUES (1, 'admin');
INSERT INTO roles (id, name) VALUES (2, 'user');
INSERT INTO roles (id, name) VALUES (3,'unauthorized');

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/createRole','POST',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/assignRole','POST',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/ListRoles','POST',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/ListPerms/param','GET',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/addPerm','POST',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/removePerm','POST',0);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/companies','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/companies/list','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/companies/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/companies/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/companies/param','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/companies/param','POST',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouses','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouses/list','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouses/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouses/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouses/param','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouses/param','POST',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehousesCells','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehousesCells/list','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehousesCells/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehousesCells/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehousesCells/param','GET',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/gtd','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/gtd/list','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/gtd/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/gtd/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/gtd/param','GET',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/countries','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/countries/list','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/countries/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/countries/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/countries/param','GET',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/payments','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/payments/list','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/payments/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/payments/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/payments/param','GET',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/documents','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/documents/list','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/documents/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/documents/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/documents/param','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/documents/hold/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/documents/unhold/param','PUT',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/compwarh/param','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/compwarh/param','POST',1);

INSERT INTO roles_permissions (roles_id, permissions_id)
    SELECT 1,id FROM permissions
