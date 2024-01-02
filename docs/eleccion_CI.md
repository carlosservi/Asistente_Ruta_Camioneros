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
- TeamCity CI.- Es propiedad de JBrains, voy a probar la opción cloud para que siempre esté disponible. Se integra perfectamente con github y se puede usar con cualquier tecnología. Sólo tengo un periodo de prueba de 14 días. Existe la version local que es gratuita pero dependerá de tenerlo encendido para que se ejecute por lo que por el momento voy a probar la versión cloud. 

## Pruebas
Por el momento descarto Circle CI ya que por el momento no puede usar la API Github Checks por lo que no me sirve, voy a probar las otras 3 opciones y decidiré.

- Githubs Actions.- He configurado un github actions para que ejecute los test automaticamente cada vez que hago un push.
- Semaphore.- Me he registrado en Semaphore, lo he conectado con el repo de github, he creado el yml y lo he conectado mediante la interfaz gráfica.
- TeamCity.- Me he conectado a github a través de TeamCity y se conecta muy fácil y es muy fácilmente configurable, simplemente añades el proyecto de github que quieras, y luego configuras una nueva tarea de compilación. Yo he creado una que ejecuta go test con la versión 1.20 y 1.21. Luego creo un trigger para que cada vez que hago un push se realicen las pruebas oportunas. Todo ello mediante una interfaz gráfica muy amigable y fácil.

## Elección CI tool
Voy a usar Githubs Action porque es muy fácil y rápido, además Semaphore solo te da 14 días gratis para usar su producto y para que me sirva para más tiempo me quedo con Githubs Actions.
