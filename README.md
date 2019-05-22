# TallerGo
Ejercicio de Go

Definir un nuevo endoint, consumir  de la APi de users, de ahi de acuerdo al site y country del mismo, consumidr de esas dos Apis,
obtener la informacionusando Gorutine, meter toda esa informacion en un canal y retornar la información en una structura que contengar user, site, 
country  a partir de su User ID respetando las reponsabilidades
-Controllers : validar parametros + Llamar servicios
-Services : Consumir la Api de Users
-Domain: Representar las estructuras de dominio

Para la implementacion me base en el repositorio de Emiliano Kohmann (https://github.com/emikohmann/academy-myml.git y
https://github.com/emikohmann/go-concurrency-patterns.git) quien nos dio el taller de GO
