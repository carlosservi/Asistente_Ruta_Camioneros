# Elección de imagen de Docker

## Criterios
- Tamaño: Debe ser lo más ligera posible para que se descargue y ejecute lo más rapido posible.
- Herramientas: A poder ser incluya las herramientas necesarias, por ejemplo el lenguaje de programación, siempre que no incluya muchas herramientas no necesarias y con ello provoque perdidas en otros de los criterios citados.
- Seguridad: Que no tenga vulnerabilidades graves.
- Mantenimiento: Tiene que estar actualizada actualmente, eso es señal de que está bien mantenida y resuelven las vulnerabilidades que van surgiendo.

## Opciones a tener en cuenta
- Golang: Es la imagen oficial de Go para Docker, está muy bien mantenida, sin fallos graves de seguridad y es muy ligera. Existe con diferentes versiones de SO subyacente, por ejemplo con Debian o con Alpine. La más ligera es Alpine y suele ser la más usada ya que Alpine es una distribución muy ligera utilizada mucho en contenedores de Docker.

- Alpine: Imagen oficial de Alpine Linux, es muy ligera, ocupando solo 5MB y con la posibilidad de instalar las herramientas necesarias, se actualiza constantemente ya que es muy utilizada para contenedores Docker.

- Otros SO: Podemos elegir entre imagenes oficiales de otros Sistemas Operativos, sería algo similar a Alpine, el problema reside en el tamaño que aunque existen versiones slim (ligeras) normalmente ocupan más tamaño.

## Decisión
Tal y como dicen los criterios de decisión, voy a usar la `imagen oficial de Golang con Alpine` ya que introduce todas las herramientas que necesito, sigue siendo muy liviana, solo 220MB con todas las herramientas instaladas y muy segura por su alto nivel de mantenimiento.

## Descarga Imagen
```bash
docker pull golang:1.21-alpine3.19
```