
using System.Threading.Tasks;
using Grpc.Net.Client;

namespace Client;

public class Client
{
    public Greeter.GreeterClient GreeterClient
    {
        get;
    }

    public Client(string addr)
    {
        using var channel = GrpcChannel.ForAddress(addr);
        GreeterClient = new Greeter.GreeterClient(channel);
    }

}