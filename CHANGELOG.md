# CHANGELOG

## 06/06/2021 Feature Prueba de conexión a spreadsheets

Realice una conexión de prueba para recuperar datos del maestro de clientes.

Utilice como referencia la documentación oficial de la API v4 que ofrece google y implemente una versión modificada del codigo de ejemplo

[https://developers.google.com/sheets/api/quickstart/go](https://developers.google.com/sheets/api/quickstart/go)

Para lograr la conexión el primer paso fue realizar la autenticación. Si bien hay dos modos de autenticación, escoji la autorizacion por auth0 recomendada por google. Esto principalmente por motivos de seguridad, el uso de una APIKey es mas simple pero requiere una exposición publica de los datos aumetnando la vulnerabilidad del sistema.

Auth0 requirio que de de alta el proyecto en google cloud. Por suerte ya existia uno, con algunas configuraciones estandar. Creo que ese proyecto se dio de alta en firebase. Esto significa que tiene algunas limitaciones, que a largo plazo podria afectar. (Recuerdo que si se crea un proyecto con firebase algunas caracteristicas de GCP no van a estar disponibles).

Por el momento es suficiente.

Se dio de alta una nueva credencial, se asignaron permisos a la aplicacion para habilitar Google API v4 y se agregaron nuevas cuentas para permitir el acceso.

Luego procedi a dar de alta el resto del codigo.

Trate de explorar la forma en que se puede oranizar el proyecto siguiendo algunas practicas del desarrollo en go (importaciones y submodulos).

- Go no permite anidamiento de modulos
- Si se quiere acceder a modulos internos, cada nivel del submodulo debe tener el archivo principal, con al menos la primer linea de package

> Para el proximo paso, se construiran dos aplicaciones:
> - la principal que levanta la configuracion existente y se conecta a los servidores de google
> - una secundaria que requiere interfaz de usuario y se usa para logearse con la plataforma de google y establecer algunas configuraciones de inicializacion

Falta definir:
- la arquitectura del proyecto inspirada en DDD
- el arbol de carpetas
- el arbol de dependencias
- el flujo principal

## 05/06/2021 Inicio

Se plantean los pasos a seguir y algunas concideraciones para el abordaje de la solución
