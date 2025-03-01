# Resumen

Esta pequeña e interesante aplicación se trata de ayudar a tres civilizaciones de una galaxia lejana, dicha galaxia esta conformada por los siguientes planetas y civilizaciones los Feregis, Vulcanos y Betasoides. Cada uno de estos planetas gira al rededor de un sol como el nuestro. 

Cada uno de estos planetas e incluso el sol impactan en el clima, causando sequías, lluvias, climas agradables pero también se pueden dar los tan temidos monzones (Lluvias intensas), capaces de causar terribles inundaciones. 

El problema que se presenta en predecir lo más posible como estará el clima, para que cada civilización pueda tomar las medidas para reducir los daños, y para ello venimos en su ayuda. 

# Datos recolectados 

Cada planeta gira en una órbita diferente, y a distancias totalmente distintas entre cada uno. Feregni esta a unos 500 kilómetros cerca de su estrella, mientras que Vulcano el planeta siguiente en la órbita esta a unos 1000 kilómetros y por último tenemos el que sería Urano para el sistema solar y es Betasoide que se encuentra a unos increíbles 2000 kilómetros.

- Ferengi: 500 kilómetros de distancia del sol y una velocidad angular de 1 grados/día.
- Vulcano: 1000 kilómetros de distancia del sol y una velocidad angular de 3 grados/día.
- Betasoide: 2000 kilómetros de distancia del sol y una velocidad angular de 5 grados/día en dirección anti-horario.

Por lo cual su movimiento de traslación varía con respecto a cada planeta, lo que quiere decir que: 

- Ferengi dura 360 días en dar una vuelta completa al rededor del sol.
- Vulcano dura 72 días en dar la vuelta completa.
- Betasoide dura 120 días en dar la vuelta completa.

# Clima a calcular

- Sequía: cuando los tres planetas están alineados entre sí y a su vez alineados con respecto al sol, el sistema solar experimenta un período de sequía.
- Lluvias: Cuando los tres planetas no están alineados, forman entre sí un triángulo. En el momento en el que el sol se encuentra dentro del triángulo, el sistema solar experimenta un período de lluvia.
- Lluvia intensas: Se presenta un pico de intensidad cuando el perímetro del triángulo está en su máximo tamaño.
- Las condiciones óptimas de presión y temperatura se dan cuando los tres planetas están alineados entre sí pero no están alineados con el sol.

# Suposiciones para dar con una posible solución
- La recolección de los datos se iniciará desde el día cero, mientras todos los planetas arrancan juntos y estarán ubicados en el plano positivo de las X en el plano cartesiano.
- El año se supone de 360 días, tiempo en que el planeta más cercano al sol(ferengi) hará un giro completo.

# Calculos utilizados para el ejercicio

- Se utilizó la fórmula para calcular movimiento circular uniforme sabiendo el radio y  los grados de desplazamiento, esto nos ayudará a saber cuanto 
se ha movido el planeta y así conocer su nueva posición y poder calcular según las reglas establecidas. 
- Para luego poder utilizar diferentes formulas para determinar si los puntos están rectos entre sí 
- Otra ecuación diferente fue realizada para poder saber si el sol se encuentra dentro del  área que forma el triángulo para poder así establecer las reglas de clima que se piden en el problema.

# Detalles técnicos

## Instalación
Se puede instalar como una librería

## Datos de conexión a la base de datos

- El manejador de base de datos utilizado fue Mysql
- La base de datos debe estar previamente creada con el nombre "galaxia"
- Los datos de conexión estarán ubicados en el archivo ".env" dentro de la carpeta "src".

# ¿Cómo utlizar el API?

El api debe ser llamada a través del puerto 9990/clima?dia=11 el cual retornará un JSON

GET → http://www.sendurtweet.com/clima?dia=1   → Respuesta: {“dia”:566, “clima”:”lluvia”}

- El parámetro día debe ser númerico 
