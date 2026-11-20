# Jenkins local para la entrega 2

Esta configuración prepara un **controlador Jenkins autenticado**, separado de los
dos servicios de la aplicación. No instala agentes, crea jobs ni registra
credenciales automáticamente. El arranque del controlador no demuestra que el
pipeline se haya ejecutado: esa evidencia sigue pendiente hasta autorizar y
configurar un agente con capacidad de construir contenedores.

## Requisitos y decisiones

| Elemento | Requisito o decisión |
| --- | --- |
| Host | Docker Engine con Compose v2 o Docker Desktop; navegador; puerto local 8081 disponible. |
| Recursos recomendados | Jenkins recomienda 4 GB o más de RAM y 50 GB o más de disco para un equipo pequeño; reservar capacidad adicional para imágenes y builds. No son mediciones de este equipo. |
| Límites de esta configuración | Controlador: máximo 2 CPU y 2 GiB de RAM; heap Java de 256 MiB a 1 GiB. Son límites iniciales de laboratorio, no requisitos mínimos ni consumo observado. |
| Software del controlador | Imagen oficial `jenkins/jenkins:lts`, que aporta Jenkins y su Java compatible. No instala Docker Engine ni Docker CLI. |
| Acceso | HTTP únicamente en `127.0.0.1:8081`, dirigido al puerto 8080 del contenedor. Sin túneles, webhook público ni publicación del puerto de agentes 50000. |
| Persistencia | Volumen nombrado `jenkins_home` para configuración, plugins, jobs e historial. También contiene información sensible de Jenkins; no se versiona ni publica. |
| Seguridad | Usuario `jenkins`, sin privilegios, sin socket Docker del host y sin credenciales preconfiguradas. Se conservan el asistente inicial y la autenticación. |
| Ejecución | El script de inicialización fija en cero los ejecutores del nodo integrado. Los builds deben usar un agente separado con etiqueta `docker`. |

La infraestructura usa un proyecto Compose independiente,
`meli-jenkins-entrega2`; no agrega un tercer servicio de aplicación ni modifica
la red `application`. El volumen de Jenkins tiene una necesidad real de
persistencia; no implica agregar almacenamiento a la aplicación en memoria.

### Plugins declarados

| ID | Función |
| --- | --- |
| `git` | Obtener el repositorio y consultar cambios SCM. |
| `github` | Integración con GitHub; instalarlo no autoriza accesos remotos ni crea un webhook. |
| `workflow-aggregator` | Componentes de Pipeline como código. |
| `pipeline-stage-view` | Visualización de etapas e historial del Pipeline. |
| `docker-workflow` | Pasos de Docker Pipeline; no proporciona un daemon Docker. |
| `credentials-binding` | Vinculación temporal de credenciales por ID cuando un paso autorizado las requiera. |

`jenkins-plugin-cli` resuelve las dependencias durante la construcción de la
imagen. La receta está versionada, pero `lts` y los plugins sin versión son
referencias móviles: **no garantiza una reconstrucción idéntica en otra fecha**.
Después de una construcción autorizada, registrar el digest de la imagen base,
el ID de la imagen resultante y las versiones instaladas para trazar la evidencia.
Fijar esas versiones requiere validar compatibilidad; no se inventaron versiones.

## Arranque local del controlador

Ejecutar desde la raíz del repositorio, únicamente cuando Docker esté disponible
y la construcción/descarga de imágenes esté autorizada:

```bash
sh tests/jenkins-config.sh
docker compose -f infra/jenkins/compose.yaml config --quiet
docker compose -f infra/jenkins/compose.yaml up --build -d
docker compose -f infra/jenkins/compose.yaml ps
```

Abrir [Jenkins local](http://127.0.0.1:8081). Completar el desbloqueo y crear un
administrador propio con el asistente oficial. Consultar la contraseña inicial
en el entorno local de forma privada: **no copiarla a logs de evidencia,
capturas, informes ni este repositorio**. Verificar en **Manage Jenkins → Nodes →
Built-In Node → Configure** que el número de ejecutores siga en `0`.

La primera construcción requiere acceso a la imagen oficial y al centro de
actualizaciones de plugins. No se ha ejecutado por el mero hecho de existir
estos comandos. Para detener sin borrar el volumen:

```bash
docker compose -f infra/jenkins/compose.yaml down
```

No usar `down -v`: elimina el estado persistente y requiere una decisión explícita
de borrado y una copia de seguridad verificada.

## Configuración pendiente: agente y repositorio

**Antes de instalar o conectar un agente**, autorizar el destino, el usuario o
sesión que se utilizará y las operaciones permitidas. Un agente que accede a un
daemon Docker puede controlar los contenedores y recursos de ese daemon; esta
capacidad no debe trasladarse al controlador montando el socket del host.

En un entorno de laboratorio aprobado, preparar un agente separado con Java
compatible con Jenkins, Git, shell POSIX, Docker CLI, Compose v2 y acceso a un
daemon Docker dedicado al laboratorio. Configurarlo con etiqueta `docker`, un
ejecutor y un directorio de trabajo propio, sin acceso a `jenkins_home`. Si se
utiliza conexión entrante WebSocket, no hace falta publicar 50000; la URL debe
ser accesible desde el agente autorizado. `127.0.0.1` solo es válido cuando el
agente está en el mismo host, no desde otro contenedor o equipo.

El acceso del agente dedicado al Docker local fue autorizado posteriormente. Su
[receta separada](agent/README.md) mantiene el socket fuera del controlador y
explica el registro del nodo después del asistente. La autorización no incluye
GitHub. Hasta contar con el agente conectado, un job dirigido a `docker`
quedará esperando ejecutor.

### Credenciales y job Pipeline

La configuración siguiente es una **guía pendiente de ejecución**, no una
integración remota ya realizada. Solo proceder después de autorizar el
repositorio remoto, la operación y la credencial o sesión que se utilizará.

1. Si el repositorio es privado, guardar en **Manage Jenkins → Credentials** una
   credencial de solo lectura, de alcance mínimo. Un ejemplo de ID no secreto
   es `github-scm-read`. Para HTTPS Git, usar el tipo compatible con usuario y
   token como contraseña; el token se introduce exclusivamente en Jenkins.
   Para un repositorio público puede no requerirse credencial.
2. Crear un job de tipo **Pipeline**, no Freestyle. En **Definition**, elegir
   **Pipeline script from SCM**, y en **SCM**, **Git**. Indicar la URL y la rama
   autorizadas que contengan los cambios de la entrega. Si corresponde,
   seleccionar el ID `github-scm-read`; nunca escribir el valor del token en la
   URL, el `Jenkinsfile` ni variables versionadas.
3. Usar **Script Path: `Jenkinsfile`**. En **Build Triggers**, habilitar **Poll SCM**
   con `H/5 * * * *`. El `Jenkinsfile` declara también esta misma programación,
   que Jenkins aplicará al cargarlo. La elección en la interfaz permite dejar
   preparado el sondeo sin esperar a un primer build exitoso.
4. Confirmar un agente autorizado disponible con etiqueta `docker` y ejecutar
   el primer build. El Pipeline versionado debe mostrar **Checkout → Build →
   Test** y su bloque `post`; debe construir `api` y `web`, no solo Jenkins.
5. Solo con autorización para la operación remota, realizar un cambio real,
   publicarlo en la rama autorizada y comprobar que el sondeo inicia otro build.

Se elige Poll SCM porque no requiere exponer Jenkins a Internet ni crear un
webhook. `H/5` distribuye el minuto inicial por hash y consulta aproximadamente
cada cinco minutos; **solo dispara un build cuando hay cambios**, no es una
ejecución incondicional cada cinco minutos. El sondeo requiere acceso al remoto
y una configuración SCM válida. GitHub Actions complementa esta entrega, pero
no reemplaza la instalación ni las evidencias de Jenkins.

## Verificación y evidencias

La prueba `tests/jenkins-config.sh` verifica declaraciones de seguridad, volumen,
puerto, usuario, ejecutores y plugins. `docker compose ... config --quiet` valida
la sintaxis Compose sin arrancar contenedores. Son controles estructurales, no
pruebas de instalación de plugins, autenticación, conexión de agente o Pipeline.
No hay una prueba RED funcional significativa antes de crear esta receta de
infraestructura; no se simula una ejecución de Jenkins para declarar GREEN.

Después del arranque y la ejecución autorizados, recoger:

- Estado real del contenedor, versión de Jenkins y listado de plugins instalados.
- Pantalla autenticada, configuración de credenciales mostrando solo IDs y agente
  conectado; ocultar cualquier secreto o dato de sesión.
- Vista de etapas, consola e historial del job, más disparo automático observado.
- Recursos realmente usados mediante `docker stats --no-stream`, diferenciados
  de las recomendaciones y límites de esta guía.

Una cola esperando `docker` indica que falta un agente compatible; **no** se
soluciona activando ejecutores del controlador ni deshabilitando seguridad.
Si el puerto 8081 está ocupado, detener la prueba y decidir otro puerto local
sin publicar en `0.0.0.0`. Un error de descarga o plugin debe conservarse como
evidencia de fallo; no ocultarlo ni sustituirlo por una captura esperada.

## Referencias oficiales

- Jenkins. (s. f.). [Docker](https://www.jenkins.io/doc/book/installing/docker/).
- Jenkins. (s. f.). [Hardware recommendations](https://www.jenkins.io/doc/book/scaling/hardware-recommendations/).
- Jenkins. (s. f.). [Offline installations: prerequisites](https://www.jenkins.io/doc/book/installing/offline/).
- Jenkins. (s. f.). [Controller isolation](https://www.jenkins.io/doc/book/security/controller-isolation/).
- Jenkins. (s. f.). [Getting started with Pipeline: in SCM](https://www.jenkins.io/doc/book/pipeline/getting-started/#in-scm).
- Jenkins. (s. f.). [Pipeline syntax: triggers](https://www.jenkins.io/doc/book/pipeline/syntax/#triggers).
- Jenkins. (s. f.). [Using credentials](https://www.jenkins.io/doc/book/using/using-credentials/).
- Jenkins project. (s. f.). [Official Jenkins Docker image](https://github.com/jenkinsci/docker).
