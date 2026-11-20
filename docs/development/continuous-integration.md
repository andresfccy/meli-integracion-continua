# Verificar la integración continua

Jenkins es el gestor de la entrega 2. GitHub Actions atiende la observación de la entrega anterior y no lo sustituye. El pipeline está en `Jenkinsfile`; el controlador se prepara según [su guía](../../infra/jenkins/README.md).

## Pruebas locales equivalentes

Desde la raíz, con Docker Engine y Compose v2 disponibles:

```sh
export CI_PROJECT="meli-check-local-$(date +%s)"
export EVIDENCE_DIR="evidence/entrega2/$CI_PROJECT"
export WEB_PORT=0
sh scripts/ci.sh build
sh scripts/ci.sh test
sh scripts/ci.sh smoke
```

`build` construye api y web. `test` ejecuta Go test y vet en una imagen Go, y Vitest y build en otra imagen Node. Son entornos de prueba, no nuevos servicios de aplicación. Se utiliza `test-frontend` porque el script combinado `test` también necesita Go. Los Dockerfiles de prueba reutilizan las versiones base fijadas de las imágenes existentes. `--no-cache` evita confundir una capa de pruebas reutilizada con una ejecución nueva.

`smoke` espera los health checks, consulta `http://api:8080/api/health` desde web, comprueba el proxy web y captura `docker network inspect`. Exige el estado healthy de ambos contenedores y comprueba que sus IDs estén en la red. Al terminar elimina únicamente el proyecto temporal indicado; no borra volúmenes ni usa prune. `WEB_PORT=0` solicita un puerto efímero enlazado a loopback.

## Jenkins y disparo automático

El Jenkinsfile requiere un agente etiquetado `docker`, con Git, shell POSIX, Docker CLI y Compose v2 y acceso autorizado a un motor Docker. Configurar un job Pipeline desde SCM con ruta `Jenkinsfile`, no un script pegado en la interfaz. Los comandos se comparten con la verificación local. Las etapas son Checkout, Build y Test; `post` limpia recursos temporales y archiva las evidencias, incluso cuando una prueba falla.

Poll SCM (`H/5 * * * *`) evita exponer un webhook. Jenkins necesita acceso autorizado al repositorio y un primer build que establezca el baseline y registre el trigger. Su presencia en el archivo no prueba un disparo real. Un cambio en GitHub, push y uso de credenciales remotas requieren autorización separada.

Para la demostración controlada, lanzar el job con `DEMONSTRATE_FAILURE=true`: después de las pruebas unitarias la etapa Test falla deliberadamente. Ejecutar luego con false y conservar ambas consolas, el historial y la vista de etapas. No se modifica código ni datos para producir el fallo. Esta explicación no es evidencia de que los builds hayan ocurrido.

## Credenciales y límites

Guardar la credencial GitHub autorizada en Jenkins y seleccionarla en SCM por ID; nunca en Jenkinsfile, .env.example ni capturas. El controlador no recibe socket Docker ni se ejecuta en modo privilegiado. Autorizar un agente con control del motor Docker es una decisión de seguridad pendiente. No usar el controlador como agente de compilación.

## Fallos y diagnóstico

- Docker inaccesible: comprobar el motor local y permisos existentes; no usar chmod sobre el socket.
- Puerto ocupado: usar un proyecto único y WEB_PORT=0 para pruebas, sin detener despliegues ajenos.
- Backend inaccesible: revisar API_UPSTREAM, health checks y red application; no usar localhost para comunicar contenedores distintos.
- Build o test fallido: consultar el archivo correspondiente bajo EVIDENCE_DIR. Se preserva su código de salida; no se fuerza éxito.
- Descargas fallidas: registrar la limitación de conectividad y no presentar pruebas en caché como nuevas.
- Job en cola: comprobar que existe un agente autorizado con etiqueta docker; el controlador por sí solo no satisface este requisito.

## Fuentes técnicas

- Jenkins. (s. f.). *Docker*. https://www.jenkins.io/doc/book/installing/docker/
- Jenkins. (s. f.). *Pipeline syntax*. https://www.jenkins.io/doc/book/pipeline/syntax/
- Jenkins. (s. f.). *Credentials*. https://www.jenkins.io/doc/book/security/credentials/
- Docker. (s. f.). *docker compose up*. https://docs.docker.com/reference/cli/docker/compose/up/
- GitHub. (s. f.). *Workflow syntax for GitHub Actions*. https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-syntax
