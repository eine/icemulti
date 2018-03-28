# [IceStrom] Extending dynamic reconfiguration beyond 4 images

## Introduction

This article is related to project [IceStorm](http://www.clifford.at/icestorm) ([cliffordwolf/icestorm](https://github.com/cliffordwolf/icestorm)):

> Project IceStorm aims at documenting the bitstream format of Lattice iCE40
FPGAs and providing simple tools for analyzing and creating bitstream files.

> See http://www.clifford.at/icestorm/ for more information.

> The focus of the project is on the iCE40 LP/HX 1K/4K/8K chips. Most of the work was done on HX1K-TQ144 and HX8K-CT256 parts.

The list of available boards with such chips comprises, but is not limited to: [iCEstick Evaluation Kit](http://www.latticesemi.com/icestick), [iCE40-HX8K Breakout Board](http://www.latticesemi.com/Products/DevelopmentBoardsAndKits/iCE40HX8KBreakoutBoard.aspx), [IceZUM Alhambra](https://github.com/FPGAwars/icezum), [icoBOARD 1.0](http://icoboard.org/about-icoboard.html), [Kéfir I iCE40-HX4K](http://fpgalibre.sourceforge.net/Kefir/), [Nandland Go board](https://www.nandland.com/goboard/introduction.html), [eCow-Logic](http://www.ecowlogic.fr/), [myStorm](https://folknologylabs.wordpress.com/2016/07/21/a-perfect-storm/), [BlackIce](https://mystorm.uk/we-forecast-blackice-this-winter-2/)...

Since these devices are SRAM-based, thus volatile, it is common practice to include a flash memory in any board design. This is used to automatically load a configuration on power-up, and SPI is one of the most used interfaces. As a result, it is common in FPGAs to find hard IP cores implementing the SPI interface.

Furthermore, according to [TN1248: iCE40 Programming and Configuration](http://www.latticesemi.com/dynamic/view_document.cfm?document_id=46502), the hard SPI interfaces in the iCE40 devices support not only the *master* mode required to load from flash. In *slave* mode,

>the iCE40 configuration data can be downloaded from an external processor, microcontroller, or DSP processor using the SPI interface.

On top of that, some devices support so-called *Cold Boot* and *Warm Boot* configuration options. This allows to write up to four images at a time to the flash memory, so that any of them can be loaded on reset, without requiring any addtional external communication.

Therefore, since all the referred boards use *HX* versions of the family, and because an external processor is included in most of them, all four configuration models are available.

NOTE: There is a fifth configuration mode: the one-time programmable NVCM (on-volatile Configuration Memory), which is, naturally, out of the scope of this article.

---

It is out of the scope of this article how to generate bitstream files since it is already explained at ADD REFERENCES. Therefore, several bitstreams of individual designs are expected to be previously generated.

To avoid mixing terms, *image* is used to reference the bitstream corresponding to a single design, and *pack* relates to multiple images packed in a single bitstream.

The naming does not matter, but for the sake of simplicity, they are expected to be named *img??_*.bin* or *pack_*.bin* where `i = 1,...,$number_of_designs` and `*` is any user defined string.

As an example, the files used in this article are available at ADD LINK.

# Scheme

Since most of the boards are based on the iCEstick, it's design is the reference for the experimental tests. As shown in the [iCEstick User Manual](http://www.latticesemi.com/view_document?document_id=50701), the FPGA, the Flash memory and the FTDI chip (which is a processor), are all connected to the same SPI bus:

- FTDI is always a master. Writes/read to/from the FPGA or the Flash.
- Flash memory is always a slave. Is read from the FTDI or the FPGA.
- FPGA is master/slave depending on the configuration mode.
  - In slave mode, is written by the FTDI.
  - In master mode, data is read from the Flash.

On top of that, the FTDI chip controls the programming reset signal of the FPGA.

# Single bitstream

As explained in [TN1248, pp 3-4],

> After exiting the Power-On Reset (POR) state or when CRESET_B returns High after being held Low, the iCE40 device samples the logical value on its SPI_SS_B pin. (...)
  - If the SPI_SS_B pin is sampled as a logic ‘1’ (High), then …
   - If enabled to configure from NVCM, the device configures itself using the Nonvolatile Configuration Memory (NVCM).
   - If not enabled to configure from NVCM, then the device configures using the SPI Master Configuration Interface.
  - If the SPI_SS_B pin is sampled as a logic ‘0’ (Low), then the device waits to be configured from an external controller or from another device in SPI Master Configuration Mode using an SPI-like interface.

Therefore a single bitstream can be directly loaded to the FPGA with:

```
```

That is, the FTDI resets the FPGA by asserting `CRESET_B` and let's it power up in slave SPI mode by keeping `SPI_SS_B` low. Then, the image is written to the SRAM directly. The flash memory ignores any command, because asserting the communication to the FPGA disables the memory's chip select. This is explained in detail in [TN1248, pp 17-20].

However, with the option above, the FPGA will loose it's functionaly as soon as it is powered off. To avoid so, the following command can be used instead:

```
```

This time, the FTDI explicitly holds the FPGA in reset state and asserts the chip select signal of the flash memory. Then, the image is written to the flash memory. When the transference is complete, the reset state is released and the FPGA is powered up in master SPI mode. Therefore, the image written just before is loaded from the flash memory. For more information check [TN1248, pp 10-13]

TODO: to pos 0? Is an applet added?

---

Instead of thoroughly analyzing the details of the format, which is explained [here](http://www.clifford.at/icestorm/format.html), a naive approach is followed:

On the one hand, checking the size reveals that all of the images require the same number of bytes: 31.4 KB (32220 bytes), although 32KB are required on disk.

On the other hand, an hexadecimal dump of the images,

```
$ hexdump -C img01_counter8.bin > img01.dump
$ hexdump -C img02_blink.bin > img02.dump
$ hexdump -C img03_led_on.bin > img03.dump
$ hexdump -C img04_pushbutton_and.bin > img04.dump
```

reveals that, as expected, the first eight bytes are the same:

```
00000000  ff 00 00 ff 7e aa 99 7e  51 00 01 05 92 00 20 62
00000010  01 4b 72 00 90 82 00 00  11 00 01 01 00 00 00 00
00000020  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00
```

NOTE: Actually, in the example images at least the first 32 bytes are the same. However, this may change if more heterogeneous designs sets are used.

# Default Cold Boot

In [TN1248, pp 14-15] the **Cold Boot Configuration Option** is explained. The procedure is roughtly the same as the second one explained in the previous section, but up to four images can be written at the same time. The advantage is that this allows the user to later change from one image to another without requiring an external processor to transfer it.

To support such a feature, an applet is written to the first addresses of the flash. Then, when the cold boot option is enabled,

> the iCE40 FPGA boots normally from power-on or a master reset (CRESET_B = Low pulse), but monitors the value on two PIO pins that are borrowed during configuration (...). These pins, labeled PIO2/CBSEL0 and PIO2/CBSEL1, tell the FPGA which of the four possible SPI configurations to load into the device.

If the applet is written, but the cold boot option is disabled:

> the FPGA configuration starts from the default location (image 0) defined in the Cold/Warm Boot applet.

## Packing images

In practice, this is achieved with a tool named [icemulti](https://github.com/cliffordwolf/icestorm/blob/master/icemulti) from the IceStrom toolchain.

> Usage: icemulti [options] input-files
 - `-c`: coldboot mode, power on reset image is selected by CBSEL0/CBSEL1
 - `-p0, -p1, -p2, -p3`: select power on reset image when not using coldboot mode
 - `-a<n>, -A<n>`: align images at 2^<n> bytes. -A also aligns image 0.
 - `-o filename`:write output image to file instead of stdout

For example, to program four images at a time, by setting the first one as the default and not enabling cold boot:

```
 $ icemulti -p0 -o pack_cp0.bin img01_counter8.bin img02_blink.bin img03_pushbutton_and.bin img04_led_on.bin
 $ iceprog pack_cp0.bin
 init..
 cdone: high
 reset..
 cdone: low
 flash ID: 0x20 0xBA 0x16 0x10 0x00 0x00 0x23 0x54 0x82 0x46 0x06 0x00 0x56 0x00 0x29 0x19 0x01 0x16 0xA4 0xB5
 file size: 130524
 erase 64kB sector at 0x000000..
 erase 64kB sector at 0x010000..
 programming..
 reading..
 VERIFY OK
 cdone: high
 Bye.
```

The same pack can be generated with the second image as the default option by changing `-p0` to `-p1`. If programmed with any, the transference will last longer (because four full packs are being written) but there will be no functional difference, since only the default image will be loaded. However, this is a good starting point to understand how packs are generated.

The size of both packs is the same: 127 KB (130524 bytes), on disk 128KB. It is difficult to see how is the applet fit in there. As done previously, an hexdump of one of the packs is generated. If we compare is with the hexdump of a single image, the starting point of each of them is easily found. Indeed, looking for `ff 7e aa 99 7e` is enough. In the following block only the most meaninful parts are shown:

```
00000000  7e aa 99 7e 92 00 00 44  03 00 01 00 82 00 00 01
00000010  08 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00
00000020  7e aa 99 7e 92 00 00 44  03 00 01 00 82 00 00 01
00000030  08 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00
00000040  7e aa 99 7e 92 00 00 44  03 00 80 00 82 00 00 01
00000050  08 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00
00000060  7e aa 99 7e 92 00 00 44  03 01 00 00 82 00 00 01
00000070  08 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00
00000080  7e aa 99 7e 92 00 00 44  03 01 80 00 82 00 00 01
00000090  08 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00
000000a0  ff ff ff ff ff ff ff ff  ff ff ff ff ff ff ff ff
*
00000100  ff 00 00 ff 7e aa 99 7e  51 00 01 05 92 00 20 62
00000110  01 4b 72 00 90 82 00 00  11 00 01 01 00 00 00 00
*
00008000  ff 00 00 ff 7e aa 99 7e  51 00 01 05 92 00 20 62
00008010  01 4b 72 00 90 82 00 00  11 00 01 01 00 00 00 00
*
00010000  ff 00 00 ff 7e aa 99 7e  51 00 01 05 92 00 20 62
00010010  01 4b 72 00 90 82 00 00  11 00 01 01 00 00 00 00
*
00018000  ff 00 00 ff 7e aa 99 7e  51 00 01 05 92 00 20 62
00018010  01 4b 72 00 90 82 00 00  11 00 01 01 00 00 00 00
```

Then, we can derive the following memory map:

- applet `0x00000000` 256 bytes
- img01 `0x00000100` 32512 bytes
- img02 `0x00008000` 32768 bytes
- img03 `0x00010000` 32768 bytes
- img04 `0x00018000` 32768 bytes

These maddresses match columns 10-12 in lines `00000020`, `00000040`, `00000060` and `00000080`.

Every image, except `img01` is placed in a 32KB section, which makes sense if no compression is used at all. The space for `img01` is smaller, because of th applet. However, since images require 32220 bytes, there are still 292 free bytes. Indeed, there are 3*548+292=1936 free bytes between images in the space `0x00000000-0x0001FFFF`, which can be used for user applications, even if cold boot is active.

Moreover, the hexdump of the other pack is generated. This should be equal to the previous one, except for the default image, which is set to `img02` instead of `img01`. Actually, they differ in a single byte:

```
00000000  7e aa 99 7e 92 00 00 44  03 00 01 00 82 00 00 01 | pack_cp0.bin
00000000  7e aa 99 7e 92 00 00 44  03 00 80 00 82 00 00 01 | pack_cp1.bin
```

It must be noted that bytes 10-12 correspond to the addresses in the memory map above. Therefore, even though *vector addresses (4)* are mentioned in [TN1248, Fig. 11], it is reasonable to think that a single one exists. Indeed, it is confirmed if the source code is checked: [icemulti.cc](https://github.com/cliffordwolf/icestorm/blob/master/icemulti/icemulti.cc) lines 138-140.

TODO: what's the byte-difference between cold-boot active/inactive?
TODO: option "-c" to activate cold boot and select with CBSELx

# Compact icemulti

Since the starting address of each image is explicitly written to an independent pointer, there is no need to reserve 32KB blocks for each image. The following expression can be used, where `x = 0,...,$number_of_images-1`: `x==? 256 ; 256+(x-1)*32220`.

TODO: try icemulti's verbose option first. then implement in go.

# Extended Cold Boot

## Hypothesis

When the applet is present and cold/warm boot is disabled, the configuration defaults to reading a pointer in a fixed position and directly jumping to it. Three bytes (24 bits) are reserved for the pointer, so `0xFFFFFF` is the largest value it can take. As a results, up to `floor((2^24-1)/2^15)=511` images can be addressed, if a memory of at least 16MB (128Mb) is used and 32KB are used for each images. The size of the flash memory in the icestick is 4MB (32Mb), so up to `floor((2^22-1)/2^15)=127` images can be addressed. The expression to compute the address corresponding to a image in position `x`, where `x = 0,...,$number_of_images-1` is `x==0 ? 0x000100 ; (x-1)*0x8000`.

If images are appended without free space between them, slightly larger packs can fit:

```
floor((2^24-1)/32220)=520
floor((2^22-1)/32220)=130
```

However, apparently, this extended memory map can not be addressed through `CBSEL`. But, if the processor is used to update just the pointers, changing between groups of four images in the extended pack can be achieved without impact in the application.

## Tests

- Write `pack_cp0.bin` to flash.
  - Change default pointer only, through FTDI.
  - Rearrange the pointers, through FTDI.
- Pack more than four images and write the binary to flash.
  - Set a fifth image as default (which is not referred by any of the four pointers).
  - Rearrange the poiners, through FTDI.

TODO: CLI to rearrange pointers. measure and compare reconfiguration time.

# Default Warm Boot

instantiate 'top' in the logic: sb_warmboot.v
test_warmboot.v a button to increment a two-bit counter and another one to trigger the reconfiguration

$ apio build
$ icemulti -p0 -o warmboot.bin hardware.bin counter8.bin led_on.bin blink.bin
$ iceprog warmboot.bin

# Extended Warm Boot

LUT and FSM in Verilog/VHDL + hard SPI

http://www.mouser.com/ds/2/225/iCE40FamilyHandbook-311139.pdf
Page 62. User SPI IP.
