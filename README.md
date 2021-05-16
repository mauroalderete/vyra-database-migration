# vyra-migration-database

Servicio de migracion de datos entre google spreadsheets y una base de datos

## Problema

El sistema de gestión del almacen se basa en una aplicación de AppSheets que construye su lote de datos en hojas de cálculo de Google.

Esto hace que resulte difícil realizar consultas, consultas dinámicas o consultas en tiempo real para conocer el estado del almacen y tomar decisiones.

Esta situación, acompañada de las limitaciones de la plataforma, generan la necesidad de montar un sistema de persistencia mas robusto que permita estructurar la información y operar los diferentes actos comerciales de manera inteligente.

## Propuesta

La solución consiste en un servicio o varios microservicios que permiten la migración y sincronización de los datos disponibles en las hojas de cálculo de Google que son alimentados por el proyecto [alMercadito](https://www.appsheet.com/start/eb2c4693-1fd2-41b0-9980-61d5b5789b40#appName=Maestros-2373988&group=%5B%5D&page=card&sort=%5B%7B%22Column%22%3A%22Nombre%22%2C%22Order%22%3A%22Ascending%22%7D%2C%7B%22Column%22%3A%22Marca%22%2C%22Order%22%3A%22Ascending%22%7D%2C%7B%22Column%22%3A%22Valor%22%2C%22Order%22%3A%22Ascending%22%7D%5D&table=Productos&view=Productos) de AppSheets, a una base de datos.

Este surge proyecto, en principio, para satisfacer la necesidad de realizar consultas dinámicas sobre el lote de información estructura utilizando SQL.

En un futuro se dejará de usar AppSheets para el ingreso de datos, debido a las limitaciones que impiden el crecimiento del sistema.

La base de datos configurará el sistema de persistencia principal del proyecto Vyra.

Para lograrlo, el proyecto de migración deberá cumplir con varias fases:

- Lectura de los datos disponibles en cada hoja de AppSheet
- Modelado de clases fuentes y  clases destinos
- Mecanismos de adaptación de las estructuras de datos de AppSheet y la planteada en la base de datos.
- Adecuación, seguimiento y persistencia de claves principales viejas y nuevas. (Tablas de compatibilidad).
- Volcado de los datos en las nuevas estructuras
- Rastreo de cambios:
    - Agregados nuevos
    - Eliminaciónes
    - Alteraciones de datos
- Persistencia de la conexión de migración.

La base de datos de destino será definida en el proyecto [vyra-database](https://gitlab.com/vyra/vyra-database.git)

La arquitectura de la solución consistira en una serie de microservicios que se ocuparan cada uno de mantener la migración de datos de cada submodulo. Se diseñará un prototipo simple de dominio sobre el cual se ejecutaran los microservicios.
Se tratará de respetar las buenas prácticas en la medida de lo posible, sin embargo, la prioridad será la disponibilidad en corto plazo de los datos sincronizados.
