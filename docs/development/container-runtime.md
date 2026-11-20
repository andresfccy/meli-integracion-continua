# Configuración de los dos contenedores de aplicación

La aplicación conserva exactamente los servicios `api` y `web`, conectados por la red `application`. El contenedor `web` recibe el destino de la API mediante una variable de entorno, sin reconstruir la imagen para cambiarlo. Jenkins es infraestructura adicional, no un tercer servicio de aplicación.

## Inicio local

Desde la raíz del repositorio:

```sh
cp .env.example .env
docker compose up --build -d --wait
curl --fail http://127.0.0.1:3000/api/health
```

`.env.example` contiene únicamente valores locales no secretos; `.env` permanece ignorado por Git. No se deben guardar credenciales en ninguno de estos valores.

| Variable | Valor predeterminado | Uso |
| --- | --- | --- |
| `WEB_PORT` | `3000` | Puerto del host, siempre enlazado a `127.0.0.1`; `0` solicita un puerto efímero para pruebas aisladas. |
| `API_UPSTREAM` | `api:8080` | Nombre DNS y puerto TCP accesibles desde `web`; no es una URL ni una dirección del navegador. |

Para cambiar la configuración, edite `.env` y vuelva a ejecutar `docker compose up -d --wait`. Si cambia solo el puerto del host, el puerto interno de Nginx continúa siendo `8080`. Para conocer un puerto efímero, ejecute `docker compose port web 8080`.

## Inicio de Nginx sin privilegios

La imagen ejecuta `frontend/docker-entrypoint.sh` como el usuario `nginx`. El script valida `API_UPSTREAM` como nombre DNS con etiquetas de hasta 63 caracteres y puerto decimal entre `1` y `65535`. Rechaza espacios, rutas, esquemas URL, credenciales, literales IPv6 y caracteres de configuración. Una dirección inválida detiene el arranque antes de generar la configuración.

El script sustituye **únicamente** `${API_UPSTREAM}` en la plantilla `frontend/nginx.conf` y escribe `/tmp/nginx.conf`. Las variables propias de Nginx, como `$host`, `$proxy_add_x_forwarded_for` y `$uri`, permanecen intactas. Después inicia Nginx con esa configuración. No depende del procesamiento automático del entrypoint original de la imagen, que aquí se reemplaza explícitamente.

## Persistencia: por qué no se agrega un volumen

`web` sirve archivos estáticos y no guarda datos de negocio. `api` utiliza el almacén en memoria del MVP; sus cambios se pierden al reiniciar el proceso. Montar un volumen no vuelve persistente ese almacén. Por eso no se agrega un volumen ni una base de datos a la aplicación. Esta decisión mantiene la arquitectura existente: [descripción de arquitectura](../architecture/README.md) y [decisiones vigentes](../decisions/README.md). El directorio de trabajo persistente de Jenkins es una necesidad distinta de infraestructura.

## Verificación

La prueba local no necesita Docker ni instala dependencias:

```sh
sh tests/web-runtime.sh
sh -n frontend/docker-entrypoint.sh
docker compose config --quiet
```

La prueba de shell usa dobles de `envsubst` y `nginx` con rutas temporales: verifica valores predeterminados, sustitución permitida, conservación de variables de Nginx, argumentos de inicio y rechazo de entradas inválidas. No sustituye una construcción real ni la prueba de comunicación entre contenedores:

```sh
docker compose exec -T web wget -qO- http://api:8080/api/health
docker compose exec -T web wget -qO- http://127.0.0.1:8080/api/health
```

El primer comando comprueba resolución por nombre de servicio; el segundo atraviesa el proxy de Nginx. Los resultados reales de estas verificaciones deben conservarse como evidencia, no inferirse de estas instrucciones.
