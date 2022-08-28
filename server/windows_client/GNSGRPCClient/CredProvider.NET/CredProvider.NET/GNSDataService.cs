namespace GNSGRPCClient
{
    using GNSRPC;
    using Grpc.Core;
    using Grpc.Net.Client;
    using System.Threading;
    public class GNSDataService
    {
        GNSBadgeData.GNSBadgeDataClient? _client;
        public delegate void Callback(GNSRPC.CardStatus status);

        public async void Init(Callback cb)
        {

            using var channel = GrpcChannel.ForAddress("http://localhost:50051");
            _client = new GNSBadgeData.GNSBadgeDataClient(channel);

            GNSRPC.CardStatus _disconnected = new GNSRPC.CardStatus();
            _disconnected.Type = CardStatus.Types.ConnectionType.Usb;
            _disconnected.Status = CardStatus.Types.ConnectionStatus.Disconnected;
            //default to disconnected
            cb(_disconnected);
            while (true)
            {
                try
                {
                    Task task = Task.Run(async () =>
                    {
                        using var stream = _client.StreamCardStatus(new GNSBadgeDataParam());

                        await foreach (var response in stream.ResponseStream.ReadAllAsync())
                        {
                            cb(response);
                        }
                    });
                    await task;
                }
                catch (Exception ex)
                {
                    cb(_disconnected);
                    Console.WriteLine("Lost card status stream. Attempting to reconnect");
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