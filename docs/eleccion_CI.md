# Criterios elección CI
- Integración con Github.- Necesito que sea compatible con la API de Github Checks
- Gratuito.- Necesito que sea gratuito ya que no quiero invertir en este momento en el proyecto.
- Compatible con el lenguaje.- Necesito que sea compatible con Golang.
- Integración con Docker.- Necesito que sea fácilmente integrable con Docker ya que el proyecto está Dockerizado.
- Este disponible Online.- Necesito que el sistema este disponible siempre para ejecutarse.

## Herramientas a considerar
- Github Actions.- Es una opción integrada en Github, gratuita e ilimitada. Tiene compatibilidad con todos los lenguajes y completamente integrable con Docker, además siempre está disponible. Es muy configurable y adaptable.
- Semaphore CI.- Se puede integrar con Github, no es gratuito pero tiene un periodo de prueba de 14 días. Admite el lenguaje Golang. Es integrable con Docker y está disponible online. Además tiene una interfaz gráfica muy amigable y sencilla para configurarlo.
- Circle CI.- No se puede integrar con la API Checks de Github por el momento (link)[https://circleci.com/docs/enable-checks/], Es compatible con Golang y Docker. Tiene un plan gratuito para un tiempo especifico al mes, por lo que podría ser útil para proyectos en desarrollo. También es Online y está siempre disponible.
- AppVeyor CI.- Es una sólida elección para tu CI, ya que es compatible con la API de GitHub Checks, gratuito y cuenta con soporte para Golang. Además, su integración con Docker es sencilla, permitiendo ejecutar tus contenedores Dockerizados. El plan gratuito permite un número limitado de minutos de construcción por mes y un número limitado de ejecuciones concurrentes
## Pruebas
Voy a probar las siguientes tres opciones:

- Githubs Actions.- He configurado un github actions para que ejecute los test automaticamente cada vez que hago un push.
- Semaphore.- Me he registrado en Semaphore, lo he conectado con el repo de github, he creado el yml y lo he conectado mediante la interfaz gráfica.
- AppVeyor.- Me he registrado en AppVeyor, he configurado el repositorio instalado la aplicación y he incluido un archivo yml de configuración en el proyecto para que cada vez que realice un push pase los test con la version 1.20 y 1.21 de Go.

## Elección CI tool
Voy a usar Githubs Action porque es muy fácil y rápido, además Semaphore solo te da 14 días gratis para usar su producto y para que me sirva para más tiempo me quedo con Githubs Actions. AppVeyor también funciona correctamente y con el plan gratuito es más que suficiente, por lo que lo dejo también conectado.
