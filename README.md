# Sistema para Optimización de ruta para Camioneros

## Indice

- Problema
- Resolución
- Tecnología

***Problema***
Por ley los camioneros tienen que hacer una serie de paradas para descansar durante el trayecto que recorren y muchas de las veces por culpa de no organizar y planificar bien la ruta, por no tener en cuenta varias variables que se deben tener en cuenta, quedan mal estacionados y otras muchas veces no les da tiempo a llegar a su destino a tiempo, con lo que el conductor tiene que quedar esperando hasta el día siguiente para entregar la mercancia con la perdida económica y de tiempo que conlleva.

***Resolución***
Teniendo en cuenta el problema, podemos crear un sistema informático que indicandole el número de ruta y la hora de llegada requerida podemos optimizar que las paradas obligatorias, que son cada 4 horas y media una parada de 45 min y cada 9 horas, 11 horas de descanso ininterrumpidas, sean los idóneos para que les de tiempo a llegar a tiempo al destino, indicando el sistema la hora de salida óptima y las paradas en estaciones de servicio idóneas.


### Información Adicional

#### OBJETIVO 0
[Configuración Objetivo 0](docs/objetivo0.md)

#### OBJETIVO 1
Se ha creado un User Journey especificando el usuario que usará la aplicación.
[User Journey](docs/user_journey.md)

En las historias de usuario hemos desgranado los deseos que espera el usuario del sistema.
[Historias de Usuario](docs/historias_de_usuario.md)

En los milestones hemos declarado los hitos que debemos cumplir para presentar cada MVP exitosamente teniendo en cuenta las historias de usuario.
[Milestones](docs/milestones.md)

#### OBJETIVO 3
En el fichero [Eleccion Gestores](docs/eleccion_gestor_tareas_y_dependencias.txt) está incluido lo que he tenido en cuenta para tomar la decisión de elegir un gestor de dependencias y un gestor de tareas.

He elegido el gestor de dependencias estándar de Golang que es Go Modules ya que funciona automáticamente y es su estándar.
Para el gestor de tareas he indagado sobre varios gestores y he decidido probar dos, Task y Mafefile y después de realizar las pruebas pertinentes, los dos funcionan muy bien y a pesar de que makefile esta muy extendido, la facilidad del archivo .yaml de Task y lo bien que funciona me hace decantarme por TASK.
En la tarea de realizar las pruebas he usado el fichero route.go que incluye algunos struct y algunas funciones para comprobar su funcionamiento y van genial los dos.

Comando para instalar Task en Ubuntu

```bash
sudo snap install task --classic
```

Makefile no necesita instalación.

Introduciendo make help o task help te muestra los comandos necesarios muy sencillos. Por ejemplo make build o task build compila el proyecto e ingresa el bin en la carpeta ./bin

Además ese archivo bin lo he incluido en gitignore para que no lo introduzca en el repositorio ya que es muy fácil generarlo.

Para realizar la comprobación **check** 
```bash
task check
```
