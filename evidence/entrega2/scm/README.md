# Jenkinsfile publicado desde SCM

La rama feature/two-containers fue publicada por SSH con autorización expresa y size:exception. El job original `meli-entrega2` obtuvo el Jenkinsfile de GitHub y el build2 terminó SUCCESS. No utilizó el archivo local ni el harness de validación anterior.

El build2 fue manual para registrar el trigger declarativo Poll SCM `H/5 * * * *`. Este cambio documental registra esa evidencia y sirve como cambio real para verificar el disparo automático posterior. No se lanzará ese siguiente build manualmente. El merge a main permanece condicionado a las verificaciones.
