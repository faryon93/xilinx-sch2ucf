# xilinx-sch2ucf
This little tool converts an **CadSoft Eagle schmeatics file (.sch)** to an **.ucf** constraint file, which can be used with Xilinx ISE.


## Requirements
- schematic file must use the xml file format, which was introduced in version 6
- the library providing the fpga device must have correctly labeled pads (like the ucf file expects)
- nets in eagle should be labeled equally to the names used in hdl code


## Setup
To install xilinx-sch2ucf only the following simple command is necessary:
```
$: go get github.com/faryon93/xilinx-sch2ucf/...
```


## Usage

The filename of the schematic and the eagle name of the fpga are mandatory. You can specify an output file with the **--out** parameter. If no out parameter is supplied the ucf file is written to stdout.
```
$: xilinx-sch2ucf fluxcapacitor.sch IC6
```


## Note
For all testing I used the library **xilinx_devices_V6.lbr**, which can be downloaded at the CadSoft website. This library seems to be pretty well done at meets the requirements for this tool.


## Todo
- support parsing of busses in schmematic