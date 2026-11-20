# Diagnóstico de entrega 2

Fecha: 20 de noviembre de 2026. Estado inicial: commit d9de619, árbol limpio. Guía local página 4 comprobada. Rúbrica y comentarios tomados del texto proporcionado por el estudiante; Canvas no consultado mediante una sesión autorizada. Esta matriz distingue diseño de comprobación ejecutada.

| Criterio o comentario | Brecha inicial | Acción | Archivo | Prueba prevista | Evidencia prevista |
|---|---|---|---|---|---|
| C1 Características Jenkins | Sin documentación específica | Documentar recursos, instalación, plugins, puertos y credenciales | infra/jenkins/README.md | Lectura estructural y fuentes oficiales | Informe |
| C2 Instalación | Sin controlador Jenkins | Imagen LTS, plugins y volumen independiente | infra/jenkins/ | Arranque local autenticado | HTTP, inventario y captura sin secretos |
| C3 Pipeline como código | Sin Jenkinsfile | Checkout, Build, Test, post | Jenkinsfile | Ejecución en agente autorizado | Etapas y consola pendientes |
| C4 GitHub | Sin trigger Jenkins | Poll SCM sin exposición pública | Jenkinsfile | Cambio remoto autorizado | Disparo automático pendiente |
| C5 Dos imágenes | Compose ya construye api y web | Reutilizar los dos servicios | scripts/ci.sh | Build de ambos | build.log |
| C6 Pruebas | Pruebas existentes, no pipeline | Reutilizar Go/Vitest y comprobar comunicación | infra/ci/, scripts/ci.sh | Pruebas reales y códigos no cero | Logs de test y smoke |
| C7 Evidencias | Solo evidencia anterior | Capturar nuevos resultados, no inventar Jenkins | evidence/entrega2/ | Fallo controlado y corrección | Builds pendientes |
| C8 Documento | Falta informe de entrega 2 | Editable, PDF, referencias, conclusiones | docs/delivery/ | Render y revisión visual | Informe y matriz final |
| O1 Variables y volumen web | Sin variable runtime; aplicación en memoria | API_UPSTREAM; justificar ausencia de volumen | compose.yaml, frontend/ | Validación, proxy y salud | Logs y decisión |
| O2 .env.example | Ausente; .env* ignorado | Ejemplo sin secretos y excepción Git | .env.example, .gitignore | Lectura y check-ignore | Archivo versionado |
| O3 GitHub Actions | Ausente | Workflow construye api y web sin push | .github/workflows/build-images.yml | Build local; ejecución remota pendiente | Configuración y log local |
| O4 network inspect | Lista textual no es inspección | Capturar salida real | scripts/ci.sh | IDs de ambos servicios en red | network-inspect.json |

Mejoras opcionales de seguridad y reproducibilidad: puertos efímeros en pruebas, proyectos aislados, timeout, controlador sin Docker socket, scripts compartidos. No son descriptores adicionales de la rúbrica. No se agregan bases de datos, Travis CI ni Codeship.
