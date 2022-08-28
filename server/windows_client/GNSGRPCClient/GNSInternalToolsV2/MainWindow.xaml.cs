using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography.X509Certificates;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;
using GNSGRPCClientCore;
using GNSRPC;
using Grpc.Core;

namespace GNSInternalToolsV2
{
    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        bool firstTime;
        GNSDataService service;
        GNSRPC.Sites sites;
        GNSRPC.WinCreds wincreds;
        public static string _uuid, _password, _username, _domain, _index, _misc;

        public static string _siteCode, _siteOffset, _siteUsername, _sitePassword;
        public MainWindow()
        {
            InitializeComponent();
            service = new GNSDataService();
            string certThumbPrint = "4867d62d9a8016a88bc46e17716e8e096c0100e9";
            service.Init(ServiceCallBack, certThumbPrint);

            sites = new GNSRPC.Sites();
            SiteUsernameInput.TextChanged += SiteUsernameInput_TextChanged;
            SitePasswordInput.TextChanged += SitePasswordInput_TextChanged;
            CodeInput.TextChanged += CodeInput_TextChanged;
            MiscInput.TextChanged += MiscInput_TextChanged;
            MiscInput.Text = null;

            wincreds = new GNSRPC.WinCreds();
            DomainInput.TextChanged += DomainInput_TextChanged;
            UsernameInput.TextChanged += UsernameInput_TextChanged;
            PasswordInput.TextChanged += PasswordInput_TextChanged;


            firstTime = true;
            checkBox.Unchecked += CheckBox_Unchecked;
            Zone2UUID.IsReadOnly = true;
            UUIDLabel.IsReadOnly = true;

        }

        private void MiscInput_TextChanged(object sender, TextChangedEventArgs e)
        {
            _misc = MiscInput.Text;
        }

        private void SitePasswordInput_TextChanged(object sender, TextChangedEventArgs e)
        {
            _sitePassword = SitePasswordInput.Text;
        }

        private void CodeInput_TextChanged(object sender, TextChangedEventArgs e)
        {
            _siteCode = CodeInput.Text;
        }

        private void SiteUsernameInput_TextChanged(object sender, TextChangedEventArgs e)
        {
            _siteUsername = SiteUsernameInput.Text;
        }

        private void checkBox_Checked(object sender, RoutedEventArgs e)
        {

              service.SetUnlockMode(true);

        }
        private void CheckBox_Unchecked(object sender, RoutedEventArgs e)
        {
            service.SetUnlockMode(false);
        }

        private void button_Click(object sender, RoutedEventArgs e)
        {
            StoreUUIDBtn.IsEnabled = false;



            try
            {
                service.StoreUUIDZ2();
                var z2uuid = service.ReadUUIDZ2();
                Zone2UUID.Text = z2uuid.Uuid;

            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
            MessageBox.Show("Store UUID completed");
            StoreUUIDBtn.IsEnabled = true;
        }

        private void AddWinCredBtn_Click(object sender, RoutedEventArgs e)
        {
            if (_domain == null) _domain = "";
            if (_password == null) _password = "";
            if (_username == null) _username = "";

            var wincred = new GNSRPC.WinCred();
            wincred.Username = _username;
            wincred.Password = _password;
            wincred.Domain = _domain;



            var freewincred = service.GetFreeWinCreds();
            if (freewincred.Idx.Count < 1)
            {
                MessageBox.Show("No more room on card");
            }
            wincred.Idx = freewincred.Idx[0];

            try
            {
                service.WriteWinCred(wincred);
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }

            MessageBox.Show("Add windows credential completed");

            //SitesResult.ItemsSource = null;
        }

        private void GetWinCredBtn_Click(object sender, RoutedEventArgs e)
        {
            Result.ItemsSource = null;

            try
            {
                wincreds = service.ReadWinCreds();
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }

            Dispatcher.BeginInvoke((Action)(() => {
                Console.WriteLine("Site cred btn clicked");
                Result.ItemsSource = wincreds.Wincreds;
            }));
        }


        private void DomainInput_TextChanged(object sender, TextChangedEventArgs e)
        {
            _domain = DomainInput.Text;
        }


        private void UsernameInput_TextChanged(object sender, TextChangedEventArgs e)
        {
            _username = UsernameInput.Text;
        }

        private void PasswordInput_TextChanged(object sender, TextChangedEventArgs e)
        {
            _password = PasswordInput.Text;
        }

        void ServiceCallBack(GNSRPC.CardStatus status)
        {
            Console.WriteLine("Service call back" + status.Status.ToString());
            Dispatcher.BeginInvoke((Action)(() =>
            {
                ReadyStatus.Content = status.Status.ToString();
                GetWinCredBtn.IsEnabled = false;
                unlockBtn.IsEnabled = false;
                FormatBtn.IsEnabled = false;
                siteCredBtn.IsEnabled = false;
                AddWinCredBtn.IsEnabled = false;
                AddSiteCredBtn.IsEnabled = false;
                StoreUUIDBtn.IsEnabled = false;
                //checkBox.IsChecked = false;

                if ((string)ReadyStatus.Content == "Authenticated")
                {
                    GetWinCredBtn.IsEnabled = true;
                    unlockBtn.IsEnabled = true;
                    FormatBtn.IsEnabled = true;
                    siteCredBtn.IsEnabled = true;
                    AddWinCredBtn.IsEnabled = true;
                    AddSiteCredBtn.IsEnabled = true;
                    StoreUUIDBtn.IsEnabled = true;
                    if (firstTime)
                    {
                        try
                        {
                            var uuid = service.ReadUUID();
                            UUIDLabel.Text = uuid.Uuid;

                            var uuidz2 = service.ReadUUIDZ2();
                            Zone2UUID.Text = uuidz2.Uuid;
                        }
                        catch (Exception ex)
                        {
                            UUIDLabel.Text = "error reading";
                            Zone2UUID.Text = "error reading";


                        }

                        firstTime = false;
                    }

                }
                else if ((string)ReadyStatus.Content == "UnlockedModeReady")
                {
                    if((bool)!checkBox.IsChecked)
                    {
                        checkBox.IsChecked = true;
                    }
                    unlockBtn.IsEnabled=true;
  
                }
                else if ((string)ReadyStatus.Content == "Disconnected")
                {
                    UUIDLabel.Text = "";
                    Zone2UUID.Text = "";
                    firstTime = true;
                }
            }));

        }

        private void FormatBtn_Click(object sender, RoutedEventArgs e)
        {
            FormatBtn.IsEnabled = false;

            try
            {
                service.Format();
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }

                MessageBox.Show("Formatting completed");

                FormatBtn.IsEnabled = true;

        }

        private void unlockBtn_Click(object sender, RoutedEventArgs e)
        {
            try
            {
                service.StartUnlock();
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
            MessageBox.Show("Unlock completed");
        }

        private void siteCredBtn_Click(object sender, RoutedEventArgs e)
        {

            
            SitesResult.ItemsSource = null;
            
            try
            {
                sites = service.ReadSiteCreds();
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }

            Dispatcher.BeginInvoke((Action)(() => {
                Console.WriteLine("Site cred btn clicked");
                SitesResult.ItemsSource = sites.Sites_;
            }));
        }

        private void AddSiteCredBtn_Click(object sender, RoutedEventArgs e)
        {
            if (_siteUsername == null) _siteUsername = "";
            if (_sitePassword == null) _sitePassword = "";
            if (_siteCode == null) _siteCode = "";
            if (_misc == null) _misc = "";

            var site = new GNSRPC.SiteCred();
            site.Username = _siteUsername;
            site.Password = _sitePassword;
            site.Code = _siteCode;
            site.Misc = _misc;
            var freesites = service.GetFreeSites();
            if (freesites.Idx.Count < 1 )
            {
                MessageBox.Show("No more room on card");
            }
            site.Idx = freesites.Idx[0];

            try
            {
                service.WriteSiteCred(site);
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }

            MessageBox.Show("Add site completed");

            SitesResult.ItemsSource = null;
            /*sites = service.ReadSiteCreds();
            Dispatcher.BeginInvoke((Action)(() => {
                Console.WriteLine("Site cred btn clicked");
                SitesResult.ItemsSource = sites.Sites_;
            }));*/
        }
    }
}
