using System;
using GNSGRPCClientNet;
class Program
{
    // The Main() function is the first function to be executed in a program
    static GNSDataService service;
    static void ServiceCallBack(GNSRPC.CardStatus status)
    {
        Console.WriteLine(status.Status.ToString());


    }
    static void Main()
    {
        // Write the string "Hello World to the 
        service = new GNSDataService();
        //service.Init();
        service.Init(ServiceCallBack);
        var test = service.ReadUUID();
        Console.WriteLine(test);
        Console.ReadLine();

    }
}
