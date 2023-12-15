# Elección de imagen de Docker

## Criterios
- Tamaño: Debe ser lo más ligera posible para que se descargue y ejecute lo más rapido posible.
- Herramientas: A poder ser incluya las herramientas necesarias, por ejemplo el lenguaje de programación, siempre que no incluya muchas herramientas no necesarias y con ello provoque perdidas en otros de los criterios citados.
- Seguridad: Que no tenga vulnerabilidades graves.
- Mantenimiento: Tiene que estar actualizada actualmente, eso es señal de que está bien mantenida y resuelven las vulnerabilidades que van surgiendo.

## Opciones a tener en cuenta
- Golang: Es la imagen oficial de Go para Docker, está muy bien mantenida, sin fallos graves de seguridad y es muy ligera. Existe con diferentes versiones de SO subyacente, como Debian, Alpine o Windowsservercore. 
    - Alpine es la más ligera y suele ser la más usada en contenedores de Docker pero tiene detectada una vulnerabilidad muy grave con los certificados ssl.
    - Con Debian se crean varias versiones como bullseye, bookworm pero son muy pesadas, entre 700MB y 800MB.

- Alpine: Imagen oficial de Alpine Linux, es muy ligera, ocupando solo 5MB y con la posibilidad de instalar las herramientas necesarias, se actualiza constantemente ya que es muy utilizada para contenedores Docker. Respecto a la seguridad tiene una vulnerabilidad grave y varias medias.

- Debian: Imagen oficial de Debian Linux, es un poco pesada ocupando 117MB, además tendría que incorporar las herramientas necesarias. Por otro lado no tiene vulnerabilidades graves conocidas y se actualiza muy a menudo. Respecto a la seguridad sólo tiene una Vulnerabilidad media detectada y algunas leves.

- Debian-slim: Existe una imagen de Debian slim que pesa menos ya que incluye el SO más básico, que incorporandole las herramientas necesarias sería mucho mas liviano que la imagen Debian normal. No tiene vulnerabilidades graves y se actualiza constantemente. Respecto a la seguridad mantiene las mismas que la versión normal.

## Decisión
Tal y como dicen los criterios de decisión, voy a usar la `imagen de Debian slim e instalaré las herramientas necesarias` instalaré las herramientas necesarias, ya que es más segura. Al usar la versión slim es muy liviana también.

## Descarga Imagen
```bash
docker pull debian:stable-slim
```