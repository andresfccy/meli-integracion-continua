# Jenkins entrega 2
Branch: feature/two-containers. Delivery: ask-on-risk; local commits only. Forecast: 900 authored lines plus report.
Checks: container builds, Go/Vitest, isolated Compose communication, Jenkins when executor consent exists, visual document QA.

## Specs
### S1
Actúa como ingeniero DevOps y asesor académico. Ayúdame a completar la Entrega proyecto 2 — Escenario 5 de Integración Continua del Politécnico Grancolombiano, con implementación funcional, pruebas reales y documentación académica en español profesional.

### S2
OBJETIVO
Preparar la entrega para el 24 de noviembre de 2026 a las 23:55, cumpliendo los ocho criterios de evaluación y corrigiendo las observaciones de la entrega anterior. No te limites a redactar una propuesta: implementa y verifica lo que permita el entorno autorizado, y señala claramente cualquier pendiente.

### S3
CONTEXTO Y FUENTES
Repositorio local:
/Users/acardenas/Workspace/Personal/Poli/2-integracion-continua/meli-integracion-continua

### S4
Guía local:
/Users/acardenas/Workspace/Personal/Poli/2-integracion-continua/descargables/Proyecto_software_herramientas_integracion_continua_Escenario5.pdf

### S5
Canvas:
- Entrega actual: https://poli.instructure.com/courses/97400/assignments/613518
- Entrega anterior: https://poli.instructure.com/courses/97400/assignments/613516
- Rúbrica publicada: https://poli.instructure.com/courses/97400/discussion_topics/955513
- Ejercicio de apoyo Jenkins + Docker y GitHub Actions: https://poli.instructure.com/courses/97400/discussion_topics/958005

### S6
Lee primero las instrucciones AGENTS.md del repositorio y la documentación que estas indiquen. Inspecciona el código y los comandos reales antes de modificarlo. Conserva el trabajo existente y los cambios ajenos.

### S7
El proyecto ya tiene dos servicios Docker Compose, api y web, conectados mediante la red application. No agregues otro servicio de aplicación por interpretar erróneamente que falta el segundo. Jenkins es infraestructura adicional.

### S8
La página 4 de la guía exige implementar Jenkins como gestor de operaciones y entregar un documento con las características necesarias para implementarlo. Travis CI y Codeship corresponden a la entrega final, no a esta.

### S9
Si tienes acceso autorizado a Canvas, contrasta las fuentes y revisa la tarea cuando esté desbloqueada. Si no puedes acceder, trabaja con estos requisitos y declara la limitación; no inventes instrucciones adicionales.

### S10
RETROALIMENTACIÓN DE LA ENTREGA ANTERIOR
Calificación: 123,75/150.

### S11
Observaciones del profesor:
1. Segundo servicio: configurarlo con variables de entorno y un volumen si guarda datos.
2. Buenas prácticas: falta un .env.example con las variables, sin valores reales.
3. CI: falta un workflow de GitHub Actions que construya las imágenes.
4. Documento: falta evidencia de docker network inspect.

### S12
Corrige esas brechas y documenta la solución de cada una. Verifica primero cuáles siguen presentes.

### S13
El almacenamiento de la aplicación está documentado como en memoria: no inventes una base de datos ni un volumen inútil. Explica si corresponde persistencia y justifica la decisión. Distingue esta condición del volumen persistente de Jenkins.

### S14
GitHub Actions debe atender el comentario anterior, pero NO reemplaza Jenkins: en la entrega 2 se evalúa Jenkins.

### S15
RÚBRICA DE LA ENTREGA 2 — 150 PUNTOS

### S16
1. Documento de características requeridas para Jenkins — 25 puntos
Incluye requisitos de hardware y software, instalación y justificación, puertos, plugins Git/GitHub, Pipeline y Docker Pipeline, y manejo de credenciales.
El profesor recomienda jenkins/jenkins:lts con un volumen. Distingue requisitos recomendados de recursos realmente usados.

### S17
2. Instalación y configuración de Jenkins — 20 puntos
Demuestra Jenkins funcionando, con los plugins necesarios y credenciales guardadas en Jenkins, nunca escritas en el código.

### S18
3. Pipeline como código — 25 puntos
Incluye un Jenkinsfile versionado con etapas Checkout, Build y Test, y un bloque post.
No sustituyas el Jenkinsfile por un job configurado exclusivamente desde la interfaz.

### S19
4. Integración con GitHub — 15 puntos
Configura ejecución automática ante cambios mediante webhook o Poll SCM: ambas opciones tienen el mismo valor según el profesor.
Justifica la opción elegida. Los builds exclusivamente manuales corresponden al nivel Básico.
Si webhook requiere exposición pública o permisos adicionales, no los habilites sin autorización; considera Poll SCM.

### S20
5. Construcción de los contenedores de la entrega 1 — 20 puntos
El pipeline debe construir las imágenes de los dos servicios existentes, no únicamente la imagen de Jenkins.

### S21
6. Pruebas y verificación automática — 15 puntos
La rúbrica admite pruebas unitarias o una prueba que levante Compose y compruebe comunicación entre los dos contenedores por nombre de servicio.
Reutiliza las pruebas existentes pertinentes y demuestra comunicación entre servicios.
Si una prueba falla, el build debe fallar; no ocultes errores ni fuerces resultados exitosos.

### S22
7. Evidencias de ejecución — 15 puntos
Incluye vista de etapas, salida de consola e historial de builds.
El profesor recomienda mostrar un build fallido y el posterior que lo corrige. Genera esa demostración de manera controlada y sin afectar datos o trabajo ajeno.

### S23
8. Calidad del documento — 15 puntos
Incluye integrantes, diagrama del flujo commit → GitHub → Jenkins → build → pruebas, enlace al repositorio, conclusiones propias y referencias APA.

### S24
Niveles publicados:
- Excelente: 100 %.
- Bueno: 75 %.
- Básico: 50 %.
- Insuficiente: 25 % o menos.

### S25
No inventes descriptores adicionales ni garantices una calificación.

### S26
FORMA DE TRABAJO
1. Diagnostica el estado actual y crea una matriz:
   criterio o comentario | brecha | acción | archivo | prueba | evidencia.
   Distingue requisitos docentes de mejoras opcionales.
2. Implementa cambios mínimos y coherentes con la arquitectura existente.
3. Verifica los comandos de prueba en los archivos reales del proyecto; no supongas que todos los comandos del README están actualizados.
4. Comprueba construcción de ambas imágenes, arranque, salud y comunicación de servicios, ejecución del pipeline y propagación de errores.
5. Recoge evidencias auténticas a medida que avanzas, incluyendo docker network inspect. Relaciona cada evidencia con un criterio.
6. Verifica el disparo automático de Jenkins con un cambio real únicamente cuando exista autorización para la operación remota. Si no la hay, deja la configuración preparada y registra la verificación como pendiente.
7. Redacta el documento a partir de los resultados observados, no de resultados deseados.
8. Realiza una revisión final contra los ocho criterios y las cuatro observaciones anteriores.

### S27
ENTREGABLES
- Cambios de código y configuración necesarios, incluyendo Jenkinsfile, configuración reproducible de Jenkins, .env.example y workflow de GitHub Actions para construir ambas imágenes.
- README actualizado con requisitos, comandos, ejecución, pruebas y solución de problemas.
- Informe académico editable y PDF revisado visualmente, salvo que Canvas indique otro formato. No atribuyas a la rúbrica una extensión o formato que no especifica.
- Evidencias reales, organizadas y referenciadas desde el informe.
- Matriz final de cumplimiento: verificado, parcial o pendiente, con referencias concretas.
- Resumen de cambios, pruebas ejecutadas, fallos, limitaciones y pasos que requieren intervención humana.

### S28
REGLAS
- No inventes capturas, logs, ejecuciones, integrantes, referencias ni resultados.
- Si falta una evidencia, indícalo claramente; una instrucción para obtenerla no equivale a haberla obtenido.
- No expongas secretos en código, documentos, capturas o consola.
- No publiques servicios, cambies permisos, uses credenciales remotas, hagas push ni envíes la tarea a Canvas sin autorización explícita para esa operación.
- No amplíes el alcance a herramientas de la entrega final.
- Explica brevemente las decisiones técnicas para que pueda comprenderlas y sustentarlas.

### S29
User confirmation: "sí" to "¿Autorizás configurar un agente Jenkins local dedicado con acceso al motor Docker de tu equipo, teniendo presente que ese acceso permite controlar todos sus contenedores?"

## Tasks
- T1 | S6-S7,S11-S13,S26-S28 | delegated: independent web runtime unit | complete | 5b8cc76; shell RED/GREEN and real smoke passed
- T2 | S8,S14-S17,S27-S28 | delegated: independent controller unit | partial | 398f441; controller built/running, setup/credentials/agent pending human action
- T3 | S18-S22,S26-S28 | inline: Jenkinsfile, Actions and shared checks | partial | 8d86bb4; local builds/tests pass; Jenkins/remote runs require authorization
- T4 | S1-S28 | delegated: independent risk items 2 and 5 verification | complete within local scope | no blockers reproduced; default nginx byte-identical to baseline; runtime proof below
- T5 | S15,S23-S28 | inline: Spanish report and evidence | partial | 2bb05fb; report DOCX/PDF visually checked, eight pages; members and Jenkins build evidence pending

- T6 | S17-S22,S28-S29 | inline: dedicated local Docker agent | partial | image/tools/socket verified; private admin creation and node connection pending

## Log
### L1 Original request
Actúa como ingeniero DevOps y asesor académico. Ayúdame a completar la Entrega proyecto 2 — Escenario 5 de Integración Continua del Politécnico Grancolombiano, con implementación funcional, pruebas reales y documentación académica en español profesional.

OBJETIVO
Preparar la entrega para el 24 de noviembre de 2026 a las 23:55, cumpliendo los ocho criterios de evaluación y corrigiendo las observaciones de la entrega anterior. No te limites a redactar una propuesta: implementa y verifica lo que permita el entorno autorizado, y señala claramente cualquier pendiente.

CONTEXTO Y FUENTES
Repositorio local:
/Users/acardenas/Workspace/Personal/Poli/2-integracion-continua/meli-integracion-continua

Guía local:
/Users/acardenas/Workspace/Personal/Poli/2-integracion-continua/descargables/Proyecto_software_herramientas_integracion_continua_Escenario5.pdf

Canvas:
- Entrega actual: https://poli.instructure.com/courses/97400/assignments/613518
- Entrega anterior: https://poli.instructure.com/courses/97400/assignments/613516
- Rúbrica publicada: https://poli.instructure.com/courses/97400/discussion_topics/955513
- Ejercicio de apoyo Jenkins + Docker y GitHub Actions: https://poli.instructure.com/courses/97400/discussion_topics/958005

Lee primero las instrucciones AGENTS.md del repositorio y la documentación que estas indiquen. Inspecciona el código y los comandos reales antes de modificarlo. Conserva el trabajo existente y los cambios ajenos.

El proyecto ya tiene dos servicios Docker Compose, api y web, conectados mediante la red application. No agregues otro servicio de aplicación por interpretar erróneamente que falta el segundo. Jenkins es infraestructura adicional.

La página 4 de la guía exige implementar Jenkins como gestor de operaciones y entregar un documento con las características necesarias para implementarlo. Travis CI y Codeship corresponden a la entrega final, no a esta.

Si tienes acceso autorizado a Canvas, contrasta las fuentes y revisa la tarea cuando esté desbloqueada. Si no puedes acceder, trabaja con estos requisitos y declara la limitación; no inventes instrucciones adicionales.

RETROALIMENTACIÓN DE LA ENTREGA ANTERIOR
Calificación: 123,75/150.

Observaciones del profesor:
1. Segundo servicio: configurarlo con variables de entorno y un volumen si guarda datos.
2. Buenas prácticas: falta un .env.example con las variables, sin valores reales.
3. CI: falta un workflow de GitHub Actions que construya las imágenes.
4. Documento: falta evidencia de docker network inspect.

Corrige esas brechas y documenta la solución de cada una. Verifica primero cuáles siguen presentes.

El almacenamiento de la aplicación está documentado como en memoria: no inventes una base de datos ni un volumen inútil. Explica si corresponde persistencia y justifica la decisión. Distingue esta condición del volumen persistente de Jenkins.

GitHub Actions debe atender el comentario anterior, pero NO reemplaza Jenkins: en la entrega 2 se evalúa Jenkins.

RÚBRICA DE LA ENTREGA 2 — 150 PUNTOS

1. Documento de características requeridas para Jenkins — 25 puntos
Incluye requisitos de hardware y software, instalación y justificación, puertos, plugins Git/GitHub, Pipeline y Docker Pipeline, y manejo de credenciales.
El profesor recomienda jenkins/jenkins:lts con un volumen. Distingue requisitos recomendados de recursos realmente usados.

2. Instalación y configuración de Jenkins — 20 puntos
Demuestra Jenkins funcionando, con los plugins necesarios y credenciales guardadas en Jenkins, nunca escritas en el código.

3. Pipeline como código — 25 puntos
Incluye un Jenkinsfile versionado con etapas Checkout, Build y Test, y un bloque post.
No sustituyas el Jenkinsfile por un job configurado exclusivamente desde la interfaz.

4. Integración con GitHub — 15 puntos
Configura ejecución automática ante cambios mediante webhook o Poll SCM: ambas opciones tienen el mismo valor según el profesor.
Justifica la opción elegida. Los builds exclusivamente manuales corresponden al nivel Básico.
Si webhook requiere exposición pública o permisos adicionales, no los habilites sin autorización; considera Poll SCM.

5. Construcción de los contenedores de la entrega 1 — 20 puntos
El pipeline debe construir las imágenes de los dos servicios existentes, no únicamente la imagen de Jenkins.

6. Pruebas y verificación automática — 15 puntos
La rúbrica admite pruebas unitarias o una prueba que levante Compose y compruebe comunicación entre los dos contenedores por nombre de servicio.
Reutiliza las pruebas existentes pertinentes y demuestra comunicación entre servicios.
Si una prueba falla, el build debe fallar; no ocultes errores ni fuerces resultados exitosos.

7. Evidencias de ejecución — 15 puntos
Incluye vista de etapas, salida de consola e historial de builds.
El profesor recomienda mostrar un build fallido y el posterior que lo corrige. Genera esa demostración de manera controlada y sin afectar datos o trabajo ajeno.

8. Calidad del documento — 15 puntos
Incluye integrantes, diagrama del flujo commit → GitHub → Jenkins → build → pruebas, enlace al repositorio, conclusiones propias y referencias APA.

Niveles publicados:
- Excelente: 100 %.
- Bueno: 75 %.
- Básico: 50 %.
- Insuficiente: 25 % o menos.

No inventes descriptores adicionales ni garantices una calificación.

FORMA DE TRABAJO
1. Diagnostica el estado actual y crea una matriz:
   criterio o comentario | brecha | acción | archivo | prueba | evidencia.
   Distingue requisitos docentes de mejoras opcionales.
2. Implementa cambios mínimos y coherentes con la arquitectura existente.
3. Verifica los comandos de prueba en los archivos reales del proyecto; no supongas que todos los comandos del README están actualizados.
4. Comprueba construcción de ambas imágenes, arranque, salud y comunicación de servicios, ejecución del pipeline y propagación de errores.
5. Recoge evidencias auténticas a medida que avanzas, incluyendo docker network inspect. Relaciona cada evidencia con un criterio.
6. Verifica el disparo automático de Jenkins con un cambio real únicamente cuando exista autorización para la operación remota. Si no la hay, deja la configuración preparada y registra la verificación como pendiente.
7. Redacta el documento a partir de los resultados observados, no de resultados deseados.
8. Realiza una revisión final contra los ocho criterios y las cuatro observaciones anteriores.

ENTREGABLES
- Cambios de código y configuración necesarios, incluyendo Jenkinsfile, configuración reproducible de Jenkins, .env.example y workflow de GitHub Actions para construir ambas imágenes.
- README actualizado con requisitos, comandos, ejecución, pruebas y solución de problemas.
- Informe académico editable y PDF revisado visualmente, salvo que Canvas indique otro formato. No atribuyas a la rúbrica una extensión o formato que no especifica.
- Evidencias reales, organizadas y referenciadas desde el informe.
- Matriz final de cumplimiento: verificado, parcial o pendiente, con referencias concretas.
- Resumen de cambios, pruebas ejecutadas, fallos, limitaciones y pasos que requieren intervención humana.

REGLAS
- No inventes capturas, logs, ejecuciones, integrantes, referencias ni resultados.
- Si falta una evidencia, indícalo claramente; una instrucción para obtenerla no equivale a haberla obtenido.
- No expongas secretos en código, documentos, capturas o consola.
- No publiques servicios, cambies permisos, uses credenciales remotas, hagas push ni envíes la tarea a Canvas sin autorización explícita para esa operación.
- No amplíes el alcance a herramientas de la entrega final.
- Explica brevemente las decisiones técnicas para que pueda comprenderlas y sustentarlas.

### L2 2026-11-20
Sí, usá gpt-6.1-sol y continuá con el diagnóstico y la implementación dentro del alcance autorizado. Y actualizaa todos los lugares donde aparezcan modelos descontinuados

### L3 Diagnosis
Initial worktree clean; all four prior gaps confirmed. Exactly api/web; memory store remains. Guide page 4 requires Jenkins. RDD off by clone; do not re-enable. Docker local daemon available via approved escalation. No authorized remote session identified; Canvas not accessed. Jenkins host socket/privileged daemon not authorized; prepare controller without either and record build execution dependency. Model audit found two unavailable gpt-6.1-luna assignments in global AGENTS; both replaced with gpt-6.1-sol, effort unchanged. No other active unavailable worker assignments found.

### L4 2026-11-20 Observed checks and human boundary
Application build exit0; fresh Go test/vet exit0 (6 packages tested, 3 without tests); frontend 7 tests/2files and build exit0; smoke exit0 with healthy api/web, direct DNS, proxy and actual network inspect. Invalid upstream actual container exit1 expected. Source test scripts RED/GREEN and error propagation verified. Independent verifier reproduced no blocker; nginx defaults byte-identical to baseline d9de619. Parent reran shell checks; Markdown links and git diff --check pass. Playwright not run (no business/login logic changes). No pipeline executed in Jenkins; no remote trigger, push, credential usage or Canvas access.
Jenkins 2.580.2, Java21.0.13, six required plugins installed. Controller meli-jenkins-entrega2-jenkins-1 remains running at 127.0.0.1:8081 with volume meli-jenkins-entrega2_jenkins_home. Setup wizard untouched; zero executors; no socket/privileged runtime. Snapshot 693.9MiB/2GiB, CPU0.50%. Screenshot authentic, empty password field. Initial HTTP reset during startup recovered without restart. Complete wizard privately; authorization for agent with Docker control is the next security decision.
Report docs/delivery/Informe_entrega2_Jenkins.docx and matching PDF: 8 pages rendered and inspected, no overlaps/clipping; APA-style references, authentic screenshot, explicit partial matrix. Members unknown and not invented. Guide p4 visually inspected. Remote repo URL taken only from local Git config.
Native review mode remains off by clone (disabled/unmanaged). Assess attempted after writers and once after source commits but unassessable because untracked report/evidence required declaration; no PASS invented, no review START attempted. Independent functional verification remains valid. Source authored delta: 682 additions/deletions across three commits; passive report/log and generated evidence additional, not a reason to publish.
Model audit: updated only 2 gpt-6.1-luna assignments in /Users/acardenas/.codex/AGENTS.md to gpt-6.1-sol high; verifier sol medium unchanged. No other active unavailable worker assignments found in scoped instructions/configs. Provider-wide retirement not established. Gentle AI global sync may overwrite generated routing; no source override located, no sync performed.
Risk: items 2 and 5 (environment-derived nginx config and Jenkins/CI runtime). Shared defaults checked against base; no application behavior or storage changes.

### L5 Local delivery record
Source commits: 5b8cc76, 398f441, 8d86bb4. Report/evidence commit: 2bb05fb. Named verification scratch cleaned by parent; runtime Jenkins container and persistent volume retained. No remote publication. Next step is explicit agent-access authorization and private initial administrator setup; remaining requirements stay partial.

### L6 2026-11-20 User authorization
sí
Context: accepts dedicated local Jenkins agent with access to local Docker engine and disclosed control over its containers. Does not authorize remote credentials, push, public exposure or changes to administrator credentials. Prepare and verify agent capabilities locally; private Jenkins setup remains required before registration/connection.

### L7 2026-11-20 Authorized agent preparation and observed boundary
Added opt-in dedicated agent image/overlay with UID1000, socket group0, no published ports, readonly dedicated secret volume, fail-closed startup and no controller socket. RED missing Dockerfile exit1; structural GREEN. Actual image build passed; Java21.0.13, Git2.47.3, Docker29.10.0, Compose5.7.1, Buildx0.39.0; authorized local server29.9.3 reachable. Missing-secret startup exit2. Evidence in evidence/entrega2/agent. Independent verifier found no blockers; args4j expansion supports @file in official sources, authenticated connection not yet observed. Parent reran both shell contracts successfully.
Privately unlocked Jenkins with existing bootstrap password, without exposing it; temporary secret file removed. Selected no additional suggested plugins, retaining preinstalled plugins. Browser now shows Create First Admin User with all fields empty; screenshot admin-setup-pending.jpg. Creating a new authentication credential requires user handoff. No node, agent secret, permanent agent, SCM credentials or job created. Existing controller remains running. No remote access or publication.
Updated partial report and evidence matrix: Docker access is authorized and capability verified, not agent connectivity or completed Jenkins pipeline. RDD remains off; no review lifecycle started. Risk: items 2 and 5 (authorized Docker engine access and agent runtime). Independent verification passed with connection explicitly pending.
