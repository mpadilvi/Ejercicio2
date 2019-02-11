# Ejercicio2 (API REST)

Pequeña Api Rest que permite crear, modificar, eliminar y listar diversos recursos Hosting, que tienen como atributos: 
* Id
* Nombre
* Cores
* Memoria
* Disco

## Cómo empezar

Descargar la imagen del [repositorio](https://cloud.docker.com/repository/docker/mpadilvi/ej2_apirest) de Docker Hub.
```
docker push mpadilvi/ej2_apirest:tagname
```
La imagen expone el puerto 8000 por lo que para crear un contenedor por ejemplo:
```
docker run -p 8080:8000 mpadilvi/ej2_apirest
```
## Uso
Una vez montado el contenedor y si se escoge el mismo puerto que en ejemplo anterior puede accederse a la API desde http://localhost:8080/hosting , donde estan listados los recursos dados por defecto.

La aplicación acepta los siguientes métodos:
* ```GET```: este método tiene dos modos. En el primero **se listan todos** los recursos:
  ```
  curl -X "GET" localhost:8080/hosting
  ```
  Por otro lado puede escogerse **mostrar un solo recurso**:
  ```
  curl -X "GET" localhost:8080/hosting/{id}
  ```
  donde ```{id}``` debe cambiarse por el Id del recurso que se desea listar.
  
* ```POST```: con este método se consigue **añadir** un recurso.
  ```
  curl -d '{"Id":{id},"Nombre":{name},"Cores":{cores},"Memoria":{mem},"Disco":{disk}}' -X "POST" localhost:8080/hosting/{id}
  ```
  Donde ```{id}``` corresponde al Id del recurso que se quiere añadir y todo lo que está entre ```{}``` debe añadirse como **strings**.

* ```PATCH```: este método **modifica** un recurso **ya añadido**. 
  ```
  curl -d '{"Id":{id},"Nombre":{name},"Cores":{cores},"Memoria":{mem},"Disco":{disk}}' -X "PATCH" localhost:8000/hosting/{id}
  ```
  Donde ```{id}``` corresponde al Id del recurso que se quiere modificar y todo lo que está entre ```{}``` debe añadirse como **strings**. En este caso no es necesario escribir todos los atributos para efectuar los cambios, con escribir el atributo a modificar basta.
  ```
  curl -d '{"Disco":{disk}}' -X "PATCH" localhost:8000/hosting/{id}
  ```
  En este ejemplo solo se modifica el atributo Disco y todo lo demás queda igual.
  
* ```DELETE```: este método **borra** el recurso que marca el ```{id}```.
  ```
  curl -X "DELETE" localhost:8000/hosting/{id}
  ```

