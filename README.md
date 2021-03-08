# go-concurrency-example

## Descripción
:bulb: - Ejemplo simple de concurrencia en Golang, donde tenemos procesos que necesitamos que se corran. Este caso poseemos "órdenes" y "reembolsos" y deben finalizar.
Creamos 2 canales (channels de tipo string), uno para cada uno.

:bulb: - Se chequea en el channel de órdenes  si tiene "lista" la información que se procesó, entonces la imprime. Si , por ejemplo, en el channel de "reembolsos" no está lista, vuelve a chequear al channel anterior para ver si posee información procesada lista; si tiene la muestra.

:bulb: - En este ejemplo, el tiempo de proceso de órdenes es 2 veces mas rápido que la de reembolsos.(es decir, debería procesar en un tiempo 2 ordenes y un reembolso)

:warning: - A tener en cuenta! => siempre debemos cerrar el channel del lado del que envía la información. Nunca del receiver (causa un panic)

![image](https://user-images.githubusercontent.com/32901911/110364877-19344a00-8023-11eb-9fe8-253c60eb57a0.png)
