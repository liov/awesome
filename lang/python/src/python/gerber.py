
from pygerber.gerberx3.api.v2 import GerberFile

GerberFile.from_file("D:/work/Gerber/m1").parse().render_raster("output.jpg")