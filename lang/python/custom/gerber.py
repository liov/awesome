
from pygerber.gerberx3.api.v2 import GerberFile
# 有问题,和gerbev不完全一次
GerberFile.from_file(r"XXX").parse().render_raster("output.png")