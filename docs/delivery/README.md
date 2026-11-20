# Entrega 2 Jenkins

Estado **parcial**, no lista para enviar todavía. Se incluyen [informe editable](Informe_entrega2_Jenkins.docx), [PDF](Informe_entrega2_Jenkins.pdf), [evidencias reales](../../evidence/entrega2/README.md) y [matriz de cumplimiento](../../evidence/entrega2/compliance.json).

Las ocho páginas del PDF se renderizaron desde el DOCX y se revisaron visualmente. Sin desbordamientos ni recortes. Los integrantes deben confirmarse; no se inventaron. El controlador Jenkins está instalado y el agente `docker-local` conectado por WebSocket. Faltan el job, las credenciales SCM autorizadas, ejecutar el pipeline y comprobar el trigger remoto. No se hicieron push ni envío a Canvas.

La prueba local pasa; no equivale a evidencia de ejecución Jenkins. Antes de entregar, actualizar ambos documentos con las capturas de etapas, consola e historial del build fallido y posterior corregido, más disparo automático observado.

La imagen del agente, sus herramientas y el acceso autorizado al motor Docker se comprobaron en contenedores efímeros. El usuario creó el administrador privadamente; el nodo está conectado con un ejecutor y etiqueta `docker` exclusiva.
