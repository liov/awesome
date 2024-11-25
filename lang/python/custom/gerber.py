
from pygerber.gerberx3.api.v2 import GerberFile

GerberFile.from_file(r"XXX").parse().render_raster("output.jpg")