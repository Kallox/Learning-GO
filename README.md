# Learning GO

GO es una tecnologia inspirada en C pero con una sintaxys mas simple y menor curva de aprendizaje.

Hay dos tipos de arquitectura:

Monoliticos: Proyecto enorme donde tenemos todo. Escalabilidad vertical.
Microservicios: Separar las funcionalidades por proyectos. Esto permite una mejor escalabilidad horizontal.

Escalabilidad: Una aplicaciones es escalable cuando es pensada para que crezca en el tiempo a nivel de codigo o infraestructura.

Ventajas de trabajar con microservicios:
- Permite una mejor mantencion del codigo.
- Facilita realizar cambios para un servicio especifico.
- Permite escalar horizontalmente.
- Podemos conectarlos con Kafka.
- Pueden estar hechos con distintos lenguajes segun los requerimientos.

Desventajas:
- Demasiados microservicios puede generar el efecto contrario y dificultar su mantencion.

Caracteristicas de GO:
- Sintaxis simple, rapido y tipado.
- Compilado
- Ideal para grandes cargas
- No orientado a objeto
- Limitado para patrones de diseno
- Aprovecha mucho el paralelismo

- Para bucles solo existe el for
- Obliga a utilizar todas las variables
- Trabaja con punteros
- El scope se resuelve con el nombre(Capital para publico, minusculas para privado)

Todas las dependencias de GO tienen que estar en un repositorio