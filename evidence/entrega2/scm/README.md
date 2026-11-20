# Jenkinsfile publicado desde SCM

La rama feature/two-containers fue publicada por SSH con autorización expresa y size:exception. El job original `meli-entrega2` obtuvo el Jenkinsfile de GitHub y el build2 terminó SUCCESS. No utilizó el archivo local ni el harness de validación anterior.

El build2 fue manual para registrar el trigger declarativo Poll SCM `H/5 * * * *`. Este cambio documental registra esa evidencia y sirve como cambio real para verificar el disparo automático posterior. No se lanzará ese siguiente build manualmente. El merge a main permanece condicionado a las verificaciones.

## Resultado del disparo automático

Build3 terminó SUCCESS sobre f2756ad0ed29bf66c35f52b416ce44646b327792. Consola: `Started by an SCM change`; causa registrada: `hudson.triggers.SCMTrigger_-SCMTriggerCause`. No se lanzó manualmente. GitHub Actions también pasó para ese commit: https://github.com/andresfccy/meli-integracion-continua/actions/runs/44809092014.

## Destino main pendiente de decisión

La consulta SSH de refs encontró únicamente feature/two-containers, que también es HEAD remoto. No existe main remoto. El main local 00bb6b4 no comparte ancestro con esta rama. No se mezclaron historiales ni se sobrescribió main local. Hace falta decidir si crear main remoto desde la rama verificada; no es un merge convencional sobre una rama remota existente.

## Destino autorizado creado

Tras autorización adicional, main remoto fue creado desde f2756ad. Verificación SSH confirmó el mismo SHA; main local sigue en 00bb6b4e41669a8d416a17cb8fa61e98c9cf3e50. No se forzó un merge de historiales no relacionados ni se cambió la rama predeterminada.
