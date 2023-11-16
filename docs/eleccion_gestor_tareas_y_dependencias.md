# Elección de Gestor de Tareas

Necesito un gestor de tareas que sea versátil y de fácil manejo. Además necesito que funcione en varios lenguajes de programación entre ellos y fundamentalmente para éste proyecto con Golang, por lo que voy a realizar una comparación de ventajas e inconvenientes y probaré dos para ver cuál me convence más.

## Herramientas para la gestión de tareas en proyectos de Go:

- Makefile: El Makefile es el gestor de tareas más comúnmente utilizado en proyectos de Go. Es una herramienta muy potente y versátil que te permite definir reglas y dependencias de construcción para tu proyecto.
        Ventajas:
        Estándar en muchos proyectos y lenguajes.
        Ampliamente conocido y utilizado en la industria del desarrollo de software.
        Permite una gran flexibilidad y control sobre las tareas de construcción.
        Desventajas:
        Puede tener una sintaxis un poco complicada para personas nuevas en desarrollo.

- Task (https://taskfile.dev/): Task es un gestor de tareas basado en Go que permite definir y ejecutar tareas en un archivo YAML. Es muy fácil de usar y proporciona una sintaxis limpia y sencilla.
        Ventajas:
        Ofrece una sintaxis simple y fácil de leer en archivos YAML.
        Está especialmente diseñado para tareas de construcción y automatización.
        No requiere conocimientos previos de Makefile.
        Desventajas:
        Aunque es popular, puede no ser tan universalmente conocido como make.

- Mage (https://magefile.org/): Mage es una herramienta de construcción para Go que permite escribir scripts de construcción usando Go. Los scripts de Mage son archivos Go regulares con funciones especiales que pueden ser invocadas desde la línea de comandos.
        Ventajas:
        Utiliza Go como lenguaje de scripting, lo que puede ser muy cómodo para los desarrolladores de Go.
        Los scripts son archivos Go regulares, lo que significa que se pueden aprovechar las capacidades de Go.
        Desventajas:
        Puede requerir un poco más de esfuerzo para configurar y aprender en comparación con otras herramientas.

- Just (https://github.com/casey/just): Just es un gestor de tareas muy sencillo que utiliza un archivo YAML para definir tareas. Aunque no está específicamente diseñado para Go, es compatible y puede ser utilizado para automatizar tareas de construcción y prueba.
        Ventajas:
        Ofrece una sintaxis sencilla en archivos YAML.
        Es fácil de configurar y usar.
        Desventajas:
        Aunque es una opción sólida, puede no ser tan conocido como make en la comunidad en general.

## Criterios
- Adopción.- Es buen indicador que un gestor de tareas sea ampliamente usado ya que puede indicar que es un gestor fiable.
- Facilidad.- Es buena elección elegir un gestor de tareas que sea fácil de usar e intuitivo.
- Fácil de configurar
- Lenguaje de uso: Es interesante que el lenguaje de uso sea lo más fácil posible y cuanto más parecido sea al lenguaje humano más fácil será, por ejemplo, Make tiene su sintaxis especial que necesitas conocer y por ejemplo Task se acerca mucho al lenguaje humano y simplemente necesitas poner la sentencias de Go para que funcione.

El más ampliamente usado es makefile, por lo que merece la pena hacer la comprobación con él y por otro lado puede ser interesante probar otro que use un archivo YAML ya que es muy fácil de utilizar, he decidido que voy a usar Task porque ofrece una sintaxis muy simple en un archivo YAML, está creado concretamente para ser un gestor de tareas como lo que necesito, además no usa ningún lenguaje de programación especifico y puedes usarlo con todos ya que simplemente pones la sentencia a ejecutar en un YAML y la ejecuta sin tener que preconfigurar nada. A diferencia de Mage por ejemplo que hace uso de la sinstaxis de Go y necesitas conocer más específicamente para utilizarlo. Just es muy similar a Task por lo que con probar Task es suficiente.

Tras estos datos voy a hacerlo con makefile y con Task y me quedaré con el que más me convenza.

# Elección de Gestor de Dependencias

Existen varios gestores de Dependencias como 'dep', 'glide', o 'govendor' que se utilizan para la gestión de dependencias en Go. Pero ninguna de ellas son parte de Go Oficial y por lo tanto, a veces se generaban algunos errores.

Desde la versión 1.11 de Go implementaron los módulos de Go que mediante el archivo go.mod puedes gestionar las depencias de forma más simple y es oficial de Go.
Por lo que la decisión es la de usar Go Mod para la gestión de dependencias, ya que es el estándar oficial y es muy sencillo.

## Algunas de sus características son:

- Versionado Semántico (SemVer)
- Comandos de gestión simplificados ('go get', 'go mod tidy')
- Descarga automática y almacenamiento en caché de las dependencias necesarias.
