# Evidencias reales de entrega 2

Fecha local: 20 de noviembre de 2026. Ejecución de aplicación: 18:52:14 a 18:53:11 UTC (13:52:14 a 13:53:11 America/Bogota). No se efectuó acceso a Canvas ni push. Las evidencias antiguas de evidence/ se conservan sin presentarlas como una ejecución actual.

| Evidencia | Resultado | Relación |
|---|---|---|
| [diagnostico.md](diagnostico.md) | Brechas iniciales verificadas en código | C1–C8 y O1–O4 |
| [verification-summary.txt](local/verification-summary.txt) | Resumen de comandos, tiempos y salidas | C5, C6 |
| [build.log](local/build.log) | Ambas imágenes construidas, exit 0; usa capas de construcción cacheadas cuando válidas | C5, O3 |
| [backend-tests.log](local/backend-tests.log) | Go test: 6 paquetes pasan, 3 sin tests; go vet pasa. Ejecución sin caché de Docker | C6 |
| [frontend-tests.log](local/frontend-tests.log) | 7 pruebas en 2 archivos; TypeScript y Vite build pasan | C6 |
| [services.log](local/services.log) | api y web healthy en proyecto aislado | C5, C6 |
| [direct-api.log](local/direct-api.log) | Desde web a api por nombre de servicio: status ok | C6 |
| [proxy-api.log](local/proxy-api.log) | Proxy nginx responde status ok | C6, O1 |
| [web-health.log](local/web-health.log) | Endpoint web responde ok | C6 |
| [network-inspect.json](local/network-inspect.json) | Red bridge con los IDs reales de api y web | O4, C6 |
| [cleanup.log](local/cleanup.log) | Eliminación solo de contenedores/red temporales, sin volúmenes | C6 |
| Carpeta jenkins/ | Evidencia del controlador independiente; no de jobs | C2 parcial |

El rechazo real de API_UPSTREAM inválido está registrado en la carpeta local. Los tests shell de propagación de fallos usan un doble de Docker: son verificaciones del script, NO un historial de builds Jenkins. La demostración DEMONSTRATE_FAILURE y su posterior build exitoso siguen pendientes.

[La matriz final](compliance.json) separa lo comprobado de lo pendiente. El informe está en [docs/delivery](../../docs/delivery/Informe_entrega2_Jenkins.docx).

## Controlador observado

Jenkins 2.580.2 y Java 21.0.13; seis plugins requeridos instalados. HTTP200 y captura inicial del asistente. [Resumen](jenkins/verification-summary.json), [recursos](jenkins/controller-resources.txt), [versiones de plugins](jenkins/plugin-versions.txt) y [captura sin contraseña](jenkins/unlock-jenkins.jpg). Medición puntual: 693.9MiB y CPU0.50%; no es una prueba de capacidad bajo carga. Contenedor local permanece encendido; volumen conservado. Posteriormente se desbloqueó el asistente de forma privada. Administrador, nodo conectado, credenciales SCM, job, builds y disparo SCM siguen pendientes.

## Revisión independiente

Sin bloqueadores reproducidos. Configuración nginx predeterminada idéntica byte a byte al baseline d9de619; rechazo de inyección; proyectos no temporales rechazados; fallos de build/test y arranque propagados aun si cleanup falla. Las pruebas shell no se presentan como ejecución Jenkins. RDD estaba deshabilitado por configuración del clon; no se activó. El assessment nativo inicial fue no evaluable por archivos sin seguimiento; esto no se convirtió en una aprobación.

## Agente autorizado

[Verificación real](agent/verification-summary.json): imagen construida, UID1000, Java21.0.13, Git2.47.3, Docker29.10.0, Compose5.7.1 y Buildx0.39.0. Consulta autorizada al motor29.9.3 correcta. Arranque sin secreto rechazado con salida2. Contenedores de prueba efímeros, sin agente permanente ni nodo conectado. [Formulario privado pendiente](agent/admin-setup-pending.jpg), sin credenciales.

## Conexión del agente observada

El 20 de noviembre de 2026 a las 14:42 America/Bogota Jenkins mostró **Agent is connected**. [Captura](agent/node-connected.jpg), [log WebSocket](agent/connection.log), [herramientas y permisos](agent/connected-tools.txt) y [aislamiento real](agent/connected-isolation.txt). Un ejecutor, etiqueta docker exclusiva; controlador sin socket; agente sin puertos publicados. Las capturas anteriores conservan el estado histórico. Job, SCM, builds y trigger siguen pendientes; no hubo acceso GitHub.
