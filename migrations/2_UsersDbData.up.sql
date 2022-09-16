INSERT INTO roles (id, name) VALUES (1, 'admin');
INSERT INTO roles (id, name) VALUES (2, 'user');
INSERT INTO roles (id, name) VALUES (3,'unauthorized');

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/createRole','POST',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/assignRole','POST',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/ListRoles','POST',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/ListPerms/param','GET',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/addPerm','POST',0);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/admin/removePerm','POST',0);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/users/list','POST',0);

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

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouseCells','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouseCells/list','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouseCells/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouseCells/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/warehouseCells/param','GET',1);

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
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/compwarh/param','DELETE',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/images','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/images','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/images/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/images/param','DELETE',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/videos','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/videos','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/videos/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/videos/param','DELETE',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/units','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/units','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/units/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/units/param','DELETE',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/categories','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/category/param','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/categories','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/categories/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/categories/param','DELETE',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/product/param','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products/param','PUT',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products','POST',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products_images','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products_images','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products_images','POST',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products_videos','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products_videos','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/products_videos','POST',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes/param','PUT',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes_values','GET',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes_values/param','DELETE',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes_values','POST',1);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes_values/param','PUT',1);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes_values_products','GET',2);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes_values_products','DELETE',2);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/attributes_values_products','POST',2);

INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/categories_products','GET',2);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/categories_products','DELETE',2);
INSERT INTO permissions (req_path, req_method, req_server_id) VALUES ('/categories_products','POST',2);


INSERT INTO roles_permissions (roles_id, permissions_id)
    SELECT 1,id FROM permissions
