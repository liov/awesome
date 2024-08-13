
from pygerber.gerberx3.api.v2 import GerberFile

GerberFile.from_file("D:/work/Gerber/Gerber_TopLayer.GTL").parse().render_svg("output.svg")