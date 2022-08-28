namespace GNSGRPCClientCore
{
    using GNSRPC;
    using Grpc.Core;
    using Grpc.Net.Client;
    using System.Security.Cryptography.X509Certificates;
    using System.Text;
    using System.Threading;
    public class GNSDataService
    {
        GNSBadgeData.GNSBadgeDataClient? _client;
        public delegate void Callback(GNSRPC.CardStatus status);
        GNSRPC.CardStatus lastStatus;
        
        //Setup a call back method to inform client about card status change.
        //It will also read in a certificate thump print from the Windows OS store to use
        //for mutual TLS authenticatoin with the gRPC server
        public async void Init(Callback cb, string certThumbPrint)
        {


            X509Store certStore = new X509Store(StoreName.My, StoreLocation.LocalMachine);
            // Try to open the store.
            certStore.Open(OpenFlags.ReadOnly);
            // Find the certificate that matches the thumbprint.
            //string certThumbPrint2 = "4867d62d9a8016a88bc46e17716e8e096c0100e9";
            X509Certificate2Collection certCollection = certStore.Certificates.Find(
                X509FindType.FindByThumbprint, certThumbPrint, false);
            certStore.Close();

            var handler = new HttpClientHandler();
            handler.ClientCertificates.Add(certCollection[0]);

            //var handler = new System.Net.Http.SocketsHttpHandler();
            //handler.SslOptions.ClientCertificates = certCollection;
            //handler.ServerCertificateCustomValidationCallback = HttpClientHandler.DangerousAcceptAnyServerCertificateValidator;
            using var channel = GrpcChannel.ForAddress("https://127.0.0.1:50051",
                                                        new GrpcChannelOptions
                                                        {
                                                            HttpHandler = handler,
                                                            //Credentials = channelCredentials,
                                                        }) ;
            
            _client = new GNSBadgeData.GNSBadgeDataClient(channel);

            GNSRPC.CardStatus _disconnected = new GNSRPC.CardStatus();
            _disconnected.Type = CardStatus.Types.ConnectionType.Usb;
            _disconnected.Status = CardStatus.Types.ConnectionStatus.Disconnected;
            //default to disconnected
            cb(_disconnected);
            lastStatus = _disconnected;

                while (true)
                {
                    try
                    {
                        Task task = Task.Run(async () =>
                        {
                            using var stream = _client.StreamCardStatus(new GNSBadgeDataParam());
           
                                await foreach (var response in stream.ResponseStream.ReadAllAsync())
                                {

                                    if (response.ToString() != lastStatus.ToString())
                                    {
                                        cb(response);
                                        lastStatus = stream.ResponseStream.Current;
                                    }
                                }

                        });
                        await task;
                    }
                    catch (Exception ex)
                    {
                        Console.WriteLine(ex.ToString());
                        //cb(_disconnected);
                        Console.WriteLine("Lost card status stream. Attempting to reconnect");
                        //break;
                    }

                }

        }

        public UUID ReadUUID()
        {
            if(_client == null)
            {
                return null;
            }
            else
            {

                return _client.ReadUUID(new GNSBadgeDataParam());

            }

        }

        public UUID ReadUUIDZ2()
        {
            if (_client == null)
            {
                return null;
            }
            else
            {
                return _client.ReadUUIDZone2(new GNSBadgeDataParam());
            }

        }


        public void StoreUUIDZ2()
        {
            if (_client == null)
            {
                return;
            }
            else
            {

                _client.StoreUUID(new GNSBadgeDataParam());

            }

        }


        public FreeSites GetFreeSites()
        {
            if (_client == null)
            {
                return null;
            }
            else
            {

                return _client.GetFreeSites(new GNSBadgeDataParam());

            }

        }
        public Sites ReadSiteCreds()
        {
            if (_client == null)
            {
                return null;
            }
            else
            {

                return _client.ReadSiteCreds(new GNSBadgeDataParam());

            }

        }

        public SiteCred ReadSiteCred(SiteCred input)
        {
            if (_client == null)
            {
                return null;
            }
            else
            {

                return _client.ReadSiteCred(input);

            }

        }

        public void WriteSiteCred(SiteCred input)
        {
            if (_client == null)
            {
                return;
            }
            else
            {

                _client.WriteSiteCred(input);

            }

        }

        public void DeleteSiteCred(SiteCred input)
        {
            if (_client == null)
            {
                return;
            }
            else
            {

                _client.DeleteSiteCred(input);

            }

        }

        public FreeWinCreds GetFreeWinCreds()
        {
            if (_client == null)
            {
                return null;
            }
            else
            {

                return _client.GetFreeWinCreds(new GNSBadgeDataParam());

            }

        }

        public WinCreds ReadWinCreds()
        {
            if (_client == null)
            {
                return null;
            }
            else
            {

                return _client.ReadWinCreds(new GNSBadgeDataParam());

            }

        }

        public WinCred ReadWinCred(WinCred input)
        {
            if (_client == null)
            {
                return null;
            }
            else
            {

                return _client.ReadWinCred(input);

            }

        }

        public void WriteWinCred(WinCred input)
        {
            if (_client == null)
            {
                return;
            }
            else
            {

                _client.WriteWinCred(input);

            }

        }

        public void DeleteWinCred(WinCred input)
        {
            if (_client == null)
            {
                return;
            }
            else
            {

                _client.DeleteWinCred(input);

            }
        }

        public void SetUnlockMode(bool enableUnlockMode)
        {
            if (_client == null)
            {
                return;
            }
            else
            {
                Text input = new Text();
                if (enableUnlockMode)
                {
                    input.Text_ = "UNLOCKED ON";
                }
                else
                {
                    input.Text_ = "UNLOCKED OFF";
                }

                _client.Execute(input);

            }
        }

        public void StartUnlock()
        {
            _client.UnlockCard(new GNSBadgeDataParam());
        }

        public void Format()
        {
            var param = new UUID();
            param.Mode = 0;
            _client.FormatCard(param);
        }
    }
}