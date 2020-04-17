MySQL

Vamos empezando por crear un usuario nuevo desde la consola de MySQL:

CREATE USER 'nombre_usuario'@'localhost' IDENTIFIED BY 'tu_contrasena';
Lamentablemente, a este punto el nuevo usuario no tiene permisos para hacer algo con las bases de datos. Por consecuencia si el usuario intenta identificarse (con la contraseña establecida), no será capaz de acceder a la consola de MySQL.

Por ello, lo primero que debemos hacer es porporcionarle el acceso requerido al usuario con la información que requiere.

GRANT ALL PRIVILEGES ON * . * TO 'nombre_usuario'@'localhost';
Los asteriscos en este comando hacen referencia a la base de datos y la tabla (respectivamente) a la cual el nuevo usuario tendrá acceso; específicamente este comando permite al usuario leer, editar, ejecutar y realizar todas las tareas en todas las bases de datos y tablas.

Una vez que has finalizado con los permisos que deseas configurar para tus nuevos usuarios, hay que asegurarse siempre de refrescar todos los privilegios.

FLUSH PRIVILEGES;
Tus cambios ahora surtirán efecto.

¿Cómo otorgar permisos de usuario diferentes?
Aquí está una pequeña lista del resto de los posibles permisos que los usuarios pueden gozar.

ALL PRIVILEGES: como mencionamos previamente esto permite a un usuario de MySQL acceder a todas las bases de datos asignadas en el sistema.
CREATE: permite crear nuevas tablas o bases de datos.
DROP: permite eliminar tablas o bases de datos.
DELETE: permite eliminar registros de tablas.
INSERT: permite insertar registros en tablas.
SELECT: permite leer registros en las tablas.
UPDATE: permite actualizar registros seleccionados en tablas.
GRANT OPTION: permite remover privilegios de usuarios.
Para proporcionar un permiso a usuario específico, puedes utilizar ésta estructura:

GRANT [permiso] ON [nombre de bases de datos].[nombre de tabla] TO ‘[nombre de usuario]’@'localhost’;
Si deseas darles acceso a cualquier base de datos o tabla, asegurate de insertar un asterisco (8) en lugar del nombre de la base de datos o tabla.

Cada vez que tu actualizas o cambias permisos, asegúrate de refrescar los privilegios mediante FLUSH PRIVILEGES;.

Si necesitas remover un permiso, la estructura es casi idéntica a la que los asigna:

REVOKE [permiso] ON [nombre de base de datos].[nombre de tabla] FROM ‘[nombre de usuario]’@‘localhost’;
Así como puedes borrar bases de datos con DROP, también puedes usar el comando DROP para borrar usuarios:

DROP USER ‘usuario_prueba’@‘localhost’;
Para probar el nuevo usaurio, debes cerrar sesión escribiendo quit y volviendo a iniciar sesión con éste comando en la consola:

mysql -u [nombre de usuario]-p

CREATE TABLE `posts` (
`id` bigint unsigned AUTO_INCREMENT,
`title` varchar(255) NOT NULL UNIQUE,
`content` varchar(255) NOT NULL,`author_id` int unsigned NOT NULL,
`created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP, 
PRIMARY KEY (`id`)
)

CREATE TABLE `users` (
`id` int unsigned AUTO_INCREMENT,
`nickname` longtext,
`email` longtext,
`password` longtext,
`created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP, 
PRIMARY KEY (`id`)
)  
 

CREATE UNIQUE INDEX idx_user_nickname_email ON `users`(`nickname`, `email`)   


Postgres

CREATE TABLE "users" 
(
 "id" bigserial,
 "nickname" text,
 "email" text,
 "password" text,
 "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP, 
 PRIMARY KEY ("id")
)  


CREATE TABLE "posts" 
(
 "id" bigserial,
 "title" varchar(255) NOT NULL UNIQUE,
 "content" varchar(255) NOT NULL,
 "author_id" bigint NOT NULL,
 "created_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 "updated_at" timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY ("id")
)  


Ejecutar Test

$ cd $GOPATH/src/github.com/<username>/fullstack
$ cd tests/modeltest
$ go test -v --run TestFindAllUsers
$ go test -v --run TestUpdateAPost
$ go test -v --run TestSignIn
$ go test -v



Ejecutar la aplicación
$ cd $GOPATH/src/github.com/<username>/fullstack
$ go run main.go	



Detener la aplicación

$ cd $GOPATH/src/github.com/<username>/fullstack
$ lsof -i :8080 | grep LISTEN
$ kill -9 <PID>	
