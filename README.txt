2020-2 San Joaquin
Rodrigo Cayazaya Marin, ROL: 201773538-4
Jean-Franco Zárate, ROL: 201773524-4

La tarea consistía en realizar 4 servicios: cliente, camiones, finanza y logistica. Cada una se encuentra en una carpeta con el mismo nombre.
Dentro de la carpeta "helloworld" se encuentran las funciones utilizadas por el servidor logistica y tambien se encuentran los .proto y .pb.go.
La conexión entre servicios se realizó de la siguiente manera:
	Logisitca es el servidor grpc para clientes y camiones.
	Logisitica es el publisher y finanza es el recibidor para RabbitMQ.

Requisitos:

	-Go
	-Grpc
	-Protocol-Buffer
	-RabbitMQ

Instrucciones:	

	Dentro de cada servicio se encuentra un makefile que se ejecuta de la siguiente manera:
	-Para cliente
	make run para ejecutar.
	-Para camiones
	make run para ejecutar y make clean para limpiar los archivos csv.
	-Para logistica
	make run para ejecutar y make clean para limpiar los archivos csv.
	-Para finanza
	make run para ejecutar y make clean para limpiar los archivos csv.

Consideraciones generales:

	-Se utiliza "protoc --go_out=plugins=grpc:helloworld helloworld/helloworld.proto" para compilar el archivo .proto.
	-Las conexiones con RabbitMQ se encuentran activadas para las máquinas virtuales que nos proporcionaron, sin embargo dentro de logisitica y finanza se pueden ejecutar make activar y make desactivar.
	-Todos los registros que piden en la tarea se encuentran en forma de csv dentro de una carpeta con el mismo nombre.
	-El orden para ejecutar este programa es: primero logistica, luego cliente/camiones y al final finanza.
	-Finanza tendrá un valor distinto dependiendo de la cantidad de paquetes que se encuentran en cola.

Consideraciones al ejecutar:

	-Al ejecutar el cliente, este le pedirá el tipo de envío (pyme/retail) o ingresar una orden de seguimiento.
	-Al ejecutar camiones, este le pedira que ingrese la cantidad de tiempo tanto para la espera del segundo paquete como para el tiempo de envío.
	*Todos los tiempos que se piden deben ser números enteros positivos.
