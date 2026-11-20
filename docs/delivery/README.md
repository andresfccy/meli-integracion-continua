# Entrega 2 Jenkins

Estado **parcial**, sin push ni envío a Canvas. [Informe editable](Informe_entrega2_Jenkins.docx), [PDF](Informe_entrega2_Jenkins.pdf), [evidencias](../../evidence/entrega2/README.md) y [matriz](../../evidence/entrega2/compliance.json).

Jenkins tiene controlador, agente conectado y credencial SCM utilizada sin revelar su valor. El job `meli-entrega2` está configurado desde GitHub, pero su build1 falla porque el remoto no contiene Jenkinsfile.

El job separado `meli-entrega2-validacion-local` probó el candidato ac45ad9: build1 falla de forma controlada y build2 pasa, con ambas imágenes, Go/vet, 7 pruebas frontend y comunicación real. El checkout remoto y el archivo local verificado por hash se mantienen diferenciados. Esto no sustituye la ejecución del Jenkinsfile publicado ni prueba el trigger automático.

Pendientes: integrantes, autorización de publicación y ejecución del job SCM original, disparo automático observado. Los documentos se actualizan con resultados reales; no se garantiza puntuación.
