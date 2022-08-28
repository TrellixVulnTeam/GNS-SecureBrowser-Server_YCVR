//
//  ContentView.swift
//  GNSInternalTools
//
//  Created by Peter Pham on 2/23/22.
//

import SwiftUI
import NIOCore

struct Credential: Hashable {
    let idx: UInt32
    let username: String
    let password: String
    let code: String
    let id = UUID()
}
struct ContentView: View {
    
    @State var cardStatus: String = "status: disconnected"
    @State var client: GNSRPC_GNSBadgeDataClient?
    @State var ready: Bool = false
    @State var siteTable: [Credential] = []
    @State var freesiteTable: [UInt32] = []
    @State private var showPopUp: Bool = false
    @State var myuuid: String = ""
    
    func saveSite() { }
    var body: some View {
        ZStack {
            NavigationView {
                ZStack(alignment: .center) {
                    Color(#colorLiteral(red: 0.737254902, green: 0.1294117647, blue: 0.2941176471, alpha: 1)).ignoresSafeArea()
                    
                    HStack(alignment: .center) {
                        Button {
                            self.disabled(true)
                            var uuid = GNSRPC_UUID.init()
                            uuid.mode = 1
                            uuid.uuid = ""
                            let result = api().formatCard(client: client!, mode: uuid)
                            result.whenSuccess { a in
                                print("Format successful")
                                self.disabled(!ready)
                            }
                        } label: {
                            Text("FormatCard")
                        }.self.disabled(!ready)
                        
                        Button {
                            let result = api().readSites(client: client!)
                            result.whenSuccess { sites in
                                print(sites.sites)
                                siteTable.removeAll()
                                for a in sites.sites{
                                    print(a)
                                    siteTable.append(Credential(idx: a.idx, username: a.username, password: a.password, code: a.code))
                                }
                                print("Site table is: ")
                                print(siteTable)
                            }
                        } label: {
                            Text("Read Sites")
                        }.self.disabled(!ready)
                        
                        Button(action: {
                            
                            let result = api().getSites(client: client!)
                            result.whenSuccess { Sites in
                                print(Sites.idx)
                                freesiteTable.removeAll()
                                for a in Sites.idx {
                                    freesiteTable.append(a)
                                }
                                withAnimation(.linear(duration: 0.3)) {
                                    showPopUp.toggle()
                                }
                            }
                            
                        }, label: {
                            Image(systemName: "plus")
                        }).self.disabled(!ready)
                        Text(cardStatus)
                            .padding().onAppear {
                                do {
                                    client = try api().Connect()
                                    let uuid = client?.readUUID(GNSRPC_GNSBadgeDataParam.init())
                                    client?.streamCardStatus(GNSRPC_GNSBadgeDataParam.init(), callOptions: nil, handler: { status in
                                        switch status.status {
                                        case .disconnected:
                                            cardStatus = "status: disconnected"
                                            ready = false
                                        case .authenticated:
                                            cardStatus = "status: authenticated"
                                            ready = true
                                            
                                            if(myuuid == "") {
                                                let result = api().readUUID(client: client!)
                                                result.whenSuccess { a in
                                                    print(a)
                                                    myuuid = a.uuid
                                                }
                                            }
                                        case .connected:
                                            cardStatus = "status: connected"
                                            ready = false
                                        case .UNRECOGNIZED(_):
                                            cardStatus = "status: unknown"
                                            ready = false
                                        case .unlockedMode:
                                            cardStatus = "status: unlockedMode"
                                        case .unlockedModeReady:
                                            cardStatus = "status: unlockedModeReady"
                                        }
                                    })
                                    
                                }
                                catch {
                                    print("We got an error calling the api")
                                    
                                }
                            }
                        
                    }.frame(width: 600, height: 50, alignment: Alignment.center)
                    
                }
                
            }.frame(minWidth: 600, idealWidth: 600, maxWidth: 800, minHeight: 50, idealHeight: 50, maxHeight: 50, alignment: Alignment.center)
            
        }
        AddSitePopup(title: "Add site", message: "", buttonText: "Save", show: $showPopUp,
                     freesiteTable: $freesiteTable, selectedIdx: 1, username: "", password: "", code: "", client: $client)
        VStack {
            Text("UUID: " + myuuid)
            
            ForEach(siteTable, id: \.self) { site in
                HStack {
                    Text("Idx: " + String(site.idx))
                    Text("Code: " + site.code)
                    Text("Username: " + site.username)
                }
            }
            
        }.frame(width: 600, height: 400, alignment: Alignment.top)
    }
    
    
    struct ContentView_Previews: PreviewProvider {
        static var previews: some View {
            ContentView()
        }
    }
    
}
