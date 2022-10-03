Para resolver esta actividad se utilizo golang como lenguaje principal en el cual se crearon varias carpetas las cuales son:

-api
-db
-handlers
-middlew
-models
-routers

La base de datos que se uso es mongoDB se estan usando de momento mi base de datos personal ya contiene el usuario y contraseña si no
se puede conectar cambiar a una base de datos de mongo a la que tengan acceso esta se encuentra el db/conexionDB.go clienteOptions agregar el usuario de la base de datos y la contraseña del usuario en la variable  y crear la base de datos pruebatecnica y la coleccion DataRepository o bien cambiar la base de datos y la coleccion en los archivos:

-eliminarArchivos.go
-nuevoArchivo.go
_verArchivos.go

para que conecte con la base de datos y la coleccion deseada para probar.

En las cuales se agregaron diferentes archivos en los cuales cada uno tiene sus funciones las cuales son creadas y llamadas en otras funciones para poder traer la informacion necesaria.

En la carpeta api se encuentra solo un archivo el cual sirve para poder conectarse con la api de github y despues regresar la informacion de requerida en por las especificaciones de la prueba.

En la carpeta db se crearon cuatro archivos los cuales se encargan de diferentes cosas se explica a continuacion:

-conexionDB.go: En este archivo estan las funciones que permiten conectarse con la base de datos y para checar si se tiene conexion con la base de datos.

-eliminarArchivos.go: En este archivo se encuentra la funcion para eliminar todos los registros de la base de datos, por si al correr la api se conecta con la base de datos que ya tiene informacion borrar la y checar que si agrega datos a una base de datos sin registros.

-nuevoArchivo.go: En este archivo esta la funcion para agregar los registros a la base de datos, en esta parte ya que habia que buscar en varias paginas y despues filtrar por las etiquetas se opto por usar la opcion de InsertMany de mongoDB para evitar meter mas ciclos e insertar de uno por uno.

-verArchivos.go: En este archivo se encuentra la funcion para leer los registros de la base de datos e im´rimirlos en la terminal uno por uno con un formato especifico.

En la carpeta handlers hay un archivo llamado handlers.go el cual se encargara de establecer las rutas de los endpoints a llamar ademas de agregar el puerto a usar una vez que se levante el servidor aqui se puede cambiar el puerto a usar si es necesario, tiene como puerto el 8080.

En la carpeta middle esta el archivo checarBD el cual checara cada vez que se llame a un endpoint si se tiene aun conexion con la base de datos.

La carpeta models tiene dos archivos modelArchivo.go y responseModel.go los cuales son estructuras para la informacion que se recibe del api de github y para la los registros a insertar en la base de datos y que sea mas facil ordenar y guardar la informacion.

La carpeta de routers tiene tres archivos que son:

-agregarArchivo.go: Este contiene la funcion para agregar archivos a la base de datos la cual llama a la funcion de la carpeta db nuevoArchivo.go la cual contiene la conexion a la base de datos y la coleccion en donde se guardara la informacio.

-eliminarArchivosRouter.go: Este contiene la funcion para eliminar todos los archivos a la base de datos la cual llama a la funcion de la carpeta db eliminarArchivos.go la cual contiene la conexion a la base de datos y la coleccion de donde se desea eliminar los registros.

-verArchivosRouter.go: Este contiene la funcion para leer los registros de la base de datos la cual llama a la funcion de la carpeta db verArchivos.go la cual contiene la conexion a la base de datos y la coleccion de donde se desean ver los registros.

Por ultimo tenemos los archivos docker-compose.yml y Dockerfile los cuales son necesarios para levantar nuestro contenedor de Docker.

Instrucciones:

Puede que sea necesario cambiar los imports de las rutas de todos los archivos para asi llamar a las funciones en otras carpetas locales como por ejemplo:

-github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db

-Checar bien el import de la ruta de los handlers en el main.go la cual es github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/handlers ya que habeces causa problemas si es asi intente con github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/Handlers

Cambiar esta ruta por la ruta en donde se guarde la api.

Correr con el comando go run main.go

Al momento de correr hace treinta llamados a la api de github ya que busca en diferentes paginas diferentes issues si desea puede modificar la cantidad de llamadas en la carpeta api en la conexionAPIGithub.go en el primer ciclo for i := 1; i <= 30; i++ puede aumentarlo a no mas de 60 o disminuirlo si asi lo desea

Si sale algun error por favor realice los siguientes pasos:

-Una vez checado que el ruteo esta bien ver si se tiene conexion al cluster de mongo de Ricardo ya que tiene permitido por el momento entrar desde cualquier direccion ip pero si no se tiene exito en la conexion cambiar la base de datos de mongo a una que puedan usar, esto lo pueden hacer en:

--github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/db/conexionDB.go en la variable clientOptions solo es necesario cambiar el usuario y contraseña

-Si se cambia cluster de mongo a usar es necesario crear una base de datos que se llame pruebatecnica junto con la coleccion DataRepository si ya se cuenta con una base de datos y una coleccion a usar por favor añadir esas conexiones en los archivos:

-eliminarArchivos.go
-nuevoArchivo.go
-verArchivos.go

Que se encuentran en la carpeta db.

Si no hay errores probar en postman las siguientes rutas con los siguentes metodos:

-/agregarArchivo Metodo POST
-/verArchivos METODO GET
-/eliminarArchivos METODO DELETE

nota: La ruta final se vera algo asi localhost:port/eliminarArchivos

nota2: El port es el puerto por donde se corre el servidor en mi caso uso el 8080 si lo cambia escriba el puerto correspondiente.

nota3: Si se cambia el port debe cambiarse en los siguientes archivos:

-handlers.go en la carpeta handlers la variable a cambiar es PORT

-docker-compose.yml en la parte que dice port

-Dockerfile en la variable de EXPOSE

Si se desea probar con docker en vez de usar el comando go run main.go usar el comando docker compose up para que funcione con docker.

Despues de usar el comando docker compose up pruebe con postman como se especifico anteriormente.

El codigo para esta prueba esta en mi github que esta en https://github.com/RicardoRomeroMedina?tab=repositories o si no lo uede ver le dejo el link https://github.com/RicardoRomeroMedina/pruebaTecnicaTredicom del repositorio.