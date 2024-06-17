
using System.Threading.Tasks;
using Grpc.Net.Client;

namespace Client;

public class Client
{
    private readonly Greeter.GreeterClient _client;

    public Client(string addr)
    {
        using var channel = GrpcChannel.ForAddress(addr);
        _client = new Greeter.GreeterClient(channel);
    }

    public Greeter.GreeterClient GetGreeterClient()
    {
        return _client;
    }
}