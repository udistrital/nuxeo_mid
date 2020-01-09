# NUXEO_MID

Api intermediaria del cliente y el api de nuxeo.

Esta api tiene inicialmente la finalidad de disparar el flujo a los documentos de nuxeo, el cual luego podra ser finalizado si se rechaza o ser continua el flujo y permitir que los documentos queden publicos.

Inicialmente es un requerimiento para el cliente de `CUMPLIDOS` , pero se espera que se pueda usar de forma generica.

Los siguientes diagramas de secuencia evidencian el proceso de disparo de flujo y de aprobacion.


# Instalación
Para instalar el proyecto de debe relizar lo siguientes pasos:

Ejecutar desde la terminal 'go get repositorio':
```shell 
go get github.com/udistrital/nuxeo_mid
```

# Ejecución del proyecto


- Ejecutar: 
```shell 
bee run
```
- O si se quiere ejecutar el swager:

```shell 
bee run -downdoc=true -gendoc=true
```

### EndPoints
|  Funcion |Tipo de peticion                  |Parametros| Endpoint |
|----------------|------------------------|---------------------|-------------------|
| **Disparar el flujo a un determinado documento** | **POST** |`DocID`   id del documento| ```workflow/[DocID]```|
| **Eliminar un flujo de un documento** | **DELETE** | `DocID`   id del documento |```workflow/[DocID]```|
| **Aprovar un documento** | **POST** | `DocID`   id del documento |```validacion/[DocID]```|


---

### Disparo de flujo

![disparo FLUJO](https://user-images.githubusercontent.com/28914781/65219366-08ce8f00-da7e-11e9-8a82-7832c0ee5384.png)

---

### Aprobación de documentos.

![aprobacion doc](https://user-images.githubusercontent.com/28914781/65219477-3e737800-da7e-11e9-8192-4600d4c8f7ef.png)

## Licencia

This file is part of cumplidos-cliente.

cumplidos-cliente is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Foobar is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Foobar. If not, see https://www.gnu.org/licenses/.