# Purchase registry

La intención de este proyecto fue hacer una api rest utilizando golang, siendo simple, pero utilizando la mayor cantidad de conceptos y conocimientos en el proyecto.
La razón de la elección este proyecto es que hace 2 años con java hice lo mismo, solo que era inexperto y no sabía lo que hacía.


# Tecnologías
Este proyecto está hecho en golang, utiliza uber-fx y gin como base. también utiliza el ORM "GORM",  y también viper.


## Arquitectura del proyecto

Este proyecto se basa principalmente en  3 capas, una de infraestructura, en donde está toda la parte de los datos en sí, la parte de servicios, que es la parte de lógica de negocios, en donde utilizo las interfaces de la capa de datos para tener un bajo nivel de dependencias. Y por último tengo una capa controlar en donde ubico los controles (endpoints, middlewares). seria la capa de presentación. Luego también hay otras carpetas como routes, utilities, models, que son para acompañar las distintas capas.
Sé que el proyecto no es perfecto, paso tiempo desde que lo cree, pero en su momento me pareció bastante correcto y estoy orgulloso de él, ya que fue una buena primera experiencia con una arquitectura bien organizada.
