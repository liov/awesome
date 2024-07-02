// See https://aka.ms/new-console-template for more information

using Gerber;
var mainProject = Gerbv.gerbv_create_project();
Gerbv.gerbv_open_layer_from_filename(mainProject, "D:/work/Gerber/m1");
Gerbv.gerbv_export_svg_file_from_project_autoscaled(mainProject, "output.svg");
Gerbv.gerbv_destroy_project(mainProject);