using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace Wpf;

    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();
        }
        
        private void PointCloud_KTV(object sender, RoutedEventArgs e)
        {
            PointCloud_VTK pointCloud = new PointCloud_VTK();
            pointCloud.Show();
        }
        
        private void PointCloud_Opencv(object sender, RoutedEventArgs e)
        {
            PointCloud_HelixToolkit pointCloud = new PointCloud_HelixToolkit();
            pointCloud.Show();
        }
        
        private void Camera(object sender, RoutedEventArgs e)
        {
            Focus focus = new Focus();
            focus.Show();
        }
    }
