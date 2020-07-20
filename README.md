
Este proyecto tiene como objeto crear un servidor de servicios web REST API con las siguientes características:

- Se envía en el body un request json un correo, dia y hora al servicio IP:PTO/auth por el método POST, este responde el token.

**Ejemplo: **
request 
POST localhost:9990/auth 
{ "email":"gxbaquero@gmail.com", "day":"2020-07-20", "hour":"12:00" } 
response 
{ "token": "MTU5NTI4MDg5NTMyODIxMDQyOT1neGJhcXVlcm9AZ21haWwuY29t" }

- con este token de validez una hora, se autentica usando la opción de 'Authorization: Bearer MTU5NTI4MDg5NTMyODIxMDQyOT1neGJhcXVlcm9AZ21haWwuY29t' el servidor valida token contra email y se llama el nombre del archivo al servicio IP:PTO/api/{archivo} por el método POST, este responde todo el contenido del archivo 

**Ejemplo: **
request 
POST localhost:9990/api/USAstates 
{ "EMAIL":"gxbaquero@gmail.com" } 
response 
{ "AL": "Alabama", "AK": "Alaska", "AS": "American Samoa", "AZ": "Arizona", "AR": "Arkansas", "CA": "California", "CO": "Colorado", "CT": "Connecticut", "DE": "Delaware", "DC": "District Of Columbia", "FM": "Federated States Of Micronesia", "FL": "Florida", "GA": "Georgia", "GU": "Guam", "HI": "Hawaii", "ID": "Idaho", "IL": "Illinois", "IN": "Indiana", "IA": "Iowa", "KS": "Kansas", "KY": "Kentucky", "LA": "Louisiana", "ME": "Maine", "MH": "Marshall Islands", "MD": "Maryland", "MA": "Massachusetts", "MI": "Michigan", "MN": "Minnesota", "MS": "Mississippi", "MO": "Missouri", "MT": "Montana", "NE": "Nebraska", "NV": "Nevada", "NH": "New Hampshire", "NJ": "New Jersey", "NM": "New Mexico", "NY": "New York", "NC": "North Carolina", "ND": "North Dakota", "MP": "Northern Mariana Islands", "OH": "Ohio", "OK": "Oklahoma", "OR": "Oregon", "PW": "Palau", "PA": "Pennsylvania", "PR": "Puerto Rico", "RI": "Rhode Island", "SC": "South Carolina", "SD": "South Dakota", "TN": "Tennessee", "TX": "Texas", "UT": "Utah", "VT": "Vermont", "VI": "Virgin Islands", "VA": "Virginia", "WA": "Washington", "WV": "West Virginia", "WI": "Wisconsin", "WY": "Wyoming" }

Para el correcto funcionamiento debe existir una carpeta services/ y dentro de ella un archivo buffer.json VACIO
