[cold/warm boot] icemulti

El objetivo de este hilo es recopilar información sobre la característica "warm/cold boot" disponible en las FPGAs de la familia ICE40 de Lattice, y compartir/coordinar las investigaciones/ñapas/contribuciones para proveer prototipos autónomos funcionales que permitan cambiar entre más de cuatro imágenes.

La fuente de información principal es el documento "iCE40 Programming and Configuration", facilitado por Lattice (Technical Note TN1248): http://www.latticesemi.com/-/media/LatticeSemi/Documents/ApplicationNotes/IK/iCE40ProgrammingandConfiguration.ashx Concretamente, las páginas 14-17.

A partir de un intercambio de comentarios en otro hilo, en marzo de 2017 iniciamos una fructífera conversación. En dicha conversación se desgranan los detalles de la TN1248 y se investiga qué funcionalidades están ya incluidas en el proyecto icestorm:

http://www.clifford.at/icestorm/format.html
https://github.com/cliffordwolf/icestorm



- Empaquetador alternativo a icemulti. Posible contribución a icestorm.

- CLI para cambiar la imagen activa por USB.
- LUT y FSM en Verilog/VHDL para cambiar la imagen activa desde la FPGA por SPI.
