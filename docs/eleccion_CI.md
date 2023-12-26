# Criterios elección CI
- Integración con Github.- Necesito que sea compatible con la API de Github Checks
- Gratuito.- Necesito que sea gratuito ya que no quiero invertir en este momento en el proyecto.
- Compatible con el lenguaje.- Necesito que sea compatible con Golang.
- Integración con Docker.- Necesito que sea fácilmente integrable con Docker ya que el proyecto está Dockerizado.
- Este disponible Online.- Necesito que el sistema este disponible siempre para ejecutarse.

## Herramientas a considerar
- Github Actions.- Es una opción integrada en Github, gratuita e ilimitada. Tiene compatibilidad con todos los lenguajes y completamente integrable con Docker, además siempre está disponible. Es muy configurable y adaptable.
- Semaphore CI.- Se puede integrar con Github, no es gratuito pero tiene un periodo de prueba de 14 días. Admite el lenguaje Golang. Es integrable con Docker y está disponible online. Además tiene una interfaz gráfica muy amigable y sencilla para configurarlo.
- Circle CI.- No se puede integrar con Github por el momento, Es compatible con Golang y Docker. Tiene un plan gratuito para un tiempo especifico al mes, por lo que podría ser útil para proyectos en desarrollo. También es Online y está siempre disponible.

## Pruebas
Por el momento descarto Circle CI ya que por el momento no puede usar la API Github Checks por lo que no me sirve, voy a probar las otras 2 opciones y decidiré.

- Githubs Actions.- He configurado un github actions para que ejecute los test automaticamente cada vez que hago un push.
- Semaphore.- Me he registrado en Semaphore, lo he conectado con el repo de github, he creado el yml y lo he conectado mediante la interfaz gráfica.

## Elección CI tool
Voy a usar Githubs Action porque es muy fácil y rápido, además Semaphore solo te da 14 días gratis para usar su producto y para que me sirva para más tiempo me quedo con Githubs Actions.
