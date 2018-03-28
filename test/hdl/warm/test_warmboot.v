//-------------------------------------------------------------------
//-- Test para arranque del tipo "warm boot" en la iCE40HX.
//--   Pulsando el botón 1 cambiamos el valor de la imagen a cargar.
//--   Pulsando el botón 2 generamos la señal "boot" para cargar dicha
//--    imagen.
//-------------------------------------------------------------------
//-- Juan Manuel Rico - Ridotech - Marzo 2017.
//-------------------------------------------------------------------
module test_warmboot (input wire btn1, btn2, output reg [7:0] data);

//-- Instanciar el bloque warm boot.
top wb (
    .boot(btn2),
    .s1(image[1]),
    .s0(image[0])
  );

// Registro del valor de la imagen a cargar.
reg [1:0] image = 2'b00;

//-- Al pulsar el botón 1 hacemos cambiar el bit 7
//   para mostrar la pulsación y elegimos la siguiente imagen.
always @(posedge(btn1)) begin
    data[7] = ~data[7];
    image = image + 1;
end

// Se muestra la imagen a cargar tras el warn boot (al pulsar el botón 2).
assign data[6:0] = {5'b00000, image[1], image[0]};

endmodule

