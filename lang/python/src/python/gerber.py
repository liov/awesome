import gerber
from gerber.render import GerberCairoContext

# Read gerber and Excellon files
top_copper = gerber.read('D:/worker/Gerber/m1')


# Rendering context
ctx = GerberCairoContext()

# Create SVG image
top_copper.render(ctx, 'composite.svg')