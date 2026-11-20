# Pruebas reales de Jenkins sin push

- `meli-entrega2` build1: FAILURE antes de ejecutar etapas; GitHub no contiene Jenkinsfile.
- `meli-entrega2-validacion-local` build1: FAILURE controlado en Test, 47.377 s.
- Validación local build2: SUCCESS, 50.334 s; ambas imágenes, Go/vet, 7 tests frontend y smoke con dos servicios saludables.

El checkout autenticado usa solo el ID `github-meli-readonly`; ningún valor secreto se incluye. Commit local ac45ad9. `candidate-provenance.json` identifica hashes; `candidate-adaptation.diff` documenta Checkout adaptado, polling retirado y valor inicial true para la demostración. Build/Test/post mantienen su lógica. `github-scm` contiene el checkout remoto separado de las fuentes locales probadas.

Las consolas fueron extraídas de los builds reales; solo se quitaron anotaciones ANSI serializadas de Jenkins para legibilidad. `build-results.json` proviene de los estados finales; un resultado provisional durante ejecución no cuenta. `stages-history.jpg` muestra las etapas y el historial. `build-2/` conserva artefactos archivados por Jenkins.

No demuestra Jenkinsfile publicado ni trigger automático. No hubo push. Los builds no cambian datos de aplicación existentes; la limpieza se limita al proyecto temporal de cada ejecución.

## Límite del log de limpieza

El trap de smoke elimina los recursos; después, post vuelve a limpiar y sobrescribe `cleanup.log` con el resultado de esa segunda llamada. La revisión independiente confirmó que no quedan contenedores ni red del proyecto del build2. Ese archivo aislado no conserva el detalle de la primera eliminación.

El archivo fuente permanece en `/tmp/meli-candidate-ac45ad99643aa39645c81eb939b13d5374eafd2a.tar` dentro del agente para reejecuciones manuales. Recrear el contenedor elimina ese archivo; debe regenerarse desde el commit y transferirse nuevamente, conservando el hash registrado. No contiene .git ni credenciales. Este mecanismo temporal es solo para validación previa a publicación.
