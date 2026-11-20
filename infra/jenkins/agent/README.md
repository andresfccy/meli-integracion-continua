# Agente Docker local autorizado

El agente es infraestructura de Jenkins, no un tercer servicio de aplicación. Su imagen incluye Java, Git, Docker CLI, Compose y Buildx. Ejecuta como `jenkins`, sin privilegios, sin puertos publicados y sin acceso directo a `jenkins_home`. La conexión usa WebSocket hacia `http://jenkins:8080/` en la red interna.

**Límite de seguridad:** el socket permite controlar todo el motor Docker local, incluidos otros contenedores y montajes. No es una frontera de aislamiento frente a jobs maliciosos. El usuario autorizó expresamente este acceso para el agente de la entrega. No ejecutar código ni PR de terceros; usar un ejecutor y el modo «solo trabajos con etiquetas coincidentes». El controlador sigue sin socket.

## Construir y comprobar herramientas

Desde la raíz:

```sh
sh tests/jenkins-agent.sh
docker compose -f infra/jenkins/compose.yaml -f infra/jenkins/compose.agent.yaml config --quiet
docker build -t meli-jenkins-docker-agent:local infra/jenkins/agent
docker run --rm --entrypoint sh meli-jenkins-docker-agent:local \
  -c 'id; java -version; git --version; docker --version; docker compose version; docker buildx version'
```

Las etiquetas base son móviles. Registrar los digests resueltos y versiones observadas tras construir. No confundir la disponibilidad de herramientas con un agente conectado.

## Registrar después del asistente inicial

1. Completar **privadamente** el asistente de Jenkins y crear el administrador. No guardar contraseñas en este repositorio ni en evidencias.
2. Crear nodo permanente `docker-local`, raíz `/home/jenkins/agent`, un ejecutor, etiqueta `docker`, uso exclusivo para jobs con esa etiqueta y lanzamiento mediante conexión del agente al controlador.
3. Provisionar el secreto emitido por ese nodo en el volumen dedicado `meli-jenkins-entrega2_agent_secret`, archivo `secret` de una sola línea, propiedad UID/GID `1000:1000`, modo `0600`. Transferir por entrada estándar, sin imprimirlo ni usarlo como argumento de comandos. Este volumen no contiene el home del controlador. El secreto de agente no es una credencial GitHub.
4. Con el volumen listo, arrancar el overlay:

```sh
docker compose -f infra/jenkins/compose.yaml -f infra/jenkins/compose.agent.yaml \
  up -d --build docker-agent
```

5. Verificar `docker-local` online en Jenkins, etiqueta `docker`, un ejecutor, y registrar evidencia sin secretos. Una instancia Docker encendida no basta para afirmar conexión.

Se pasa `-secret @/run/jenkins-agent/secret`, no el valor del secreto en variables de entorno o argumentos. El archivo se monta de solo lectura; falta o vacío provoca salida 2, sin reintento infinito. El GID adicional por defecto es 0 para el socket del Docker Desktop local; si no coincide, comprobar el GID real y ajustar `DOCKER_SOCKET_GID` sin cambiar permisos del socket.

## Detener el acceso

```sh
docker compose -f infra/jenkins/compose.yaml -f infra/jenkins/compose.agent.yaml stop docker-agent
```

No borrar volúmenes ni ejecutar prune. La autorización del agente no autoriza GitHub, push, uso de SSH ni exposición pública. La ejecución del pipeline sigue pendiente hasta disponer del nodo conectado y SCM autorizado.

Referencias: [imagen oficial](https://github.com/jenkinsci/docker-agent/blob/master/README_inbound-agent.md), [agentes Jenkins](https://www.jenkins.io/doc/book/using/using-agents/), [seguridad de Docker](https://docs.docker.com/engine/security/).

## Verificación observada

La imagen se construyó y sus herramientas se ejecutaron como UID1000 con acceso autorizado al motor local; el arranque sin secreto devolvió 2. Evidencia: `evidence/entrega2/agent/`. La expansión `@archivo` corresponde al parser args4j utilizado por [Remoting](https://github.com/jenkinsci/remoting/blob/master/src/main/java/hudson/remoting/Launcher.java), no a una opción independiente de `-secret`. La conexión autenticada sigue pendiente.
